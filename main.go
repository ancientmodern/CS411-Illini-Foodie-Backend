package main

import (
	. "example/web-service-gin/api"
	. "example/web-service-gin/database"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := InitDB()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}

	router := gin.Default()

	router.Use(cors.Default())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/searchRestaurant", SearchRestaurant)
		v1.GET("/searchDish", SearchDish)
		v1.POST("/placeOrder", PlaceOrder)
		v1.DELETE("/deleteOrder", DeleteOrder)
		v1.POST("/updateComment", UpdateComment)
		v1.GET("/getComment", GetComment)
		v1.DELETE("/deleteComment", DeleteComment)
		v1.GET("/historyOrders", SearchOrderHistory)
		v1.POST("/updateDishPrice", UpdateDishPrice)
		v1.GET("/advancedCustomers", AdvancedCustomers)
		v1.GET("/advancedRestaurants", AdvancedRestaurants)
	}

	router.Run("0.0.0.0:80")
}
