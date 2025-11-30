package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type FileManager struct {
	Dir   string
	Files []string
}

func NewFileManager(dir string) *FileManager {
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

	var errs []error

	for _, file := range fm.Files {
		full := filepath.Join(fm.Dir, file)

		info, err := os.Stat(full)
		if err != nil {
			logErrorf("cannot stat %s: %v", file, err)
			errs = append(errs, fmt.Errorf("stat failed for %s: %w", file, err))
			continue
		}

		if info.IsDir() {
			err = os.RemoveAll(full)
		} else {
			err = os.Remove(full)
		}

		if err != nil {
			logErrorf("failed to remove %s: %v", file, err)
			errs = append(errs, fmt.Errorf("remove failed for %s: %w", file, err))
			continue
		}

		logErrorf("  removed: %s", file)
	}

	fm.Files = nil
	log("clean: finished\n\n")

	// Log collected errors, if any
	if len(errs) > 0 {
		logErrorf("clean: encountered %d errors:", len(errs))
		for _, err := range errs {
			logErrorf("  %v", err)
		}
	}
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
