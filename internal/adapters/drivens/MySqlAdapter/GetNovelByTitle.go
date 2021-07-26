package mysqladapter

import "github.com/AlejandroWaiz/novels-box/internal/domain/structs"

func (db *MySqlAdapter) GetNovelByTitleInDB(title string) (novel structs.Novel, err error) {

	q := "SELECT * FROM `allnovels` WHERE `title` LIKE = ?"

	row := db.dbconn.QueryRow(q, title)

	err = row.Scan(&novel)

	if err != nil {

		var empty structs.Novel

		return empty, err

	}

	return novel, nil

	/* resp, err := db.dbconn.Query(q, title)

	if err != nil {

		return novels, err

	}

	for resp.Next() {

		var searchedNovel structs.Novel

		err := resp.Scan(searchedNovel.ID, searchedNovel.Title, searchedNovel.Author, searchedNovel.Genres)

		if err != nil {

			return novels, err

		}

		novels = append(novels, searchedNovel)

		break

	}

	return novels, err */

}
