package crawlers

import (
	"errors"
	"io"
	"net/http"
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
