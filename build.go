//go:generate go run generate.go

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type buildTarget struct {
	os   string
	arch string
}

var targets = []buildTarget{
	{"windows", "amd64"},
	{"linux", "amd64"},
	{"darwin", "amd64"},
	{"darwin", "arm64"},
}

func main() {
	// Create dist directory if it doesn't exist
	if err := os.MkdirAll("dist", 0755); err != nil {
		fmt.Printf("Error creating dist directory: %v\n", err)
		os.Exit(1)
	}

	// Get the module name
	moduleName := "github.com/5e-bits/dndimg"
	outputName := "dndimg"

	for _, target := range targets {
		// Set environment variables for cross-compilation
		os.Setenv("GOOS", target.os)
		os.Setenv("GOARCH", target.arch)

		// Determine the output file extension
		ext := ""
		if target.os == "windows" {
			ext = ".exe"
		}

		// Create the output filename
		outputFile := filepath.Join("dist", fmt.Sprintf("%s-%s-%s%s", outputName, target.os, target.arch, ext))

		// Build the command
		cmd := exec.Command("go", "build", "-o", outputFile, moduleName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Printf("Building for %s/%s...\n", target.os, target.arch)
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error building for %s/%s: %v\n", target.os, target.arch, err)
			os.Exit(1)
		}
	}

	fmt.Println("Build completed successfully!")
}
