package place

type Park struct {
	Lat float32
	Lon float32
}

func (p *Park) Accept(v exportVisitor) {
	v.visitForPark(p)
}

func (p *Park) GetType() string {
	return "Park"
}
