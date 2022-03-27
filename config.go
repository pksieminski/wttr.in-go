package wttr

import "net/url"

const URL = "https://wttr.in"

type Configuration struct {
	URL url.URL
}

func NewConfiguration() (*Configuration, error) {
	u, err := url.Parse(URL)

	if err != nil {
		return nil, err
	}

	return &Configuration{URL: *u}, nil
}
