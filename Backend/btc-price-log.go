package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type BtcPriceRequest struct {
	Time struct {
		Updated string `json:updated`
	} `json:'time'`
	Bpi struct {
		USD struct {
			Code      string `json:"code"`
			Rate      string `json:"rate"`
			RateFloat string `json:"rate_float"`
		} `json:"usd"`
		EUR struct {
			Code      string `json:"code"`
			Rate      string `json:"rate"`
			RateFloat string `json:"rate_float"`
		} `json:"eur"`
	} `json:'bpi'`
}

type BtcPriceRecord struct {
	gorm.Model
	Value     float64 `json:"value"`
	Updated   string  `json:"updated"`
	Timestamp int64   `json:"timestamp"`
}

type ChartAxis struct {
	X string  `json:"x"`
	Y float64 `json:"y"`
}
type PriceRecordResponse struct {
	TableData []BtcPriceRecord `json:"tableData"`
	ChartData []ChartAxis      `json:"chartData"`
}

func getBtcPrice(w http.ResponseWriter, r *http.Request) {
	var btcPrice BtcPriceRecord
	db.Order("ID ASC").Find(&btcPrice)
	json.NewEncoder(w).Encode(btcPrice)
}
func formatStringValue(value string) float64 {
	value = strings.Replace(value, ",", "", 1)

	floatValue, err := strconv.ParseFloat(value, 64)
	fmt.Println(floatValue)
	if err != nil {
		logger(err.Error(), "Converting price to Integer")
	}
	return floatValue
}

// func formatDate(date string) string {
// 	stringDate := fmt.Sprintf("%s", date)
// 	splittedString := strings.Split(stringDate, " ")
// 	splittedDate := fmt.Sprintf("%s %s %s%s", splittedString[0], splittedString[1], splittedString[2], splittedString[3])
// 	return splittedDate
// }
func getBtcPrices(w http.ResponseWriter, r *http.Request) {
	var btcPrices []BtcPriceRecord
	db.Order("ID ASC").Limit(1000).Find(&btcPrices)

	var chartData []ChartAxis
	for i := 0; i < len(btcPrices); i++ {
		if err != nil {
			logger(err.Error(), "Converting price to Integer")
		}
		chartAxis := ChartAxis{
			X: btcPrices[i].Updated,
			Y: btcPrices[i].Value,
		}
		chartData = append(chartData, chartAxis)
	}
	response := PriceRecordResponse{
		TableData: btcPrices,
		ChartData: chartData,
	}
	json.NewEncoder(w).Encode(response)
}

func requestCurrentBtcPrice() {
	for true {
		t := &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   65 * time.Second,
				KeepAlive: 60 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 60 * time.Second,
		}
		c := &http.Client{
			Transport: t,
		}
		resp, err := c.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger(err.Error(), "Btc Price Request")
		}
		if resp.StatusCode == 200 {
			var btcPrice BtcPriceRequest
			json.Unmarshal(body, &btcPrice)
			db.Create(&BtcPriceRecord{
				Value:     formatStringValue(btcPrice.Bpi.USD.Rate),
				Updated:   formatDate(btcPrice.Time.Updated),
				Timestamp: time.Now().Unix(),
			})
		}
		time.Sleep(3600 * time.Second)
	}
}
