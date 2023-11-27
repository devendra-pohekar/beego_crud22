package models

import (
	"time"

	"github.com/beego/beego/orm"
	// "github.com/astaxie/beego/orm"
)

type UserMasterTable struct {
	UserId      int    `orm:"auto"`
	FirstName   string `orm:"size(255)"`
	LastName    string `orm:"size(255)"`
	Email       string `orm:"size(255)"`
	Password    string `orm:"size(255)"`
	Mobile      string `orm:"size(255)"`
	IsVerified  int
	OtpCode     string    `orm:size(255)`
	CreatedDate time.Time `orm:"type(datetime)"`
}
type CarsTypes string

const (
	Suv       CarsTypes = "suv"
	Sedan     CarsTypes = "sedan"
	HatchBack CarsTypes = "hatchback"
)

type CarsMasterTable struct {
	CarsId      int    `orm:"auto"`
	Name        string `orm:"size(255)"`
	Description string `orm:"size(255)"`
	CreatedBy   int
	UpdatedBy   int
	CarsImage   string
	CarTypes    CarsTypes `orm:column(car_types)`
	CreatedDate time.Time `orm:"type(date)"`
	UpdatedDate time.Time `orm:"type(date)"`
}

type HomePagesSettingTable struct {
	PageSettingId int `orm:"auto"`
	Section       string
	DataType      string `orm:"size(255)"`
	UniqueCode    string
	SettingData   string `orm:"type(text)"`
	CreatedDate   time.Time
	UpdatedDate   time.Time
	CreatedBy     int
	UpdatedBy     int
}
type CountriesTable struct {
	CountryId   int `orm:"auto"`
	CountryCode string
	CountryName string
}

type StatesTable struct {
	CountryId int
	StateId   int `orm:"auto"`
	StateCode string
	StateName string
}

type CitiesTable struct {
	CountryId int
	StateId   int
	CityId    int `orm:"auto"`
	CityCode  string
	CityName  string
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=Dev@123 host=localhost port=5432 dbname=golang_practice sslmode=disable")
	orm.RegisterModel(new(UserMasterTable), new(CarsMasterTable), new(HomePagesSettingTable), new(CountriesTable), new(StatesTable), new(CitiesTable))
	// orm.Debug = true

	orm.RunSyncdb("default", false, true)
}
