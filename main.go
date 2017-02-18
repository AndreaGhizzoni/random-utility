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
	flag.StringVar(&sampleGeneration, "-generate", "", generateUsage)
	flag.StringVar(&sampleGeneration, "g", "", generateUsage)

	flag.StringVar(&out, "-output", "", outUsage)
	flag.StringVar(&out, "o", "", outUsage)

	flag.Int64Var(&m, "-min", 0, minUsage)
	flag.Int64Var(&m, "m", 0, minUsage)
	flag.Int64Var(&M, "-max", math.MaxInt64, maxUsage)
	flag.Int64Var(&M, "M", math.MaxInt64, maxUsage)

	flag.Int64Var(&cols, "-columns", 0, colsUsage)
	flag.Int64Var(&cols, "c", 0, colsUsage)
	flag.Int64Var(&cols, "-length", 0, lengthUsage)
	flag.Int64Var(&cols, "l", 0, lengthUsage)

	flag.Int64Var(&rows, "-rows", 0, rowsUsage)
	flag.Int64Var(&rows, "r", 0, rowsUsage)
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Printf(helpHeader, os.Args[0])
		flag.PrintDefaults()
	}
}
