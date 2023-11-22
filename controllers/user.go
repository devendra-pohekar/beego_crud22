package controllers

import (
	"crudDemo/helpers"
	"crudDemo/models"
	requestStruct "crudDemo/requstStruct"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

type Users struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type JwtClaim struct {
	Email  string `json:"user_email"`
	UserID int    `json:"user_id"`
	jwt.StandardClaims
}

var secretKey = []byte("devendra_secretkey")

func (c *UserController) Login() {
	var user requestStruct.LoginUser
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		c.CustomAbort(http.StatusBadRequest, "Invalid JSON format")
		return
	}
	loginUserData := models.LoginUsers(user)
	tokenExpire := time.Now().Add(1 * time.Hour)
	claims := &JwtClaim{Email: loginUserData.Email, UserID: loginUserData.UserId, StandardClaims: jwt.StandardClaims{
		ExpiresAt: tokenExpire.Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, fmt.Sprintf("Error signing token: %s", err.Error()))
		return
	}

	data := map[string]interface{}{"User_Data": token.Claims, "Token": tokenString}

	c.Data["json"] = map[string]interface{}{"data": data}
	c.ServeJSON()
}

func (u *UserController) RegisterUser() {
	var user requestStruct.InsertUser
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	result, _ := models.RegisterUser(user)
	if result != nil {
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Register Successfully User Please Login Now")
	}
	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")
}

// func (u *UserController) LoginUser() {
// 	var user requestStruct.LoginUser
// 	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
// 	result := models.LoginUser(user)
// 	if result != 0 {
// 		u.Data["json"] = result
// 		// helpers.ApiSuccessResponse(u.Controller, "", "Login Successfully User")
// 	} else {
// 		// helpers.ApiFailedResponse(u.Controller, "Invalid Email and Password Please Try Again")
// 	}
// }
