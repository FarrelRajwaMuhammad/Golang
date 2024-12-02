package usercontroller

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FetchAndSaveUsers(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from JSONPlaceholder"})
		return
	}
	defer resp.Body.Close()

	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode JSON data"})
		return
	}

	for _, user := range users {
		if err := models.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data to database"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data successfully fetched from JSONPlaceholder and saved to database"})
}

func Index(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func Show(C *gin.Context) {
	userId := C.Param("id")
	var users []models.User

	if err := models.DB.Where("user_id = ?", userId).Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			C.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
			return
		}
		C.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	C.JSON(http.StatusOK, gin.H{"posts": users})
}

func Create(C *gin.Context) {
	var user models.User

	if err := C.ShouldBindJSON(&user); err != nil {
		C.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	models.DB.Create(&user)
	C.JSON(http.StatusCreated, user)
}

func Update(C *gin.Context) {
	var user models.User
	id := C.Param("id")

	if err := C.ShouldBindJSON(&user); err != nil {
		C.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		C.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed update data"})
		return
	}

	C.JSON(http.StatusOK, gin.H{"message": "Data Updated"})
}

func Delete(C *gin.Context) {
	var input struct {
		ID int64 `json:"id"`
	}

	if err := C.ShouldBindJSON(&input); err != nil {
		C.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	fmt.Println("Deleting user with ID:", input.ID)

	if err := models.DB.Where("id = ?", input.ID).Delete(&models.User{}).Error; err != nil {
		C.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user", "error": err.Error()})
		return
	}

	C.JSON(http.StatusOK, gin.H{"message": "Data has been Deleted"})
}
