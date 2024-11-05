package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Example document path in Docker container
	inputPath := "/app/src/example.pptx"

	pdfBytes, err := convertToPDF(inputPath)
	if err != nil {
		log.Fatalf("PDF conversion failed: %v", err)
	}

	fmt.Printf("PDF successfully generated. Size: %d bytes\n", len(pdfBytes))
}

func convertToPDF(inputPath string) ([]byte, error) {
	outputPath := inputPath[:len(inputPath)-len(filepath.Ext(inputPath))] + ".pdf"

	// Create the command with logging enabled
	cmd := exec.Command("unoconv", "-f", "pdf", "-e", "PageRange=1-1", inputPath)

	// Capture stdout and stderr for debugging
	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error in generating pdf from file: %v; stdout: %s; stderr: %s", err, out.String(), errOut.String())
	}

	pdfBytes, err := os.ReadFile(outputPath)
	if err != nil {
		return nil, fmt.Errorf("error reading generated pdf: %v", err)
	}

	if len(pdfBytes) == 0 {
		return nil, fmt.Errorf("failed reading generated pdf")
	}

	err = os.Remove(outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to delete generated pdf: %v", err)
	}

	return pdfBytes, nil
}
