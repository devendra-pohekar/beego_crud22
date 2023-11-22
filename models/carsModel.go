package models

import (
	requestStruct "crudDemo/requstStruct"
	"log"
	"time"

	"github.com/beego/beego/orm"
	_ "github.com/lib/pq"
)

func RegisterCar(c requestStruct.CarsInsert, user_id float64) (interface{}, error) {
	db := orm.NewOrm()
	res := CarsMasterTable{
		Name:        c.Name,
		CreatedBy:   int(user_id),
		UpdatedBy:   0,
		CarsImage:   "stt",
		UpdatedDate: time.Now(),
		Description: c.Description,
		CreatedDate: time.Now(),
	}

	_, err := db.Insert(&res)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func UpdateCars(c requestStruct.UpdateCars, user_id float64) (interface{}, error) {
	db := orm.NewOrm()

	cars := CarsMasterTable{CarsId: c.CarsId}
	if db.Read(&cars) == nil {
		cars.Name = c.Name
		cars.UpdatedBy = int(user_id)
		cars.UpdatedDate = time.Now()
		cars.Description = c.Description
		if num, err := db.Update(&cars); err == nil {
			return num, nil
		}

	}
	return nil, orm.ErrArgs

}

func DeleteCar(c requestStruct.DeleteCar) int {
	db := orm.NewOrm()
	cars := CarsMasterTable{CarsId: c.CarsId}
	if _, err := db.Delete(&cars); err == nil {
		return 1
	}
	return 0

}
func FetchCars() (interface{}, error) {
	db := orm.NewOrm()
	var cars []struct {
		Name        string    `json:"car_name"`
		Description string    `json:"description"`
		CreatedDate time.Time `json:"created_date"`
		UpdatedDate time.Time `json:"updated_date"`
		CreatedBy   string    `json:"created_by"`
	}
	_, err := db.Raw(`SELECT name, description, created_date, updated_date  FROM cars_master_table `).QueryRows(&cars)
	log.Print(cars)

	if err != nil {
		return nil, err
	}

	if len(cars) == 0 {
		return "Not Found Cars", nil
	}
	return cars, nil
}
