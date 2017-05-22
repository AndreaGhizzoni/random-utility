package main

// constants for flags
const (
	// name of utility
	name = "zenium"

	// usage text
	usage     = "utility program to generate random data structures"
	usageText = " [rslice|oslice|matrix|bound] [OPTIONS]"

	// flag usage
	outFlag  = "out, o"
	outUsage = "`FILE` to output the generate sample. stdout is the default."

	minFlag  = "min, m"
	minUsage = "`m` is a integer number to set the minimum value of random " +
		"generator."

	maxFlag  = "max, M"
	maxUsage = "`M` is a integer number to set the maximum value of random " +
		"generator. "

	colsFlag  = "columns, c"
	colsUsage = "`c` is a integer number to set the matrix columns."

	rowsFlag  = "rows, r"
	rowsUsage = "`r` is a integer number to set the matrix rows."

	lengthFlag  = "length, l"
	lengthUsage = "`l` is a integer number to set the slice length."

	widthFlag  = "width, w"
	widthUsage = "`w` is a integer number to set the bound width."

	amountFlag  = "amount, a"
	amountUsage = "`a` is the amount of bounds to generate."

	// command usage
	rsliceCommand = "rslice"
	rsliceUsage   = "command to generate random slice (aka vector)." + commonPart

	osliceCommand = "oslice"
	osliceUsage   = "command to generate ordered slice (aka vector)." + commonPart

	matrixCommand = "matrix"
	matrixUsage   = "command to generate random matrix. " + commonPart

	boundCommand = "bound"
	boundUsage   = "command to generate random bound. " + commonPart

	commonPart = "See '" + name + " [COMMAND] help' for more options."
)
