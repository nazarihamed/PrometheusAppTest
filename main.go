package main

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Device struct {
	ID       int    `json:"id"`
	MAC      string `json:"mac"`
	FIRMWARE string `json:"firmware"`
}

var dvs []Device

func init() {
	dvs = []Device{
		{1, "5F-33-CC-1F-43-82", "2.1.6"},
		{2, "EF-2B-C4-F5-D6-34", "2.1.6"},
	}
}

func getDevices(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(dvs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

type Metrics struct {
	timestamp prometheus.Gauge
	devices   prometheus.Gauge
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "PrometheusAppTest",
			Name:      "Connected_Devices",
			Help:      "Number of currently connected devices.",
		}),
		timestamp: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "TimeStamp",
			Help: "Timestamp for collecting stats",
		}),
	}
	reg.MustRegister(m.timestamp)
	reg.MustRegister(m.devices)
	return m
}

func main() {
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())
	m := NewMetrics(reg)
	m.devices.Set(float64(len(dvs)))

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	http.Handle("/metrics", promHandler)
	http.HandleFunc("/devices", getDevices)
	http.ListenAndServe(":8081", nil)

}
