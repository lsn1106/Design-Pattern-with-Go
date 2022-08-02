package furniture

type VictorianFurnitureFactory struct {
}

func (f *VictorianFurnitureFactory) CreateChair() IChair {
	return &ModernChair{
		Chair: Chair{
			style:  "Victorian",
			legCnt: 0,
			weight: 10,
		},
	}
}

func (f *VictorianFurnitureFactory) CreateSofa() ISofa {
	return &ModernSofa{
		Sofa: Sofa{
			style:  "Victorian",
			cost:   100,
			weight: 50,
		},
	}
}

func (f *VictorianFurnitureFactory) CreateDesk() IDesk {
	return &ModernDesk{
		Desk: Desk{
			style:  "Victorian",
			legCnt: 2,
			cost:   100,
			weight: 100,
		},
	}
}
