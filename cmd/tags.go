package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func (cmd *Cmd) tagnew(cCtx *cli.Context) error {
	tagname := cCtx.Args().Get(0)
	if tagname == "" {
		return fmt.Errorf("empty tagname")
	}
	return cmd.tdb.TagNew(tagname, -1)
}
