package routes

import (
	"github.com/labstack/echo/v4"
	views "github.com/vukovlevi/vukovlevidev/views/auth"
)

func HandleLogin(c echo.Context) error {
    return render(views.Auth(), c)
}

func HandleLoginForm(c echo.Context) error {
    return render(views.Login(), c)
}

func HandleRegister(c echo.Context) error {
    return render(views.Register(), c)
}
