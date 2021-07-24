package mysqladapter

func (db *MySqlAdapter) RemoveNovelByNameInDB(name string) error {

	q := "DELETE FROM `allnovels` WHERE `name` LIKE = ?"

	_, err := db.dbconn.Exec(q, name)

	if err != nil {

		return err

	}

	return nil

}
