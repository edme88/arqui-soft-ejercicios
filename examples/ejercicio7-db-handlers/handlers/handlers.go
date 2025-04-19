package handlers

import (
	"net/http"
	"vinyl-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAlbums(db *gorm.DB, ctx *gin.Context){
	var albums []models.Album
	db.Find(&albums)
	ctx.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(db *gorm.DB, ctx *gin.Context){
	var newAlbum models.Album
	if err := ctx.BindJSON(&newAlbum); err != nil{
		return
	}
	db.Create(&newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumsByID(db *gorm.DB, ctx *gin.Context){
	id := ctx.Param("id")

	var album models.Album

	result := db.First(&album, "id = ?", id)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, album)
}

func PutAlbumsByID(db *gorm.DB, ctx *gin.Context){
	id := ctx.Param("id")
	var modifyAlbum models.Album
	var album models.Album

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

func DeleteByID(db *gorm.DB, ctx *gin.Context){
	id := ctx.Param("id")
	result := db.Delete(&models.Album{}, "id = ?", id)
	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "álbum eliminado"})
}