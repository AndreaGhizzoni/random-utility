package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"sort"
)

var version = "0.0.1"

func main() {
	app := cli.NewApp()

	app.Name = name
	app.Version = version
	app.Usage = usage
	app.UsageText = name + usageText
	app.Authors = []cli.Author{
		{Name: "Andrea Ghizzoni", Email: "andrea.ghz@gmail.com"},
	}

	// TODO maybe move all flags to separate file
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
		Value: 0,
		Usage: colsUsage,
	}
	rowsFlag := cli.Int64Flag{
		Name:  rowsFlag,
		Value: 0,
		Usage: rowsUsage,
	}
	lengthFlag := cli.Int64Flag{
		Name:  lengthFlag,
		Value: 0,
		Usage: lengthUsage,
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
			Flags:  []cli.Flag{outFlag, minFlag, maxFlag},
			Action: generateBound,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}

func generateRSlice(c *cli.Context) error {
	fmt.Println("generate rslice command")
	fmt.Printf("out = %v\n", c.String("out"))
	fmt.Printf("min = %v\n", c.Int64("min"))
	fmt.Printf("max = %v\n", c.Int64("max"))
	fmt.Printf("length = %v\n", c.Int64("length"))
	return nil
}

func generateOSlice(c *cli.Context) error {
	fmt.Println("generate oslice command")
	fmt.Printf("out = %v\n", c.String("out"))
	fmt.Printf("min = %v\n", c.Int64("min"))
	fmt.Printf("max = %v\n", c.Int64("max"))
	fmt.Printf("length = %v\n", c.Int64("length"))
	return nil
}

func generateMatrix(c *cli.Context) error {
	fmt.Println("generate matrix command")
	fmt.Printf("out = %v\n", c.String("out"))
	fmt.Printf("min = %v\n", c.Int64("min"))
	fmt.Printf("max = %v\n", c.Int64("max"))
	fmt.Printf("rows = %v\n", c.Int64("rows"))
	fmt.Printf("cols = %v\n", c.Int64("columns"))
	return nil
}

func generateBound(c *cli.Context) error {
	fmt.Println("generate bound command")
	fmt.Printf("out = %v\n", c.String("out"))
	fmt.Printf("min = %v\n", c.Int64("min"))
	fmt.Printf("max = %v\n", c.Int64("max"))
	return nil
}
