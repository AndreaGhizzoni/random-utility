package out

import (
	"fmt"
	"os"
)

// TODO add doc
func Write(slice []int64, path string) error {
	// TODO check args
	// TODO !!! check if file is not a binary file || directory || permissions

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// create and write file header
	s := fmt.Sprintf("%d\n", len(slice))
	_, err = file.WriteString(s)
	if err != nil {
		return err
	}

	// create and write file body
	err = writeSingleSlice(slice, file)
	if err != nil {
		return err
	}

	// Issue a Sync to flush writes to stable storage
	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}

// reusable method to write a single slice on given file
func writeSingleSlice(slice []int64, file *os.File) error {
	var s string
	for _, v := range slice {
		s = fmt.Sprintf("%d ", v)
		_, err := file.WriteString(s)
		if err != nil {
			return err
		}
	}
	_, err := file.WriteString("\n")
	return err
}
