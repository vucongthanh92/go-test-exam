package models

type CalculateDistanceStockCityReq struct {
	ClientID  string `json:"client_id"`
	StockCity string `json:"city" form:"city"`
}

type CalculateDistanceStockCityRes struct {
	StockCity string  `json:"stock_city"`
	Distance  float64 `json:"distance"`
	Unit      string  `json:"unit"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
