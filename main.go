package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	SourceFile     string `short:"s" long:"source-file" description:"Source file containing all lines"`
	DeleteTemplate string `short:"d" long:"delete-template" description:"Delete template containing all lines that should be removed"`
	OutputFile     string `short:"o" long:"output-file" description:"Path of the file containing result"`
}

func main() {
	var options Options
	_, optionsErr := flags.Parse(&options)

	if optionsErr != nil {
		panic(optionsErr)
	}

	fmt.Printf("Reading lines from %v \n", options.SourceFile)
	fmt.Printf("Removing lines from %v \n", options.DeleteTemplate)
	fmt.Printf("Saving output to %v \n", options.OutputFile)
}
