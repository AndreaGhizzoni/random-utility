package out

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Utility function to check if given file is readable or writable.
// To this function must be passed a file path as argument.
// nil, err is occurred otherwise.
func openFileIfCanRW(parentDir, file string) (*os.File, error) {
	var f *os.File
	var e error
	_, err := os.Stat(parentDir + file)

	// if err != nil and file doesn't exists create a new file
	if os.IsNotExist(err) {
		if parentDir != "" {
			os.MkdirAll(parentDir, os.ModePerm)
		}
		f, e = os.Create(parentDir + file)
	}

	// check if file is readable or writable.
	f, e = os.OpenFile(parentDir+file, os.O_RDWR, 0666)
	// e is the error that comes from the creation or opening a new file.
	if e != nil || os.IsPermission(e) {
		return nil, e
	}
	return f, nil
}

// reusable method to sanitize given path and split it in absolute path and
// file name. err != nil if givenPath is empty string or error occurs while
// processing it.
func sanitizePath(givenPath string) (dir, file string, err error) {
	if givenPath == "" {
		return "", "", errors.New("Given path can not be empty string")
	}

	// sanitize path
	path, err := filepath.Abs(givenPath)
	if err != nil {
		return "", "", err
	}
	// split dir and file
	dir, file = filepath.Split(path)
	// if file is empty string means that givenPath leads to a folder
	if file == "" {
		return "", "", errors.New("Given path is a directory")
	}
	return dir, file, nil
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
func WriteSlice(slice []int64, givenPath string) error {
	if slice == nil {
		return errors.New("Given slice can not be nil")
	}

	// sanitize the given path
	dir, file, err := sanitizePath(givenPath)
	if err != nil {
		return err
	}

	// open the file (dir+file) if only and only if passes all checks
	openedFile, err := openFileIfCanRW(dir, file)
	if err != nil {
		return err
	}
	defer openedFile.Close()

	// create and write openedFile header for slice. Check README
	s := fmt.Sprintf("%d\n", len(slice))
	if _, err := openedFile.WriteString(s); err != nil {
		return err
	}
	// create and write openedFile body
	if err := writeSingleSlice(slice, openedFile); err != nil {
		return err
	}

	// Issue a Sync to flush writes to stable storage
	if err := openedFile.Sync(); err != nil {
		return err
	}

	return nil
}

// TODO add doc
func WriteMatrix(matrix [][]int64, givenPath string) error {
	if matrix == nil {
		return errors.New("Given matrix can not be nil")
	}

	// sanitize the given path
	dir, file, err := sanitizePath(givenPath)
	if err != nil {
		return err
	}

	// open the file (dir+file) if only and only if passes all checks
	openedFile, err := openFileIfCanRW(dir, file)
	defer openedFile.Close()
	if err != nil {
		return err
	}

	// TODO check matrix has at least one row

	// write header for matrix. Check README
	s := fmt.Sprintf("%d %d\n", len(matrix), len(matrix[0]))
	if _, err := openedFile.WriteString(s); err != nil {
		return err
	}

	// write every matrix row into the file
	for _, slice := range matrix {
		if err := writeSingleSlice(slice, openedFile); err != nil {
			return err
		}
	}

	// Issue a Sync to flush writes to stable storage
	if err := openedFile.Sync(); err != nil {
		return err
	}
	return nil
}
