package clients

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpenWetherClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "/weather?lat=55.7558&lon=37.6173&appid=testkey&units=metric", r.URL.String())

		resp := openWeatherResponse{
			Main: struct {
				Temp decimal.Decimal `json:"temp"`
			}{Temp: decimal.NewFromFloat(20.5)},
		}

		json.NewEncoder(w).Encode(resp)
	}))

	cli := NewOpenWeatherClient("testkey", ts.URL+"/weather")
	temp, err := cli.LocationCurrentTemperature(decimal.NewFromFloat(55.7558), decimal.NewFromFloat(37.6173))
	require.NoError(t, err)
	assert.Equal(t, temp, decimal.NewFromFloat(20.5))
}

func TestGoogleWetherClient(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/currentConditions:lookup?key=testkey&location.latitude=55.7558&location.longitude=37.6173", r.URL.String())

		resp := googleWeatherResponse{
			Temperature: struct {
				Degrees decimal.Decimal `json:"degrees"`
			}{Degrees: decimal.NewFromFloat(20.5)}}

		json.NewEncoder(w).Encode(resp)
	}))

	cli := NewGoogleWetherClient("testkey", ts.URL+"/currentConditions")
	temp, err := cli.LocationCurrentTemperature(decimal.NewFromFloat(55.7558), decimal.NewFromFloat(37.6173))
	require.NoError(t, err)
	assert.Equal(t, temp, decimal.NewFromFloat(20.5))
}
