package wttr

import (
	"reflect"
	"testing"
)

func TestParseWeather(t *testing.T) {
	body := `stockholm

     \  /       Partly cloudy
   _ /"".-.     -3(-6) °C      
     \_(   ).   → 4 km/h       
     /(___(__)  10 km          
                0.0 mm         
`
	expected := &Weather{
		Location:    "stockholm",
		Description: "Partly cloudy",
		Temperature: -3,
		WindSpeed:   4,
	}

	compareWeather(t, body, expected)
}

func TestParseWeather_Mist(t *testing.T) {
	body := `szczecin

                Mist
   _ - _ - _ -  -1(-3) °C      
    _ - _ - _   ← 7 km/h       
   _ - _ - _ -  3 km           
                0.0 mm         
`
	expected := &Weather{
		Location:    "szczecin",
		Description: "Mist",
		Temperature: -1,
		WindSpeed:   7,
	}

	compareWeather(t, body, expected)
}

func TestParseWeather_Clear(t *testing.T) {
	body := `berlin

      \   /     Clear
       .-.      +5(3) °C       
    ― (   ) ―   ↘ 4 km/h       
       ` + "`" + `-’      10 km
	/   \     0.0 mm
	`
	expected := &Weather{
		Location:    "berlin",
		Description: "Clear",
		Temperature: 5,
		WindSpeed:   4,
	}

	compareWeather(t, body, expected)
}

func TestParseWeather_Drizzle(t *testing.T) {
	body := `tokyo

   _` + "`" + `/"".-.     Patchy light drizzle
	,\_(   ).   17 °C
	/(___(__)  ↗ 40 km/h
	‘ ‘ ‘ ‘  5 km
	‘ ‘ ‘ ‘   0.2 mm
	`
	expected := &Weather{
		Location:    "tokyo",
		Description: "Patchy light drizzle",
		Temperature: 17,
		WindSpeed:   40,
	}

	compareWeather(t, body, expected)
}

func compareWeather(t *testing.T, body string, expected *Weather) {
	weather, err := ParseWeather(body)

	if err != nil {
		t.Fatalf("ParseWeather: unexpected error %s", err)
	}
	if !reflect.DeepEqual(weather, expected) {
		t.Fatalf("ParseWeather: incorrectly parsed weather")
	}
}

func TestParseWeather_Empty(t *testing.T) {
	_, err := ParseWeather("")

	if err == nil {
		t.Fatalf("ParseWeather: empty body error did not ocur")
	}
}

func TestParseWeather_MissingLines(t *testing.T) {
	body := `stockholm

`
	_, err := ParseWeather(body)

	if err == nil {
		t.Fatalf("ParseWeather: missing lines error did not ocur")
	}
}

func TestParseWeather_MalformedTemperature(t *testing.T) {
	body := `stockholm

     \  /       Partly cloudy
   _ /"".-.     xxx °C      
     \_(   ).   → 4 km/h       
     /(___(__)  10 km          
                0.0 mm         
`
	_, err := ParseWeather(body)

	if err == nil {
		t.Fatalf("ParseWeather: malformed temperature error did not ocur")
	}
}

func TestParseWeather_MalformedWindSpeed(t *testing.T) {
	body := `stockholm

     \  /       Partly cloudy
   _ /"".-.     -3(-6) °C        
     \_(   ).   → xxxxx    
     /(___(__)  10 km          
                0.0 mm         
`
	_, err := ParseWeather(body)

	if err == nil {
		t.Fatalf("ParseWeather: malformed temperature error did not ocur")
	}
}
