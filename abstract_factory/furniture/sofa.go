package furniture

type ISofa interface {
	setSofaStyle(style string) //억지메소드..
}

type Sofa struct {
	style  string
	cost   int
	weight int
}

func (s *Sofa) setSofaStyle(style string) {
	s.style = style
}

type ArtDecoSofa struct {
	Sofa
}

type VictoriaSofa struct {
	Sofa
}

type ModernSofa struct {
	Sofa
}
