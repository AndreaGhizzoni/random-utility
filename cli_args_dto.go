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

func NewCliArgsDTO_(context *cli.Context) (*cliArgsDTO, error) {
	dto := new(cliArgsDTO)
	for _, flagName := range context.FlagNames() {
		if flagName == "out" {
			dto.OutPath = context.String("out")
		} else {
			bigIntOfString, err := util.FromStringToBigInt(flagName, flagName)
			if err != nil {
				return nil, err
			}
			dto.placeInCorrectField(flagName, bigIntOfString)
		}
	}
}

func (dto *cliArgsDTO) placeInCorrectField(fieldName string, value *big.Int) {
	switch fieldName {
	case "length":
		dto.Length = value
	case "min":
		dto.Min = value
	case "max":
		dto.Max = value
	case "rows":
		dto.Rows = value
	case "columns":
		dto.Columns = value
	case "width":
		dto.Width = value
	case "amount":
		dto.Amount = value
	}
}

func NewCliArgsDTO(c *cli.Context) (*cliArgsDTO, error) {
	var err error
	dto := new(cliArgsDTO)
	dto.OutPath = c.String("out")

	dto.Length, err = util.FromStringToBigInt(c.String("length"), "Length")
	if err != nil {
		return nil, err
	}
	dto.Min, err = util.FromStringToBigInt(c.String("min"), "Minimum")
	if err != nil {
		return nil, err
	}
	dto.Max, err = util.FromStringToBigInt(c.String("max"), "Maximum")
	if err != nil {
		return nil, err
	}
	dto.Rows, err = util.FromStringToBigInt(c.String("rows"), "Rows")
	if err != nil {
		return nil, err
	}
	dto.Columns, err = util.FromStringToBigInt(c.String("columns"), "Columns")
	if err != nil {
		return nil, err
	}
	dto.Width, err = util.FromStringToBigInt(c.String("width"), "Width")
	if err != nil {
		return nil, err
	}
	dto.Amount, err = util.FromStringToBigInt(c.String("amount"), "Amount")
	if err != nil {
		return nil, err
	}
	return dto, nil
}
