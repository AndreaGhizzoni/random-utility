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

func NewCliArgsDTO(context *cli.Context) (*cliArgsDTO, error) {
	dto := new(cliArgsDTO)
	for _, flagName := range context.FlagNames() {
		if flagName == outFlagLong {
			dto.OutPath = context.String(outFlagLong)
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
	case lengthFlagLong:
		dto.Length = value
	case minFlagLong:
		dto.Min = value
	case maxFlagLong:
		dto.Max = value
	case rowsFlagLong:
		dto.Rows = value
	case colsFlagLong:
		dto.Columns = value
	case widthFlagLong:
		dto.Width = value
	case amountFlagLong:
		dto.Amount = value
	}
}
