package regroup

import (
	"errors"
	"regexp"
	"strconv"
)

type MatchedGroups struct {
	matches map[string]string
}

func MatchGroups(re *regexp.Regexp, s string) *MatchedGroups {
	groups := make(map[string]string)
	match := re.FindStringSubmatch(s)

	if len(match) > 0 {
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				groups[name] = match[i]
			}
		}
	}

	return &MatchedGroups{matches: groups}
}

func (m *MatchedGroups) Get(key string) (string, error) {
	val, ok := m.matches[key]

	if !ok {
		return "", errors.New("value not matched")
	}

	return val, nil
}

func (m *MatchedGroups) GetInt(key string) (int, error) {
	value, err := m.Get(key)
	if err != nil {
		return 0, err
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New("value cast error")
	}

	return intVal, nil
}
