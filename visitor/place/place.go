package place

type place interface {
	GetType() string
	Accept(exportVisitor)
}
