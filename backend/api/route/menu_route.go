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

func NewMenuRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	mr := repository.NewMenuItemRepository(db)
	mt := &controller.MenuController{
		MenuItemUsecase: usecase.NewMenuUsecase(mr, timeout),
	}

	MenuGroup := group.Group("/menu")

	MenuGroup.POST("/", mt.Create)
	MenuGroup.GET("/", mt.Fetch)
	MenuGroup.GET("/price-range", mt.GetByPriceRange)
	MenuGroup.GET("/category/:categoryID", mt.GetByCategory)
	MenuGroup.GET("/:id", mt.GetByID)
	MenuGroup.PUT("/:id", mt.Update)
	MenuGroup.DELETE("/:id", mt.Delete)
}
