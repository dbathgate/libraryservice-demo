package route

import (
	"../metrics"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupHttpRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/library", GetLibrary).Methods("GET")
	router.HandleFunc("/version", GetVersion).Methods("GET")
	router.HandleFunc("/version/all", GetVersionAll).Methods("GET")
	router.Handle("/metrics", promhttp.HandlerFor(metrics.MetricRegistry, promhttp.HandlerOpts{})).Methods("GET")

	return router
}
