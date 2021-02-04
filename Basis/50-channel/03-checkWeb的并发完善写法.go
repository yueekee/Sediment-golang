package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize2(url string, channel chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}

func main() {
	sizes := make(chan Page)
	urls := []string{
		"https://example.com/", "https://baidu.com/", "https://www.bing.com/",
	}
	for _, url := range urls {
		go responseSize2(url, sizes)
	}

	for i := 0; i < len(urls); i++ {
		page := <-sizes
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}
