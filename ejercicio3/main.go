package main

import (
	"fmt"
	"math"
)

func main(){
	var lado1, lado2 float64

	//Entrada de Datos
	fmt.Print("Ingrese lado 1: ")
	fmt.Scanln(&lado1)
	fmt.Print("Ingrese lado 2: ")
	fmt.Scanln(&lado2)

	//Cálculo
	area := (lado1 * lado2)/2
	fmt.Printf("El área del triángulo es: %.2f \n", area)
	hipotenusa := math.Sqrt(math.Pow(lado1,2)+math.Pow(lado2,2))
	perimetro := lado1 + lado2 + hipotenusa
	fmt.Printf("El perímetro del triángulo es: %.2f", perimetro)
}