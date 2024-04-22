package tagdb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type TagDB struct {
	db *sql.DB
}

func Open(dbpath string) (*TagDB, error) {
	tdb := TagDB{}
	db, err := sql.Open("sqlite3", dbpath)
	tdb.db = db
	return &tdb, err
}

func (tdb *TagDB) Close() {
	tdb.db.Close()
}

func Initdb(dbpath string) error {
	tdb, err := Open(dbpath)
	if err != nil {
		return err
	}
	_, err = tdb.db.Exec(`
	CREATE TABLE tags(
		tagid INTEGER NOT NULL,
		tagname TEXT PRIMARY KEY
	);
	CREATE TABLE files(
		fileid INTEGER PRIMARY KEY,
		path TEXT NOT NULL,
		timestamp INTEGER NOT NULL,
		description TEXT
	);
	CREATE TABLE albums(
		albumid INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT
	);
	CREATE TABLE albumentries(
		albumid INTEGER NOT NULL REFERENCES albums(albumid),
		fileid INTEGER NOT NULL REFERENCES files(fileid),
		ordernr INTEGER NOT NULL,
		PRIMARY KEY(albumid, ordernr)
	);
	CREATE TABLE taggings(
		fileid INTEGER NOT NULL REFERENCES files(fileid),
		tagid INTEGER NOT NULL REFERENCES tags(tagid),
		PRIMARY KEY(fileid, tagid)
	);
	`)
	tdb.Close()
	return err
}
