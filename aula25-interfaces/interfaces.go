package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2) //Raio ao quadrado (raio * raio)
}

type forma interface {
	area() float64
}

func calcArea(f forma) float64 {
	return f.area()
}

func main() {
	r := retangulo{10, 15}
	fmt.Printf("A area da retangulo é %0.2f \n", calcArea(r))

	c := circulo{10}
	fmt.Printf("A area da circulo é %0.2f \n", calcArea(c))

}
