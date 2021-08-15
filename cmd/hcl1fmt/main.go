package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/maxb/hcl/hcl/fmtcmd"
)

func main() {
	var options fmtcmd.Options
	var extensionsCSV string

	flag.BoolVar(&options.List, "list", false,
		"List filenames in need of formatting")
	flag.BoolVar(&options.Write, "write", false,
		"Write reformatted content back to files in place")
	flag.BoolVar(&options.Diff, "diff", false,
		"Display diff of formatting changes")
	flag.IntVar(&options.SpacesWidth, "indent", 2,
		"Number of spaces per indent level")
	flag.StringVar(&extensionsCSV, "extensions", "hcl",
		"File extensions (comma separated) to operate on when processing a directory")

	extensions := strings.Split(extensionsCSV, ",")

	flag.Parse()

	err := fmtcmd.Run(flag.Args(), extensions, os.Stdin, os.Stdout, options)
	if err != nil {
		fmt.Println(err)
	}
}
