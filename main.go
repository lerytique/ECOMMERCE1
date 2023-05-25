package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lerytique/ECOMMERCE1/controllers"
	"github.com/lerytique/ECOMMERCE1/database"
	"github.com/lerytique/ECOMMERCE1/middleware"
	"github.com/lerytique/ECOMMERCE1/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products")(database.Userdata(database.Client, "Users")))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentification())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
