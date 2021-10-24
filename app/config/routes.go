package config

import (
	"app/controllers"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.Path("/api/v1/{exchanges}/").
		Queries("market", "{market}", "period", "{period:1m|5m|15m|30m|1h|2h|4h|1d}").
		HandlerFunc(controllers.Signal).
		Name("SmaSignal").
		Methods("GET")
}
