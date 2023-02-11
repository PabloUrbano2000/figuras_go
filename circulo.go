package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (c *Circulo) area() float64 {
	return c.Radio * c.Radio * math.Pi
}

func (c *Circulo) perimetro() float64 {
	return 2 * math.Pi * c.Radio
}
