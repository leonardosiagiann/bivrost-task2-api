package controllers

import (
	"bivrost-task2/database"
	"bivrost-task2/models"
	"log"
	"net/http"

	bv "github.com/koinworks/asgard-bivrost/service"
)

func CreateItem(ctx *bv.Context) bv.Result {
	db := database.GetDB()

	var item models.Item
	err := ctx.BodyJSONBind(&item)
	if err != nil {
		log.Println(err)
		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
			Message: map[string]string{
				"err": "error when binding data",
			},
		})
	}

	err = db.Debug().Create(&item).Error
	if err != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
			Message: map[string]string{
				"err": "error when creating new items",
			},
		})
	}

	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Data: item,
	})
}

func GetItems(ctx *bv.Context) bv.Result {
	db := database.GetDB()

	var items []models.Item

	err := db.Debug().Find(&items).Error
	if err != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
			Message: map[string]string{
				"err": "error, cant find item list on database",
			},
		})
	}

	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Data: items,
	})

}

func PingHandler(ctx *bv.Context) bv.Result {

	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "Welcome to Ping API",
			"id": "Selamat datang di Ping API",
		},
	})

}
