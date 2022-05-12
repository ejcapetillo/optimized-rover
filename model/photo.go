package model

type Photo struct {
	Id        int64  `json:"id"`
	Sol       int    `json:"sol"`
	ImgSrc    string `json:"img_src"`
	EarthDate string `json:"earth_date"`
	Rover     Rover  `json:"rover"`
}

type PhotoWrapper struct {
	Photos []*Photo `json:"photos"`
}

type DailyPhoto struct {
	EarthDate string
	Count     int
}
