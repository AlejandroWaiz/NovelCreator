package mysqladapter

func (db *MySqlAdapter) RemoveNovelByTitleInDB(title string) error {

	q := "DELETE FROM `allnovels` WHERE `title` LIKE = ?"

	_, err := db.dbconn.Exec(q, title)

	if err != nil {

		return err

	}

	return nil

}
