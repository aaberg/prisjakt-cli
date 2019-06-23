package main

import (
	"flag"
)

// UserFlags contains the parsed CLI flags
type UserFlags struct {
	maxNumberOfResults int
}

// GetUserFlags returns the parsed CLI flags
func GetUserFlags() *UserFlags {
	var maxNumberOfResultsFlag = flag.Int("n", 10, "Max number of returned values")

	flag.Parse()

	var userFlags = UserFlags{*maxNumberOfResultsFlag}
	return &userFlags
}
