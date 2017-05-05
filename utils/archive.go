package gocopy

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CreateArchive creates zip archive for input files
func CreateArchive(path string, files []string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	w := zip.NewWriter(out)
	defer w.Close()
	for _, file := range files {
		err = addToArchive(w, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func addToArchive(w *zip.Writer, fpath string) error {
	info, err := os.Stat(fpath)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(fpath)
	}

	filepath.Walk(fpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, fpath))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := w.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
