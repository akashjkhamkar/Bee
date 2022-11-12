package utils

import (
	"io"
	"log"
	"os"
)

func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		log.Fatal(err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Fatalf("%s is not a regular file", src)
	}

	source, err := os.Open(src)

	if err != nil {
		log.Fatal(err)
	}

	defer source.Close()

	destination, err := os.Create(dst)

	if err != nil {
		log.Fatal(err)
	}

	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}