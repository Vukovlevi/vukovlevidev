package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/vukovlevidev/auth"
	"github.com/vukovlevi/vukovlevidev/models"
	"github.com/vukovlevi/vukovlevidev/views"
	v "github.com/vukovlevi/vukovlevidev/views/auth"
)

func HandleGetLogin(c echo.Context) error {
    return render(v.Auth(), c)
}

func HandlePostLogin(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    _, err := models.GetUserByUsernameAndPassword(username, password)
    if err != nil {
        c.Response().WriteHeader(http.StatusUnauthorized)
        return render(views.Error("hibás felhasználónév vagy jelszó"), c)
    }

    return nil
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

    err := auth.ValidatePassword(password, passwordRepeat)
    if err != nil {
        c.Response().WriteHeader(http.StatusBadRequest)
        return render(views.Error(err.Error()), c)
    }

    user := models.User{
        Username: username,
        Password: password,
        Role: "user",
    }

    err = user.CreateUser()
    if err != nil {
        c.Response().WriteHeader(http.StatusInternalServerError)
        return render(views.Error("a felhasználót nem sikerült létrehozni"), c)
    }

    c.Response().WriteHeader(http.StatusCreated)
    return render(views.Error("a felhasználó sikeresen létrejött"), c) //TODO: change this to not render error
}
