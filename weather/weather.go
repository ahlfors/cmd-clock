package weather

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Weather weather info
type Weather struct {
	City string
	Tem  string
	Wea  string
}

// GetWeather get weather info
func GetWeather() (*Weather, error) {
	req, _ := http.NewRequest("GET", "http://www.baidu.com/s?word=天气", nil)
	req.Header["User-Agent"] = []string{"Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1"}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}
	const selectPrefix = "article > section > div.c-row-tile > div.ms-weather-top.c-gap-inner-bottom"
	tem := doc.Find(selectPrefix + " > div.ms-weather-main.ms-weather-main-wrapper > div.ms-weather-main-temp").Text()
	wea := doc.Find(selectPrefix + " > div.ms-weather-main.ms-weather-main-wrapper > div.ms-weather-main-wind.c-line-clamp1 > span:nth-child(1)").Text()
	city := doc.Find(selectPrefix + " > section.ms-weather-btns > div.c-row.c-row-tight > div.WA_LOG_OTHER.c-span7 > div > span").Text()
	if city == "" {
		return nil, nil
	}
	return &Weather{city, tem, wea}, nil
}
