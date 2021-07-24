package mysqladapter

import "github.com/AlejandroWaiz/novels-box/internal/domain/structs"

func (db *MySqlAdapter) CreateNovelInDB(novels structs.Novel) error {

	q := "INSERT INTO `allnovels` (ID, Title, Author, Genres) VALUES (?,?,?,?)"

	insert, err := db.dbconn.Prepare(q)

	defer insert.Close()

	if err != nil {

		return nil

	}

	_, err = insert.Exec(novels.ID, novels.Title, novels.Author, novels.Genres)

	if err != nil {

		return err

	}

	return nil

}
