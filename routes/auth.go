package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/vukovlevidev/models"
	"github.com/vukovlevi/vukovlevidev/views"
	v "github.com/vukovlevi/vukovlevidev/views/auth"
)

func HandleGetLogin(c echo.Context) error {
    return render(v.Auth(), c)
}

func HandleLoginForm(c echo.Context) error {
    return render(v.Login(), c)
}

func HandleGetRegister(c echo.Context) error {
    return render(v.Register(), c)
}

func HandlePostRegister(c echo.Context) error {
    c.Response().Header().Add("HX-Retarget", "#result")
    c.Response().Header().Add("HX-Reswap", "innerHTML")

    username := c.FormValue("username")
    password := c.FormValue("password")
    passwordRepeat := c.FormValue("password-repeat")

    if len(password) < 8 {
        c.Response().WriteHeader(http.StatusBadRequest)
        return render(views.Error("a jelszó túl rövid"), c)
    }

    if password != passwordRepeat {
        c.Response().WriteHeader(http.StatusBadRequest)
        return render(views.Error("a jelszavak nem egyeznek"), c)
    }

    user := models.User{
        Username: username,
        Password: password,
        Role: "user",
    }

    err := user.CreateUser()
    if err != nil {
        c.Response().WriteHeader(http.StatusInternalServerError)
        return render(views.Error("a felhasználót nem sikerült létrehozni"), c)
    }

    c.Response().WriteHeader(http.StatusCreated)
    return render(views.Error("a felhasználó sikeresen létrejött"), c) //TODO: change this to not render error
}
