package main

import (
	"log"
	"vinyl-api/database"
	"vinyl-api/handler"

	"github.com/gin-gonic/gin"
)

func main(){
	db, err := database.InitDB()

	if err != nil{
		log.Fatal(err)
	}
	router := gin.Default()

	router.GET("/albums", func(ctx *gin.Context) {
		handler.GetAlbums(ctx, db)
	})

	router.POST("/albums", func(ctx *gin.Context) {
		handler.PostAlbums(ctx, db)
	})

	router.GET("/albums/:id", func(ctx *gin.Context) {
		handler.GetAlbumsByID(ctx, db)
	})
	router.PUT("/albums/:id", func(ctx *gin.Context) {
		handler.PutAlbumByID(ctx, db)
	})

	router.DELETE("/albums/:id", func(ctx *gin.Context) {
		handler.DeleteAlbumByID(ctx, db)
	})
	// router.GET("/albums", getAlbums)
	// router.POST("/albums", postAlbums)
	// router.GET("/albums/:id", getAlbumsByID)
	// router.PUT("/albums/:id", putAlbumByID)
	// router.DELETE("/albums/:id", deleteAlbumByID)
	router.Run("localhost:8080")
}