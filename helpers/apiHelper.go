package helpers

import (
	"github.com/astaxie/beego/context"

	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type ResponseSuccess struct {
	ResStatus int         `json:"status"`
	Message   string      `json:"message"`
	Result    interface{} `json:"data"`
}

type ResponseFailed struct {
	ResStatus int    `json:"status"`
	Message   string `json:"message"`
}

func ApiSuccessResponse(w http.ResponseWriter, result interface{}, message string) {
	if result == "" {
		result = map[int]int{}
	}
	response := ResponseSuccess{
		Message:   message,
		ResStatus: 1,
		Result:    result,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func ApiFailedResponse(w http.ResponseWriter, message string) {
	response := ResponseFailed{
		Message:   message,
		ResStatus: 0,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetTokenClaims(c *context.Context) map[string]interface{} {
	token_claims := c.Input.GetData("LoginUserData")
	user_id := token_claims.(jwt.MapClaims)["user_id"]
	user_email := token_claims.(jwt.MapClaims)["user_email"]
	response := map[string]interface{}{"User_id": user_id, "User_Email": user_email}
	return response
}
