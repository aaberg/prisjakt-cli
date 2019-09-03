package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello bloody world")

	var userFlags = GetUserFlags()

	fmt.Println("Max number of results is", userFlags.maxNumberOfResults)
	fmt.Println("Search term is", userFlags.searchTerm)

	products, err := Search(userFlags, userFlags.searchTerm)

	if err != nil {
		fmt.Println("Error!!")
		fmt.Println(err)
		return
	}

	for _, element := range products.Items {
		fmt.Printf("Product %s with id %d\n", element.Name, element.Id)
	}
}
