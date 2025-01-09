package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sammig6i/sydneys-sourdough-co/bootstrap"
	"github.com/sammig6i/sydneys-sourdough-co/database"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db database.Database, gin *gin.Engine) {
	// publicRouter := gin.Group("")
	// TODO maybe update to handle verification of allowed auth users in users table in DB
	// 1. Sign up
	// 2 Login

	// TODO fix context deadline exceeded when fetching menu items
	protectedRouter := gin.Group("")
	NewCategoryRouter(env, timeout, db, protectedRouter)
	NewMenuRouter(env, timeout, db, protectedRouter)
	NewSearchRouter(env, timeout, db, protectedRouter)
}
