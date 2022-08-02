package infra

import (
	"mc-payment-api/handler"
	"mc-payment-api/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app usecase.UsecaseInterface) {
	bookSrv := handler.NewHttpHandler(app)
	CategorySrv := handler.NewHttpHandler(app)
	AuthorSrv := handler.NewHttpHandler(app)
	book := r.Group("/book")
	{
		book.POST("", bookSrv.PostBook)
		book.GET("", bookSrv.GetAllBooks)
		book.PUT("/:id", bookSrv.UpdateBookByID)
		book.GET("/:id", bookSrv.GetBookByID)
		book.DELETE("/:id", bookSrv.DeleteBookByID)
	}

	category := r.Group("/category")
	{
		category.POST("", CategorySrv.PostCategory)
		category.GET("", CategorySrv.GetAllCategories)
		category.PUT("/:id", CategorySrv.UpdateCategoryByID)
		category.GET("/:id", CategorySrv.GetCategoryByID)
		category.DELETE("/:id", CategorySrv.DeleteCategoryByID)
	}

	author := r.Group("/author")
	{
		author.POST("", AuthorSrv.PostAuthor)
		author.GET("", AuthorSrv.GetAllAuthors)
		author.PUT("/:id", AuthorSrv.UpdateAuthorByID)
		author.GET("/:id", AuthorSrv.GetAuthorByID)
		author.DELETE("/:id", AuthorSrv.DeleteAuthorByID)
	}

}
