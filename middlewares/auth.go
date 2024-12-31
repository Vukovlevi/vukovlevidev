package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/vukovlevidev/auth"
)

func AuthorizeAdminUser(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cookie, err := c.Cookie("session_token")
        if err != nil {
            return c.Redirect(http.StatusPermanentRedirect, "/login")
        }

        err = auth.AuthorizeAdminUser(cookie.Value)
        if err != nil {
            return c.Redirect(http.StatusPermanentRedirect, "/login")
        }

        return next(c)
    }
}
