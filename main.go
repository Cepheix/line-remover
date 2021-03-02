package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	sourceLines, _ := readLines(options.SourceFile)
	deleteLines, _ := readLines(options.DeleteTemplate)

	result := make(map[string]bool)

	for _, sourceLine := range sourceLines {
		result[sourceLine] = true

		for _, deleteLine := range deleteLines {
			if strings.Contains(sourceLine, deleteLine) {
				result[sourceLine] = false
			}
		}
	}

	output, err := os.Create(options.OutputFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	for resultLine, shouldPrint := range result {
		if shouldPrint {
			fmt.Fprintln(output, resultLine)
		}
	}

	err = output.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			lines = append(lines, text)
		}
	}
	return lines, scanner.Err()
}
