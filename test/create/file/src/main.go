// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	input1 := os.Getenv("BUCKET_NAME")
	outputDir := os.Getenv("OUTPUT_DIR")
	if outputDir == "-" {
		outputDir = ""
	}

	formula.Formula{
		Name:      input1,
		OutputDir: outputDir,
	}.Run(os.Stdout)
}
