package prom

import (
	"fmt"
	"log"

	forecast "github.com/billykwooten/darksky/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/robfig/cron"
)

var (
	temperatureGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_temperature",
			Help: "Temperature in degree Celsius or Fahrenheit",
		},
		[]string{"city","latitude", "longitude"},
	)
	precipIntensity = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_precipitation_intensity",
			Help: "Precipitation intensity",
		},
		[]string{"city","latitude", "longitude"},
	)
	precipProbability = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_precipitation_probability",
			Help: "Precipitation probability",
		},
		[]string{"city","latitude", "longitude"},
	)
	apparentTemperature = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_apparent_temperature",
			Help: "Apparent temperature in degree Celsius or Fahrenheit",
		},
		[]string{"city","latitude", "longitude"},
	)
	dewPointTemperature = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_dew_point",
			Help: "Dew point in degree Celsius or Fahrenheit",
		},
		[]string{"city","latitude", "longitude"},
	)
	humidity = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_humidity",
			Help: "Humidity",
		},
		[]string{"city","latitude", "longitude"},
	)
	pressure = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_pressure_mbar",
			Help: "Pressure in mB",
		},
		[]string{"city","latitude", "longitude"},
	)
	windSpeed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_wind_speed",
			Help: "Wind speed in km/h or mph",
		},
		[]string{"city","latitude", "longitude"},
	)
	windGust = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_wind_gust",
			Help: "Wind gust in km/h or mph",
		},
		[]string{"city","latitude", "longitude"},
	)
	windBearing = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_bearing_degree",
			Help: "Wind bearing",
		},
		[]string{"city","latitude", "longitude"},
	)
	cloudCover = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_cloud_cover",
			Help: "Cloud cover",
		},
		[]string{"city","latitude", "longitude"},
	)
	uvIndex = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_uv_index",
			Help: "UV index",
		},
		[]string{"city","latitude", "longitude"},
	)
	visibility = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_visibility",
			Help: "Visibility km or miles",
		},
		[]string{"city","latitude", "longitude"},
	)
	ozone = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_ozone_dobson",
			Help: "Ozone in dobson",
		},
		[]string{"city","latitude", "longitude"},
	)
	nearestStormDistance = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "darksky_nearestStormDistance",
			Help: "Nearest storm in km or miles",
		},
		[]string{"city","latitude", "longitude"},
	)
)

func init() {
	prometheus.MustRegister(temperatureGauge)
	prometheus.MustRegister(precipIntensity)
	prometheus.MustRegister(precipProbability)
	prometheus.MustRegister(apparentTemperature)
	prometheus.MustRegister(dewPointTemperature)
	prometheus.MustRegister(humidity)
	prometheus.MustRegister(pressure)
	prometheus.MustRegister(windSpeed)
	prometheus.MustRegister(windGust)
	prometheus.MustRegister(windBearing)
	prometheus.MustRegister(cloudCover)
	prometheus.MustRegister(visibility)
	prometheus.MustRegister(ozone)
	prometheus.MustRegister(nearestStormDistance)
}

func f2s(f float64) string {
	return fmt.Sprintf("%f", f)
}

func CollectSample(apikey string, latitude string, longitude string, city string) {
	log.Println("Collecting sample...")
	f, err := forecast.Get(apikey, latitude, longitude, "now", forecast.AUTO, forecast.English)
	if err != nil {
		log.Println(err)
		log.Fatalf("Skipping measurement due to error.")
		return
	}

	temperatureGauge.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.Temperature)
	precipIntensity.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.PrecipIntensity)
	precipProbability.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.PrecipProbability)
	apparentTemperature.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.ApparentTemperature)
	dewPointTemperature.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.DewPoint)
	humidity.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.Humidity)
	pressure.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.Pressure)
	windSpeed.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.WindSpeed)
	windBearing.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.WindBearing)
	cloudCover.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.CloudCover)
	visibility.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.Visibility)
	ozone.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.Ozone)
	nearestStormDistance.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.NearestStormDistance)
	windGust.With(prometheus.Labels{"latitude": f2s(f.Latitude), "longitude": f2s(f.Longitude), "city": city}).Set(f.Currently.WindGust)

}

func StartCron(apikey string, latitude string, longitude string, interval string, city string) {
	c := cron.New()
	c.AddFunc(fmt.Sprintf("@every %s", interval), func() { CollectSample(apikey, latitude, longitude, city) })
	c.Start()
}
