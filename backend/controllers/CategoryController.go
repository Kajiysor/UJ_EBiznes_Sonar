package controllers

import (
	scopes "go_shop/gormScopes"
	"go_shop/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

func (cc CategoryController) GetAllCategories(c echo.Context) error {
	var categories []models.Category
	cc.DB.Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func (cc CategoryController) GetCategoryById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category
	cc.DB.Scopes(scopes.WithId(uint(id))).First(&category)
	return c.JSON(http.StatusOK, category)
}

func (cc CategoryController) GetCategoryByName(c echo.Context) error {
	name := c.Param("name")
	var category models.Category
	cc.DB.Scopes(scopes.WithName(name)).First(&category)
	return c.JSON(http.StatusOK, category)
}

func (cc CategoryController) CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return err
	}
	cc.DB.Create(&category)
	return c.JSON(http.StatusCreated, category)
}

func (cc CategoryController) UpdateCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category
	cc.DB.Scopes(scopes.WithId(uint(id))).First(&category)
	if err := c.Bind(&category); err != nil {
		return err
	}
	cc.DB.Save(&category)
	return c.JSON(http.StatusOK, category)
}

func (cc CategoryController) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category
	cc.DB.Scopes(scopes.WithId(uint(id))).First(&category)
	if category.ID == 0 {
		return c.JSON(http.StatusNotFound, "Category not found")
	}
	cc.DB.Delete(&category)
	return c.NoContent(http.StatusNoContent)
}