package main

// constants for flags
const (
	helpHeader = "Usage of Zenium sample generator: %s [OPTIONS]\n"

	generateUsage = "`rvector|ovector|matrix|bound` one of these options to " +
		"generate a random or ordered vector, a random matrix or random bound. "

	outUsage = "`/path/to/file` to output the generate sample. " +
		"" +
		"stdout is the default"

	minUsage = "`X` is a integer numeber to set the minimum value of random " +
		"generator. (default 0)"
	maxUsage = "`X` is a integer numeber to set the maximum value of random " +
		"generator. "

	colsUsage = "`X` is a integer number to set the matrix columns if -g matrix " +
		"is set. (default 0)"

	rowsUsage = "`X` is a integer number to set the matrix rows if -g matrix " +
		"is set. (default 0)"

	lengthUsage = "`X` is a integer number to set the matrix length if " +
		"-g rvector|ovector is set. (default 0)"
)
