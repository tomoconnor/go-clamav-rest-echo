package main

import (
	"github.com/dutchcoders/go-clamd"
	"github.com/labstack/echo/v4"
)

func pingHandler(clam *clamd.Clamd) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := clam.Ping()
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(500, "Could not ping clamd")
		}
		return c.JSON(200, "OK")
	}
}

func scanHandler(clam *clamd.Clamd) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		file, err := c.FormFile("file")
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(500, "Could not get file")
		}
		src, err := file.Open()
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(500, "Could not open file")
		}
		defer src.Close()
		response, err := clam.ScanStream(src, make(chan bool))
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(500, "Could not scan file")
		}
		result := <-response
		if result.Status == "FOUND" {
			c.Logger().Errorf("Malware detected in file %v", name)
			return echo.NewHTTPError(451, "Malware detected")
		} else {
			c.Logger().Infof("No malware detected in file %v", name)
			return c.JSON(200, "OK")
		}
	}
}

func scanResponseHandler(clam *clamd.Clamd) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		file, err := c.FormFile("file")
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(500, "Could not get file")
		}
		src, err := file.Open()
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(500, "Could not open file")
		}
		defer src.Close()
		response, err := clam.ScanStream(src, make(chan bool))
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(500, "Could not scan file")
		}
		result := <-response
		if result.Status == "FOUND" {
			c.Logger().Errorf("Malware detected in file %v -- %v", name, result.Raw)
			return echo.NewHTTPError(451, result)
		} else {
			c.Logger().Infof("No malware detected in file %v", name)
			return c.JSON(200, result)
		}
	}
}
