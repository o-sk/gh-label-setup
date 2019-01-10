package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/jinzhu/configor"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

var Config = struct {
	Github struct {
		AccessToken string
	}
}{}

func main() {
	configor.Load(&Config, "config.toml")
	app := cli.NewApp()

	app.Name = "gh-label-setup"
	app.Usage = "Setup Github label in the repository"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		var repository string
		if c.NArg() > 0 {
			repository = c.Args().Get(0)
		} else {
			return errors.New("Repository not given")
		}

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: Config.Github.AccessToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

		labels, _, err := client.Issues.ListLabels(ctx, "o-sk", repository, nil)
		if err != nil {
			return err
		}

		for _, label := range labels {
			fmt.Printf("%s\n", github.Stringify(label.Name))
		}

		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
