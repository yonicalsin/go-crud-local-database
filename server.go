package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

// User is ...
type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Init is ...
type Init struct {
	Message   string `json:"message"`
	SatusCode int    `json:"statusCode"`
}

func saveUser(ctx echo.Context) error {
	user := &User{
		Name:     ctx.FormValue("name"),
		Username: ctx.FormValue("username"),
		Password: ctx.FormValue("password"),
	}

	return ctx.JSONPretty(http.StatusOK, user, "  ")
}
func getUser(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err == nil {
		fmt.Println("Error")
	}

	return ctx.JSONPretty(http.StatusOK, &User{
		ID:       id,
		Name:     "Yoni Calsin",
		Username: "yonicalsin",
		Password: "yonicalsin_password",
	}, "  ")

}

func main() {
	var server = echo.New()

	// Configuration
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// Middlewares
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// e.GET("/users/:id", getUser)
	server.GET("/", func(ctx echo.Context) error {
		return ctx.JSONPretty(http.StatusOK, &Init{
			Message:   "Bienvenido a la api de yoni calsin",
			SatusCode: 200,
		}, "  ")
	})

	server.POST("/user", saveUser)
	server.GET("/user/:id", getUser)

	server.Logger.Fatal(server.Start(":4000"))
}
