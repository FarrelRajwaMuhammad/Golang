package usercontroller

import (
	"backend/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FetchAndSaveUsers(c *gin.Context) {
	// Mengambil data dari JSONPlaceholder
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from JSONPlaceholder"})
		return
	}
	defer resp.Body.Close()

	// Parsing JSON response dari JSONPlaceholder
	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode JSON data"})
		return
	}

	// Menyimpan data ke database MySQL
	for _, user := range users {
		if err := models.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data to database"})
			return
		}
	}

	// Mengembalikan response bahwa data berhasil disimpan
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
		C.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := models.DB.Create(&user).Error; err != nil {
		C.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	C.JSON(http.StatusCreated, user)
}

func Update(C *gin.Context) {
	id := C.Param("id")
	var input struct {
		UserID uint   `json:"userId"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	if err := C.ShouldBindJSON(&input); err != nil {
		C.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := models.DB.Model(&models.User{}).
		Where("id = ?", id).
		Updates(models.User{
			UserID: input.UserID,
			Title:  input.Title,
			Body:   input.Body,
		}).Error; err != nil {
		C.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	if models.DB.RowsAffected == 0 {
		C.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	C.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func Delete(C *gin.Context) {
	id := C.Param("id")

	if err := models.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		C.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	if models.DB.RowsAffected == 0 {
		C.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	C.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
