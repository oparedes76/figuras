package figuras

// struct de la clase figura
type Rectangulo struct {
	Base   float64
	Altura float64
}

// implementación del método de la interface para el area
func (r Rectangulo) area() float64 {
	return r.Base * r.Altura
}

// implementación del método de la interface para el perímetro
func (r Rectangulo) perimetro() float64 {
	return (2 * r.Base) + (2 * r.Altura)
}
