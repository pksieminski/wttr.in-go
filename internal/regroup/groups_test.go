package regroup

import (
	"reflect"
	"regexp"
	"testing"
)

func TestMatchGroups(t *testing.T) {
	// Single match
	re := regexp.MustCompile("(?P<value>\\d)+.+km/h")
	groups := MatchGroups(re, "     \\_(   ).   → 4 km/h       ")

	if len(groups.matches) != 1 {
		t.Fatalf("MatchGroups: expected 1 match, actual %d", len(groups.matches))
	}

	expected := map[string]string{"value": "4"}
	if !reflect.DeepEqual(groups.matches, expected) {
		t.Fatalf("MatchGroups: matched values invalid")
	}
}

func TestMatchGroups_MultipleMatches(t *testing.T) {
	re := regexp.MustCompile("(?P<sign>[+-])?(?P<value>\\d)+(\\([+-]?\\d+\\)).+°C")
	groups := MatchGroups(re, "   _ /\"\".-.     -3(-6) °C      ")

	if len(groups.matches) != 2 {
		t.Fatalf("MatchGroups: expected 0 matches, actual %d", len(groups.matches))
	}

	expected := map[string]string{"value": "3", "sign": "-"}
	if !reflect.DeepEqual(groups.matches, expected) {
		t.Fatalf("MatchGroups: matched values invalid")
	}
}

func TestMatchGroups_NoMatches(t *testing.T) {
	re := regexp.MustCompile("(?P<value>\\d)+.+km/h")
	groups := MatchGroups(re, "")

	if len(groups.matches) != 0 {
		t.Fatalf("MatchGroups: expected 0 matches, actual %d", len(groups.matches))
	}
}

func TestMatchedGroups_Get(t *testing.T) {
	groups := &MatchedGroups{matches: map[string]string{"value": "4"}}

	t.Run("key=value", func(t *testing.T) {
		value, err := groups.Get("value")

		if err != nil {
			t.Fatalf("MatchedGroups_Get: expected 'value' but got error %s", err)
		}

		if value != "4" {
			t.Fatalf("MatchedGroups_Get: expected 'value' equal '4' but got %s", value)
		}
	})
	t.Run("key=sign", func(t *testing.T) {
		sign, err := groups.Get("sign")

		if err == nil {
			t.Fatalf("MatchedGroups_Get: expected error, but got value %s", sign)
		}
	})
}

func TestMatchedGroups_GetInt(t *testing.T) {
	groups := &MatchedGroups{matches: map[string]string{"value": "4", "unit": "km/h"}}

	t.Run("key=value", func(t *testing.T) {
		value, err := groups.GetInt("value")

		if err != nil {
			t.Fatalf("MatchedGroups_Get: expected 'value' but got error %s", err)
		}

		if value != 4 {
			t.Fatalf("MatchedGroups_Get: expected 'value' equal 4 but got %d", value)
		}
	})
	t.Run("key=sign", func(t *testing.T) {
		sign, err := groups.GetInt("sign")

		if err == nil {
			t.Fatalf("MatchedGroups_Get: expected error, but got value %d", sign)
		}
	})
	t.Run("key=unit", func(t *testing.T) {
		unit, err := groups.GetInt("unit")

		if err == nil {
			t.Fatalf("MatchedGroups_Get: expected cast error, but got value %d", unit)
		}
	})
}
