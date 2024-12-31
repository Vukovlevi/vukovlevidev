package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/vukovlevidev/views/admin"
)

func HandleGetAdmin(c echo.Context) error {
    return render(views.AdminIndex(), c)
}
