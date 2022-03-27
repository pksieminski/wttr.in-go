package wttr

import (
	"io/ioutil"
	"net/url"
	"strings"
)

const (
	english = "en"
)

const (
	onlyCurrent = "0" // only current weather
	forceAnsi   = "A" // force ANSI output
	quiet       = "q" // quiet output without "Weather report"
	noColors    = "T" // quiet output without "Weather report"
)

type GetWeatherParams struct {
	Location string
	Language string
	Flags    []string
}

func (c *Client) GetCurrentWeather(location string) (*Weather, error) {
	weather, err := c.getCurrentWeather(location)

	if err != nil {
		return nil, err
	}

	return ParseWeather(weather)
}

func (c *Client) getCurrentWeather(location string) (string, error) {
	params := GetWeatherParams{
		Location: location,
		Language: english,
		Flags:    []string{onlyCurrent, forceAnsi, quiet, noColors},
	}

	return c.getWeather(params)
}

func (c *Client) getWeather(params GetWeatherParams) (string, error) {
	weatherURL := c.getWeatherURL(params)
	resp, err := c.client.Get(weatherURL.String())

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *Client) getWeatherURL(params GetWeatherParams) *url.URL {
	u := url.URL{Path: params.Location}
	query := u.Query()
	query.Set("", strings.Join(params.Flags, ""))

	if params.Language != "" {
		query.Set("lang", params.Language)
	}

	u.RawQuery = query.Encode()[1:]

	return c.config.URL.ResolveReference(&u)
}
