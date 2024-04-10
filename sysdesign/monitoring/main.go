package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "The current CPU temperature in celsius",
	})

	pingPathTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "ping_path_total",
		Help: "The total number of requests to /ping",
	})

	pingPathBucketMilliseconds = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "ping_path_response_bucket",
		Help:    "Histogram of /ping response times in milliseconds",
		Buckets: []float64{10, 100, 1000},
	})

	homePathTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "home_path_total",
		Help: "The total number of requests to /",
	}, []string{"path"})

	homePathBucketMilliseconds = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "home_path_response_bucket",
		Help:    "Histogram of / response times in milliseconds",
		Buckets: []float64{10, 100, 1000},
	})
)

func init() {
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(homePathTotal)
	prometheus.MustRegister(homePathBucketMilliseconds)
	prometheus.MustRegister(pingPathTotal)
	prometheus.MustRegister(pingPathBucketMilliseconds)
}

func main() {
	randCpuTemp()

	r := gin.Default()

	// Here's the magic where we hook up the prometheus to
	// a /metrics endpoint via gin and handlers.
	// We configure prometheus to hit this endpoint via
	// /prometheus/prometheus.yml.
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.LoadHTMLGlob("templates/*")

	r.GET("/ping", func(c *gin.Context) {
		start := time.Now()

		randCpuTemp()
		pingPathTotal.Inc()

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

		pingPathBucketMilliseconds.
			Observe(float64(time.Since(start).Milliseconds()))
	})

	r.GET("/", func(c *gin.Context) {
		start := time.Now()

		randCpuTemp()
		homePathTotal.WithLabelValues("/").Add(1)

		c.HTML(200, "index.tmpl", gin.H{
			"title": "Main website",
		})

		homePathBucketMilliseconds.
			Observe(float64(time.Since(start).Milliseconds()))
	})

	r.Run(":8080")
}

// randCpuTemp sets a random CPU temperature.
func randCpuTemp() {
	cpuTemp.Set(12.0 + rand.Float64()*2)
}
