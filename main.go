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
	minFlag := cli.Int64Flag{
		Name:  minFlag,
		Value: -922337203685477580,
		Usage: minUsage,
	}
	maxFlag := cli.Int64Flag{
		Name:  maxFlag,
		Value: 922337203685477580,
		Usage: maxUsage,
	}
	columnFlag := cli.Int64Flag{
		Name:  colsFlag,
		Value: 1,
		Usage: colsUsage,
	}
	rowsFlag := cli.Int64Flag{
		Name:  rowsFlag,
		Value: 1,
		Usage: rowsUsage,
	}
	lengthFlag := cli.Int64Flag{
		Name:  lengthFlag,
		Value: 1,
		Usage: lengthUsage,
	}
	widthFlag := cli.Int64Flag{
		Name:  widthFlag,
		Value: 1,
		Usage: widthUsage,
	}
	amountFlag := cli.Int64Flag{
		Name:  amountFlag,
		Value: 1,
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

func generateRSlice(c *cli.Context) error {
	output := c.String("out")
	l := c.Int64("length")
	min := c.Int64("min")
	max := c.Int64("max")

	p, err := out.NewPrinter(output)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	gen := samples.NewGenerator()
	slice, err := gen.Slice(l, min, max)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	return p.WriteSlice(slice)
}

func generateOSlice(c *cli.Context) error {
	fmt.Printf("TODO")
	return nil
}

func generateMatrix(c *cli.Context) error {
	output := c.String("out")
	min := c.Int64("min")
	max := c.Int64("max")
	rows := c.Int64("rows")
	cols := c.Int64("columns")

	p, err := out.NewPrinter(output)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	gen := samples.NewGenerator()
	matrix, err := gen.Matrix(rows, cols, min, max)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	return p.WriteMatrix(matrix)
}

func generateBound(c *cli.Context) error {
	output := c.String("out")
	min := c.Int64("min")
	max := c.Int64("max")
	width := c.Int64("width")
	amount := c.Int64("amount")

	p, err := out.NewPrinter(output)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	}

	gen := samples.NewGenerator()
	bounds := make([]samples.Bound, amount)
	for i := range bounds {
		if bound, err := gen.Bound(min, max, width); err != nil {
			os.Stderr.WriteString(err.Error())
			return err
		} else {
			bounds[i] = *bound
		}
	}

	return p.WriteBounds(bounds)
}
