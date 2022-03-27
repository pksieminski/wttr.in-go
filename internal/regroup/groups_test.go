package regroup

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestMatchGroups(t *testing.T) {
	// Single match
	re := regexp.MustCompile("(?P<value>\\d)+.+km/h")
	groups := MatchGroups(re, "     \\_(   ).   → 4 km/h       ")

	assert.Equal(t, 1, len(groups.matches))

	expected := map[string]string{"value": "4"}
	assert.Equal(t, expected, groups.matches)
}

func TestMatchGroups_MultipleMatches(t *testing.T) {
	re := regexp.MustCompile("(?P<sign>[+-])?(?P<value>\\d)+(\\([+-]?\\d+\\)).+°C")
	groups := MatchGroups(re, "   _ /\"\".-.     -3(-6) °C      ")

	assert.Equal(t, 2, len(groups.matches))

	expected := map[string]string{"value": "3", "sign": "-"}
	assert.Equal(t, expected, groups.matches)
}

func TestMatchGroups_NoMatches(t *testing.T) {
	re := regexp.MustCompile("(?P<value>\\d)+.+km/h")
	groups := MatchGroups(re, "")

	assert.Equal(t, 0, len(groups.matches))
}

func TestMatchedGroups_Get(t *testing.T) {
	groups := &MatchedGroups{matches: map[string]string{"value": "4"}}

	t.Run("key=value", func(t *testing.T) {
		value, err := groups.Get("value")

		if assert.Nil(t, err, "MatchedGroups_Get: unexpected error %s", err) {
			assert.Equal(t, "4", value)
		}
	})
	t.Run("key=sign", func(t *testing.T) {
		_, err := groups.Get("sign")

		assert.NotNil(t, err, "MatchedGroups_Get: expected missing key error did not occur")
	})
}

func TestMatchedGroups_GetInt(t *testing.T) {
	groups := &MatchedGroups{matches: map[string]string{"value": "4", "unit": "km/h"}}

	t.Run("key=value", func(t *testing.T) {
		value, err := groups.GetInt("value")

		if assert.Nil(t, err, "MatchedGroups_GetInt: unexpected error %s", err) {
			assert.Equal(t, 4, value)
		}
	})
	t.Run("key=sign", func(t *testing.T) {
		_, err := groups.GetInt("sign")

		assert.NotNil(t, err, "MatchedGroups_Get: expected missing key error did not occur")
	})
	t.Run("key=unit", func(t *testing.T) {
		_, err := groups.GetInt("unit")

		assert.NotNil(t, err, "MatchedGroups_Get: expected cast error did not occur")
	})
}
