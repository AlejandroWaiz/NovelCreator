package outadapters

import (
	"github.com/AlejandroWaiz/novels-box/internal/domain/structs"
)

type OutPort interface {
	CreateNovelInDB(structs.Novel) error
	RemoveNovelByNameInDB(name string) error
	GetNovelByNameInDB(name string) ([]structs.Novel, error)
}
