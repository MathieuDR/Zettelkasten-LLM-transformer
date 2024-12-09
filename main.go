package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Version information
// these will be set during build time using ldflags
var (
  ProgramName = "ZK LLM Transformer"
	Version   = "development"
	BuildTime = "unknown"
)

func main() {
	exec, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting executable path: %v\n", err)
		os.Exit(1)
	}
	
	exe := filepath.Base(exec)

	fmt.Printf("%s version %s (built on %s, exe is %s, exec is %s\n", ProgramName, Version, BuildTime, exe, exec)
}
