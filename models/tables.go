package models

import (
	"time"

	"github.com/beego/beego/orm"
)

type UserMasterTable struct {
	UserId      int       `orm:"auto"`
	FirstName   string    `orm:"size(255)"`
	LastName    string    `orm:"size(255)"`
	Email       string    `orm:"size(255)"`
	Password    string    `orm:"size(255)"`
	Mobile      string    `orm:"size(255)"`
	CreatedDate time.Time `orm:"type(datetime)"`
}

type CarsMasterTable struct {
	CarsId      int    `orm:"auto"`
	Name        string `orm:"size(255)"`
	Description string `orm:"size(255)"`
	CreatedBy   int
	UpdatedBy   int
	CarsImage   string    `orm:"null"`
	CreatedDate time.Time `orm:"type(date)"`
	UpdatedDate time.Time `orm:"type(date);default:null"`
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=Dev@123 host=localhost port=5432 dbname=golang_practice sslmode=disable")
	orm.RegisterModel(new(UserMasterTable), new(CarsMasterTable))
	orm.RunSyncdb("default", false, true)
}
