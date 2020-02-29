package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

type CryptoPriceRequest struct {
	Rates struct {
		Btc struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"btc"`
		Eth struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"eth"`
		Ltc struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"ltc"`
		Bch struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"bch"`
		Xrp struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"xrp"`
		Eos struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"eos"`
		Xlm struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"xlm"`
		Bnb struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"bnb"`
		Usd struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"usd"`
	} `json:'rates'`
}

type CryptoRecord struct {
	gorm.Model
	Eth       float64 `json:"eth"`
	Ltc       float64 `json:"ltc"`
	Xrp       float64 `json:"xrp"`
	Eos       float64 `json:"eos"`
	Xlm       float64 `json:"xlm"`
	Bnb       float64 `json:"bnb"`
	Updated   string  `json:"updated"`
	Timestamp int64   `json:"timestamp"`
}

func Round(x, unit float64) float64 {
	return math.Round(x*unit) / unit
}
func requestCurrentCryptoPrices() {
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
		resp, err := c.Get("https://api.coingecko.com/api/v3/exchange_rates")
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger(err.Error(), "Crypto Price Request")
		}
		if resp.StatusCode == 200 {
			var cryptoPrice CryptoPriceRequest
			json.Unmarshal(body, &cryptoPrice)
			btcPrice := cryptoPrice.Rates.Usd.Value
			unit := float64(10000)
			db.Create(&CryptoRecord{
				Eth:       Round(btcPrice/cryptoPrice.Rates.Eth.Value, unit),
				Ltc:       Round(btcPrice/cryptoPrice.Rates.Ltc.Value, unit),
				Xrp:       Round(btcPrice/cryptoPrice.Rates.Xrp.Value, unit),
				Eos:       Round(btcPrice/cryptoPrice.Rates.Eos.Value, unit),
				Xlm:       Round(btcPrice/cryptoPrice.Rates.Xlm.Value, unit),
				Bnb:       Round(btcPrice/cryptoPrice.Rates.Bnb.Value, unit),
				Updated:   formatDate(fmt.Sprintf("%s", time.Now())),
				Timestamp: time.Now().Unix(),
			})
		}
		time.Sleep(43200 * time.Second)
	}
}

func getLastCryptoPrice(w http.ResponseWriter, r *http.Request) {
	var cryptoRecord CryptoRecord
	db.Order("ID ASC").Find(&cryptoRecord)
	json.NewEncoder(w).Encode(cryptoRecord)
}

func getCryptoRecords(w http.ResponseWriter, r *http.Request) {
	var cryptoRecords []CryptoRecord
	db.Order("ID ASC").Find(&cryptoRecords)
	json.NewEncoder(w).Encode(cryptoRecords)
}
