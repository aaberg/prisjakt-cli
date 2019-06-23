package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello bloody world")

	var userFlags = GetUserFlags()

	fmt.Println("Max number of results is", userFlags.maxNumberOfResults)
}
