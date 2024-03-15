package main

import (
	"assignment-2/controller"
	"assignment-2/lib"
	"assignment-2/model"
	"assignment-2/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := lib.InitDatabase()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Order{}, &model.Items{})
	if err != nil {
		panic(err)
	}
	orderRepository := repository.NewOrderRepository(db)
	orderController := controller.NewOrderController(orderRepository)

	itemRepository := repository.NewItemRepository(db)
	itemController := controller.NewItemController(itemRepository)

	ginEngine := gin.Default()
	ginEngine.GET("/orders", orderController.GetAll)
	ginEngine.POST("/orders", orderController.Create)
	ginEngine.DELETE("/order", orderController.Delete)
	ginEngine.PUT("/orders/:id", orderController.Update)

	ginEngine.GET("/items", itemController.GetAll)
	ginEngine.POST("/items", itemController.Create)
	ginEngine.DELETE("/item", itemController.Delete)
	ginEngine.PUT("items/:id", itemController.Update)

	err = ginEngine.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}