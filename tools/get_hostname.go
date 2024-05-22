package tools

import (
	"net/url"
)

func GetHostname(inputUrl string) string {
	u, err := url.Parse(inputUrl)
	if err != nil {
		return ""
	}
	hostName := u.Hostname()
	return hostName
}
