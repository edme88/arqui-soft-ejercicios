package main

import (
	"fmt"
	"os"
)

func readFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main(){
	texto,_ := readFile("hola.txt")
	fmt.Println(string(texto))
}