package handlers

import (
	"net/http"
	"vinyl-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAlbums(ctx *gin.Context, db *gorm.DB){
	var albums []models.Album
	db.Find(&albums)
	ctx.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(ctx *gin.Context, db *gorm.DB){
	var newAlbum models.Album
	if err := ctx.BindJSON(&newAlbum); err != nil{
		return
	}
	db.Create(&newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumsByID(ctx *gin.Context, db *gorm.DB){
	id := ctx.Param("id")

	var album models.Album

	result := db.First(&album, "id = ?", id)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, album)
}

func PutAlbumsByID(ctx *gin.Context, db *gorm.DB){
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

func DeleteByID(ctx *gin.Context, db *gorm.DB){
	id := ctx.Param("id")
	result := db.Delete(&models.Album{}, "id = ?", id)
	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "álbum eliminado"})
}