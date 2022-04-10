package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "DNS Lookup Utility"
	app.Usage = "Lets you query a hostname for its IP, NS, CNAME and MX records."

	cliFlags := []cli.Flag{
		&cli.StringFlag{
			Name:    "url",
			Aliases: []string{"u"},
		},
		&cli.StringFlag{
			Name:    "ip",
			Aliases: []string{"i"},
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ip",
			Usage: "Looks up a hostname's IP/s.",
			Flags: cliFlags,
			Action: func(ctx *cli.Context) error {
				ip, err := net.LookupIP(ctx.String("url"))
				if err != nil {
					fmt.Println(err)
				}
				for i := range ip {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "ns",
			Usage: "Looks up a hostname's Name Servers.",
			Flags: cliFlags,
			Action: func(ctx *cli.Context) error {
				ns, err := net.LookupNS(ctx.String("url"))
				if err != nil {
					fmt.Println(err)
				}
				for i := range ns {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up a hostname's CNAMEs.",
			Flags: cliFlags,
			Action: func(ctx *cli.Context) error {
				cname, err := net.LookupCNAME(ctx.String("url"))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up a hostname's MX records.",
			Flags: cliFlags,
			Action: func(ctx *cli.Context) error {
				mx, err := net.LookupMX(ctx.String("url"))
				if err != nil {
					fmt.Println(err)
				}
				for i := range mx {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
		{
			Name:  "ptr",
			Usage: "Reverse lookup of an IP",
			Flags: cliFlags,
			Action: func(ctx *cli.Context) error {
				ptr, err := net.LookupAddr(ctx.String("ip"))
				if err != nil {
					fmt.Println(err)
				}
				for i := range ptr {
					fmt.Println(ptr[i])
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
