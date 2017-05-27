package main

import (
	"fmt"
	"github.com/AndreaGhizzoni/zenium/out"
	"github.com/AndreaGhizzoni/zenium/samples"
	"github.com/urfave/cli"
	"os"
	"sort"
)

var version = "0.0.2"

func main() {
	app := cli.NewApp()

	app.Name = name
	app.Version = version
	app.Usage = usage
	app.UsageText = name + usageText
	app.Authors = []cli.Author{
		{Name: "Andrea Ghizzoni", Email: "andrea.ghz@gmail.com"},
	}

	outFlag := cli.StringFlag{
		Name:  outFlag,
		Value: "",
		Usage: outUsage,
	}
	minFlag := cli.StringFlag{
		Name:  minFlag,
		Value: "-922337203685477580",
		Usage: minUsage,
	}
	maxFlag := cli.StringFlag{
		Name:  maxFlag,
		Value: "922337203685477580",
		Usage: maxUsage,
	}
	columnFlag := cli.StringFlag{
		Name:  colsFlag,
		Value: "1",
		Usage: colsUsage,
	}
	rowsFlag := cli.StringFlag{
		Name:  rowsFlag,
		Value: "1",
		Usage: rowsUsage,
	}
	lengthFlag := cli.StringFlag{
		Name:  lengthFlag,
		Value: "1",
		Usage: lengthUsage,
	}
	widthFlag := cli.StringFlag{
		Name:  widthFlag,
		Value: "1",
		Usage: widthUsage,
	}
	amountFlag := cli.StringFlag{
		Name:  amountFlag,
		Value: "1",
		Usage: amountUsage,
	}

	app.Commands = []cli.Command{
		{
			Name:   rsliceCommand,
			Usage:  rsliceUsage,
			Flags:  []cli.Flag{outFlag, minFlag, maxFlag, lengthFlag},
			Action: generateRSlice,
		},
		{
			Name:   osliceCommand,
			Usage:  osliceUsage,
			Flags:  []cli.Flag{outFlag, minFlag, maxFlag, lengthFlag},
			Action: generateOSlice,
		},
		{
			Name:  matrixCommand,
			Usage: matrixUsage,
			Flags: []cli.Flag{outFlag, minFlag, maxFlag, rowsFlag,
				columnFlag},
			Action: generateMatrix,
		},
		{
			Name:   boundCommand,
			Usage:  boundUsage,
			Flags:  []cli.Flag{outFlag, minFlag, maxFlag, widthFlag, amountFlag},
			Action: generateBound,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}

// call back on rslice command
func generateRSlice(c *cli.Context) error {
	dto, err := NewCliArgsDTO(c)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	printer, err := out.NewPrinter(dto.OutPath)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	gen, err := samples.NewGenerator(dto.Min, dto.Max)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}
	slice, err := gen.Slice(dto.Length)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	return printer.WriteSlice(slice)
}

// call back on oslice command
func generateOSlice(c *cli.Context) error {
	fmt.Printf("TODO")
	return nil
}

// call back on matrix command
func generateMatrix(c *cli.Context) error {
	cliArgsDTO, err := NewCliArgsDTO(c)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	printer, err := out.NewPrinter(cliArgsDTO.OutPath)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	gen, err := samples.NewGenerator(cliArgsDTO.Min, cliArgsDTO.Max)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}
	matrix, err := gen.Matrix(cliArgsDTO.Rows, cliArgsDTO.Columns)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	return printer.WriteMatrix(matrix)
}

// call back on bound command
func generateBound(c *cli.Context) error {
	cliArgsDTO, err := NewCliArgsDTO(c)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	printer, err := out.NewPrinter(cliArgsDTO.OutPath)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	gen, err := samples.NewGenerator(cliArgsDTO.Min, cliArgsDTO.Max)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}
	bounds, err := gen.Bounds(cliArgsDTO.Width, cliArgsDTO.Amount)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	return printer.WriteBounds(bounds)
}
