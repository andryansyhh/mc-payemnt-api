package handler

import (
	"database/sql"
	"mc-payment-api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (b *handlerHttpServer) PostAuthor(c *gin.Context) {
	var req models.AuthorRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	_, err = time.Parse("02-01-2006", req.TangalLahir)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	err = b.app.PostAuthor(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (b *handlerHttpServer) GetAllAuthors(c *gin.Context) {
	res, err := b.app.GetAllAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (b *handlerHttpServer) GetAuthorByID(c *gin.Context) {
	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	res, err := b.app.GetAuthorByID(idInt)
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

func (b *handlerHttpServer) UpdateAuthorByID(c *gin.Context) {
	var req models.AuthorRequest
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

	err = b.app.UpdateAuthorByID(idInt, req)
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

func (b *handlerHttpServer) DeleteAuthorByID(c *gin.Context) {
	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	err := b.app.DeleteAuthorByID(idInt)
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
