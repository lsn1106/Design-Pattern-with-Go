package place

type Building struct {
	Lat float32
	Lon float32
}

func (b *Building) Accept(v exportVisitor) {
	v.visitForBuilding(b)
}

func (b *Building) GetType() string {
	return "Building"
}
