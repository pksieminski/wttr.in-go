package wttr

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestClient_GetCurrentWeather(t *testing.T) {
	mux := setupMux(t)

	s := httptest.NewServer(mux)
	defer s.Close()

	expected := &Weather{
		Location:    "fakecity",
		Description: "Partly cloudy",
		Temperature: -3,
		WindSpeed:   4,
	}

	c := getClient(t, s)
	weather, err := c.GetCurrentWeather("fakecity")

	if err != nil {
		t.Fatalf("Client_GetCurrentWeather: unexpected fetch error %s", err)
	}
	if !reflect.DeepEqual(weather, expected) {
		t.Fatalf("ParseWeather: incorrectly parsed weather")
	}
}

func TestClient_GetCurrentWeather_404(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/notfoundcity", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, err := io.WriteString(w, "404 Page Not Found")

		if err != nil {
			t.Fatalf("ParseWeather: unexpected error %s", err)
		}
	})

	s := httptest.NewServer(mux)
	defer s.Close()

	c := getClient(t, s)
	_, err := c.GetCurrentWeather("fakecity")

	if err == nil {
		t.Fatalf("Client_GetCurrentWeather: expected 404 error did not occur")
	}
}

func setupMux(t *testing.T) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/fakecity", func(w http.ResponseWriter, r *http.Request) {
		response := `fakecity

     \  /       Partly cloudy
   _ /"".-.     -3(-6) °C      
     \_(   ).   → 4 km/h       
     /(___(__)  10 km          
                0.0 mm         
`
		_, err := io.WriteString(w, response)

		if err != nil {
			t.Fatalf("ParseWeather: unexpected error %s", err)
		}
	})

	return mux
}

func getClient(t *testing.T, s *httptest.Server) *Client {
	serverUrl, err := url.Parse(s.URL)
	if err != nil {
		t.Error(err)
	}

	return NewClientWithConfiguration(&Configuration{URL: *serverUrl})

}
