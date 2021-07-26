package mysqladapter

import "github.com/AlejandroWaiz/novels-box/internal/domain/structs"

func (db *MySqlAdapter) CreateNovelInDB(novel structs.Novel) error {

	q := "INSERT INTO `allnovels` (ID, Title, Author, Genres) VALUES (?,?,?,?)"

	insert, err := db.dbconn.Prepare(q)

	defer insert.Close()

	if err != nil {

		return nil

	}

	_, err = insert.Exec(novel.ID, novel.Title, novel.Author, novel.Genres)

	if err != nil {

		return err

	}

	return nil

}
