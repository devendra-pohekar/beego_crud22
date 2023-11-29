package routers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

// func init() {

// 	orm.RegisterDriver("postgres", orm.DRPostgres)
// 	orm.RegisterDataBase("default", "postgres", "user=postgres password=Dev@123 host=localhost port=5432 dbname=golang_practice sslmode=disable")
// 	// orm.RegisterModel(new(UserMasterTable), new(UserMasterTableTest), new(HomePagesSettingTable), new(CountriesTable), new(StatesTable), new(CitiesTable))
// 	orm.RunSyncdb("default", false, true)
// 	_, file, _, _ := runtime.Caller(0)
// 	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
// 	beego.TestBeegoInit(apppath)
// }

// func init() {
// 	beego.BeeApp = beego.NewApp()
// 	orm.RegisterDriver("postgres", orm.DRPostgres)
// 	orm.RegisterDataBase("default", "postgres", "user=postgres password=Dev@123 host=localhost port=5432 dbname=golang_practice sslmode=disable")
// 	// orm.RegisterModel(new(UserMasterTable), new(UserMasterTableTest), new(HomePagesSettingTable), new(CountriesTable), new(StatesTable), new(CitiesTable))
// 	orm.RunSyncdb("default", false, true)

// }

// func TestUserModelTesting(t *testing.T) {
// 	t.Run("register_user", func(t *testing.T) {
// 		userdata := requestStruct.InsertUser{
// 			FirstName: "testing user first name",
// 			LastName:  "testing User Last Name",
// 			Email:     "devendrapohekar.siliconithub@gmail.com",
// 			Password:  "dev@123",
// 		}
// 		result, err := models.RegisterUser(userdata)
// 		if err != nil {
// 			log.Print("inside the error section")
// 		}
// 		log.Print(result, "inside the result")

// 	})

// 	t.Run("Login_user", func(t *testing.T) {
// 		user_data := requestStruct.LoginUser{
// 			Email:    "devendrapohekar.siliconithub@gmail.com",
// 			Password: "Dev@123",
// 		}
// 		result, err := models.LoginUsers(user_data)
// 		if err != nil {
// 			log.Print(err)
// 		}
// 		log.Print(result, "fetch user data by useing email and password")
// 	})
// 	t.Run("update_user", func(t *testing.T) {

// 	})
// }

// func RequestTestingFunction(t *testing.T, method, endPontURL string, requestedData []byte, expectedStatusCode int) interface{} {
// 	req, err := http.NewRequest(method, endPontURL, bytes.NewBuffer(requestedData))
// 	w := httptest.NewRecorder()
// 	if err != nil {
// 		return w.Body
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return w.Body
// 	}

// 	if resp.StatusCode != expectedStatusCode {
// 		log.Print("Expected Status code", expectedStatusCode, "but return status code", resp.StatusCode)

// 	}

// 	return w.Body
// }

func RequestTestingFunction(t *testing.T, method, endPointURL string, requestData []byte, expectedStatusCode int) ([]byte, error) {
	req, err := http.NewRequest(method, endPointURL, bytes.NewBuffer(requestData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatusCode {
		return nil, fmt.Errorf("Expected status code %d but received %d", expectedStatusCode, resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func TestFetchAllSettings(t *testing.T) {
	t.Run("tesst All user Controller", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "http://localhost:8080/v1/homepage/fetch_settings", nil)

		client := &http.Client{}
		resp, _ := client.Do(req)

		body, _ := io.ReadAll(resp.Body)
		log.Print(string(body))
	})

}

func TestRegisterUser(t *testing.T) {
	endPoint := "http://localhost:8080/v1/user/add_user"
	var jsonStr = []byte(`{"first_name":"Devendra", "last_name":"Pohekar", "email":"devendrapohekar30@gmail.com", "mobile":"9516263597", "password":"1234567"}`)
	result, _ := RequestTestingFunction(t, "POST", endPoint, jsonStr, 200)
	log.Print(string(result))

}

func TestLoginUser(t *testing.T) {
	endPoint := "http://localhost:8080/v1/user/login"
	var jsonStr = []byte(`{"email":"devendrapohekar.siliconithub@gmail.com","password":"Devendra@123"}`)
	result, _ := RequestTestingFunction(t, "POST", endPoint, jsonStr, 200)
	log.Print(string(result))

}

func TestSendOTPonMails(t *testing.T) {
	endPoint := "http://localhost:8080/v1/user/send_otp"
	var jsonStr = []byte(`{"email":"devendrapohekar.siliconithub@gmail.com", "user_name":"Devendra Pohekar"}`)
	result, _ := RequestTestingFunction(t, "POST", endPoint, jsonStr, 200)
	log.Print(string(result))

}

func TestVerifyEmailOTP(t *testing.T) {
	endPoint := "http://localhost:8080/v1/user/verify_email"
	var jsonStr = []byte(`{"email":"devendrapohekar.siliconithub@gmail.com", "otp":"3D61qW1u"}`)
	result, _ := RequestTestingFunction(t, "POST", endPoint, jsonStr, 200)
	log.Print(string(result))

}

func TestForgotPasswordSendMail(t *testing.T) {
	endPoint := "http://localhost:8080/v1/user/send_otp_forgot"
	var jsonStr = []byte(`{"email":"devendrapohekar.siliconithub@gmail.com"}`)
	result, _ := RequestTestingFunction(t, "POST", endPoint, jsonStr, 200)
	log.Print(string(result))

}

func TestForgotPassword(t *testing.T) {
	endPoint := "http://localhost:8080/v1/user/verify_otp_forgot"
	var jsonStr = []byte(`{"otp":"Kmqqlvyb","new_password":"Devendra@123"}`)
	result, _ := RequestTestingFunction(t, "POST", endPoint, jsonStr, 200)
	log.Print(string(result))

}
