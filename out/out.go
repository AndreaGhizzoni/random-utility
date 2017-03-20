package out

import (
	"errors"
	"fmt"
	"os"
)

// utility function to check if open file given is a directory or not.
func isDirectory(file *os.File) (error, bool) {
	fileStat, errs := os.Stat(file)
	if errs != nil {
		return nil, errors.New("Error while retriving stats from file")
	}
	return nil, fileStat.IsDir()
}

// utilty function to check if given file is readable or writable.
// nil, err is occurred otherwise.
func openIfCanRW(file string) (*os.File, error) {
	if file == nil {
		return nil, errors.New("Given string cao not be nil.")
	}

	// check if file is readable or writable.
	openFile, err := os.OpenFile(file, os.O_RDWR, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return nil, err
		}
	}

	// check fi given path points to directory
	if err, isDir := isDirectory(openFile); err != nil {
		return nil, err
	} else if isDir {
		return nil, errors.New("Given path is a directory")
	}

	return openFile, nil
}

// reusable method to write a single slice on given file
func writeSingleSlice(slice []int64, file *os.File) error {
	var s string
	for _, v := range slice {
		s = fmt.Sprintf("%d ", v)
		if _, err := file.WriteString(s); err != nil {
			return err
		}
	}
	_, err := file.WriteString("\n")
	return err
}

// TODO add doc
func Write(slice []int64, path string) error {
	if slice == nil {
		return errors.New("Given slice can not be nil")
	}

	var file *os.File
	if file, err := openIfCanRW(path); err != nil {
		return err
	}
	defer file.Close()

	// create and write file header
	s := fmt.Sprintf("%d\n", len(slice))
	if _, err = file.WriteString(s); err != nil {
		return err
	}

	// create and write file body
	if err := writeSingleSlice(slice, file); err != nil {
		return err
	}

	// Issue a Sync to flush writes to stable storage
	if err := file.Sync(); err != nil {
		return err
	}

	return nil
}
