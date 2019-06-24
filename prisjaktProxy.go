package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Search(userflags *UserFlags, searchTerm string) {

	urlParts := []string{
		"https://search.prisjakt.nu/classic?class=Search_Supersearch&method=search&market=no&skip_login=1&modes=product,raw_sorted,raw&limit=",
		strconv.Itoa(userflags.maxNumberOfResults),
		"&q=",
		searchTerm}

	var url = strings.Join(urlParts, "")

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var o interface{}

	json.Unmarshal(body, &o)
	fmt.Println(o)
}
