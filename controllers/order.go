package controllers

import (
	"bivrost-task2/database"
	"bivrost-task2/models"
	"log"
	"net/http"

	bv "github.com/koinworks/asgard-bivrost/service"
)

func CreateOrder(ctx *bv.Context) bv.Result {
	db := database.GetDB()

	var order models.Order
	err := ctx.BodyJSONBind(&order)
	if err != nil {
		log.Println(err)
		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
			Message: map[string]string{
				"err": "error when binding data",
			},
		})
	}

	err = db.Debug().Create(&order).Error
	if err != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
			Message: map[string]string{
				"err": "error when creating new items",
			},
		})
	}

	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Data: order,
	})
}

func GetOrders(ctx *bv.Context) bv.Result {
	db := database.GetDB()

	var orders []models.Order

	err := db.Debug().Preload("Item").Find(&orders).Error
	if err != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
			Message: map[string]string{
				"err": "error, cant find item list on database",
			},
		})
	}

	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Data: orders,
	})

}
