package controllers

import (
	"Quiz-3/database"
	"Quiz-3/repository"
	"Quiz-3/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetAllBook(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book structs.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		panic(err)
	}

	page := book.TotalPage
	if page <= 100 {
		book.Thickness = "tipis"
	} else if page >= 101 && page <= 200 {
		book.Thickness = "sedang"
	} else if page >= 201 {
		book.Thickness = "tebal"
	}

	book.Created_at = time.Now()
	book.Updated_at = time.Now()

	err = repository.InsertBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Book",
	})
}

func UpdateBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.Id = id
	book.Updated_at = time.Now()
	page := book.TotalPage
	if page <= 100 {
		book.Thickness = "tipis"
	} else if page >= 101 && page <= 200 {
		book.Thickness = "sedang"
	} else if page >= 201 {
		book.Thickness = "tebal"
	}

	err = repository.UpdateBook(database.DbConnection, book)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Book",
	})
}

func DeleteBook(c *gin.Context) {
	var book structs.Book
	id, err := strconv.Atoi(c.Param("id"))

	book.Id = id

	err = repository.DeleteBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Book",
	})
}
