package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lerytique/ECOMMERCE1/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



	type Application struct {
		prodCollection *mongo.Collection
		userCollection *mongo.Collection
	}

	func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
		return &Application{
			prodCollection: prodCollection,
			userCollection: userCollection,
		}
	}



func (app *Application) AddToCart() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		productQueryID :=c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")
			_ = c.AbortWithError (http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		userQueryID := c.Query ("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_   = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return 
	}

	productID, err := primitive.ObjectIDFromHex(productQueryID)

	if err!= nil{
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError)
		return 
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = database.AddProductToCart (ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err!= nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully added to the cart ")
	}
}
func (app *Application) RemoveItem() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			productQueryID :=c.Query("id")
			if productQueryID == "" {
				log.Println("product id is empty")
				_ = c.AbortWithError (http.StatusBadRequest, errors.New("product id is empty"))
				return
			}
			userQueryID := c.Query ("userID")
			if userQueryID == "" {
				log.Println("user id is empty")
				_   = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
				return 
		}
	
		productID, err := primitive.ObjectIDFromHex(productQueryID)
	
		if err!= nil{
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError)
			return 
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	
		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
	
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(200, "Successfully removed item from cart")
	}

}


func GetItemFRomCard() gin.HandlerFunc{

}

func (app *Application) BuyFromCard() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		userQueryID := c.Query == "" {
			log.Panicln("user id is empty")
			_ = c.AbortWithError (http.StatusBadRequest, errors.New("UserID is empty"))
		}
		var ctx, cancel = context.WithTimeout(ccontext.Background(), 100*time.Second)

		defer cancel()

		err  := database.BuyItemFromCart(ctx, app.userCollection, app.prodCollection, userQueryID)
		if err!= nil{
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

			c.IndentedJSON ("Successfully placed the order")
		}

	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productQueryID :=c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")
			_ = c.AbortWithError (http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		userQueryID := c.Query ("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_   = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return 
	}

	productID, err := primitive.ObjectIDFromHex(productQueryID)

	if err!= nil{
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError)
		return 
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = database.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

	if err != nil {
		c.IndentedJSON (http.StatusInternalServerError, err)
		return
	}
		c.IndentedJSON (200, "Succsessfully placed the order")
		
}
}