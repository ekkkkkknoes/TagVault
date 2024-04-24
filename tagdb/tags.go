package tagdb

func (tdb *TagDB) NewTag(tagname string, id int) error {
	var err error
	if id == -1 {
		_, err = tdb.db.Exec("INSERT INTO tags VALUES((SELECT COALESCE(max(tagid), 0) + 1 FROM tags), ?);", tagname)
	} else {
		_, err = tdb.db.Exec("INSERT INTO tags VALUES(?, ?);", id, tagname)
	}
	return err
}

func (tdb *TagDB) GetTagID(tagname string) (int, error) {
	var tagid int
	err := tdb.db.QueryRow("SELECT tagid FROM tags WHERE tagname IS ?;", tagname).Scan(&tagid)
	return tagid, err
}

func (tdb *TagDB) GetTagNames(tagid int) ([]string, error) {
	rows, err := tdb.db.Query("SELECT tagname FROM tags WHERE tagid IS ?;", tagid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tagnames := make([]string, 0)
	for rows.Next() {
		var tagname string
		if err := rows.Scan(&tagname); err != nil {
			return nil, err
		}
		tagnames = append(tagnames, tagname)
	}
	return tagnames, rows.Err()
}

func (tdb *TagDB) GetTagAliasCount(tagid int) (int, error) {
	var aliascount int
	err := tdb.db.QueryRow("SELECT COUNT(tagname) FROM tags WHERE tagid IS ?;", tagid).Scan(&aliascount)
	return aliascount, err
}

func (tdb *TagDB) RemoveTag(tagname string) error {
	tagid, err := tdb.GetTagID(tagname)
	if err != nil {
		return err
	}
	idcount, err := tdb.GetTagAliasCount(tagid)
	if err != nil {
		return err
	}
	if _, err := tdb.db.Exec("DELETE FROM tags WHERE tagname IS ?;", tagname); err != nil {
		return err
	}
	if idcount == 1 {
		_, err = tdb.db.Exec("DELETE FROM taggings WHERE tagid IS ?;", tagid)
	}
	return err
}

func (tdb *TagDB) GetTagUsageCount(tagid int) (int, error) {
	var usagecount int
	err := tdb.db.QueryRow("SELECT COUNT(fileid) FROM taggings WHERE tagid IS ?;", tagid).Scan(&usagecount)
	return usagecount, err
}
