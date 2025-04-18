package main

import "os"

const path = "hola.txt"
 
func CreateFile(path string) error {
  file, err := os.Create(path)
  defer file.Close()
 
  if err != nil {
     return err
  }
 
  return nil
}

func WriteFile(path string, content []byte) error {
	if err := os.WriteFile(path, content, 0644); err != nil {
	   return err
	}
	return nil
  }

func main(){
	data := []byte("Hola mundo!")
	CreateFile("hola.txt")
	WriteFile("hola.txt", data)
}