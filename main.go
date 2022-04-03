package main

import(
	"gorm.io/gorm"
	"sample/database/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"gorm.io/driver/sqlite"
)

func main() {
	e := echo.New()

	db, err := gorm.Open(sqlite.Open("skelp.db"))

	if err != nil {
		panic("Error")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})

	e.GET("/users", func(c echo.Context) error {
		var users[] models.User
		db.Find(&users)

		return c.JSON(http.StatusOK, users)
	})

	e.GET("/products", func(c echo.Context) error {
		var products[] models.Product
		db.Find(&products)

		return c.JSON(http.StatusOK, products)
	})

	e.GET("/categories", func(c echo.Context) error {
		var categories[] models.Category
		db.Find(&categories)

		return c.JSON(http.StatusOK, categories)
	})

	e.Start(":8080")
}
