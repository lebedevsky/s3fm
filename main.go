package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/larrabee/s3fm/fmdb"
	"github.com/jinzhu/gorm"
	"github.com/namsral/flag"
)

var db *gorm.DB

func main() {
	e := echo.New()
	e.HideBanner = true
	flag := parseFlags()
	db, err := fmdb.OpenDB("test.db")
	if err != nil {
		e.Logger.Fatalf("Cannot open DB. err: %s", err)
	}

	db.Close()
	// Middleware
	if flag.accessLogs {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	// Routes
	addRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(flag.listen))
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