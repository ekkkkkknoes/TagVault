package cmd

import (
	"github.com/urfave/cli/v2"
)

type Cmd struct {
	app      *cli.App
	vaultdir string
}

func (cmd *Cmd) Run(arguments []string) error {
	return cmd.app.Run(arguments)
}

func CreateApp() Cmd {
	var cmd Cmd
	cmd.app = &cli.App{
		Name: "TagVault",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "vaultdir",
				Value:       "",
				Aliases:     []string{"d"},
				Destination: &cmd.vaultdir,
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "init",
				Action: cmd.initcmd,
			},
			{
				Name: "query",
			},
			{
				Name: "querydir",
			},
			{
				Name: "list",
				Subcommands: []*cli.Command{
					{
						Name: "tags",
					},
					{
						Name: "untagged",
					},
					{
						Name: "albums",
					},
					{
						Name: "files",
					},
				},
			},
			{
				Name: "album",
				Subcommands: []*cli.Command{
					{
						Name: "new",
					},
					{
						Name: "add",
					},
					{
						Name: "remove",
					},
					{
						Name: "show",
					},
					{
						Name: "dir",
					},
				},
			},
			{
				Name: "file",
				Subcommands: []*cli.Command{
					{
						Name: "new",
					},
					{
						Name: "modify",
					},
					{
						Name: "addtags",
					},
					{
						Name: "removetags",
					},
					{
						Name: "remove",
					},
					{
						Name: "show",
					},
				},
			},
			{
				Name: "tag",
				Subcommands: []*cli.Command{
					{
						Name: "new",
					},
					{
						Name: "alias",
					},
					{
						Name: "unalias",
					},
					{
						Name: "remove",
					},
				},
			},
		},
	}
	return cmd
}
