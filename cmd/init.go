package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ekkkkkknoes/TagVault/tagdb"

	"github.com/urfave/cli/v2"
)

func (cmd *Cmd) initcmd(cCtx *cli.Context) error {
	dbpath := filepath.Join(cmd.vaultdir, tagdb.DBNAME)
	if _, err := os.Stat(dbpath); err == nil {
		return fmt.Errorf("%s already exists", dbpath)
	}
	checkdbpath, _ := tagdb.GetDBPath(cmd.vaultdir)
	if checkdbpath != "" {
		fmt.Printf("Warning: found vaultdb at %s", checkdbpath)
	}
	return tagdb.Initdb(dbpath)
}
