package manager

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/prometheus/common/log"
	"github.com/shao1f/app/model"
)

const (
	GaodeURL = "https://restapi.amap.com/v3"
	GaodeKey = "9e5e2d9521574c9ee9008e7879571d55"
)

func GetGaodeLoca(info model.LocationInfo) (string, error) {
	param := fmt.Sprintf("%s,%s", info.Lat, info.Lon)
	url := fmt.Sprintf("%v/assistant/coordinate/convert?locations=%v&key=%v", GaodeURL, param, GaodeKey)
	var gaode model.GaodeLocaResp
	if err := httpGet(url, &gaode); err != nil {
		log.Errorf("get gaode loca error,err=%v", err)
		return "", err
	}
	return gaode.Locations, nil
}

func GetGaodeCity(loca string) (string, error) {
	url := fmt.Sprintf("%v/geocode/regeo?location=%s&key=%v&radius=100&extensions=base", GaodeURL, loca, GaodeKey)
	var gaode model.GaodeRes
	if err := httpGet(url, &gaode); err != nil {
		log.Errorf("get gaode loca error,err=%v", err)
		return "", err
	}
	return gaode.Regeocode.AddressComponent.Adcode, nil
}

func GetGodeWeather(cityCode string) (*model.WeatherInfo, error) {
	url := fmt.Sprintf("%v/weather/weatherInfo?city=%v&key=%v", GaodeURL, cityCode, GaodeKey)
	weather := &model.WeatherResp{}
	if err := httpGet(url, weather); err != nil {
		log.Errorf("get gaode weather error,err=%v", err)
		return nil, err
	}
	if len(weather.Lives) == 0 {
		return nil, errors.New("weather err")
	}
	return weather.Lives[0], nil
}
