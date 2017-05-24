package out

import (
	"errors"
	"fmt"
	"github.com/AndreaGhizzoni/zenium/samples"
	"os"
)

// Printer is responsible to know how to write custom structure to file or
// console. Every Printer.WriteX method, where X is a structure, produce
// representation of X in format specified in README.md.
type Printer struct {
	file *os.File
}

// NewPrinter instance a new Printer. Argument path must be a file path
// (relative/absolute) or empty string to print on standard out.
// error is returned if: path leads to a non-readable/non-writable file or
// leads to a directory
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

// reusable method to write a single slice
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

// reusable method to check write synchronization.
func (p *Printer) synchronizeFileWrite () error{
	if err := p.file.Sync(); err != nil {
		return err
	}
    return nil
}

// WriteSlice write the given slice in specified format according to instanced
// out.Printer.
// error is returned if: slice == nil or there is an i/o error.
func (p *Printer) WriteSlice(slice []int64) error {
	if slice == nil {
		return errors.New("Given slice can not be nil.")
	}

	sliceHeader := fmt.Sprintf("%d\n", len(slice))
	if _, err := p.file.WriteString(sliceHeader); err != nil {
		return err
	}

	if err := p.writeSingleSlice(slice); err != nil {
		return err
	}

	return p.synchronizeFileWrite()
}

// WriteMatrix write a given matrix in specified format according to instanced
// out.Printer.
// This method assumes that every rows in the matrix has the same number of
// elements of matrix[0].
// error is returned if: matrix == nil or there is an i/o error.
func (p *Printer) WriteMatrix(matrix [][]int64) error {
	if matrix == nil {
		return errors.New("Given matrix can not be nil.")
	}

	if rows := len(matrix); rows != 0 {
		matrixHeader := fmt.Sprintf("%d %d\n", rows, len(matrix[0]))
		if _, err := p.file.WriteString(matrixHeader); err != nil {
			return err
		}
	} else {
		return errors.New("Given matrix has zero rows.")
	}

	for _, rows := range matrix {
		if err := p.writeSingleSlice(rows); err != nil {
			return err
		}
	}

	return p.synchronizeFileWrite()
}

// WriteBound writes a single sample.Bound structure in specified format
// according to instanced out.Printer. error is returned if there is a i/o error.
func (p *Printer) WriteBound(bound samples.Bound) error {
	return p.WriteBounds([]samples.Bound{bound})
}

// WriteBounds writes a slice of sample.Bound in specified format according to
// instanced out.Printer.
// error is returned if there is a i/o error or if slice is nil.
func (p *Printer) WriteBounds(bounds []samples.Bound) error {
	if bounds == nil {
		return errors.New("Given bound can not be nil")
	}

	boundHeader := fmt.Sprintf("%d\n", len(bounds))
	if _, err := p.file.WriteString(boundHeader); err != nil {
		return err
	}

	var boundAsString string
	for _, b := range bounds {
		boundAsString = fmt.Sprintf("%d %d\n", b.Low(), b.Up())
		if _, err := p.file.WriteString(boundAsString); err != nil {
			return err
		}
	}

	return p.synchronizeFileWrite()
}
