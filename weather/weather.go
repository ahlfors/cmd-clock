package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// Weather weather info
type Weather struct {
	City string
	Tem1 string
	Tem2 string
	Wea  string
}

type jsonTemplate struct {
	City string `json:"city"`
	Data []struct {
		Wea  string `json:"wea"`
		Tem1 string `json:"tem1"`
		Tem2 string `json:"tem2"`
	} `json:"data"`
}

// GetWeather get weather info
func GetWeather() (*Weather, error) {
	ip, err := getCurrentIP()
	if err != nil {
		return nil, err
	}
	return getWeatherByIP(ip)
}

func getCurrentIP() (string, error) {
	b, err := httpGet("http://whatismyip.akamai.com")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func getWeatherByIP(ip string) (*Weather, error) {
	jsonData, err := httpGet(fmt.Sprintf("https://www.tianqiapi.com/api/?version=v1&ip=%s", ip))
	if err != nil {
		return nil, err
	}
	t := jsonTemplate{}
	err = json.Unmarshal(jsonData, &t)
	if err != nil {
		return nil, err
	}
	return &Weather{
		City: quoteToCharacter(t.City),
		Tem1: quoteToCharacter(t.Data[0].Tem1),
		Tem2: quoteToCharacter(t.Data[0].Tem2),
		Wea:  quoteToCharacter(t.Data[0].Wea),
	}, nil
}

func httpGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func quoteToCharacter(str string) string {
	r, err := regexp.Compile("\\\\u[a-fA-F0-9]{4}")
	if err != nil {
		panic(err)
	}
	return r.ReplaceAllStringFunc(str, func(s string) string {
		unicodeStr := strings.TrimPrefix(s, "\\u")
		c, err := strconv.ParseInt(unicodeStr, 16, 32)
		if err != nil {
			return "-"
		}
		return fmt.Sprintf("%c", c)
	})
}
