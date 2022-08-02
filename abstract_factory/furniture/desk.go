package furniture

type IDesk interface {
	setDeskStyle(style string) //억지메소드..
}

type Desk struct {
	style  string
	legCnt int
	cost   int
	weight int
}

func (d *Desk) setDeskStyle(style string) {
	d.style = style
}

type ArtDecoDesk struct {
	Desk
}

type VictorianDesk struct {
	Desk
}

type ModernDesk struct {
	Desk
}
