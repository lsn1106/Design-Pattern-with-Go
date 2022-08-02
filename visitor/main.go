package main

import (
	"fmt"
	"visitor/place"
)

func main() {
	empireStateBuilding := &place.Building{Lat: 100.0, Lon: 32.0}
	lotteWorld := &place.Park{Lat: 101.0, Lon: 34.0}

	exportJSON := &place.ExportJSON{}
	empireStateBuilding.Accept(exportJSON)
	lotteWorld.Accept(exportJSON)

	fmt.Println()
	exportXML := &place.ExportXML{}
	empireStateBuilding.Accept(exportXML)
	lotteWorld.Accept(exportXML)
}
