package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"sort"
)

func main() {
	app := cli.NewApp()

	// TODO move constants to separate file
	app.Name = "zenium"
	app.Version = "0.1"
	app.Usage = "utility program to generate random data structures"
	app.UsageText = app.Name + " [rslice|oslice|matrix|bound]"
	app.Authors = []cli.Author{
		{Name: "Andrea Ghizzoni", Email: "andrea.ghz@gmail.com"},
	}

	// TODO maybe move all flags to separate file
	outFlag := cli.StringFlag{
		Name:  "out, o",
		Value: "",
		Usage: "usage flag out `FILE`",
	}
	minFlag := cli.Int64Flag{
		Name:  "min, m",
		Value: -922337203685477580,
		Usage: "usage flag min",
	}
	maxFlag := cli.Int64Flag{
		Name:  "max, M",
		Value: 922337203685477580,
		Usage: "usage flag max",
	}
	columnFlag := cli.Int64Flag{
		Name:  "columns, c",
		Value: 0,
		Usage: "usage flag column",
	}
	rowsFlag := cli.Int64Flag{
		Name:  "rows, r",
		Value: 0,
		Usage: "usage flag rows",
	}
	lengthFlag := cli.Int64Flag{
		Name:  "length, l",
		Value: 0,
		Usage: "usage flag length",
	}

	app.Commands = []cli.Command{
		{
			Name:  "rslice",
			Usage: "usage",
			Flags: []cli.Flag{outFlag, minFlag, maxFlag, lengthFlag},
            Action:generateRSlice,
		},
		{
			Name:   "oslice",
			Usage:  "usage",
			Flags:  []cli.Flag{outFlag, minFlag, maxFlag, lengthFlag},
			Action: generateOSlice,
		},
		{
			Name:  "matrix",
			Usage: "usage",
			Flags: []cli.Flag{outFlag, minFlag, maxFlag, rowsFlag,
				columnFlag},
			Action: generateMatrix,
		},
		{
			Name:   "bound",
			Usage:  "usage",
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
