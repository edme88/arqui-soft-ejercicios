package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Album struct {
	ID string `json:"id" gorm:"primarykey;autoIncrement"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Year uint `json:"year"`
	Price float64  `json:"price"`
}

var db *gorm.DB

func initDB(){
	dsn := "root:GoLang123#@tcp(localhost:3306)/vi" //datos de la BD para poder conectarnos usuario:Password@tcp(localhost:port)
	// //Abrir la conexión con la BD
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //Se utiliza el pacquete de forma indirecta
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Conexión a la BD exitosa") //Sin el fmt y usando log indica la fecha y hora
}

func getAlbums(ctx *gin.Context){
	var albums []Album
	db.Find(&albums)
	ctx.IndentedJSON(http.StatusOK, albums)
}

func postAlnums(ctx *gin.Context){
	var newAlbum Album
	if err := ctx.BindJSON(&newAlbum); err != nil{
		return
	}
	db.Create(&newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByID(ctx *gin.Context){
	id := ctx.Param("id")

	var album Album

	result := db.First(&album, "id = ?", id)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, album)
}

func putAlbumsByID(ctx *gin.Context){
	id := ctx.Param("id")
	var modifyAlbum Album
	var album Album

	if err := db.First(&album, "id = ?", id).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
		return
	}

	if err := ctx.BindJSON(&modifyAlbum); err != nil{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Datos incorrectos."})
		return
	}

	album.ID = id
	db.Save(&album)
	ctx.IndentedJSON(http.StatusOK, album)
}

func deleteByID(ctx *gin.Context){
	id := ctx.Param("id")
	result := db.Delete(&Album{}, "id = ?", id)
	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "álbum eliminado"})
}

func main(){

	initDB()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlnums)
	router.GET("/albums/:id", getAlbumsByID)
	router.PUT("/albums/:id", putAlbumsByID)
	router.DELETE("/albums/:id", deleteByID)
	router.Run("localhost:8080")
}