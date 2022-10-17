package ippublico

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// init retorna a aplicação pronta para ser executada
func Run() (app *cli.App) {
	app = &cli.App{
		Name:  "ippublico",
		Usage: "Retorna o ip publico do site",
		//? Função padrão quando o cli é iniciado
		Action: func(*cli.Context) error {
			fmt.Println("Aplicação funcionando corretamente")
			return nil
		},
		//? Criando os comandos
		Commands: []*cli.Command{
			Search(),
			Servers(),
		},
	}
	return
}
