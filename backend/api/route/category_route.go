package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sammig6i/sydneys-sourdough-co/api/controller"
	"github.com/sammig6i/sydneys-sourdough-co/bootstrap"
	"github.com/sammig6i/sydneys-sourdough-co/database"
	"github.com/sammig6i/sydneys-sourdough-co/repository"
	"github.com/sammig6i/sydneys-sourdough-co/usecase"
)

func NewCategoryRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	cr := repository.NewCategoryRepository(db)
	ct := &controller.CategoryController{
		CategoryUsecase: usecase.NewCategoryUsecase(cr, timeout),
	}

	categoryGroup := group.Group("/categories")

	categoryGroup.POST("/", ct.Create)
	categoryGroup.GET("/", ct.Fetch)
	categoryGroup.GET("/:id", ct.GetByID)
	categoryGroup.PUT("/:id", ct.Update)
	categoryGroup.DELETE("/:id", ct.Delete)
}
