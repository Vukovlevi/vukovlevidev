package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vukovlevi/battleship/server/logger"
	"github.com/vukovlevi/vukovlevidev/db"
	"github.com/vukovlevi/vukovlevidev/middlewares"
	"github.com/vukovlevi/vukovlevidev/models"
	"github.com/vukovlevi/vukovlevidev/routes"
)

func main() {
    debugMode := false
    flag.BoolVar(&debugMode, "debug", false, "debug mode enables the logger to write to specified output instead of default out.log and debug.txt")

    log := logger.CreateLogger(os.Stdout, os.Stdout, debugMode)
    models.SetLogger(&log)

    err := db.Connect()
    if err != nil {
        log.Error("could not connect to database", "err", err)
        os.Exit(1)
    }

    err = godotenv.Load()
    if err != nil {
        log.Error("could not load .env file")
        os.Exit(1)
    }

    e := echo.New()

    e.GET("/", routes.HandleIndex)
    e.GET("/battleship", routes.HandleBattleship)

    e.GET("/login", routes.HandleGetLogin)
    e.POST("/login", routes.HandlePostLogin)
    e.GET("/login-form", routes.HandleLoginForm)
    e.GET("/register", routes.HandleGetRegister)
    e.POST("/register", routes.HandlePostRegister)

    adminGroup := e.Group("/admin")
    adminGroup.Use(middlewares.AuthorizeAdminUser)
    adminGroup.GET("", routes.HandleGetAdmin)

    e.Static("/", "public")

    port := os.Getenv("PORT")
    log.Info("server started", "port", port)

    s := http.Server{
        Addr: fmt.Sprintf(":%s", port),
        Handler: e,
    }

    env := os.Getenv("ENVIRONMENT")

    if env == "PROD" {
        if err = s.ListenAndServeTLS("cert/certificate.crt", "cert/private.key"); err != nil {
            log.Error("server could not be started", "err", err)
            os.Exit(1)
        }
    } else {
        if err = s.ListenAndServe(); err != nil {
            log.Error("server could not be started", "err", err)
            os.Exit(1)
        }
    }
}
