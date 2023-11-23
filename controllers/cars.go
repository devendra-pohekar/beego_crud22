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

// func (u *CarsControllers) RegisterCar() {
// 	var cars requestStruct.CarsInsert

// 	if err := u.ParseForm(&cars); err != nil {
// 		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Parsing Data Error")
// 		return
// 	}

// 	json.Unmarshal(u.Ctx.Input.RequestBody, &cars)

// 	file, fileHeader, err := u.GetFile("file")
// 	if err != nil {
// 		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "File Getting Error")
// 		return
// 	}

// 	// uploadDir := "./uploads/images"
// 	uploadDir := os.Getenv("uploadCarDirectory")

// 	filePath, err := helpers.UploadFile(file, fileHeader, uploadDir)
// 	if err != nil {
// 		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "File Uploading Error")
// 		return
// 	}

// 	token_data := helpers.GetTokenClaims(u.Ctx)
// 	user_id := token_data["User_id"]
// 	result, _ := models.RegisterCar(cars, user_id.(float64), filePath)
// 	if result != nil {
// 		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Car Register Successfully")
// 		return
// 	}
// 	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")

// }

// func (u *CarsControllers) UpdateCar() {
// 	var cars requestStruct.UpdateCars
// 	if err := u.ParseForm(&cars); err != nil {
// 		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Parsing Data Error")
// 		return
// 	}
// 	json.Unmarshal(u.Ctx.Input.RequestBody, &cars)
// 	file, fileHeader, err := u.GetFile("file")
// 	if err != nil {
// 		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "File Getting Error")
// 		return
// 	}

// 	// uploadDir := "./uploads/images"
// 	uploadDir := os.Getenv("uploadCarDirectory")
// 	filePath, err := helpers.UploadFile(file, fileHeader, uploadDir)
// 	if err != nil {
// 		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "File Uploading Error")
// 		return
// 	}

// 	token_data := helpers.GetTokenClaims(u.Ctx)
// 	user_id := token_data["User_id"]
// 	result, _ := models.UpdateCars(cars, filePath, user_id.(float64))
// 	if result != nil {
// 		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Car Updated Successfully")
// 		return
// 	}
// 	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")
// }

// func (u *CarsControllers) DeleteCar() {
// 	var cars requestStruct.DeleteCar
// 	json.Unmarshal(u.Ctx.Input.RequestBody, &cars)
// 	result := models.CarExistsDelete(cars)
// 	if result != 0 {
// 		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Car Deleted Successfully")
// 		return
// 	}
// 	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")
// }

func (u *CarsControllers) FetchCar() {
	var search requestStruct.SearchigData
	if err := u.ParseForm(&search); err != nil {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}
	json.Unmarshal(u.Ctx.Input.RequestBody, &search)

	result, totalPages, currentPage, _ := models.Searching(search.FilterString, 5, 1)
	if result != nil {
		u.Data["json"] = result
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, result, "Cars Found Successfully", totalPages, currentPage)
		return
	}
	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Not Found Data Please Try Again")
}
