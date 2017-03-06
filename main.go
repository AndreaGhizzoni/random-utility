package main

import (
	"flag"
	"fmt"
	"github.com/AndreaGhizzoni/zenium/samples"
	"os"
)

var (
	// this var is used for '-generate|-g xxx' flag, where xxx is:
	//  - rslice
	//  - oslice
	//  - matrix
	//  - bound
	sampleGeneration string

	// this var is used for '-output|-o xxx' flag, where xxx is by default
	// std out or:
	//  - path/to/file
	out string

	// those two flags are used for '-min|-m' and '-max|-M' bounds respectively
	m, M int64

	// those two flags are used for '--row|-r' and '--col|-c' in matrix
	// generation. cols is also used when `--length|-l` is set in slice
	// generation.
	cols, rows int64
)

func init() {
	flag.StringVar(&sampleGeneration, "generate", "", generateUsage)
	flag.StringVar(&sampleGeneration, "g", "", generateUsage)

	flag.StringVar(&out, "output", "", outUsage)
	flag.StringVar(&out, "o", "", outUsage)

	flag.Int64Var(&m, "min", -922337203685477580, minUsage)
	flag.Int64Var(&m, "m", -922337203685477580, minUsage)
	flag.Int64Var(&M, "max", 922337203685477580, maxUsage)
	flag.Int64Var(&M, "M", 922337203685477580, maxUsage)

	flag.Int64Var(&cols, "columns", 0, colsUsage)
	flag.Int64Var(&cols, "c", 0, colsUsage)
	flag.Int64Var(&cols, "length", 0, lengthUsage)
	flag.Int64Var(&cols, "l", 0, lengthUsage)

	flag.Int64Var(&rows, "rows", 0, rowsUsage)
	flag.Int64Var(&rows, "r", 0, rowsUsage)
}

// utility function to print a message and helper, than exit 1
func printDefaults(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()
	// no flag set, print defaults
	if flag.NFlag() == 0 {
		printDefaults(fmt.Sprintf(helpHeader, os.Args[0]))
	}

	// missing -g|-generate cli flag
	if sampleGeneration == "" {
		printDefaults("What to generate missinig: check --generate|-g " +
			"argument.")
	}

	gen := samples.New()
	// check if -g|-generate flag has an accepted value
	switch sampleGeneration {
	case "rslice":
		s, err := gen.Slice(cols, m, M)
		if err != nil {
			printDefaults(err.Error())
		}
		fmt.Println(s) // TODO change this when output package is ready
	case "oslice":
		// TODO change this when output package is ready
		fmt.Println("Not implemented yet :P")
	case "matrix":
		m, err := gen.Matrix(rows, cols, m, M)
		if err != nil {
			printDefaults(err.Error())
		}
		fmt.Println(m) // TODO change this when output package is ready
	case "bound":
		// TODO change this when output package is ready
		fmt.Println("Not implemented yet :P")
	default:
		printDefaults("What to generate not recognized. Check --generate|-g " +
			"argument.")
	}
}
