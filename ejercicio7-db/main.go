package main

import (
	"log"
	"vinyl-api/database"
	handlers "vinyl-api/handler"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func wrap(f func(ctx *gin.Context, db *gorm.DB), db db *gorm.DB)

func main(){

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos: ", err)
	}

	router := gin.Default()

	router.GET("/albums", func(ctx *gin.Context) {
		handlers.GetAlbums(ctx, db)
	})

	router.POST("/albums", func(ctx *gin.Context) {
		handlers.PostAlbums(ctx, db)
	})

	router.GET("/albums/:id", func(ctx *gin.Context) {
		handlers.GetAlbumsByID(ctx, db)
	})

	router.PUT("/albums/:id", func(ctx *gin.Context) {
		handlers.PutAlbumsByID(ctx, db)
	})

	router.DELETE("/albums/:id", func(ctx *gin.Context) {
		handlers.DeleteByID(ctx, db)
	})

	// router.POST("/albums", PostAlnums)
	// router.GET("/albums/:id", GetAlbumsByID)
	// router.PUT("/albums/:id", PutAlbumsByID)
	// router.DELETE("/albums/:id", DeleteByID)
	router.Run("localhost:8080")
}