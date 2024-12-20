package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/battleship/server/logger"
	"github.com/vukovlevi/vukovlevidev/routes"
)

func main() {
    debugMode := false
    flag.BoolVar(&debugMode, "debug", false, "debug mode enables the logger to write to specified output instead of default out.log and debug.txt")

    log := logger.CreateLogger(os.Stdout, os.Stdout, debugMode)

    err := godotenv.Load()
    if err != nil {
        log.Error("could not load .env file")
        os.Exit(1)
    }

    e := echo.New()

    e.GET("/", routes.HandleIndex)
    e.GET("/battleship", routes.HandleBattleship)

    e.Static("/", "public")

    port := os.Getenv("PORT")
    log.Info("server started", "port", port)

    s := http.Server{
        Addr: fmt.Sprintf(":%s", port),
        Handler: e,
    }

    if err = s.ListenAndServeTLS("cert/certificate.crt", "cert/private.key"); err != nil {
        log.Error("server could not be started", "err", err)
        os.Exit(1)
    }
}
