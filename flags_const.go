package main

// constants for flags
const (
	// name of utility
	name = "zenium"

	// usage text
	usage     = "utility program to generate random data structures"
	usageText = " [rslice|oslice|matrix|bound] [OPTIONS]"

	// flag usage
	outFlagShort = "o"
	outFlagLong  = "out"
	outFlag      = outFlagLong + ", " + outFlagShort
	outUsage     = "`FILE` to output the generate sample. stdout is the default."

	minFlagShort = "m"
	minFlagLong  = "min"
	minFlag      = minFlagLong + ", " + minFlagShort
	minUsage     = "`m` is a integer number to set the minimum value of random " +
		"generator."

	maxFlagShort = "M"
	maxFlagLong  = "max"
	maxFlag      = maxFlagLong + ", " + maxFlagShort
	maxUsage     = "`M` is a integer number to set the maximum value of random " +
		"generator. "

	colsFlagShort = "c"
	colsFlagLong  = "columns"
	colsFlag      = colsFlagLong + ", " + colsFlagShort
	colsUsage     = "`c` is a integer number to set the matrix columns."

	rowsFlagShort = "r"
	rowsFlagLong  = "rows"
	rowsFlag      = rowsFlagLong + ", " + rowsFlagShort
	rowsUsage     = "`r` is a integer number to set the matrix rows."

	lengthFlagShort = "l"
	lengthFlagLong  = "length"
	lengthFlag      = lengthFlagLong + ", " + lengthFlagShort
	lengthUsage     = "`l` is a integer number to set the slice length."

	widthFlagShort = "w"
	widthFlagLong  = "width"
	widthFlag      = widthFlagLong + ", " + widthFlagShort
	widthUsage     = "`w` is a integer number to set the bound width."

	amountFlagShort = "a"
	amountFlagLong  = "amount"
	amountFlag      = amountFlagLong + ", " + amountFlagShort
	amountUsage     = "`a` is the amount of bounds to generate."

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
