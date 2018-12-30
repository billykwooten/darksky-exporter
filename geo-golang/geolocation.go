package geo_golang

import (
	"github.com/codingsince1985/geo-golang"
	"log"
	"strconv"
)

func Get_coords(geocoder geo.Geocoder, city string) (string, string){
	location, _ := geocoder.Geocode(city)
	if location != nil {
		log.Printf("Location given was %s", city)
		log.Printf("Coords are (%.6f, %.6f)", location.Lat, location.Lng)
	} else {
		log.Fatalf("got <nil> location")
	}

	latitudef := strconv.FormatFloat(location.Lat, 'f', 6, 64)
	longitudef := strconv.FormatFloat(location.Lng, 'f', 6, 64)

	return latitudef, longitudef
}