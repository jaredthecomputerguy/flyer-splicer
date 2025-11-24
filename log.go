package main

import (
	"fmt"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	reset       = "\033[0m"
)

func log(format string, v ...any) {
	fmt.Printf(format, v...)
}

func logError(err error) {
	fmt.Printf(colorRed+"%s"+reset+"\n", err)
}

func logErrorf(format string, v ...any) {
	fmt.Printf(colorRed+format+reset+"\n", v...)
}

func logWarning(format string, v ...any) {
	fmt.Printf(colorYellow+format+reset+"\n", v...)
}

func logSuccess(format string, v ...any) {
	fmt.Printf(colorGreen+format+reset+"\n", v...)
}
