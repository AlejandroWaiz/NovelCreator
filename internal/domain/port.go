package domain

import "github.com/AlejandroWaiz/novels-box/internal/domain/structs"

type Port interface {
	Add(structs.Novel) error
	DeleteByName(string) error
	GetByName(name string) ([]structs.Novel, error)
}
