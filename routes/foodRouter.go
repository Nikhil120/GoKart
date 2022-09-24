package routes

import (
	"GoCart/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFoods())
	incomingRoutes.POST("/foods", controllers.CreateFoods())
	incomingRoutes.PATCH("/foods/:food_id", controllers.UpdateFood())
}
