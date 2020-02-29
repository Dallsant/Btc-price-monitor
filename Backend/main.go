package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error
var postgressPassword = os.Getenv("PGP")

var (
	outfile, _ = os.OpenFile("./err.log", os.O_RDWR|os.O_APPEND, 0755)
	l          = log.New(outfile, "", 0)
)

func main() {
	fmt.Println("Starting App")
	pgurl := fmt.Sprintf("host=localhost port=5432 user=postgres dbname=btclogger password=%s", postgressPassword)
	db, err = gorm.Open("postgres", pgurl)
	if err != nil {
		logger(err.Error(), "PostgresDB conexion")
	} else if err == nil {
		fmt.Println("Postgress: Connected")
	}
	db.SingularTable(true)
	InitialMigration()
	time.Sleep(1 * time.Second)

	// go func() {
	// 	requestCurrentBtcPrice()
	// }()
	// go func() {
	// 	requestCurrentbsPrice()
	// }()
	go func() {
		requestCurrentCryptoPrices()
	}()
	handleRequests()
}

func InitialMigration() {
	db.AutoMigrate(&BtcPriceRecord{}, &BsPriceRecord{}, &CryptoRecord{})
}

func logger(err string, domain string) {
	fmt.Println(fmt.Sprintf("$timestamp: %v $at: %s $message: %v", time.Now().Unix(), domain, err))

	l.Println(fmt.Sprintf("$timestamp: %v $at: %s $message: %v", time.Now().Unix(), domain, err))
}
