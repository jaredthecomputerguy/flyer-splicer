package main

import (
	"fmt"
	"os"
)

func parseArgs() (inDir, outDir, volDir string) {
	args := os.Args

	if len(args) < 3 {
		printUsage()
		os.Exit(1)
	}

	args = os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Error: input_dir and output_dir are required.")
		fmt.Println()
		printUsage()
		os.Exit(1)
	}

	inDir = args[0]
	outDir = args[1]

	if len(args) >= 3 {
		volDir = args[2]
	}

	return
}

func printUsage() {
	fmt.Println("Flyer Splicer")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  flyer-splicer <input_dir> <output_dir> [volume_dir]")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Println("  input_dir    Directory containing the source files")
	fmt.Println("  output_dir   Directory where processed files will be written")
	fmt.Println("  volume_dir   Optional directory to copy the final output into")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  flyer-splicer ./raw ./out /Volumes/USBDrive")
}
