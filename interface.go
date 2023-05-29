package figuras

import "fmt"

// interface de figuras geometricas
type FiguraGeometrica interface {
	area() float64
	perimetro() float64
}

// método de la interface para imprimir el area de las figuras
func ImprimirArea(figura FiguraGeometrica) {
	fmt.Printf("El area de la figura es: %.2f\n", figura.area())
}

// método de la interface para imprimir el perímetro de las figuras
func ImprimirPerimetro(figura FiguraGeometrica) {
	fmt.Printf("El perímetro de la figura es: %.2f\n", figura.perimetro())
}
