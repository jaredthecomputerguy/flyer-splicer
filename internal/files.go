package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"slices"
	"strings"
)

func ProcessFiles(in *FileManager, out *FileManager) {
	if len(in.Files) < 1 {
		logWarning("copy: input dir is empty. exiting...")
		os.Exit(1)
	}

	log("copy: copying non-flyer files to `%s`\n", out.Dir)

	var flyerFile string
	var prefixes []string

	for _, file := range in.Files {
		if strings.Contains(file, "flyer") {
			flyerFile = file
		} else {
			copy(in.Dir, out.Dir, file, file)
			logSuccess("  copied: %s", file)

			firstLetter := string(file[0])

			if slices.Contains(prefixes, firstLetter) {
				prefixes = append(prefixes, firstLetter+string(file[1]))
			} else {
				prefixes = append(prefixes, firstLetter)
			}
		}
	}
	log("copy: finished\n\n")

	log("copy: copying flyer files to `%s`\n", out.Dir)
	for _, prefix := range prefixes {
		flyerFilePath := prefix + ".mp4"
		copy(in.Dir, out.Dir, flyerFile, flyerFilePath)
		logSuccess("  copied: %s", flyerFilePath)
	}
	log("copy: finished\n\n")

	log("finished processing files!\n\n")

	err := out.getFiles()
	handleErr(err)

	log("final output directory:\n")
	for _, file := range out.Files {
		logSuccess("  - %s", file)
	}

}

func CopyToVolume(vol string, out *FileManager) {
	if vol == "" {
		log("volume: no volume passed, skipping copy\n")
		return
	}
	info, err := os.Stat(vol)
	if err != nil {
		handleErr(fmt.Errorf("volume: cannot stat volume path: %w", err))
	}
	if !info.IsDir() {
		handleErr(fmt.Errorf("volume: provided path is not a directory: %s", vol))
	}

	log("volume: copying files to volume `%s`\n", vol)

	for _, file := range out.Files {
		copy(out.Dir, vol, file, file)
		logSuccess("  copied to volume: %s", file)
	}

	log("volume: finished copying to volume\n\n")

	logSuccess("completed file processing and copy to volume successfully!")
}

func AskForConfirmation(s string, v ...any) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		logWarning("\n%s [(y)es or (n)o]: ", fmt.Sprintf(s, v...))

		response, err := reader.ReadString('\n')
		if err != nil {
			logError(err)
			handleErr(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		switch response {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		}
	}
}

func UnmountVolume(vol string) {
	run := func(cmd string, args ...string) {
		c := exec.Command(cmd, args...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			logError(err)
			handleErr(err)
		}
	}

	// disable indexing
	run("mdutil", "-i", "off", vol)

	// unmount disk
	run("diskutil", "eject", vol)
}

func copy(inDir, outDir, inFile, outFile string) {
	src, err := os.Open(path.Join(inDir, inFile))
	handleErr(err)
	defer func() {
		if err := src.Close(); err != nil {
			logError(err)
			handleErr(err)
		}
	}()

	dest, err := os.Create(path.Join(outDir, outFile))
	handleErr(err)

	defer func() {
		if err := dest.Close(); err != nil {
			logError(err)
			handleErr(err)
		}
	}()

	_, err = io.Copy(dest, src)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		logError(err)
		os.Exit(1)
	}
}
