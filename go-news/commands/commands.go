package commands

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:  "top",
		Usage: "Top Headlines",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "c",
				Usage: "Country",
				Value: "ie",
			},
		},
		Action: headlines,
	},
	{
		Name:  "everything",
		Usage: "Everything",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "l",
				Usage: "Language",
				Value: "en",
			},
			cli.StringFlag{
				Name:  "s",
				Usage: "SortBy",
				Value: "popularity",
			},
		},
		Action: everything,
	},
}
