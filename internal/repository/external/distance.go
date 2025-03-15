package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/umahmood/haversine"
	"github.com/vucongthanh92/go-test-exam/helper/constants"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

func GetCoordinatesFromIP(ip string) (res models.Location, err error) {

	url := fmt.Sprintf("%s%s", constants.UrlCheckLocationIP, ip)

	resp, err := http.Get(url)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	lat, ok1 := result["lat"].(float64)
	lon, ok2 := result["lon"].(float64)
	if !ok1 || !ok2 {
		return res, errors.New("unable to get coordinates from IP")
	}

	res.Lat = lat
	res.Lon = lon

	return res, nil
}

type Location struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func GetCoordinatesFromCity(city string) (res models.Location, err error) {

	params := url.Values{}
	params.Set("q", city)
	params.Set("format", "json")
	params.Set("limit", "1")

	apiURL := fmt.Sprintf("%s?%s", constants.UrlCheckLocationCity, params.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	var locations []Location
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		return res, err
	}

	if len(locations) == 0 {
		return res, fmt.Errorf("no coordinates found for city: %s", city)
	}

	res.Lat, err = strconv.ParseFloat(locations[0].Lat, 64)
	res.Lon, err = strconv.ParseFloat(locations[0].Lon, 64)

	return res, nil
}

// Tính khoảng cách bằng Haversine
func CalculateDistance(src, dst models.Location) float64 {
	p1 := haversine.Coord{Lat: src.Lat, Lon: src.Lon}
	p2 := haversine.Coord{Lat: dst.Lat, Lon: dst.Lon}
	_, km := haversine.Distance(p1, p2)
	return km
}
