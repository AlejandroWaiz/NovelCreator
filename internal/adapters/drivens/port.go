package outadapters

import (
	"github.com/AlejandroWaiz/novels-box/internal/domain/structs"
)

type OutPort interface {
	CreateNovelInDB(structs.Novel) error
	RemoveNovelByTitleInDB(title string) error
	GetNovelByTitleInDB(title string) (structs.Novel, error)
	GetAllNovelsInDB() ([]structs.Novel, error)
}
