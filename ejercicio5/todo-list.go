package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Tarea struct{
	nombre string
	descripcion string
	completada bool
}

type ListaTareas struct{
	tareas []Tarea
}

func (l *ListaTareas) agregarTarea(t Tarea){
	l.tareas = append(l.tareas, t)
	fmt.Println("Tarea agregada correctamente")
}

func (l *ListaTareas) marcarCompletado(indice int){
	l.tareas[indice].completada = true
}

func (l *ListaTareas) eliminarTarea(indice int){
	l.tareas = append(l.tareas[:indice], l.tareas[indice+1:]...)
}

func (l *ListaTareas) editarTarea(indice int, t Tarea){
	l.tareas[indice] = t
}

func (lista *ListaTareas) MostrarTareas(){
	fmt.Println("Listado de Tareas: ")
	fmt.Println("================================")
	for ind, t:= range lista.tareas{
		fmt.Printf("%d. %s - %s - completada: %t \n", ind, t.nombre, t.descripcion, t.completada)
	}
	fmt.Println("================================")
}

func crearTarea(action string) Tarea{
	leer := bufio.NewReader(os.Stdin)
	fmt.Printf("Ingrese nombre de la tarea que desea %s: ", action)
	nombre, _ := leer.ReadString('\n')
	fmt.Printf("Ingrese descripcion de la tarea que desea %s: ", action)
	descr, _ := leer.ReadString('\n')
	return Tarea{nombre: nombre, descripcion: descr}
}

func (lista *ListaTareas) obtenerIndice(action string) (int, error) {
	fmt.Printf("Ingresar indice de la tarea %s: ", action)
	var indice int
	fmt.Scanln(&indice)

	if indice<0 || indice>=len(lista.tareas){
		return -1, errors.New("Valor de íncide incorrecto")
	}

	return indice, nil
}

func main(){
	lista := ListaTareas{}

	for{
		fmt.Println("Seleccione una opción: ")
		fmt.Println("1. Agregar tarea")
		fmt.Println("2. Marcar tarea como completada")
		fmt.Println("3. Editar Tarea")
		fmt.Println("4. Eliminar Tarea")
		fmt.Println("5. Salir")
		fmt.Println("Ingrese la opción:")
		var opcion int
		fmt.Scanln(&opcion)
	
		switch opcion{
		case 1:
			//agregar
			t := crearTarea("agregar")
			lista.agregarTarea(t)
		case 2:
			//marcar tarea como completada
			indice, err := lista.obtenerIndice("marcar como completada")
			if err != nil{
				fmt.Println("Error: ", err)
				break
			}
			lista.marcarCompletado(indice)
		case 3:
			//editar una tarea
			indice, err := lista.obtenerIndice("editar")
			if err != nil{
				fmt.Println("Error: ", err)
				break
			}
			t := crearTarea("editar")

			lista.editarTarea(indice, t)
		case 4:
			//Eliminar tarea
			indice, err := lista.obtenerIndice("eliminar")
			if err != nil{
				fmt.Println("Error: ", err)
				break
			}
			lista.eliminarTarea(indice)
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opcion incorrecta.")
		}
	
		lista.MostrarTareas()
	}
	
}