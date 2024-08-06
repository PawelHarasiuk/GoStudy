package crawlers

import (
	"errors"
	"io"
	"net/http"
	"regexp"
)

func GetHtml(urlString string) (string, error) {
	resp, err := http.Get(urlString)
	if err != nil {
		return "", errors.New("invalid link")
	}

	defer resp.Body.Close()
	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(html), nil
}

func Find(html string) ([]string, error) {
	emailRegex := "\\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}\\b"
	re := regexp.MustCompile(emailRegex)
	emails := re.FindAllString(html, -1)
	return emails, nil
}
