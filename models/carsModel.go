package models

import (
	"crudDemo/helpers"
	requestStruct "crudDemo/requstStruct"
	"time"

	// "github.com/astaxie/beego/orm"

	"github.com/beego/beego/orm"
	_ "github.com/lib/pq"
)

func RegisterCar(c requestStruct.CarsInsert, user_id float64, file_path string) (interface{}, error) {
	db := orm.NewOrm()
	res := CarsMasterTable{
		Name:        c.Name,
		CreatedBy:   int(user_id),
		UpdatedBy:   0,
		CarsImage:   file_path,
		UpdatedDate: time.Now(),
		CarTypes:    CarsTypes(c.CarType),
		Description: c.Description,
		CreatedDate: time.Now(),
	}

	_, err := db.Insert(&res)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func UpdateCars(c requestStruct.UpdateCars, file_path string, user_id float64) (interface{}, error) {
	db := orm.NewOrm()
	carID := c.CarsId
	car, err := FetchCarByID(carID)
	if err != nil {
		return nil, err
	}
	file_name, file_directory := helpers.SplitFilePath(car)
	helpers.RemoveFile(file_name, file_directory)

	cars := CarsMasterTable{CarsId: c.CarsId}

	if db.Read(&cars) == nil {
		cars.Name = c.Name
		cars.UpdatedBy = int(user_id)
		cars.UpdatedDate = time.Now()
		cars.Description = c.Description
		cars.CarsImage = file_path
		if num, err := db.Update(&cars); err == nil {
			return num, nil
		}

	}
	return nil, orm.ErrArgs

}

func DeleteCar(car_id int) int {
	db := orm.NewOrm()
	cars := CarsMasterTable{CarsId: car_id}
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
		CarTypes    string    `json:"car_types"`
		CreatedDate time.Time `json:"created_date"`
		UpdatedDate time.Time `json:"updated_date"`
		CarsImage   string    `json:"cars_image"`
		CreatedBy   string    `json:"created_by"`
	}
	_, err := db.Raw(`SELECT name, description, car_types,created_date, updated_date ,cars_image,created_by FROM cars_master_table `).QueryRows(&cars)

	if err != nil {
		return nil, err
	}

	if len(cars) == 0 {
		return "Not Found Cars", nil
	}
	return cars, nil
}

func CarExistsDelete(u requestStruct.DeleteCar) int {
	carID := u.CarsId
	car, err := FetchCarByID(carID)
	if err != nil {
		return 0
	}

	file_name, file_directory := helpers.SplitFilePath(car)
	helpers.RemoveFile(file_name, file_directory)
	DeleteCar(carID)
	return 1

}

func FetchCarByID(carID int) (string, error) {
	db := orm.NewOrm()
	var car CarsMasterTable
	err := db.Raw(`SELECT  cars_image FROM cars_master_table WHERE cars_id = ?`, carID).QueryRow(&car)
	if err != nil {
		return "errror", err
	}
	return car.CarsImage, nil
}

func Searching(searchString string, curentPage, pageSize int) (interface{}, int, int, error) {
	db := orm.NewOrm()
	var cars []struct {
		Name        string    `json:"car_name"`
		Description string    `json:"description"`
		CarTypes    string    `json:"car_types"`
		CreatedDate time.Time `json:"created_date"`
		UpdatedDate time.Time `json:"updated_date"`
		CarsImage   string    `json:"cars_image"`
		CreatedBy   string    `json:"created_by"`
	}

	query := "SELECT name, description, car_types, created_date, updated_date, cars_image, created_by FROM cars_master_table"

	if searchString != "" {
		query += " WHERE UPPER(name) LIKE UPPER(?) OR UPPER(description) LIKE UPPER(?) OR UPPER(car_types) LIKE UPPER(?)"
		searchString = "%" + searchString + "%"
	}

	offset := (curentPage - 1) * pageSize

	query += " LIMIT ? OFFSET ?"

	_, err := db.Raw(query, searchString, searchString, searchString, pageSize, offset).QueryRows(&cars)
	if err != nil {
		return nil, 0, 0, err
	}

	totalCountQuery := "SELECT COUNT(*) FROM cars_master_table"
	if searchString != "" {
		totalCountQuery += " WHERE UPPER(name) LIKE UPPER(?) OR UPPER(description) LIKE UPPER(?) OR UPPER(car_types) LIKE UPPER(?)"
	}

	var totalCount int
	err = db.Raw(totalCountQuery, searchString, searchString, searchString).QueryRow(&totalCount)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := (totalCount + (pageSize - 1)) / pageSize

	return cars, totalPages, curentPage, nil
}
