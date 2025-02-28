package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {

	if _, ok := os.LookupEnv("PORT"); !ok {
		os.Setenv("PORT", "9501")
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		escolha := strings.ToLower(c.QueryParams().Get("escolha"))

		novoJogo := NovoJogo()

		r := novoJogo.IniciarUmaPartidaContraMaquina(escolha)

		var statusCode int

		if r.ErrorMessage == "" {
			statusCode = http.StatusOK
		}

		statusCode = http.StatusBadRequest

		return c.JSON(statusCode, r)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
