package internal

import (
	"fmt"
	"os"
	"path"
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
