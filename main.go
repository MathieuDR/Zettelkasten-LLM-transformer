// main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Version information - these will be set during build time
var (
	Version   = "development"
	BuildTime = "unknown"
)

func main() {
	// Get the executable name
	exec, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting executable path: %v\n", err)
		os.Exit(1)
	}
	
	// This will print as 'zk-transformer' when properly built
	programName := filepath.Base(exec)
	
	fmt.Printf("%s version %s (built on %s)\n", programName, Version, BuildTime)
	
	// Your existing program logic here...
}
