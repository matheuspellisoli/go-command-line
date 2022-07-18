package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a app de linha de comando pronta para ser executada
func Generate() *cli.App {

	app := cli.NewApp()

	app.Name = "Command line app"
	app.Usage = "Busca IPs e nomes de servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca o IPs de endereços na internet",
			Flags:  flags,
			Action: findIps,
		}, {
			Name:   "servers",
			Usage:  "Busca o nome de servidores na internet",
			Flags:  flags,
			Action: findServers,
		},
	}

	return app
}

func findServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func findIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
