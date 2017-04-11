package out

import (
	"errors"
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
