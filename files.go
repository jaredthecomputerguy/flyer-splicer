package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"slices"
	"sort"
	"strings"
)

type FileManager struct {
	Dir   string
	Files []string
}

func newFileManager(dir string) *FileManager {
	fm := &FileManager{
		Dir: dir,
	}
	if err := fm.getFiles(); err != nil {
		handleErr(err)
	}
	return fm
}

func (fm *FileManager) Clean() {
	log("clean: cleaning `%s`\n", fm.Dir)
	if len(fm.Files) < 1 {
		logWarning("clean: directory %s is empty\n", fm.Dir)
		return
	}

	var err error

	for _, file := range fm.Files {
		err = os.Remove(path.Join(fm.Dir, file))
		logErrorf("  removed: %s", file)
	}

	fm.Files = nil
	log("clean: finished\n\n")
	handleErr(err)
}

func (fm *FileManager) Sort() {
	sort.Strings(fm.Files)
}

func (fm *FileManager) getFiles() error {
	dirEntries, err := os.ReadDir(fm.Dir)
	if err != nil {
		return fmt.Errorf("getFiles: reading directory: %w", err)
	}

	files := make([]string, len(dirEntries))
	for i, entry := range dirEntries {
		files[i] = entry.Name()
	}

	fm.Files = files

	return nil
}

func processFiles(in *FileManager, out *FileManager) {
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

func copy(inDir, outDir, inFile, outFile string) {
	src, err := os.Open(path.Join(inDir, inFile))
	handleErr(err)
	defer src.Close()

	dest, err := os.Create(path.Join(outDir, outFile))
	handleErr(err)
	defer dest.Close()

	_, err = io.Copy(dest, src)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		logError(err)
		os.Exit(1)
	}
}
