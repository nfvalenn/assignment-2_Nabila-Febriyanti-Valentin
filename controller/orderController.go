package controller

import (
	"assignment-2/model"
	"assignment-2/repository"
	"assignment-2/util"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)




type orderController struct {
	orderRepository repository.IOrderRepository
}

func NewOrderController(orderRepository repository.IOrderRepository) *orderController {
	return &orderController{
		orderRepository: orderRepository,
	}
}

func (o *orderController) Create(ctx *gin.Context)  {
	var newOrder model.Order

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	*newOrder.OrderedAT = time.Now()

	createdOrder, err := o.orderRepository.Create(newOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.CreateResponse(true, createdOrder, ""))
}

func (o *orderController) GetAll(ctx *gin.Context)  {
	orders,err := o.orderRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.CreateResponse(true, orders, ""))
}

func (o *orderController) Delete(ctx *gin.Context)  {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	fmt.Println(id)

	err := o.orderRepository.Delete(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.CreateResponse(true, nil, ""))
}

func (o *orderController) Update(ctx *gin.Context) {
	var updatedOrder model.Order

	err := ctx.ShouldBindJSON(&updatedOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	updatedOrder, err = o.orderRepository.Update(updatedOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, updatedOrder, ""))
}
