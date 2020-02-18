package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type BsPriceRequest struct {
	USD struct {
		Dolartoday float64 `json:dolartoday`
	} `json:'dolartoday'`
}

type BsPriceRecord struct {
	gorm.Model
	Value     float64 `json:"value"`
	Updated   string  `json:"updated"`
	Timestamp int64   `json:"timestamp"`
}

type BsPriceRecordResponse struct {
	TableData []BsPriceRecord `json:"tableData"`
	ChartData []ChartAxis     `json:"chartData"`
}

func getBsPrice(w http.ResponseWriter, r *http.Request) {
	var bsPrice BsPriceRecord
	db.Order("ID ASC").Find(&bsPrice)
	json.NewEncoder(w).Encode(bsPrice)
}

func formatDate(date string) string {
	stringDate := fmt.Sprintf("%s", date)
	splittedString := strings.Split(stringDate, " ")
	splittedDate := fmt.Sprintf("%s %s %s%s", splittedString[0], splittedString[1], splittedString[2], splittedString[3])
	return splittedDate
}
func getBsPrices(w http.ResponseWriter, r *http.Request) {
	var bsPrices []BsPriceRecord
	db.Order("ID ASC").Limit(1000).Find(&bsPrices)

	var chartData []ChartAxis
	for i := 0; i < len(bsPrices); i++ {
		if err != nil {
			logger(err.Error(), "Converting price to Integer")
		}
		chartAxis := ChartAxis{
			X: bsPrices[i].Updated,
			Y: bsPrices[i].Value,
		}
		chartData = append(chartData, chartAxis)
	}
	response := BsPriceRecordResponse{
		TableData: bsPrices,
		ChartData: chartData,
	}
	json.NewEncoder(w).Encode(response)
}

func requestCurrentbsPrice() {
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
		resp, err := c.Get("https://s3.amazonaws.com/dolartoday/data.json")
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger(err.Error(), "BS Price Request")
		}
		if resp.StatusCode == 200 {
			var bsPrice BsPriceRequest
			json.Unmarshal(body, &bsPrice)
			db.Create(&BsPriceRecord{
				Value:     bsPrice.USD.Dolartoday,
				Updated:   formatDate(fmt.Sprintf("%s", time.Now())),
				Timestamp: time.Now().Unix(),
			})
		}
		time.Sleep(3600 * time.Second)
	}
}
