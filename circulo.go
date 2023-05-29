package figuras

import "math"

// estructura de la clase
type Circulo struct {
	Radio float64
}

// implementación del método de la interface para el area
func (c Circulo) area() float64 {
	return math.Pow(c.Radio, 2) * math.Pi
}

// implementación del método de la interface para el perímetro
func (c Circulo) perimetro() float64 {
	return 2 * math.Pi * c.Radio
}
