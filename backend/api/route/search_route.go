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

func NewSearchRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	sr := repository.NewSearchRepository(db)
	st := &controller.SearchController{
		SearchUsecase: usecase.NewSearchUsecase(sr, timeout),
	}

	searchGroup := group.Group("/search")

	searchGroup.GET("/", st.Search)
}
