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

// VersionInfo version info strut
type VersionInfo struct {
	Version string `json:"version"`
}

type VersionInfoAll struct {
	LibraryVersion string `json:"libraryVersion"`
	BookVersion    string `json:"bookVersion"`
}

// GetVersion version get rest handler
func GetVersion(w http.ResponseWriter, r *http.Request) {

	metrics.HttpRequestsTotal.With(prometheus.Labels{"api": "getVersion", "method": "GET", "status": "200"}).Inc()

	version := VersionInfo{Version: config.Version}
	json.NewEncoder(w).Encode(version)
}

func GetVersionAll(w http.ResponseWriter, r *http.Request) {

	resp, getErr := http.Get(fmt.Sprintf("%s%s", config.Config.Bookservice.Host, "/version"))

	if getErr != nil {
		http.Error(w, getErr.Error(), 500)
		metrics.HttpRequestsTotal.With(prometheus.Labels{"api": "getVersionAll", "method": "GET", "status": "500"}).Inc()
		return
	}

	jsonBody, jsonErr := ioutil.ReadAll(resp.Body)

	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), 500)
		metrics.HttpRequestsTotal.With(prometheus.Labels{"api": "getVersionAll", "method": "GET", "status": "500"}).Inc()
		return
	}

	var versionInfo VersionInfo
	json.Unmarshal(jsonBody, &versionInfo)

	versionInfoAll := VersionInfoAll{BookVersion: versionInfo.Version, LibraryVersion: config.Version}
	metrics.HttpRequestsTotal.With(prometheus.Labels{"api": "getVersionAll", "method": "GET", "status": "200"}).Inc()

	json.NewEncoder(w).Encode(versionInfoAll)
}
