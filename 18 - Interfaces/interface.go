package main

import (
	"fmt"
	"math"
)

type forma interface {
	calculaArea() float32
}

type retangulo struct {
	altura  float32
	largura float32
}

type circulo struct {
	raio float32
}

func escreverArea(f forma) {
	fmt.Printf("Área é : %0.2f", f.calculaArea())
}

func (r retangulo) calculaArea() float32 {
	return r.altura * r.largura
}

func (c circulo) calculaArea() float32 {
	return float32(math.Pi * math.Pow(float64(c.raio), 2))
}

func main() {
	retang := retangulo{10, 15}
	escreverArea(retang)

	circ := circulo{10}
	escreverArea(circ)
}
