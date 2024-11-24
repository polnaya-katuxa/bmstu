package rand

type Computational struct {
	n          int
	x          int
	a, b, c, m int
}

func NewComputational(n, x0, a, b, c, m int) *Computational {
	return &Computational{
		n: n,
		x: x0,
		a: a,
		b: b,
		c: c,
		m: m,
	}
}

func (c *Computational) Rand() int {
	c.x = (c.a*c.x*c.x + c.b*c.x + c.c) % c.m

	r := Ranges[c.n]

	return c.x%(r.End-r.Start) + r.Start
}
