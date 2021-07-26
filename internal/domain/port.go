package domain

import "github.com/AlejandroWaiz/novels-box/internal/domain/structs"

type Port interface {
	Add(structs.Novel) error
	DeleteByTitle(string) error
	GetByTitle(title string) (structs.Novel, error)
	GetAllNovels() ([]structs.Novel, error)
}
