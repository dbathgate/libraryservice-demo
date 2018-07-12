package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../config"
	"../metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type Book struct {
	Id       string   `json:"id"`
	Title    string   `json:"title"`
	Author   string   `json:"author`
	Category []string `json:"category"`
}

type Library struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

func GetLibrary(w http.ResponseWriter, r *http.Request) {
	var books []Book

	resp, getErr := http.Get(fmt.Sprintf("%s%s", config.Config.Bookservice.Host, "/books"))

	if getErr != nil {
		http.Error(w, getErr.Error(), 500)
		metrics.HttpRequestsTotal.With(prometheus.Labels{"api": "getLibrary", "method": "GET", "status": "500"}).Inc()
		return
	}

	jsonBody, jsonErr := ioutil.ReadAll(resp.Body)

	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), 500)
		metrics.HttpRequestsTotal.With(prometheus.Labels{"api": "getLibrary", "method": "GET", "status": "500"}).Inc()
		return
	}

	json.Unmarshal(jsonBody, &books)

	library := Library{Id: "1", Name: "Dr. B's Library", Books: books}
	json.NewEncoder(w).Encode(library)
	metrics.HttpRequestsTotal.With(prometheus.Labels{"api": "getLibrary", "method": "GET", "status": "200"}).Inc()
}
