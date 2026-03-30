package savers

import (
	"archive/zip"
	"io"
	"os"
)

type Compressor struct {
}

func NewCompressor() *Compressor {
	return &Compressor{}
}

func (c *Compressor) Compress(name string) error {
	archive, err := os.Create(name + ".zip")
	if err != nil {
		return err
	}
	defer archive.Close()

	zipper := zip.NewWriter(archive)

	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	w1, err := zipper.Create(name)
	if err != nil {
		return err
	}
	_, err = io.Copy(w1, f)
	if err != nil {
		return err
	}
	zipper.Close()
	return nil
}
