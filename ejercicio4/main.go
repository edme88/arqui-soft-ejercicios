package main

import (
	"fmt"
	"math/rand"
)

func main(){
	play()
}

func play(){
	numAleatorio := rand.Intn(100)
	var canIntentos int
	var numIngresado int

	for canIntentos < 10 {
		canIntentos++
		fmt.Printf("Ingrese un número (intentos restantes %d): ", 10-canIntentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("Felicitaciones Ganaste!")
			displayMenu()
			return
		} else if numIngresado>numAleatorio{
			fmt.Println("No acertaste. El número es menor.")
		}else {
			fmt.Println("No acertaste. El número es mayor.")
		}
	}
	fmt.Println("Se acabaron los intentos. El número es: ", numAleatorio)
	displayMenu()
}

func displayMenu(){
	var option string
	fmt.Print("Desea jugar nuevamente? (s/n): ")
	fmt.Scanln(&option)

	switch option{
	case "s":
		play()
	case "n":
		fmt.Println("Gracias por jugar!")
	default:
		fmt.Println("Opción incorrecta. Intente de Nuevo!")
		displayMenu()
	}
}