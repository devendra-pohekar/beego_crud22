package models

import (
	requestStruct "crudDemo/requstStruct"
	"time"

	"github.com/beego/beego/orm"
)

func RegisterUser(u requestStruct.InsertUser) (interface{}, error) {
	db := orm.NewOrm()
	res := UserMasterTable{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		Password:    u.Password,
		Mobile:      u.Mobile,
		CreatedDate: time.Now(),
	}
	_, err := db.Insert(&res)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func LoginUsers(u requestStruct.LoginUser) UserMasterTable {
	db := orm.NewOrm()

	res := UserMasterTable{
		Email:    u.Email,
		Password: u.Password,
	}
	result := db.Read(&res, "Email", "Password")
	if result != nil {
		panic("result")
	}
	return res
}
