package main

import (
	"flag"
	"fmt"
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

    port := os.Getenv("PORT")
    if err = e.Start(fmt.Sprintf(":%s", port)); err != nil {
        log.Error("server could not be started", "err", err)
        os.Exit(1)
    }
}
