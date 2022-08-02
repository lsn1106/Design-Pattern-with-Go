package furniture

import "fmt"

type FurnitureFactory interface {
	CreateChair() IChair
	CreateSofa() ISofa
	CreateDesk() IDesk
}

func GetFurnitureFactory(concept string) (FurnitureFactory, error) {
	if concept == "art_deco" {
		return &ArtDecoFurnitureFactory{}, nil
	} else if concept == "victorian" {
		return &VictorianFurnitureFactory{}, nil
	} else if concept == "modern" {
		return &ModernFurnitureFactory{}, nil
	}

	return nil, fmt.Errorf("unknown furniture concept passed")
}
