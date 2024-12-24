package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sammig6i/sydneys-sourdough-co/bootstrap"
	"github.com/sammig6i/sydneys-sourdough-co/database"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db database.Database, gin *gin.Engine) {
	// publicRouter := gin.Group("")
	// TODO Add Supabase Auth:
	// 1. Sign up
	// 2 Login

	protectedRouter := gin.Group("/protected")
	// TODO Add middleware for protected routes
	NewCategoryRouter(env, timeout, db, protectedRouter)
	NewMenuRouter(env, timeout, db, protectedRouter)
	NewSearchRouter(env, timeout, db, protectedRouter)
}
