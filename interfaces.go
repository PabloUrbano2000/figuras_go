package figuras

import "fmt"

type Forma interface {
	area() float64
	perimetro() float64
}

func Medidas(g Forma) {
	fmt.Println(g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perímetro:", g.perimetro())
}
