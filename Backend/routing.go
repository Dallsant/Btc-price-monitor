package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logger(err.Error(), "IP address retrievement")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	IP := GetOutboundIP()
	port := ":8000"
	router.HandleFunc("/btc-prices", getBtcPrices).Methods("GET")
	router.HandleFunc("/btc-price", getBtcPrice).Methods("GET")

	fmt.Println(fmt.Sprintf(fmt.Sprintf("Starting Server at http://%s%s", IP, port)))
	log.Fatal(http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}
