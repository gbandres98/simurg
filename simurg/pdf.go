package simurg

import (
	"log"
	"os"
	"path"

	"github.com/signintech/gopdf"
)

func GeneratePdfs(pdfDir string, tempDir string) {
	files, err := os.ReadDir(tempDir)
	if err != nil {
		log.Println(err)
	}

	for _, f := range files {
		if !f.Type().IsDir() {
			continue
		}

		joinImages(path.Join(tempDir, f.Name()), path.Join(pdfDir, f.Name()+".pdf"))
	}
}

func joinImages(dir string, file string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Println(err)
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	for _, f := range files {
		pdf.AddPage()
		pdf.Image(path.Join(dir, f.Name()), 0, 0, gopdf.PageSizeA4)
	}

	err = pdf.WritePdf(file)
	if err != nil {
		log.Println(err)
	}
}
