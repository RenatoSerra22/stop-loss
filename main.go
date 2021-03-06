package main

import (
	"fmt"
	"github.com/RenatoSerra22/stop-loss/commsclients"
	"github.com/RenatoSerra22/stop-loss/core"
	"github.com/RenatoSerra22/stop-loss/exchangeclients"
	"github.com/RenatoSerra22/stop-loss/handlers"
	log "github.com/cihub/seelog"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	core.LoadEnvVars()
	commsclients.Load()
	exchangeclients.Load()
	addr, _ := determineListenAddress()

	commsclients.Telegram.Send(commsclients.TelegramUser, "Bot Started up!")

	go core.CheckPrices()
	go commsclients.Telegram.Start()

	router := mux.NewRouter()
	router.HandleFunc("/health", handlers.HealthEndpointHandler).Methods("GET")
	log.Critical(http.ListenAndServe(addr, router))
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
