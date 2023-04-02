package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:                 "shook",
		Version:              "0.1",
		Usage:                "make shell could be invoke by webhook",
		EnableBashCompletion: true,
		Suggest:              true,
		Commands: []*cli.Command{
			{
				Name:        "init",
				Description: "set base url for shook",
				UsageText:   "shook init <base_url>",
				Action: func(context *cli.Context) error {
					cmdInit(context.Args().First(), context.Args().Get(1))
					return nil
				},
			},
			{
				Name:        "pwd",
				Description: "print the config file's dir",
				UsageText:   "shook pwd",
				Action: func(context *cli.Context) error {
					cmdPwd()
					return nil
				},
			},
			{
				Name:        "create",
				Aliases:     []string{"c"},
				UsageText:   "shook create <name> <shell>",
				Description: "create a new webhook",
				Action: func(context *cli.Context) error {
					wd, err := os.Getwd()
					if err != nil {
						return err
					}
					cmdCreate(context.Args().First(), wd, context.Args().Get(1))
					return nil
				},
			},
			{
				Name:        "del",
				Aliases:     []string{"d"},
				Description: "delete a webhook",
				UsageText:   "shook del <name>",
				Action: func(context *cli.Context) error {
					cmdDel(context.Args().First())
					return nil
				},
			},
			{
				Name:        "run",
				Aliases:     []string{"r"},
				Description: "trigger a webhook",
				UsageText:   "shook run <name>",
				Action: func(context *cli.Context) error {
					cmdRun(context.Args().First())
					return nil
				},
			},
			{
				Name:        "ls",
				Aliases:     []string{"l"},
				Description: "list all created webhooks",
				UsageText:   "shook ls",
				Action: func(context *cli.Context) error {
					cmdLs()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
