package controller

import (
	"assignment-2/model"
	"assignment-2/repository"
	"assignment-2/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



type itemController struct {
	itemRepository repository.IItemRepository
}

func NewItemController(itemRepository repository.IItemRepository) *itemController  {
	return &itemController {
		itemRepository: itemRepository,
	}
}

func (i *itemController) Create(ctx *gin.Context) {
	var newItem model.Items

	err := ctx.ShouldBindJSON(&newItem)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error: err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	createItem, err := i.itemRepository.Create(newItem)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.CreateResponse(true, createItem, ""))
}

func (i *itemController) GetAll(ctx *gin.Context)  {
	items, err := i.itemRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.CreateResponse(true, items, ""))
}

func (i *itemController) Delete(ctx *gin.Context)  {
	idString := ctx.Param("id")
	idInt, _ := strconv.Atoi(idString)
	fmt.Println(idInt)

	err := i.itemRepository.Delete(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, nil, ""))
}

func (i *itemController) Update(ctx *gin.Context) {
	var updatedItem model.Items

	err := ctx.ShouldBindJSON(&updatedItem)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	updatedItem, err = i.itemRepository.Update(updatedItem)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, updatedItem, ""))
}
