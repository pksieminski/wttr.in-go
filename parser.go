package wttr

import (
	"errors"
	"fmt"
	"github.com/pksieminski/wttr.in-go/internal/regroup"
	"regexp"
	"strings"
)

type ParseError struct {
	Err error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("weather response parsing error: %s", e.Err)
}

func ParseWeather(resp string) (*Weather, error) {
	lines := strings.Split(resp, "\n")

	if len(lines) < 8 {
		return nil, &ParseError{Err: errors.New("returned weather body is too short")}
	}

	temp, err := parseTemperature(lines[3])

	if err != nil {
		return nil, err
	}

	speed, err := parseWindSpeed(lines[4])

	if err != nil {
		return nil, err
	}

	return &Weather{
		Location:    lines[0],
		Description: parseDescription(lines[2]),
		Temperature: temp,
		WindSpeed:   speed,
	}, nil
}

func parseDescription(line string) string {
	return removeIndent(line)
}

func parseTemperature(line string) (int, error) {
	re := regexp.MustCompile("(?P<sign>[+-])?(?P<value>\\d+)(?:\\([+-]?\\d+\\))?.+°C")
	groups := regroup.MatchGroups(re, line)

	temp, err := groups.GetInt("value")
	if err != nil {
		return 0, &ParseError{Err: errors.New("temperature parse error")}
	}

	if sign, _ := groups.Get("sign"); sign == "-" {
		temp = -temp
	}

	return temp, nil
}

func parseWindSpeed(line string) (int, error) {
	re := regexp.MustCompile("(?P<value>\\d+).+km/h")
	groups := regroup.MatchGroups(re, line)

	speed, err := groups.GetInt("value")
	if err != nil {
		return 0, &ParseError{Err: errors.New("wind speed parse error")}
	}

	return speed, nil
}

func removeIndent(line string) string {
	return line[16:]
}
