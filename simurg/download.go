package simurg

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"
	"sync/atomic"
)

var queue uint64
var downloaded uint64

func DownloadImages(file string, dir string) {
	lines, _ := readLines(file)

	wg := sync.WaitGroup{}

	for _, line := range lines {
		log.Println(line)

		code := path.Base(line)
		log.Println(code)

		manifestUrl := "http://simurg.csic.es/iiif/collection/object/" + code

		lineDir := filepath.Join(dir, code)
		err := os.MkdirAll(lineDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		processManifest(manifestUrl, lineDir, &wg)
	}

	wg.Wait()
	log.Printf("Descargadas %d imágenes", downloaded)
}

func readLines(path string) (lines []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func processManifest(url string, dir string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var manifest Manifest
	if err := json.NewDecoder(resp.Body).Decode(&manifest); err != nil {
		log.Fatal(err)
	}

	for i, sequence := range manifest.Sequences {
		for j, canvas := range sequence.Canvases {
			for k, image := range canvas.Images {
				atomic.AddUint64(&queue, 1)
				wg.Add(1)
				go downloadImage(
					image.Resource.Service.ID+"/full/full/0/default.jpg",
					filepath.Join(dir, strconv.Itoa(i)+strconv.Itoa(j)+strconv.Itoa(k)+".jpg"),
					wg,
				)
				log.Printf("Descargando %d/%d", downloaded, queue)
			}
		}
	}

}

func downloadImage(url string, path string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	atomic.AddUint64(&downloaded, 1)
	log.Printf("Descargando %d/%d", downloaded, queue)
}
