package ippublico

import (
	"fmt"
	"log"
	"net"

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
			{
				Name:    "search",
				Aliases: []string{"s"},
				Usage:   "Busca o IP de endereços na internet",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Value: "www.google.com",
					},
				},
				Action: search,
			},
		},
	}
	return
}

func search(context *cli.Context) error {
	fmt.Println("*** SEARCH FOR PUBLIC IP ***")

	// Pega o valor da Flag host
	host := context.String("host")
	fmt.Printf("Host: %s\n", host)

	// LookupIp retorna um slice com os ips
	if ips, err := net.LookupIP(host); err != nil {
		log.Fatalln(err)
	} else {
		for _, ip := range ips {
			fmt.Println("Ip: ", ip)
		}
	}

	return nil
}
