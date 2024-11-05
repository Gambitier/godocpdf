package main

import (
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

	cmd := exec.Command("unoconv", "-f", "pdf", "-e", "PageRange=1-1", "-P", "PaperOrientation=landscape", inputPath)

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error in generating pdf from file: %v", err)
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
