package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhil6392/go-backend/config"
	"github.com/nikhil6392/go-backend/models"
)

func SignUp(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return 
	}

	var exists bool
	err := config.DB.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", user.Email).Scan(&exists)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return 
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exits"})
		return 
	}

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, created_at`
	err = config.DB.QueryRow(context.Background(), query, user.Name, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": 			user.ID,
		"name":			user.Name,
		"email":		user.Email,
		"created_at":	user.CreatedAt,	
})
}