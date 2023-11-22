package controllers

import (
	"crudDemo/helpers"
	"crudDemo/models"
	requestStruct "crudDemo/requstStruct"
	"encoding/json"

	"github.com/astaxie/beego"
)

type CarsControllers struct {
	beego.Controller
}

func (u *CarsControllers) RegisterCar() {
	var cars requestStruct.CarsInsert
	json.Unmarshal(u.Ctx.Input.RequestBody, &cars)

	token_data := helpers.GetTokenClaims(u.Ctx)
	user_id := token_data["User_id"]
	result, _ := models.RegisterCar(cars, user_id.(float64))
	if result != nil {
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Car Register Successfully")
		return
	}
	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")

}

func (u *CarsControllers) UpdateCar() {
	var cars requestStruct.UpdateCars
	json.Unmarshal(u.Ctx.Input.RequestBody, &cars)
	token_data := helpers.GetTokenClaims(u.Ctx)
	user_id := token_data["User_id"]
	result, _ := models.UpdateCars(cars, user_id.(float64))
	if result != nil {
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Car Updated Successfully")
		return
	}
	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")
}

func (u *CarsControllers) DeleteCar() {
	var cars requestStruct.DeleteCar
	json.Unmarshal(u.Ctx.Input.RequestBody, &cars)
	result := models.DeleteCar(cars)
	if result != 0 {
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Car Deleted Successfully")
		return
	}
	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")
}

func (u *CarsControllers) FetchCar() {
	result, _ := models.FetchCars()
	if result != nil {
		u.Data["json"] = result
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, result, "Cars Found Successfully")
		return
	}
	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Something Wrong Please Try Again")
}
