package controllers

import (
	scopes "go_shop/gormScopes"
	"go_shop/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CartController struct {
	DB *gorm.DB
}

func (cc CartController) GetCart (c echo.Context) error {
	var cart []models.CartItem
	cc.DB.Find(&cart)
	return c.JSON(http.StatusOK, cart)
}

func (cc CartController) GetCartByProductId (c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("productId"))
	var cart []models.CartItem
	cc.DB.Scopes(scopes.WithProductId(uint(productId))).Find(&cart)
	return c.JSON(http.StatusOK, cart)
}

func (cc CartController) AddToCart (c echo.Context) error {
	// check if product already exists in cart if not add it, else update quantity, productID is in the body of the request
    cartItem := new(models.CartItem)
    if err := c.Bind(cartItem); err != nil {
        return err
    }

    // if is present, increase quantity otherwise just create it
    existingCartItem := new(models.CartItem)
    if err := cc.DB.Scopes(scopes.WithProductId(cartItem.ProductID)).First(existingCartItem).Error; err != nil{
        cc.DB.Create(cartItem)
    } else {
        existingCartItem.Quantity++
        cc.DB.Save(existingCartItem)
    }

	return c.JSON(http.StatusCreated, cartItem)
}

func (cc CartController) UpdateCart (c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("productId"))
	var cart models.CartItem
	cc.DB.Scopes(scopes.WithProductId(uint(productId))).First(&cart)
	if err := c.Bind(&cart); err != nil {
		return err
	}
	cc.DB.Save(&cart)
	return c.JSON(http.StatusOK, cart)
}

func (cc CartController) DeleteFromCart (c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("productId"))
	var cart models.CartItem
	cc.DB.Scopes(scopes.WithProductId(uint(productId))).First(&cart)
	cc.DB.Delete(&cart)
	return c.JSON(http.StatusOK, cart)
}

func (cc CartController) DeleteCart (c echo.Context) error {
	var cart []models.CartItem
	cc.DB.Find(&cart)
	cc.DB.Delete(&cart)
	return c.JSON(http.StatusOK, cart)
}

func (cc CartController) GetCartTotal (c echo.Context) error {
	var cart []models.CartItem
	cc.DB.Find(&cart)
	var total float64
	for _, item := range cart {
		var product models.Product
		cc.DB.Scopes(scopes.WithId(item.ProductID)).First(&product)
		total += float64(item.Quantity) * product.Price
	}
	return c.JSON(http.StatusOK, total)
}

func (cc CartController) Confirmation (c echo.Context) error {
	var confirmation models.Confirmation
	if err := c.Bind(&confirmation); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, confirmation)
}