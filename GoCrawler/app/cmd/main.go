package main

import (
	"GoCrawler/app/crawlers"
	"fmt"
)

func main() {
	urlString := "https://regexr.com/3e48o"
	html, err := crawlers.GetHtml(urlString)
	if err != nil {
		return
	}
	emails, err := crawlers.Find(html)
	if err != nil {
		return
	}

	for _, email := range emails {
		fmt.Println(email)
	}
}
