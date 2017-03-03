package main

// constants for flags
const (
	helpHeader = "Usage of Zenium samples generator: %s [OPTIONS]\n"

	generateUsage = "`rslice|oslice|matrix|bound` one of these options to " +
		"generate a random or ordered slice (aka vector), a random matrix or " +
		"random bound. Multiple argument of this flag pipe-separated is not " +
		"supported yet."

	outUsage = "`/path/to/file` to output the generate sample. stdout is the " +
		"default."

	minUsage = "`m` is a integer numeber to set the minimum value of random " +
		"generator."
	maxUsage = "`M` is a integer numeber to set the maximum value of random " +
		"generator. "

	colsUsage = "`c` is a integer number to set the matrix columns if " +
		"-g matrix is set. (default 0)"
	rowsUsage = "`r` is a integer number to set the matrix rows if -g matrix " +
		"is set. (default 0)"
	lengthUsage = "`l` is a integer number to set the slice length if " +
		"-g rslice|oslice is set. (default 0)"
)
