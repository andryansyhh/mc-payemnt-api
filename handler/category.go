package handler

import (
	"database/sql"
	"mc-payment-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (b *handlerHttpServer) PostCategory(c *gin.Context) {
	var req models.CategoryRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	err = b.app.PostCategory(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (b *handlerHttpServer) GetAllCategories(c *gin.Context) {
	res, err := b.app.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (b *handlerHttpServer) GetCategoryByID(c *gin.Context) {
	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	res, err := b.app.GetCategoryByID(idInt)
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

func (b *handlerHttpServer) UpdateCategoryByID(c *gin.Context) {
	var req models.CategoryRequest
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

	err = b.app.UpdateCategoryByID(idInt, req)
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

func (b *handlerHttpServer) DeleteCategoryByID(c *gin.Context) {
	ID := c.Param("id")
	idInt, _ := strconv.Atoi(ID)

	err := b.app.DeleteCategoryByID(idInt)
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
