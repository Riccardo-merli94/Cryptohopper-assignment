package controllers

import (
	"app/exchanges"
	"app/exchanges/coinbase"
	"app/indicators/sma"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var SupportedExchanges = map[string]exchange.Exchange{
	"coinbasepro": coinbase.New(),
}

func Signal(w http.ResponseWriter, req *http.Request) {
	requestedExchange := mux.Vars(req)["exchanges"]
	market := mux.Vars(req)["market"]
	period := mux.Vars(req)["period"]

	exchangeService, isSupported := SupportedExchanges[requestedExchange]

	if !isSupported {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{error: Exchange not supported}")
		return
	}

	indicator, err := sma.New(exchangeService,market, period)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{error: Invalid period}")
		return
	}

	signal, err := indicator.Calculate()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{error: Could not compute the request}")
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(signal)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{error: Could not compute the request}")
	}
}