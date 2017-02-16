package main

import (
	"flag"
	"fmt"
	"math"
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
	// c [std out] or:
	//  - path/to/file
	out string

	// those two flags are used for '-min|-m' and '-max|-M' bounds respectively
	m, M int64

	// those two flags are used for '-row' and '-col' in matrix generation mode.
	cols, rows int64
)

func init() {
	flag.StringVar(&sampleGeneration, "generate", "", "`type_of_sample` TODO USAGE")
	flag.StringVar(&sampleGeneration, "g", "", "TODO USAGE")

	flag.StringVar(&out, "output", "c", "TODO USAGE")
	flag.StringVar(&out, "o", "c", "TODO USAGE")

	flag.Int64Var(&m, "min", math.MinInt64, "TODO USAGE")
	flag.Int64Var(&m, "m", math.MinInt64, "TODO USAGE")
	flag.Int64Var(&M, "max", math.MaxInt64, "TODO USAGE")
	flag.Int64Var(&M, "M", math.MaxInt64, "TODO USAGE")

	flag.Int64Var(&cols, "columns", 0, "TODO USAGE")
	flag.Int64Var(&cols, "c", 0, "TODO USAGE")
	flag.Int64Var(&cols, "length", 0, "TODO USAGE")
	flag.Int64Var(&cols, "l", 0, "TODO USAGE")

	flag.Int64Var(&rows, "rows", 0, "TODO USAGE")
	flag.Int64Var(&rows, "r", 0, "TODO USAGE")
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Println("Usage of Zenium sample generator:", os.Args[0], "[OPTIONS]")
		flag.PrintDefaults()
	}
}
