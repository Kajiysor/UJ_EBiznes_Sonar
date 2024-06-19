package controllers

import (
	scopes "go_shop/gormScopes"
	"go_shop/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func (pc ProductController) GetAllProducts(c echo.Context) error {
	var products []models.Product
	pc.DB.Find(&products)
	return c.JSON(http.StatusOK, products)
}

func (pc ProductController) GetProductById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	pc.DB.Scopes(scopes.WithId(uint(id))).First(&product)
	return c.JSON(http.StatusOK, product)
}

func (pc ProductController) GetProductsByCategory(c echo.Context) error {
	categoryId, _ := strconv.Atoi(c.Param("categoryId"))
	var products []models.Product
	pc.DB.Scopes(scopes.WithCategoryId(uint(categoryId))).Find(&products)
	return c.JSON(http.StatusOK, products)
}

func (pc ProductController) GetProductsByName(c echo.Context) error {
	name := c.Param("name")
	var products []models.Product
	pc.DB.Scopes(scopes.WithName(name)).Find(&products)
	return c.JSON(http.StatusOK, products)
}

func (pc ProductController) GetProductsByPrice(c echo.Context) error {
	price, _ := strconv.ParseFloat(c.Param("price"), 64)
	var products []models.Product
	pc.DB.Scopes(scopes.WithPrice(price)).Find(&products)
	return c.JSON(http.StatusOK, products)
}

func (pc ProductController) CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	pc.DB.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

func (pc ProductController) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	pc.DB.Scopes(scopes.WithId(uint(id))).First(&product)
	if err := c.Bind(&product); err != nil {
		return err
	}
	pc.DB.Save(&product)
	return c.JSON(http.StatusOK, product)
}

func (pc ProductController) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	pc.DB.Scopes(scopes.WithId(uint(id))).First(&product)
	if product.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}
	pc.DB.Delete(&product)
	return c.NoContent(http.StatusNoContent)
}