package clients

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewOpenWeatherClient_Connection(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	server.Close()

	cli := NewOpenWeatherClient("testKey", server.URL+"/weather")

	_, err := cli.LocationCurrentTemperature(decimal.NewFromFloat(1), decimal.NewFromFloat(1))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "connection refused")
}

func TestNewOpenWeatherClient_BadTocken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	cli := NewOpenWeatherClient("", server.URL+"/weather")
	_, err := cli.LocationCurrentTemperature(decimal.NewFromFloat(55.7558), decimal.NewFromFloat(37.6173))
	assert.Error(t, err)
}

func TestNewGoogleWeatherClient_Connection(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	server.Close()

	cli := NewGoogleWetherClient("testKey", server.URL+"/weather")

	_, err := cli.LocationCurrentTemperature(decimal.NewFromFloat(1), decimal.NewFromFloat(1))
	assert.Error(t, err)

}

func TestNewGoogleWeatherClient_BadTocken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	cli := NewGoogleWetherClient("", server.URL+"/weather")
	_, err := cli.LocationCurrentTemperature(decimal.NewFromFloat(55.7558), decimal.NewFromFloat(37.6173))
	assert.Error(t, err)
}

func TestNewOpenWeatherClientCoords_Connection(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	server.Close()

	cli := NewOpenWeatherCoords("testKey", server.URL)

	_, err := cli.GetLocation("London")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "connection refused")
}

func TestNewOpeeWeatherClientCoords_BadTocken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	cli := NewOpenWeatherCoords("", server.URL)
	_, err := cli.GetLocation("London")
	assert.Error(t, err)
}
