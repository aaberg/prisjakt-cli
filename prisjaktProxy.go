package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Search(userflags *UserFlags, searchTerm string) (Product, error) {

	urlParts := []string{
		"https://search.prisjakt.nu/classic?class=Search_Supersearch&method=search&market=no&skip_login=1&modes=product,raw_sorted,raw&limit=",
		strconv.Itoa(userflags.maxNumberOfResults),
		"&q=",
		url.QueryEscape(searchTerm)}

	var url = strings.Join(urlParts, "")

	resp, err := http.Get(url)

	var product Product

	if err != nil {
		fmt.Println(err)
		return product, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return product, err
	}

	var o Response

	json.Unmarshal(body, &o)
	product = o.Message.Product

	return product, nil
}
