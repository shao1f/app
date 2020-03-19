package model

import "time"

type Response struct {
	ErrCode  int    `json:"err_code"`
	ErrorMsg string `json:"error_msg"`
}

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

type LocationInfo struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type GaodeLocaResp struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	Infocode  string `json:"infocode"`
	Locations string `json:"locations"`
}

type GaodeCityResp struct {
	Status    string `json:"status"`
	Regeocode struct {
		AddressComponent struct {
			Adcode string `json:"adcode"`
		} `json:"addressComponent"`
	} `json:"regeocode"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
}

type GaodeRes struct {
	Status    string `json:"status"`
	Regeocode struct {
		AddressComponent struct {
			Adcode string `json:"adcode"`
		} `json:"addressComponent"`
	} `json:"regeocode"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
}

type WeatherResp struct {
	Status   string         `json:"status"`
	Count    string         `json:"count"`
	Info     string         `json:"info"`
	Infocode string         `json:"infocode"`
	Lives    []*WeatherInfo `json:"lives"`
}

type WeatherInfo struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	Adcode        string `json:"adcode"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	Winddirection string `json:"winddirection"`
	Windpower     string `json:"windpower"`
	Humidity      string `json:"humidity"`
	Reporttime    string `json:"reporttime"`
}

type WeatherResult struct {
	Response
	Data *WeatherInfo `json:"data"`
}

type SearchReq struct {
	City string `form:"city" json:"city"`
}

type SearchResp struct {
	Response
	Data *WeatherInfo `json:"data"`
}

type City struct {
	City   string `json:"city"`
	CityId int    `json:"city_id"`
}

type CityToId struct {
	City      string    `gorm:"city"`
	CityId    int       `gorm:"city_id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
