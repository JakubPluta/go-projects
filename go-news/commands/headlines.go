package commands

import (
	"os"
	"time"

	"github.com/JakubPluta/go-projects/go-news/newsapi"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func headlines(c *cli.Context) error {

	key := c.GlobalString("key")

	countryCode := c.String("c")

	query := c.Args().Get(0)

	client, err := newsapi.New(key)

	if err != nil {
		return err
	}

	opts := newsapi.Options{Country: countryCode, Q: query}

	resp, err := client.GetTopHeadlines(opts)

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "Description", "Source", "Posted"})
	table.SetRowLine(true)

	for _, v := range resp.Articles {
		publishedAt := v.PublishedAt.Format(time.RFC1123)
		table.Append([]string{v.Title, v.Description, v.Source.Name, publishedAt})
	}

	table.Render()

	return err
}
