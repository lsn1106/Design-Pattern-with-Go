package furniture

type ArtDecoFurnitureFactory struct {
}

func (f *ArtDecoFurnitureFactory) CreateChair() IChair {
	return &ArtDecoChair{
		Chair: Chair{
			style:  "Art Deco",
			legCnt: 0,
			weight: 10,
		},
	}
}

func (f *ArtDecoFurnitureFactory) CreateSofa() ISofa {
	return &ArtDecoSofa{
		Sofa: Sofa{
			style:  "Art Deco",
			cost:   100,
			weight: 50,
		},
	}
}

func (f *ArtDecoFurnitureFactory) CreateDesk() IDesk {
	return &ArtDecoDesk{
		Desk: Desk{
			style:  "Art Deco",
			legCnt: 2,
			cost:   100,
			weight: 100,
		},
	}
}
