package mysqladapter

import "github.com/AlejandroWaiz/novels-box/internal/domain/structs"

func (db *MySqlAdapter) GetAllNovelsInDB() (novels []structs.Novel, err error) {

	q := "SELECT * FROM `allnovels`"

	query, err := db.dbconn.Query(q)

	if err != nil {

		return nil, err

	}

	for query.Next() {

		var novel structs.Novel

		err = query.Scan(&novel.ID, &novel.Title, &novel.Author, &novel.Genres)

		if err != nil {

			return nil, err

		}

		novels = append(novels, novel)

	}

	return novels, nil

}
