package main

import (
	"fmt"
	_ "main/docs"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {

	if _, ok := os.LookupEnv("PORT"); !ok {
		os.Setenv("PORT", "9999")
	}

	e := echo.New()
	e.GET("/doc/*any", echoSwagger.WrapHandler)
	e.GET("/", Handler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

// @Summary		Faz a jogada do jokenpo
// @Description	Retorna um resultado de quem ganhou
// @Tags			result
// @Accept			json
// @Produce		json
// @Param			escolha			query		string													true	"escolha de qual jogada vai fazer (pedra,papel,tesoura)"
// @Success		200				{object}	[]Resultado	"Resultado"
// @Router			/ [get]
func Handler(c echo.Context) error {
	escolha := strings.ToLower(c.QueryParams().Get("escolha"))

	novoJogo := NovoJogo()

	r := novoJogo.IniciarUmaPartidaContraMaquina(escolha)

	var statusCode int

	if r.ErrorMessage == "" {
		statusCode = http.StatusOK
	}

	statusCode = http.StatusBadRequest

	return c.JSON(statusCode, r)
}
