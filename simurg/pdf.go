package simurg

import (
	"log"
	"os"
	"path"

	"github.com/signintech/gopdf"

	"image"
	_ "image/jpeg"
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
		imgPath := path.Join(dir, f.Name())
		r, err := os.Open(imgPath)
		if err != nil {
			log.Println(err)
			continue
		}

		i, _, err := image.Decode(r)
		if err != nil {
			log.Println(err)
			continue
		}

		pageSize := &gopdf.Rect{W: float64(i.Bounds().Dx()), H: float64(i.Bounds().Dy())}

		pdf.AddPageWithOption(gopdf.PageOption{PageSize: pageSize})
		pdf.Image(imgPath, 0, 0, pageSize)
	}

	err = pdf.WritePdf(file)
	if err != nil {
		log.Println(err)
	}
}
