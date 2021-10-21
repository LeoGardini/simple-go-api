package controllers

import (
	"strconv"

	"github.com/LeoGardini/simple-go-api/db"
	"github.com/LeoGardini/simple-go-api/models"
	"github.com/gin-gonic/gin"
)

func ListBooks(c *gin.Context) {
	db := db.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot list books (" + err.Error() + ")"})
		return
	}

	c.JSON(200, books)
}

func ShowBook(c *gin.Context) {
	id := c.Param("id")
	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "ID has to be integer"})
		return
	}

	db := db.GetDatabase()

	var book models.Book
	err = db.First(&book, newId).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find book (" + err.Error() + ")"})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := db.GetDatabase()

	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot find JSON (" + err.Error() + ")"})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		return
	}

	c.JSON(201, book)
}

func UpdateBook(c *gin.Context) {
	db := db.GetDatabase()

	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find JSON (" + err.Error() + ")"})
		return
	}

	err = db.Save(&book).Error
	if err != nil {
		return
	}

	c.JSON(201, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID has to be integer"})
		return
	}

	db := db.GetDatabase()

	var book models.Book
	err = db.Find(&book).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find book (" + err.Error() + ")"})
		return
	}

	err = db.Delete(book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot delete book (" + err.Error() + ")"})
		return
	}

	c.JSON(200, book)
}
