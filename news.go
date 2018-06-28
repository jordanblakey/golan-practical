package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(bytes)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	resp.Body.Close()
}
