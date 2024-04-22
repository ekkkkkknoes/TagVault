package tagdb

import (
	"os"
	"path/filepath"
)

const DBNAME string = ".tagvault.db"

func GetDBPath(vaultdir string) (string, error) {
	var dbpath string
	if vaultdir != "" {
		dbpath, err := filepath.Abs(filepath.Join(vaultdir, DBNAME))
		if err != nil {
			return "", err
		}
		if _, err := os.Stat(dbpath); err == nil {
			return dbpath, nil
		}
		return "", nil
	}
	vaultdir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	var newdir string
	for {
		dbpath = filepath.Join(vaultdir, DBNAME)
		if _, err := os.Stat(dbpath); err == nil {
			return dbpath, nil
		}
		newdir = filepath.Dir(vaultdir)
		if newdir == vaultdir {
			break
		}
		vaultdir = newdir
	}
	return "", nil
}
