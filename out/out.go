package out

import (
	"errors"
	"fmt"
	"github.com/AndreaGhizzoni/zenium/structures"
	"github.com/AndreaGhizzoni/zenium/util"
	"math/big"
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
func (p *Printer) writeSingleSlice(slice []*big.Int) error {
	var stringOfElement string
	for _, element := range slice {
		stringOfElement = fmt.Sprintf("%v ", element)
		if _, err := p.file.WriteString(stringOfElement); err != nil {
			return err
		}
	}
	_, err := p.file.WriteString("\n")
	return err
}

// reusable method to check write synchronization.
func (p *Printer) synchronizeFileWrite() error {
	if err := p.file.Sync(); err != nil {
		return err
	}
	return nil
}

// WriteSlice write the given slice in specified format according to instanced
// out.Printer. error is returned if: slice == nil, contain a nil pointer or
// there is an i/o error.
func (p *Printer) WriteSlice(slice []*big.Int) error {
	if slice == nil {
		return errors.New("Given slice can not be nil.")
	}

	lenSlice, err := util.CountSliceIfNotNil(slice)
	if err != nil {
		return err
	}

	sliceHeader := fmt.Sprintf("%v\n", lenSlice)
	if _, err := p.file.WriteString(sliceHeader); err != nil {
		return err
	}

	if err := p.writeSingleSlice(slice); err != nil {
		return err
	}

	return p.synchronizeFileWrite()
}

// WriteMatrix write a given matrix in specified format according to instanced
// out.Printer. error is returned if: matrix == nil, contain a nil pointer,
// not contain any row or there is an i/o error.
func (p *Printer) WriteMatrix(matrix [][]*big.Int) error {
	if matrix == nil {
		return errors.New("Given matrix can not be nil.")
	}

	rows, columns, err := util.CountMatrixIfNotNil(matrix)
	if err != nil {
		return err
	}

	if rows.Cmp(big.NewInt(0)) == 0 {
		return errors.New("Given matrix has zero rows.")
	}

	matrixHeader := fmt.Sprintf("%v %v\n", rows, columns)
	if _, err := p.file.WriteString(matrixHeader); err != nil {
		return err
	}
	for _, rows := range matrix {
		if err := p.writeSingleSlice(rows); err != nil {
			return err
		}
	}

	return p.synchronizeFileWrite()
}

// WriteBound writes a single structures.Bound in specified format according to
// instanced out.Printer. error is returned if there is a i/o error.
func (p *Printer) WriteBound(bound *structures.Bound) error {
	return p.WriteBounds([]*structures.Bound{bound})
}

// WriteBounds writes a slice of structures.Bound in specified format according
// to instanced out.Printer. error is returned if: slice is nil, contain a nil
// pointer or there is a i/o error.
func (p *Printer) WriteBounds(bounds []*structures.Bound) error {
	if bounds == nil {
		return errors.New("Given bound can not be nil")
	}

	lenBounds, err := util.CountBoundsIfNotNil(bounds)
	if err != nil {
		return err
	}

	boundHeader := fmt.Sprintf("%v\n", lenBounds)
	if _, err := p.file.WriteString(boundHeader); err != nil {
		return err
	}

	var boundAsString string
	for _, b := range bounds {
		boundAsString = fmt.Sprintf("%v %v\n", b.Low(), b.Up())
		if _, err := p.file.WriteString(boundAsString); err != nil {
			return err
		}
	}

	return p.synchronizeFileWrite()
}
