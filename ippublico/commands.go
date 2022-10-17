package ippublico

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli/v2"
)

// * Comando search busca pelos IPs
func Search() *cli.Command {
	return &cli.Command{
		Name:    "search",
		Aliases: []string{"s"},
		Usage:   "Busca o IP de endereços na internet",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "google.com",
			},
		},
		Action: searchIp,
	}

}

func searchIp(ctx *cli.Context) error {
	fmt.Println("*** SEARCH FOR PUBLIC IP ***")

	// Pega o valor da Flag host
	host := ctx.String("host")
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

// * Comando serves busca pelos servidores
func Servers() *cli.Command {
	return &cli.Command{
		Name:  "servers",
		Usage: "Busca o nome dos servidores na internet",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "google.com",
			},
		},
		Action: searchServers,
	}
}

func searchServers(ctx *cli.Context) error {
	fmt.Println("*** SEARCH FOR SERVERS ***")

	host := ctx.String("host")
	fmt.Println("Host: " + host)

	// ns => Name Server
	if servers, err := net.LookupNS(host); err != nil {
		log.Fatalln(err)
	} else {
		for _, server := range servers {
			fmt.Println("Server: ", server.Host)
		}
	}
	return nil
}
