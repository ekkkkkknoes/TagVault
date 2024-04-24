package cmd

import (
	"fmt"

	"github.com/ekkkkkknoes/TagVault/tagdb"
	"github.com/urfave/cli/v2"
)

type Cmd struct {
	app      *cli.App
	vaultdir string
	tdb      *tagdb.TagDB
}

func (cmd *Cmd) opendb(cCtx *cli.Context) error {
	dbpath, err := tagdb.GetDBPath(cmd.vaultdir)
	if err != nil {
		return err
	}
	if dbpath == "" {
		return fmt.Errorf("no tagdb found, use TagVault init")
	}
	cmd.tdb, err = tagdb.Open(dbpath)
	return err
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
						Name:      "new",
						Action:    cmd.tagnew,
						Args:      true,
						ArgsUsage: "<tagname>",
						Before:    cmd.opendb,
					},
					{
						Name:      "alias",
						Action:    cmd.tagalias,
						Args:      true,
						ArgsUsage: "<tagname> <aliasname> [aliasname]...",
						Before:    cmd.opendb,
					},
					{
						Name:      "remove",
						Action:    cmd.removetag,
						Args:      true,
						ArgsUsage: "<tagname> [tagname]...",
						Before:    cmd.opendb,
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:    "force",
								Aliases: []string{"f"},
								Usage:   "Remove the tag(s) even if it's the last alias and a file has been tagged with it",
							},
						},
					},
				},
			},
		},
	}
	return cmd
}
