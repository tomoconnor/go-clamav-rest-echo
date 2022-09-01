package main

import (
	"fmt"
	"os"

	"github.com/dutchcoders/go-clamd"
	promMW "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	filesScanned = promauto.NewCounter(prometheus.CounterOpts{
		Name: "files_scanned",
		Help: "Number of files scanned",
	})
	filesPositive = promauto.NewCounter(prometheus.CounterOpts{
		Name: "files_positive",
		Help: "Number of files with a detected malware signature",
	})
	filesNegative = promauto.NewCounter(prometheus.CounterOpts{
		Name: "files_negative",
		Help: "Number of files with no detected malware signature",
	})
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {

	clamHost := getEnv("CLAMD_HOST", "localhost")
	clamPort := getEnv("CLAMD_PORT", "3310")
	listenPort := getEnv("LISTEN_PORT", "8080")

	clamConnection := clamd.NewClamd(fmt.Sprintf("tcp://%v:%v", clamHost, clamPort))

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	p := promMW.NewPrometheus("echo", nil)
	p.Use(e)
	e.GET("/", pingHandler(clamConnection))
	e.GET("/healthz", pingHandler(clamConnection))

	e.POST("/scan", scanHandler(clamConnection))
	e.POST("/scanResponse", scanResponseHandler(clamConnection))

	e.Logger.Fatal(e.Start(":" + listenPort))
}
