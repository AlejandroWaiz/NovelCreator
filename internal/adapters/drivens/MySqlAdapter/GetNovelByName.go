package mysqladapter

import "github.com/AlejandroWaiz/novels-box/internal/domain/structs"

func (db *MySqlAdapter) GetNovelByNameInDB(name string) (novels []structs.Novel, err error) {

	q := "SELECT * FROM `allnovels` WHERE `name` LIKE = ?"

	resp, err := db.dbconn.Query(q, name)

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

	return novels, err

}
