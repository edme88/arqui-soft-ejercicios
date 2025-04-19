package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"` //backtick
	Title string `json:"title"`
	Artist string `json:"artist"`
	Year uint `json:"year"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1",Title: "The Number of the Beast",Artist: "Iron Maiden",Year: 1982,Price: 25.19},
	{ID: "2",Title: "Youthanasia",Artist: "Medadeth",Year: 1994,Price: 13.65},
	{ID: "3",Title: "Master of Puppets",Artist: "Metallica",Year: 1986,Price: 20.97},
}

//1
func getAlbums(c *gin.Context) {
	//responde al cliente con la data serizalizada y con el status de HTTP
	c.IndentedJSON(http.StatusOK, albums)
}

//2
func postAlbums(c *gin.Context) {
	var newAlmbum album
	//recibo json y lo voy a transformar en la estructura tipo album
	if err := c.BindJSON(&newAlmbum); err != nil{
		return
	}
	albums = append(albums, newAlmbum)
	//responde al cliente con la data serizalizada y con el status de HTTP
	c.IndentedJSON(http.StatusCreated, newAlmbum)
}

//3
func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

//4
func putAlbumsByID(ctx *gin.Context){
	id := ctx.Param("id")
	var modifyAlbum album
	if err := ctx.BindJSON(&modifyAlbum); err != nil{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Datos inválidos"})
		return
	}
	
	for i, a := range albums {
		if a.ID == id {
			albums[i] = modifyAlbum
			albums[i].ID = id
			ctx.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

//5
func deleteAlbumsByID(ctx *gin.Context){
	id := ctx.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "album eliminado"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

func main(){
	router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// *gin.Context representa toda la información relacionada con una solicitud HTTP: 
	// lo que el cliente pidió, las cabeceras, los parámetros, el cuerpo, la IP, etc. 
	// Y también te permite construir la respuesta.
	//respuesta
	// 	c.JSON(200, gin.H{
	// 		"message": "hola mundo",
	// 	})
	// })

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", putAlbumsByID) //Quedo de Tarea
	router.DELETE("/albums/:id", deleteAlbumsByID) //Quedo de Tarea

	//levantar sevidor
	router.Run("localhost:8080")
}