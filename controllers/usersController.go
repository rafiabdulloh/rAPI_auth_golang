package usersControllers

import (
	"auth_with_token/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	auth "auth_with_token/auth"

	log "github.com/sirupsen/logrus"
)

func GetAll(ctx *gin.Context, db *gorm.DB) {
	var data []models.Users

	db.Find(&data)

	var count int64
	db.Model(&models.Users{}).Count(&count)
	if count == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": data,
	})

}

func GetById(c *gin.Context, db *gorm.DB) {
	var data models.Users
	idProduct := c.Param("id")
	// id, _ := strconv.ParseInt(idProduct, 10, 64)

	// dataId := models.Users{Id: id}
	db.Find(&data, db.Where("id=?", idProduct))

	var count int64
	db.Model(&models.Users{}).Count(&count)
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})

}

func Post(ctx *gin.Context, db *gorm.DB) {
	var data models.Users

	err := ctx.Bind(&data)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}
	if err == nil {
		data.Password = auth.DoHash(data.Password)

		db.Create(&data)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success create to database",
			"data":    data,
		})
	}
}

func Update(c *gin.Context, db *gorm.DB) {

	id, _ := strconv.Atoi(c.Param("id")) // convert id because by default c.Param function is return string

	var data models.Users

	if err := db.First(&data, id).Error; err != nil { //find just first row
		c.JSON(http.StatusBadRequest, gin.H{"error": "data not found"})
		return
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.Password = auth.DoHash(data.Password)

	db.Save(&data) // Save updates value in database. If value doesn't contain a matching primary key, value is inserted.
	fmt.Println(data.Password)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func Delete(c *gin.Context, db *gorm.DB) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data models.Users
	if err := db.First(&data, id).Error; err != nil {
		log.Error("data with this id not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	db.Delete(&data)
	c.JSON(http.StatusOK, gin.H{"data": "User deleted successfully"})
}
