package main

import (
	"github.com/AndreaGhizzoni/zenium/util"
	"github.com/urfave/cli"
	"math/big"
)

type cliArgsDTO struct {
	OutPath                      string
	Min, Max                     *big.Int
	Columns, Rows, Length, Width *big.Int
	Amount                       *big.Int
}

func NewCliArgsDTO(c *cli.Context) (*cliArgsDTO, error) {
	output := c.String("out")
	stringLength := c.String("length")
	stringMin := c.String("min")
	stringMax := c.String("max")
	stringRows := c.String("rows")
	stringCols := c.String("columns")
	stringWidth := c.String("width")
	stringAmount := c.String("amount")
	var err error

	dto := new(cliArgsDTO)

	dto.OutPath = output
	dto.Length, err = util.FromStringToBigInt(stringLength, "Length")
	if err != nil {
		return nil, err
	}
	dto.Min, err = util.FromStringToBigInt(stringMin, "Minimum")
	if err != nil {
		return nil, err
	}
	dto.Max, err = util.FromStringToBigInt(stringMax, "Maximum")
	if err != nil {
		return nil, err
	}
	dto.Rows, err = util.FromStringToBigInt(stringRows, "Rows")
	if err != nil {
		return nil, err
	}
	dto.Columns, err = util.FromStringToBigInt(stringCols, "Columns")
	if err != nil {
		return nil, err
	}
	dto.Width, err = util.FromStringToBigInt(stringWidth, "Width")
	if err != nil {
		return nil, err
	}
	dto.Amount, err = util.FromStringToBigInt(stringAmount, "Amount")
	if err != nil {
		return nil, err
	}
	return dto, nil
}
