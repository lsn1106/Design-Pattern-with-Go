package main

import (
	"abstract_factory/furniture"
	"fmt"
)

func main() {
	artDecoFactory, _ := furniture.GetFurnitureFactory("art_deco")
	victorianFactory, _ := furniture.GetFurnitureFactory("victorian")
	modernFactory, _ := furniture.GetFurnitureFactory("modern")

	artDecoChair := artDecoFactory.CreateChair()
	artDecoSofa := artDecoFactory.CreateSofa()
	artDecoDesk := artDecoFactory.CreateDesk()

	victorianChair := victorianFactory.CreateChair()
	victorianSofa := victorianFactory.CreateSofa()
	victorianDesk := victorianFactory.CreateDesk()

	modernChair := modernFactory.CreateChair()
	modernSofa := modernFactory.CreateSofa()
	modernDesk := modernFactory.CreateDesk()

	fmt.Println(artDecoChair)
	fmt.Println(artDecoSofa)
	fmt.Println(artDecoDesk)

	fmt.Println(victorianChair)
	fmt.Println(victorianSofa)
	fmt.Println(victorianDesk)

	fmt.Println(modernChair)
	fmt.Println(modernSofa)
	fmt.Println(modernDesk)
}
