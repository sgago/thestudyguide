package main

import (
	"embed"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	templPath = "templates/*"

	indexPath   = "/"
	pingPath    = "/ping"
	metricsPath = "/metrics"
	healthzPath = "/healthz"

	indexTempl = "index.html"
)

var (
	//go:embed templates
	embeddedFiles embed.FS

	request_buckets = []float64{
		0.000,
		0.005,
		0.010,
		0.015,
		0.020,
		0.025,
		0.030,
		0.035,
		0.040,
		0.045,
		0.050,
		0.055,
	}

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
		Help:    "Histogram of /ping response times in seconds",
		Buckets: request_buckets,
	})

	indexPathTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "index_path_total",
		Help: "The total number of requests to /",
	}, []string{"path"})

	indexPathBucketMilliseconds = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "index_path_response_bucket",
		Help:    "Histogram of / response times in seconds",
		Buckets: request_buckets,
	})
)

func init() {
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(indexPathTotal)
	prometheus.MustRegister(indexPathBucketMilliseconds)
	prometheus.MustRegister(pingPathTotal)
	prometheus.MustRegister(pingPathBucketMilliseconds)
}

func main() {
	templ := template.
		Must(template.New("").
			ParseFS(embeddedFiles, templPath))

	go func() {
		for {
			randCpuTemp()
			time.Sleep(time.Second * 5)
		}
	}()

	r := gin.Default()

	r.SetHTMLTemplate(templ)

	r.GET(healthzPath, healthz)

	r.GET(indexPath, index)
	r.GET(pingPath, ping)

	// Here's the magic where we hook up the prometheus to
	// a /metrics endpoint via gin and handlers.
	// We configure prometheus to hit this endpoint via
	// /prometheus/prometheus.yml.
	r.GET(metricsPath, gin.WrapH(promhttp.Handler()))

	r.Run()
}

func index(c *gin.Context) {
	start := time.Now()

	indexPathTotal.WithLabelValues("/").Add(1)

	c.HTML(200, indexTempl, gin.H{
		"title": "GoGinAir",
	})

	indexPathBucketMilliseconds.
		Observe(float64(time.Since(start).Seconds()))
}

func ping(c *gin.Context) {
	start := time.Now()

	pingPathTotal.Inc()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

	pingPathBucketMilliseconds.
		Observe(float64(time.Since(start).Seconds()))
}

func healthz(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

// randCpuTemp sets a random CPU temperature.
func randCpuTemp() {
	cpuTemp.Set(12.0 + rand.Float64()*2)
}
