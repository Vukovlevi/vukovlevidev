package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/vukovlevidev/views"
)

func render(component templ.Component, c echo.Context) error {
    return component.Render(c.Request().Context(), c.Response())
}

func HandleIndex(c echo.Context) error {
    return c.Redirect(http.StatusPermanentRedirect, "/battleship")
}

func HandleBattleship(c echo.Context) error {
    return render(views.Battleship(), c)
}
