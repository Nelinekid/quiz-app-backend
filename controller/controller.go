package controller

import (
	"net/http"

	"github.com/Gambi18/Quizzo/config"
	"github.com/Gambi18/Quizzo/models"
	"github.com/gin-gonic/gin"
)

// Functions for the user authentication
func GetUsers(c *gin.Context) {
	var users []models.User

	// Find users from the database
	if err := config.DB.Find(&users).Error; err != nil {
		// Handle error if the database query fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	// If no error, return the list of users
	c.JSON(http.StatusOK, &users)
}

func GetUser(c *gin.Context) {
	user := []models.User{}

	if err := config.DB.Where("id = ?", c.Param("id")).Find(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, &user)
}
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(201, &user)
}
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(200, &user)
}
func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}

// Functions for the score management
func GetScores(c *gin.Context) {
	scores := []models.Score{}
	if err := config.DB.Find(&scores).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve scores"})
		return
	}

	c.JSON(200, &scores)
}

func GetScore(c *gin.Context) {
	score := []models.Score{}
	if err := config.DB.Where("id = ?", c.Param("id")).Find(&score).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Score not found"})
		return
	}

	c.JSON(200, &score)
}

// func CreateScore(c *gin.Context) {
// 	var score models.Score
// 	c.BindJSON(&score)
// 	config.DB.Create(&score)
// 	c.JSON(200, &score)
// }

func UpdateScore(c *gin.Context) {
	var score models.Score
	id := c.Param("id")
	subject := c.Query("subject") //Get the subject from query parameters

	// Find the existing record by ID
	if err := config.DB.Where("id = ?", id).First(&score).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Score not found"})
		return
	}

	var input struct {
		Increment int `json:"increment"`
	}

	// Bind JSON body to input struct
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Increment the correct score based on the subject
	switch subject {
	case "subject1":
		score.Score1 += input.Increment
	case "subject2":
		score.Score2 += input.Increment
	case "subject3":
		score.Score3 += input.Increment
	case "subject4":
		score.Score4 += input.Increment
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject"})
		return
	}

	// Update the total score
	score.ScoreT = score.Score1 + score.Score2 + score.Score3 + score.Score4

	// Save the updated score
	if err := config.DB.Save(&score).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update score"})
		return
	}

	c.JSON(http.StatusOK, &score)
}
