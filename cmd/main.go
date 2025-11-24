package main

import (
	"github.com/jaredthecomputerguy/flyer-splicer/internal"
)

const (
	DEFAULT_INPUT_DIR = "./in"
	DEFAULT_OUT_DIR   = "./out"
)

func main() {
	inDir, outDir, volDir := parseArgs()

	inFM := internal.NewFileManager(inDir)
	outFM := internal.NewFileManager(outDir)

	internal.ProcessFiles(inFM, outFM)

	if volDir != "" {
		confirmed := internal.AskForConfirmation("confirm: do you want to erase the files in volume %s?", volDir)
		if confirmed {
			volFM := internal.NewFileManager(volDir)
			volFM.Clean()
		}
		internal.CopyToVolume(volDir, outFM)
	}
}
