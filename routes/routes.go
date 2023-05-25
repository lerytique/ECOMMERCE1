package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(IncomingRoutes *gin.Engine) {
	IncomingRoutes.POST("/user/signup", controllers.SignUp())
	IncomingRoutes.POST("/user/login", controllers.Login())
	IncomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	IncomingRoutes.POST("/user/productview", controllers.SearchProduct())
	IncomingRoutes.POST("/user/search", controllers.SearchProductByQuery())

}
