package main

import (
	"log"
	"os"

	"github.com/JakubPluta/go-projects/go-news/commands"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "newsapi.org Golang client"
	app.Usage = "a cli for newsapi.org"
	app.UsageText = "news [global opts] top | everything [command opts] <QUERY>"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "key",
			Usage:  "API Key for Authentication",
			EnvVar: "NEWS_API_KEY",
		},
		cli.StringFlag{
			Name:  "url",
			Usage: "Base API url",
			Value: "https://newsapi.org",
		},
		cli.BoolFlag{
			Name:   "k",
			Usage:  "Skip SSL verify",
			Hidden: true,
		},
	}

	app.Commands = commands.Commands

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
