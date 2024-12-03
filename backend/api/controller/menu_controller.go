package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sammig6i/sydneys-sourdough-co/domain"
)

type MenuController struct {
	MenuItemUsecase domain.MenuItemUsecase
}

func (mc *MenuController) Create(c *gin.Context) {
	var menuIem domain.MenuItem

	if err := c.ShouldBindJSON(&menuIem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := mc.MenuItemUsecase.Create(c, &menuIem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, menuIem)
}

func (mc *MenuController) Fetch(c *gin.Context) {
	menuItems, err := mc.MenuItemUsecase.Fetch(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menuItems)
}

func (mc *MenuController) Update(c *gin.Context) {
	var menuItem domain.MenuItem

	if err := c.ShouldBindJSON(&menuItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := mc.MenuItemUsecase.Update(c.Request.Context(), &menuItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menuItem)
}

func (mc *MenuController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = mc.MenuItemUsecase.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu item deleted successfully."})
}

func (mc *MenuController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	menuItem, err := mc.MenuItemUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menuItem)
}

func (mc *MenuController) GetByCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	menuItems, err := mc.MenuItemUsecase.GetByCategory(c.Request.Context(), categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categoryID": categoryID,
		"items":      menuItems,
	})
}

func (mc *MenuController) GetByPriceRange(c *gin.Context) {
	minPriceStr := c.Query("minPrice")
	maxPriceStr := c.Query("maxPrice")

	if minPriceStr == "" || maxPriceStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "minPrice and maxPrice are required."})
		return
	}

	minPrice, err := strconv.ParseFloat(minPriceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minPrice value."})
		return
	}

	maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maxPrice value."})
		return
	}

	menuItems, err := mc.MenuItemUsecase.GetByPriceRange(c.Request.Context(), minPrice, maxPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menuItems)
}
