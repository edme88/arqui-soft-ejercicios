package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct{
	ID string `json:"id"`
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

func getAlbums(ctx *gin.Context){
	ctx.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(ctx *gin.Context){
	var newAlbum album

	if err := ctx.BindJSON(&newAlbum); err != nil{
		return
	}
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByID(ctx *gin.Context){
	id := ctx.Param("id")

	for _, a := range albums{
		if a.ID == id{
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

func putAlbumByID(ctx *gin.Context){
	id := ctx.Param("id")
	var modifyAlbum album

	if err := ctx.BindJSON(&modifyAlbum); err != nil{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Datos incorrectos"})
		return
	}

	for index, a:= range albums {
		if a.ID == id {
			albums[index] = modifyAlbum
			albums[index].ID = id
			ctx.IndentedJSON(http.StatusOK, albums[index])
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

func deleteAlbumByID(ctx *gin.Context){
	id := ctx.Param("id")
	for index, a := range albums {
		if a.ID == id {
			albums = append(albums[:index], albums[index+1:]...)
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "album eliminado"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.PUT("/albums/:id", putAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.Run("localhost:8080")
}