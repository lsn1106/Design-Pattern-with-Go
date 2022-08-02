package furniture

type ModernFurnitureFactory struct {
}

func (f *ModernFurnitureFactory) CreateChair() IChair {
	return &ModernChair{
		Chair: Chair{
			style:  "modern",
			legCnt: 0,
			weight: 10,
		},
	}
}

func (f *ModernFurnitureFactory) CreateSofa() ISofa {
	return &ModernSofa{
		Sofa: Sofa{
			style:  "modern",
			cost:   100,
			weight: 50,
		},
	}
}

func (f *ModernFurnitureFactory) CreateDesk() IDesk {
	return &ModernDesk{
		Desk: Desk{
			style:  "modern",
			legCnt: 2,
			cost:   100,
			weight: 100,
		},
	}
}
