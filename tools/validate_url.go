package tools

import (
	"net/url"
)

func ValidateURL(inUrl string) bool {
	_, err := url.Parse(inUrl)
	if err != nil {
		return false
	}
	return true
}
