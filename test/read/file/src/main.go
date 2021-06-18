// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	outputDir := os.Getenv("OUTPUT_DIR")
	if outputDir == "-" {
		outputDir = ""
	}

	formula.Formula{
		OutputDir: outputDir,
	}.Run(os.Stdout)
}
