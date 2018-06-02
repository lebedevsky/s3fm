package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
	"github.com/namsral/flag"
	"github.com/lebedevsky/s3fm/fmdb"
	"github.com/labstack/gommon/log"
)

var db *gorm.DB

func main() {
	var err error
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)
	flags := parseFlags()

	if db, err = fmdb.OpenDB("test.db"); err != nil {
		e.Logger.Fatalf("Cannot open DB. err: %s", err)
	}

	// Middleware
	if flags.accessLogs {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	// Routes
	addRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(flags.listen))
}

func parseFlags() (flags Flags) {
	flag.StringVar(&flags.listen, "CFG_LISTEN", "127.0.0.1:7075", "Listen interface and port")
	flag.BoolVar(&flags.accessLogs, "CFG_ACCESS_LOGS", false, "Enable access logs")
	flag.Parse()
	return
}

type Flags struct {
	listen string
	accessLogs bool
}
