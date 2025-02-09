package color

type Okhsl struct {
	H, S, L float64
}

func (c Okhsl) HSL() (h, s, l float64) {
	return c.H, c.S, c.L
}
