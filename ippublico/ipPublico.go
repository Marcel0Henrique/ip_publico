package ippublico

import "github.com/urfave/cli/v2"

// init retorna a aplicação pronta para ser executada
func Run() (app *cli.App) {
	app = &cli.App{
		Name:  "ippublico",
		Usage: "Retorna o ip publico do site",
	}
	return
}
