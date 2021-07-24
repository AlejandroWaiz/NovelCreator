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

func (nc *NovelsDomain) Add(novels structs.Novel) error {

	err := nc.Database.CreateNovelInDB(novels)

	if err != nil {

		return err

	}

	return nil

}

func (nc *NovelsDomain) DeleteByName(name string) error {

	err := nc.Database.RemoveNovelByNameInDB(name)

	if err != nil {

		return err

	}

	return nil
}

func (nc *NovelsDomain) GetByName(name string) ([]structs.Novel, error) {

	novels, err := nc.Database.GetNovelByNameInDB(name)

	if err != nil {

		return nil, err

	}

	return novels, nil

}
