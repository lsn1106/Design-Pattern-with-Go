package place

type exportVisitor interface {
	visitForBuilding(*Building)
	visitForPark(*Park)
}
