package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func (cmd *Cmd) tagnew(cCtx *cli.Context) error {
	tagname := cCtx.Args().Get(0)
	if tagname == "" {
		return fmt.Errorf("empty tagname")
	}
	return cmd.tdb.NewTag(tagname, -1)
}

func (cmd *Cmd) tagalias(cCtx *cli.Context) error {
	if cCtx.Args().Len() < 2 {
		return fmt.Errorf("not enough arguments")
	}
	args := cCtx.Args().Slice()
	tagid, err := cmd.tdb.GetTagID(args[0])
	if err != nil {
		return err
	}
	for _, tagname := range args[1:] {
		if err := cmd.tdb.NewTag(tagname, tagid); err != nil {
			return err
		}
	}
	return nil
}

func (cmd *Cmd) removetag(cCtx *cli.Context) error {
	if cCtx.Args().Len() < 1 {
		return fmt.Errorf("not enough arguments")
	}
	for _, tagname := range cCtx.Args().Slice() {
		tagid, err := cmd.tdb.GetTagID(tagname)
		if err != nil {
			return err
		}
		tagcount, err := cmd.tdb.GetTagAliasCount(tagid)
		if err != nil {
			return err
		}
		tagusage, err := cmd.tdb.GetTagUsageCount(tagid)
		if err != nil {
			return err
		}
		if tagcount < 2 && tagusage > 0 && !cCtx.Bool("force") {
			fmt.Fprintf(os.Stderr, "Warning: %s has no other aliases, is still in use, and -f is not specified. Skipping.\n", tagname)
			continue
		}
		if err := cmd.tdb.RemoveTag(tagname); err != nil {
			return err
		}
	}
	return nil
}
