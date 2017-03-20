package out

import (
	"errors"
	"fmt"
	"os"
)

// utility function to check if given file is readable or writable.
// nil, err is occurred otherwise.
func openIfCanRW(file string) (*os.File, error) {
	var f *os.File
	var e error
	fileInfo, err := os.Stat(file)

	// if err != nil and file doesn't exists then skip directory check, and
	// create a new file
	if os.IsNotExist(err) {
		f, e = os.Create(file)
	} else {
		// check fi given path points to directory.
		if fileInfo.IsDir() {
			return nil, errors.New("Given path is a directory")
		}

		// check if file is readable or writable.
		f, e = os.OpenFile(file, os.O_RDWR, 0666)
	}

	// e is the error that comes from the creation or opening a new file.
	if e != nil {
		if os.IsPermission(e) {
			return nil, e
		}
	}
	return f, nil
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

// this method write a given slice in path file given. error is returned if:
// slice == nil, can not read/write (or is a directory) to file path or there
// is an i/o error.
func Write(slice []int64, path string) error {
	if slice == nil {
		return errors.New("Given slice can not be nil")
	}

	// TODO check illegal character like "", " ", "\n" ecc

	// open the file if only if passes all checks
	file, err := openIfCanRW(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// create and write file header for slice. Check README
	s := fmt.Sprintf("%d\n", len(slice))
	if _, err := file.WriteString(s); err != nil {
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
