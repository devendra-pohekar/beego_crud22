package routers

import (
	"bytes"
	"crudDemo/controllers"
	"crudDemo/models"
	requestStruct "crudDemo/requstStruct"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	"github.com/beego/beego/orm"
	"github.com/stretchr/testify/assert"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=Dev@123 host=localhost port=5432 dbname=golang_practice sslmode=disable")
	// orm.RegisterModel(new(UserMasterTable), new(UserMasterTableTest), new(HomePagesSettingTable), new(CountriesTable), new(StatesTable), new(CitiesTable))
	orm.RunSyncdb("default", false, true)

}

func TestUserModelTesting(t *testing.T) {
	t.Run("register_user", func(t *testing.T) {
		userdata := requestStruct.InsertUser{
			FirstName: "testing user first name",
			LastName:  "testing User Last Name",
			Email:     "devendrapohekar.siliconithub@gmail.com",
			Password:  "dev@123",
		}
		result, err := models.RegisterUser(userdata)
		if err != nil {
			log.Print("inside the error section")
		}
		log.Print(result, "inside the result")

	})

	t.Run("Login_user", func(t *testing.T) {
		user_data := requestStruct.LoginUser{
			Email:    "devendrapohekar.siliconithub@gmail.com",
			Password: "Dev@123",
		}
		result, err := models.LoginUsers(user_data)
		if err != nil {
			log.Print(err)
		}
		log.Print(result, "fetch user data by useing email and password")
	})
	t.Run("update_user", func(t *testing.T) {

	})
}

func TestRegisterUser(t *testing.T) {

	t.Run("user_registration", func(t *testing.T) {
		Ctrl := &controllers.UserController{}
		endPoint := "/v1/user/add_user"
		userdata := requestStruct.InsertUser{
			FirstName: "testing user first name",
			LastName:  "testing User Last Name",
			Email:     "devendra.siliconithub@gmail.com",
			Password:  "dev@123",
			Mobile:    "1234567890",
		}

		requestBody, _ := json.Marshal(userdata)

		req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(requestBody))

		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router := beego.NewControllerRegister()
		router.Add(endPoint, Ctrl)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code-204, "Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	})
}
