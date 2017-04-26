package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"sort"
)

func main() {
	app := cli.NewApp()
	app.Name = "zenium"             // TODO move in separate file
	app.Version = "app.Version"     // TODO move in separate file
	app.Usage = "app.Usage"         // TODO move in separate file
	app.UsageText = "app.UsageText" // TODO move in separate file
	// TODO move in separate file
	app.ArgsUsage = "generate [rslice|oslice|matrix|bound]"
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
			Name:    "generate",
			Aliases: []string{"g"},
			//Category:    "generate command",
			Usage: "generate usage",
			//UsageText:   "generate usage text",
			Description: "generate description",
			ArgsUsage:   "generate args usage",
			Subcommands: cli.Commands{
				cli.Command{
					Name:  "rslice",
					Usage: "rslice usage",
					Flags: []cli.Flag{outFlag, minFlag, maxFlag, lengthFlag},
					Action: func(c *cli.Context) error {
						fmt.Println("generate rslice command")
						fmt.Printf("out flag is set ? %v\n", c.IsSet("out"))
						fmt.Printf("min flag is set ? %v\n", c.IsSet("min"))
						fmt.Printf("max flag is set ? %v\n", c.IsSet("max"))
						fmt.Printf("length flag is set ? %v\n",
							c.IsSet("length"))
						return nil
					},
				},
				cli.Command{
					Name:  "oslice",
					Usage: "oslice usage",
					Flags: []cli.Flag{outFlag, minFlag, maxFlag, lengthFlag},
					Action: func(c *cli.Context) error {
						fmt.Println("generate oslice command")
						fmt.Printf("out flag is set ? %v\n", c.IsSet("out"))
						fmt.Printf("min flag is set ? %v\n", c.IsSet("min"))
						fmt.Printf("max flag is set ? %v\n", c.IsSet("max"))
						fmt.Printf("length flag is set ? %v\n",
							c.IsSet("length"))
						return nil
					},
				},
				cli.Command{
					Name:  "matrix",
					Usage: "matrix usage",
					Flags: []cli.Flag{outFlag, minFlag, maxFlag, rowsFlag,
						columnFlag},
					Action: func(c *cli.Context) error {
						fmt.Println("generate matrix command")
						fmt.Printf("out flag is set ? %v\n", c.IsSet("out"))
						fmt.Printf("min flag is set ? %v\n", c.IsSet("min"))
						fmt.Printf("max flag is set ? %v\n", c.IsSet("max"))
						fmt.Printf("rows flag is set ? %v\n", c.IsSet("rows"))
						fmt.Printf("cols flag is set ? %v\n",
							c.IsSet("columns"))
						return nil
					},
				},
				cli.Command{
					Name:  "bound",
					Usage: "bound usage",
					Flags: []cli.Flag{outFlag, minFlag, maxFlag},
					Action: func(c *cli.Context) error {
						fmt.Println("generate bound command")
						fmt.Printf("out flag is set ? %v\n", c.IsSet("out"))
						fmt.Printf("min flag is set ? %v\n", c.IsSet("min"))
						fmt.Printf("max flag is set ? %v\n", c.IsSet("max"))
						fmt.Printf("rows flag is set ? %v\n", c.IsSet("rows"))
						return nil
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}
