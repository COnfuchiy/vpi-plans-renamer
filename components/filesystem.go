package components

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func GetFilesFromDirectory(path string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files = append(files, path)
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func CopyFile(source string, destination string) error {
	inputReader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer func(inputReader *os.File) {
		err = inputReader.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(inputReader)
	outputReader, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer func() {
		err = outputReader.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err = io.Copy(outputReader, inputReader); err != nil {
		return err
	}
	return outputReader.Sync()
}
