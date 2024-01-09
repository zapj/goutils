package fileutils

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
)

func CreateZipFile(source, target string) (err error) {
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer func() {
		err = errors.Join(zipFile.Close())
	}()

	archive := zip.NewWriter(zipFile)
	defer func() {
		err = archive.Close()
	}()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) (currentErr error) {
		if info.IsDir() {
			return
		}

		if currentErr = errors.Join(currentErr, err); currentErr != nil {
			return
		}

		header, currentErr := zip.FileInfoHeader(info)
		if currentErr != nil {
			return currentErr
		}

		header.Method = zip.Deflate
		writer, currentErr := archive.CreateHeader(header)
		if currentErr != nil {
			return currentErr
		}

		file, currentErr := os.Open(path)
		if currentErr != nil {
			return currentErr
		}
		defer func() {
			currentErr = errors.Join(file.Close())
		}()
		_, currentErr = io.Copy(writer, file)
		return
	})
}
