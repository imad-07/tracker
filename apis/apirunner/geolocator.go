package apirunner

import (
	"api/apis/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func Geolocator(city string) string {
	latitude, longitude, err := getCityCoordinates(city)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// Format the coordinates with more precision
	coords := fmt.Sprintf("%.6f,%.6f", latitude, longitude)
	// Use the correct Google Maps URL format
	mapURL := fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%s", url.QueryEscape(coords))
	return mapURL
}

func getCityCoordinates(city string) (float64, float64, error) {
	baseURL := "https://nominatim.openstreetmap.org/search"
	params := url.Values{}
	params.Add("q", city)
	params.Add("format", "json")
	params.Add("limit", "1")

	resp, err := http.Get(baseURL + "?" + params.Encode())
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var results []structs.NominatimResponse
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return 0, 0, err
	}

	if len(results) == 0 {
		return 0, 0, fmt.Errorf("no results found for city: %s", city)
	}

	lat, err := strconv.ParseFloat(results[0].Lat, 64)
	if err != nil {
		return 0, 0, err
	}

	lon, err := strconv.ParseFloat(results[0].Lon, 64)
	if err != nil {
		return 0, 0, err
	}

	return lat, lon, nil
}
