package place

import (
	"fmt"
)

type ExportXML struct {
	xml interface{} //any
}

func (x *ExportXML) visitForBuilding(b *Building) {
	fmt.Println("exporting Building to XML")
}

func (x *ExportXML) visitForPark(p *Park) {
	fmt.Println("exporting Park to XML")
}
