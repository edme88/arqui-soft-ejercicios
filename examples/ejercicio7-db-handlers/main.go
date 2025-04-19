package main

import (
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"vinyl-api/database"
	"vinyl-api/handlers"
)

func main(){

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	router := gin.Default()
	router.GET("/albums", func(ctx *gin.Context) {
		handlers.GetAlbums(db, ctx)
	})
	router.POST("/albums", func(ctx *gin.Context) {
		handlers.PostAlbums(db, ctx)
	})
	router.GET("/albums/:id", func(ctx *gin.Context) {
		handlers.GetAlbumsByID(db, ctx)
	})
	router.PUT("/albums/:id", func(ctx *gin.Context) {
		handlers.PutAlbumsByID(db, ctx)
	})
	router.DELETE("/albums/:id", func(ctx *gin.Context) {
		handlers.DeleteByID(db, ctx)
	})
	router.Run("localhost:8080")
}