package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
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

		fmt.Printf("%s", repository)
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
