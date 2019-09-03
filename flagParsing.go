package main

import (
	"flag"
)

// UserFlags contains the parsed CLI flags
type UserFlags struct {
	maxNumberOfResults int
	searchTerm         string
}

// GetUserFlags returns the parsed CLI flags
func GetUserFlags() *UserFlags {
	var maxNumberOfResultsFlag = flag.Int("n", 10, "Max number of returned values")
	var searchTerm = flag.String("s", "", "Search term")

	flag.Parse()

	var userFlags = UserFlags{*maxNumberOfResultsFlag, *searchTerm}
	return &userFlags
}
