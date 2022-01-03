package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	imdb_scraper "imdb-scraper"
)

func main() {
	var URL, result string
	var positionStart, positionEnd int
	var fmtResp map[string]interface{}

	fmt.Scanln(&URL)
	fmt.Println(URL)

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}

	resultHttp, _ := ioutil.ReadAll(resp.Body)
	t := string(resultHttp)

	positionStart = strings.LastIndex(t, "IMDbReactInitialState.push({")
	result = t[positionStart+len("IMDbReactInitialState.push({")-1:]
	positionEnd = strings.Index(result, ");")
	result = result[:positionEnd]

	err = json.Unmarshal([]byte(result), &fmtResp)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, _ := json.Marshal(fmtResp["titles"])
	films := make(map[string]imdb_scraper.Film)
	err = json.Unmarshal(res, &films)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(films))

	for _, value := range films {
		fmt.Printf("%s || https://www.imdb.com%s\n", value.Primary.Title, value.Primary.Href)
	}
}
