package domain

import (
	outadapters "github.com/AlejandroWaiz/novels-box/internal/adapters/drivens"
	"github.com/AlejandroWaiz/novels-box/internal/domain/structs"
)

type NovelsDomain struct {
	Database outadapters.OutPort
}

func CreateNewNovelsDomain(database outadapters.OutPort) *NovelsDomain {
	return &NovelsDomain{Database: database}
}

func (nd *NovelsDomain) Add(novels structs.Novel) error {

	err := nd.Database.CreateNovelInDB(novels)

	if err != nil {

		return err

	}

	return nil

}

func (nd *NovelsDomain) DeleteByTitle(title string) error {

	err := nd.Database.RemoveNovelByTitleInDB(title)

	if err != nil {

		return err

	}

	return nil
}

func (nd *NovelsDomain) GetByTitle(title string) (structs.Novel, error) {

	novel, err := nd.Database.GetNovelByTitleInDB(title)

	if err != nil {

		var empty structs.Novel

		return empty, err

	}

	return novel, nil

}

func (nd *NovelsDomain) GetAllNovels() ([]structs.Novel, error) {

	novels, err := nd.Database.GetAllNovelsInDB()

	if err != nil {

		return nil, err

	}

	return novels, err

}
