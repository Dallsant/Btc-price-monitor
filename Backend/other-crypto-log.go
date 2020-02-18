package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
			logger(err.Error(), "BS Price Request")
		}
		if resp.StatusCode == 200 {
			var cryptoPrice CryptoPriceRequest
			json.Unmarshal(body, &cryptoPrice)
			db.Create(&CryptoRecord{
				Eth:       cryptoPrice.Rates.Eth.Value,
				Ltc:       cryptoPrice.Rates.Ltc.Value,
				Xrp:       cryptoPrice.Rates.Xrp.Value,
				Eos:       cryptoPrice.Rates.Eos.Value,
				Xlm:       cryptoPrice.Rates.Xlm.Value,
				Bnb:       cryptoPrice.Rates.Bnb.Value,
				Updated:   formatDate(fmt.Sprintf("%s", time.Now())),
				Timestamp: time.Now().Unix(),
			})
		}
		time.Sleep(43200 * time.Second)
	}
}
