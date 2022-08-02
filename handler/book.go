package handler

import (
	"database/sql"
	"mc-payment-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (b *handlerHttpServer) PostBook(c *gin.Context) {
	var req models.BookRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	err = b.app.PostBook(req)
	if err == models.ErrCategoryIDNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	if err == models.ErrAuthorIDNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (b *handlerHttpServer) GetAllBooks(c *gin.Context) {
	res, err := b.app.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (b *handlerHttpServer) GetBookByID(c *gin.Context) {
	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	res, err := b.app.GetBookByID(idInt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (b *handlerHttpServer) UpdateBookByID(c *gin.Context) {
	var req models.BookRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	err = b.app.UpdateBookByID(idInt, req)
	if err == models.ErrIdNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (b *handlerHttpServer) DeleteBookByID(c *gin.Context) {
	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	err := b.app.DeleteBookByID(idInt)
	if err == models.ErrIdNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}
