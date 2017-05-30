package main

import (
	"fmt"
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

func NewCliArgsDTO(context *cli.Context) (*cliArgsDTO, error) {
	dto := new(cliArgsDTO)
	for _, flagName := range context.FlagNames() {
		if flagName == "out" {
			dto.OutPath = context.String("out")
		} else {
			bigIntOfString, err := util.FromStringToBigInt(
				context.String(flagName),
				flagName,
			)
			if err != nil {
				return nil, err
			}
			dto.placeInCorrectField(flagName, bigIntOfString)
		}
	}
	return dto, nil
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
