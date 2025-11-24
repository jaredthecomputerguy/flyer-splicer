package main

import "os"

const (
	DEFAULT_INPUT_DIR = "./in"
	DEFAULT_OUT_DIR   = "./out"
)

func main() {

	inputDir := DEFAULT_INPUT_DIR
	outputDir := DEFAULT_OUT_DIR

	args := os.Args[1:] // skip program name
	if len(args) >= 1 {
		inputDir = args[0]
	}
	if len(args) >= 2 {
		outputDir = args[1]
	}

	inputFileManager := newFileManager(inputDir)
	outputFileManager := newFileManager(outputDir)
	outputFileManager.Clean()

	inputFileManager.Sort()

	processFiles(inputFileManager, outputFileManager)
}
