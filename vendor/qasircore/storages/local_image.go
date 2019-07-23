package storages

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func CreateImageLocal(dec []byte, image_name string, path_image string) image.Image {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path_image = dir + "/public/" + path_image

	createDirIfNotExist(path_image)

	f, err := os.Create(path_image + image_name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
	image_old := bytes.NewReader(dec)
	image, _, err_decode := image.Decode(image_old)
	if err_decode != nil {
		log.Fatalf("failed to decode image: %v", err_decode)
	}

	return image
}

func ResizeImage(image image.Image, width, height int, path_image, name_image string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path_image = dir + "/public" + path_image
	createDirIfNotExist(path_image)
	dstImage128 := imaging.Resize(image, width, height, imaging.Lanczos)

	// Save the resulting image as JPEG.
	err = imaging.Save(dstImage128, path_image+name_image)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
