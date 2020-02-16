package main

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

type BtcPriceRequest struct {
	Time struct {
		Updated string `json:updated`
	} `json:'time'`
	Bpi struct {
		USD struct {
			Code string `json:"code"`
			Rate string `json:"rate"`
		} `json:"usd"`
	} `json:'bpi'`
}

type BtcPriceRecord struct {
	gorm.Model
	Value     string `json:"value"`
	Updated   string `json:"updated"`
	Timestamp int64  `json:"timestamp"`
}

func getBtcPrice(w http.ResponseWriter, r *http.Request) {
	var btcPrice BtcPriceRecord
	db.Order("ID ASC").Find(&btcPrice)
	json.NewEncoder(w).Encode(btcPrice)
}
func getBtcPrices(w http.ResponseWriter, r *http.Request) {
	var btcPrices []BtcPriceRecord
	db.Order("ID DESC").Limit(1000).Find(&btcPrices)
	json.NewEncoder(w).Encode(btcPrices)
}

func requestCurrentBtcPrice() {
	for true {
		t := &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 60 * time.Second,
		}
		c := &http.Client{
			Transport: t,
		}
		resp, err := c.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
		if err != nil {
			logger(err.Error(), "Btc Price Request")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger(err.Error(), "Btc Price Request")
		}
		if resp.StatusCode == 200 {
			var btcPrice BtcPriceRequest
			json.Unmarshal(body, &btcPrice)
			db.Create(&BtcPriceRecord{
				Value:     btcPrice.Bpi.USD.Rate,
				Updated:   btcPrice.Time.Updated,
				Timestamp: time.Now().Unix(),
			})
		}
		time.Sleep(3600 * time.Second)
	}
}
