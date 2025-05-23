package main

import (
	"log"
	"time"
	"vinyl-api/database"
	handlers "vinyl-api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main(){

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos: ", err)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET","POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
  	}))

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