package main

import (
	"GoCrawler/app/crawlers"
	"GoCrawler/app/regex"
	"fmt"
	"sync"
)

func main() {
	urlString := "https://regexr.com/3e48o"
	html, err := crawlers.GetHtml(urlString)
	var wg sync.WaitGroup

	if err != nil {
		fmt.Println(err)
	}

	emailFinder := regex.EmailFinder{
		EmailRegex: "\\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}\\b",
	}

	customFinder := regex.CustomFinder{
		CustomRegex: "ar.",
	}

	wg.Add(1)
	go func() {
		emails, err := emailFinder.Find(html)
		defer wg.Done()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(emails)
	}()

	wg.Add(1)
	go func() {
		custom, err := customFinder.Find(html)
		defer wg.Done()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(custom)
	}()

	wg.Wait()
}
