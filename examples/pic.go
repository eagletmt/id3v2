package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	".."
)

func main() {
	exitcode := 0
	for _, path := range os.Args[1:] {
		err := dump(path)
		if err != nil {
			log.Printf("%s: %s", path, err)
			exitcode++
		}
	}
	os.Exit(exitcode)
}

func dump(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	tag, err := id3v2.Read(file)
	if err != nil {
		return err
	}

	return dumpPicture(path, tag)
}

func dumpPicture(path string, tag *id3v2.Tag) error {
	pic := tag.AttachedPicture()
	if pic == nil {
		return fmt.Errorf("Cannot find APIC frame")
	}

	extension, ok := extensionByType[pic.MimeType]
	if !ok {
		// Assume JPEG
		extension = ".jpg"
	}
	picPath := filepath.Join(filepath.Dir(path), filepath.Base(path)+extension)
	file, err := os.Create(picPath)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Printf("Write picture data to %s\n", picPath)
	file.Write(pic.PictureData)
	return nil
}

var extensionByType = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
}
