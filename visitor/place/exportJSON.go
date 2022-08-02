package place

import (
	"fmt"
)

type ExportJSON struct {
	json interface{} //any
}

func (j *ExportJSON) visitForBuilding(b *Building) {
	fmt.Println("exporting Building to JSON")
}

func (j *ExportJSON) visitForPark(p *Park) {
	fmt.Println("exporting Park to JSON")
}
