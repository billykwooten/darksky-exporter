package main

import (
	"log"
	"net/http"
	"os"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/darksky_exporter/geo-golang"
	"github.com/darksky_exporter/prom"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app      = kingpin.New("darksky_exporter", "DarkSky Exporter for DarkSky Weather API").Author("Billy Wooten")
	addr     = app.Flag("listen-address", "HTTP port to listen on").Envar("LISTEN_ADDRESS").Default(":9091").String()
	apikey   = app.Flag("api-key", "DarkSky API Key").Envar("APIKEY").Required().String()
	city     = app.Flag("city", "City for DarkSky to gather metrics from.").Envar("CITY").Default("New York, NY").String()
	interval = app.Flag("interval", "Interval to poll the DarkSky API.").Envar("INTERVAL").Default("2m").String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	latitude, longitude := geo_golang.Get_coords(openstreetmap.Geocoder(), *city)
	prom.CollectSample(*apikey, latitude, longitude)
	prom.StartCron(*apikey, latitude, longitude, *interval)

	http.Handle("/metrics", prometheus.Handler())

	log.Printf("Listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
