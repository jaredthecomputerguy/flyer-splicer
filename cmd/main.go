package main

import (
	"os"

	"github.com/jaredthecomputerguy/flyer-splicer/internal"
)

const (
	DEFAULT_INPUT_DIR = "./in"
	DEFAULT_OUT_DIR   = "./out"
)

func main() {
	var volume string

	inputDir := DEFAULT_INPUT_DIR
	outputDir := DEFAULT_OUT_DIR

	args := os.Args[1:]
	if len(args) >= 1 {
		inputDir = args[0]
	}
	if len(args) >= 2 {
		outputDir = args[1]
	}
	if len(args) >= 3 {
		volume = args[2]
	}

	inputFileManager := internal.NewFileManager(inputDir)
	outputFileManager := internal.NewFileManager(outputDir)
	outputFileManager.Clean()

	inputFileManager.Sort()

	internal.ProcessFiles(inputFileManager, outputFileManager)
	internal.CopyToVolume(volume, outputFileManager)
}
