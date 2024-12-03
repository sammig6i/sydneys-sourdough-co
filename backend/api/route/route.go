package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sammig6i/sydneys-sourdough-co/bootstrap"
	"github.com/sammig6i/sydneys-sourdough-co/database"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db database.Database, gin *gin.Engine) {
	// publicRouter := gin.Group("")
	// TODO Supabase Auth for Sign up and Login

	protectedRouter := gin.Group("/protected")
	NewCategoryRouter(env, timeout, db, protectedRouter)
	NewMenuRouter(env, timeout, db, protectedRouter)
	// NewSearchRouter(env, timeout, db, protectedRouter)
}
