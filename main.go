package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"github.com/jinzhu/configor"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

var Config = struct {
	Github struct {
		AccessToken string
	}
	Label []struct {
		Name        string
		Color       string
		Description string
	}
}{}

func main() {
	configor.Load(&Config, "config.toml")
	app := cli.NewApp()

	app.Name = "gh-label-setup"
	app.Usage = "Setup Github label in the repository"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		var owner, repository string
		if c.NArg() > 0 {
			owner = strings.Split(c.Args().Get(0), "/")[0]
			repository = strings.Split(c.Args().Get(0), "/")[1]
		} else {
			return errors.New("Repository not given")
		}

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: Config.Github.AccessToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

		for _, label := range Config.Label {
			l := &github.Label{
				Name:        &label.Name,
				Color:       &label.Color,
				Description: &label.Description,
			}
			_, _, err := client.Issues.CreateLabel(ctx, owner, repository, l)
			if err != nil {
				fmt.Printf("%s\n", err)
			}
		}

		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
