package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sammig6i/sydneys-sourdough-co/domain"
)

type SearchController struct {
	SearchUsecase domain.SearchUsecase
}

func (sc *SearchController) Search(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	results, err := sc.SearchUsecase.Search(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
