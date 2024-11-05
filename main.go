package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Input and output paths in Docker container
	inputPath := "/app/src/example.pptx"
	outputPath := "/app/src/example.pdf"

	err := convertToPDF(inputPath, outputPath)
	if err != nil {
		log.Fatalf("PDF conversion failed: %v", err)
	}

	fmt.Printf("PDF successfully generated and saved to %s\n", outputPath)
}

// convertToPDF converts the input file to a single-page PDF and saves it to outputPath.
func convertToPDF(inputPath, outputPath string) error {
	// Construct the command with PageRange option
	cmd := exec.Command("unoconv", "-f", "pdf", "-e", "PageRange=1-1", "-o", outputPath, inputPath)

	// Capture stdout and stderr for debugging
	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error generating PDF from file: %v; stdout: %s; stderr: %s", err, out.String(), errOut.String())
	}

	// Verify that the output file was created successfully
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		return fmt.Errorf("PDF file not created at expected location: %s", outputPath)
	}

	return nil
}
