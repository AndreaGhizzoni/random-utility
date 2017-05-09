package out

import (
	"errors"
	"fmt"
	"os"
)

// This structure is responsible to know how to write custom structure to file.
type Printer struct {
	file *os.File
}

// This method instance a new printer that knows how to write some data
// structures to given path, according to README.md file.
// Argument path must be a file path (relative/absolute) or empty string to
// print on standard out.
// error is returned if: path leads to a non-readable/non-writable file or
// leads to a directory o
func NewPrinter(path string) (*Printer, error) {
	if path != "" {
		dir, file, errS := sanitizePath(path)
		if errS != nil {
			return nil, errS
		}
		oFile, errF := openFileIfCanRW(dir, file)
		if errF != nil {
			return nil, errF
		}
		return &Printer{oFile}, nil
	}
	return &Printer{os.Stdout}, nil
}

// reusable method to write a single slice on given file
func (p *Printer) writeSingleSlice(slice []int64) error {
	var s string
	for _, v := range slice {
		s = fmt.Sprintf("%d ", v)
		if _, err := p.file.WriteString(s); err != nil {
			return err
		}
	}
	_, err := p.file.WriteString("\n")
	return err
}

// This method write a given slice according to instanced out.Printer.
// error is returned if: slice == nil or there is an i/o error.
func (p *Printer) WriteSlice(slice []int64) error {
	if slice == nil {
		return errors.New("Given slice can not be nil")
	}

	// write header for slice. Check README
	s := fmt.Sprintf("%d\n", len(slice))
	if _, err := p.file.WriteString(s); err != nil {
		return err
	}
	// write slice
	if err := p.writeSingleSlice(slice); err != nil {
		return err
	}

	// Issue a Sync to flush writes to stable storage
	if err := p.file.Sync(); err != nil {
		return err
	}
	return nil
}

// This method write a given matrix according to instanced out.Printer.
// error is returned if: matrix == nil or there is an i/o error.
// This method assumes that every rows in the matrix has the same number of
// elements of matrix[0]
func (p *Printer) WriteMatrix(matrix [][]int64) error {
	if matrix == nil {
		return errors.New("Given matrix can not be nil")
	}

	// write header for matrix. Check README
	// check matrix has at least one row
	if rows := len(matrix); rows != 0 {
		s := fmt.Sprintf("%d %d\n", rows, len(matrix[0]))
		if _, err := p.file.WriteString(s); err != nil {
			return err
		}
	} else {
		return errors.New("Given Matrix has zero rows")
	}

	// write every matrix row into the file
	for _, slice := range matrix {
		if err := p.writeSingleSlice(slice); err != nil {
			return err
		}
	}

	// Issue a Sync to flush writes to stable storage
	if err := p.file.Sync(); err != nil {
		return err
	}
	return nil
}
