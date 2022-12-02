package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gbandres98/simurg/simurg"
)

func main() {
	tempDir := filepath.Join(".", "temp")
	err := os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	pdfDir := filepath.Join(".", "pdf")
	err = os.MkdirAll(pdfDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	simurg.DownloadImages("enlaces.txt", tempDir)
	simurg.GeneratePdfs(pdfDir, tempDir)
}
