package regex

import (
	"errors"
	"regexp"
	"time"
)

type Set map[string]struct{}

type finder interface {
	Find() (Set, error)
}

type EmailFinder struct {
	EmailRegex string
}

type CustomFinder struct {
	CustomRegex string
}

func (emailCrawler EmailFinder) Find(html string) (Set, error) {
	re := regexp.MustCompile(emailCrawler.EmailRegex)
	emails := re.FindAllString(html, -1)
	if len(emails) == 0 {
		return Set{}, errors.New("could not find any emails")
	}
	time.Sleep(1 * time.Second)

	uniqueEmails := make(Set)

	for _, val := range emails {
		uniqueEmails[val] = struct{}{}
	}

	return uniqueEmails, nil
}

// two diffrent methods do not make sense for this usecase but i wanted to practice it
func (customCrawler CustomFinder) Find(html string) (Set, error) {
	re := regexp.MustCompile(customCrawler.CustomRegex)
	el := re.FindAllString(html, -1)
	time.Sleep(2 * time.Second)

	uniqueElements := make(Set)

	for _, val := range el {
		uniqueElements[val] = struct{}{}
	}

	return uniqueElements, nil
}
