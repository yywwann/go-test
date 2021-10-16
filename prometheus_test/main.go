package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"strconv"
	"time"
)

// 记录 cpu 温度
var cpuTemp = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	},
)

// 统计所有 url 访问次数
var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

// 统计所有 url 访问次数
var requestSource = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_source",
		Help: "Number of requests from ip.",
	},
	[]string{"ip"},
)

//
var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

//
var httpDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
		// 自定义 buckets 范围
		//Buckets: []float64{0.005, 0.02, 0.20, 0.40, 0.80},
	},
	[]string{"path"},
)

func init() {
	prometheus.Register(cpuTemp)
	prometheus.Register(totalRequests)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)
	prometheus.Register(requestSource)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		c.Next()
		//statusCode := c.Writer.Status()
		statusCode, ok := c.Get("code")
		if !ok {
			statusCode = 500
		}
		responseStatus.WithLabelValues(strconv.Itoa(statusCode.(int))).Inc()

		ip := c.ClientIP()
		requestSource.WithLabelValues(ip).Inc()

		totalRequests.WithLabelValues(path).Inc()
		timer.ObserveDuration()
	}
}

func main() {
	cpuTemp.Set(65.3)

	r := gin.New()
	e := r.Group("/", prometheusMiddleware())
	e.GET("/", HelloWorldHandler)

	r.GET("/prometheus", prometheusHandler())

	r.Run()
}

func HelloWorldHandler(c *gin.Context) {
	nowTime := time.Now().UnixNano()
	rand.Seed(nowTime)
	code := rand.Intn(10)
	c.Set("code", code)
	c.JSON(200, fmt.Sprintf("Hello world!+%d", code))
}
