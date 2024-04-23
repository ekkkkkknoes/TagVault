package tagdb

func (tdb *TagDB) TagNew(tagname string, id int) error {
	var err error
	if id == -1 {
		_, err = tdb.db.Exec("INSERT INTO tags VALUES((SELECT COALESCE(max(tagid), 0) + 1 FROM tags), ?);", tagname)
	} else {
		_, err = tdb.db.Exec("INSERT INTO tags VALUES(?, ?);", id, tagname)
	}
	return err
}
