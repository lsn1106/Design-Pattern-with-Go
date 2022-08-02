package furniture

type IChair interface {
	setChairStyle(style string) //억지메소드..
}

type Chair struct {
	style  string
	legCnt int
	weight int
}

func (c *Chair) setChairStyle(style string) {
	c.style = style
}

type ArtDecoChair struct {
	Chair
}

type VictorianChair struct {
	Chair
}

type ModernChair struct {
	Chair
}
