package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate return line command app
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Line command application"
	app.Usage = "Get IPs and server names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search internet address IPs",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(context *cli.Context) {
	host := context.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(context *cli.Context) {
	host := context.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
