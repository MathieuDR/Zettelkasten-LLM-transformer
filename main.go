package main

import (
  "errors"
	"fmt"
	"os"
  "flag"
	"path/filepath"
)

// Version information
// these will be set during build time using ldflags
var (
  ProgramName = "ZK LLM Transformer"
	Version   = "development"
	BuildTime = "unknown"
)

type Options struct {
  Summary bool
  Output string
  SingleFile bool
}

const debugTemplate = `
Summary? %t
SingleFile? %t
Output %s
`

func main() {
  opts, optsError := parseFlags()
  if optsError != nil{
    fmt.Fprintln(os.Stderr, "Error:", optsError)
    os.Exit(1)
  }

  fmt.Printf(debugTemplate, opts.Summary,  opts.SingleFile, opts.Output)

  if opts.Summary {
    printSummary()
  }
}

func validateFlags(options Options) error{
  if options.Output == ""{
    return errors.New("No output option defined")
  }

  return nil
}

func parseFlags() (Options, error){
  opts := Options{}

  flag.BoolVar(&opts.Summary, "summary", false, "Show summary")
  flag.BoolVar(&opts.Summary, "s", false, "Show summary")

  flag.BoolVar(&opts.SingleFile, "single-file", true, "Single file output, default true")
  flag.BoolVar(&opts.SingleFile, "f", true, "Single file output, default true")

  flag.StringVar(&opts.Output, "o", "", "Required output destination path")
  flag.StringVar(&opts.Output, "output", "", "Required output destination path")
  flag.Parse()

  if err := validateFlags(opts); err != nil {
    return Options{}, err
  }

  return opts, nil
}


const summaryTemplate = `
%s
Version: %s
Built:   %s
Binary:  %s
Path:    %s
Args:    %v
`

func printSummary() {
	exec, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting executable path: ", err)
		os.Exit(1)
	}
	
	bin := filepath.Base(exec)
  fmt.Printf(summaryTemplate, ProgramName, Version, BuildTime, bin, exec, os.Args[1:])
}
