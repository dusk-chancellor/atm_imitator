package main

// пакет запуска (основная горутина)
// go run main.go

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dusk-chancellor/atm_imitator/atm"
	"github.com/dusk-chancellor/atm_imitator/configs"
	"github.com/gorilla/mux"
)

// ...

func main() {
	cfg := configs.ReadConfig()
	atm.InitLogger()

	accountService := atm.NewAccountService()

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	r := mux.NewRouter()
	r.HandleFunc("/accounts", atm.CreateAccountHandler(accountService)).Methods("POST")
	r.HandleFunc("/accounts/{id}/deposit", atm.DepositHandler(accountService)).Methods("POST")
	r.HandleFunc("/accounts/{id}/withdraw", atm.WithdrawHandler(accountService)).Methods("POST")
	r.HandleFunc("/accounts/{id}/balance", atm.GetBalance(accountService)).Methods("GET")

	log.Printf("Starting server on %s\n", addr)
	go log.Fatal(http.ListenAndServe(addr, r))	
}