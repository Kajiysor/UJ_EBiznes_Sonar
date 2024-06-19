package main

import (
	controllers "go_shop/controllers"
	models "go_shop/models"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	Echo *echo.Echo
}

func (s *Server) InitializeRoutings() {
	productController := &controllers.ProductController{DB: s.DB}
	categoryController := &controllers.CategoryController{DB: s.DB}
	cartController := &controllers.CartController{DB: s.DB}

	productsId := "/products/:id"
	categoriesId := "/categories/:id"
	cartProtuctId := "/cart/product/:productId"

	s.Echo.GET("/products", productController.GetAllProducts)
	s.Echo.GET(productsId, productController.GetProductById)
	s.Echo.GET("/products/category/:categoryId", productController.GetProductsByCategory)
	s.Echo.GET("/products/name/:name", productController.GetProductsByName)
	s.Echo.GET("/products/price/:price", productController.GetProductsByPrice)
	s.Echo.POST("/products", productController.CreateProduct)
	s.Echo.PUT(productsId, productController.UpdateProduct)
	s.Echo.DELETE(productsId, productController.DeleteProduct)

	s.Echo.GET("/categories", categoryController.GetAllCategories)
	s.Echo.GET(categoriesId, categoryController.GetCategoryById)
	s.Echo.GET("/categories/name/:name", categoryController.GetCategoryByName)
	s.Echo.POST("/categories", categoryController.CreateCategory)
	s.Echo.PUT(categoriesId, categoryController.UpdateCategory)
	s.Echo.DELETE(categoriesId, categoryController.DeleteCategory)

	s.Echo.GET("/cart", cartController.GetCart)
	s.Echo.GET(cartProtuctId, cartController.GetCartByProductId)
	s.Echo.POST("/cart", cartController.AddToCart)
	s.Echo.PUT(cartProtuctId, cartController.UpdateCart)
	s.Echo.DELETE(cartProtuctId, cartController.DeleteFromCart)
	s.Echo.DELETE("/cart", cartController.DeleteCart)
	s.Echo.GET("/cart/total", cartController.GetCartTotal)
	s.Echo.POST("/cart/confirmation", cartController.Confirmation)
}

func (s *Server) Initialize() error {
	err := s.DB.AutoMigrate(&models.Category{})
	if err != nil {
		return err
	}
	err = s.DB.AutoMigrate(&models.Product{})
	if err != nil {
		return err
	}
	err = s.DB.AutoMigrate(&models.CartItem{})
	if err != nil {
		return err
	}
	s.InitializeRoutings()
	return nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("store.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	server := &Server{
		DB: db,
		Echo: echo.New(),
	}

	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:3000", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	
	err = server.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	server.Echo.Logger.Fatal(server.Echo.Start(":8080"))
}