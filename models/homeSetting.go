package models

import (
	"crudDemo/helpers"
	requestStruct "crudDemo/requstStruct"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/beego/beego/orm"
)

func RegisterSetting(c requestStruct.HomeSeetingInsert, user_id float64, file_path interface{}) (int, error) {
	db := orm.NewOrm()
	if file_path == "" {

		file_path, _ = json.Marshal(c.SettingData)
		file_path = string(file_path.([]byte))
		log.Print("----------------------------------------------", file_path)
	}

	res := HomePagesSettingTable{
		Section:     c.Section,
		DataType:    c.DataType,
		UniqueCode:  "",
		SettingData: file_path.(string),
		CreatedBy:   int(user_id),
		UpdatedBy:   0,
		UpdatedDate: time.Now(),
		CreatedDate: time.Now(),
	}

	_, err := db.Insert(&res)
	if err != nil {
		return 0, err
	}
	log.Print("----------------------44444444444------------------------", file_path)
	lastInsertID := res.PageSettingId
	UpdateUniqueCode(lastInsertID)
	return lastInsertID, nil
}

func UpdateUniqueCode(user_id int) (int64, error) {
	db := orm.NewOrm()

	unique_codes := helpers.UniqueCode(user_id, os.Getenv("homePageModule"))
	home_page_setting := HomePagesSettingTable{PageSettingId: user_id}
	if db.Read(&home_page_setting) == nil {
		home_page_setting.UniqueCode = unique_codes
		if num, err := db.Update(&home_page_setting); err == nil {
			return num, nil
		}
	}
	return 1, nil
}

func UpdateSetting(c requestStruct.HomeSeetingUpdate, file_path string, user_id float64) (int64, error) {
	db := orm.NewOrm()
	log.Print("-------------------gggggggggggggggg", file_path)
	page_setting_id := c.SettingId
	/* if setting data found than already exists file in the give directory will be remove and new updated file will insert into folder and also update in database column */
	homePageSetting, err := FetchPageSettingByID(page_setting_id)
	log.Print("-------------------dddddddddddddddd")

	if err != nil {
		return 0, err
	}

	if file_path == "" {
		file_path = c.SettingData
	} else {
		file_name, file_directory := helpers.SplitFilePath(homePageSetting)
		helpers.RemoveFile(file_name, file_directory)
		log.Print("-------------------eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", file_name, file_directory)

	}

	homePageData := HomePagesSettingTable{PageSettingId: page_setting_id}
	log.Print("-------------------jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj", homePageData)

	if db.Read(&homePageSetting) == nil {
		homePageData.UpdatedBy = int(user_id)
		homePageData.UpdatedDate = time.Now()
		homePageData.DataType = c.DataType
		homePageData.Section = c.Section
		homePageData.SettingData = file_path
		if num, err := db.Update(&homePageData); err == nil {
			return num, nil
		}
	}
	log.Print("-------------------qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq")

	return 1, nil

}

func FetchPageSettingByID(pageSettingID int) (string, error) {
	db := orm.NewOrm()
	var pageSetting HomePagesSettingTable
	err := db.Raw(`SELECT  setting_data FROM home_pages_setting_table WHERE page_setting_id = ?`, pageSettingID).QueryRow(&pageSetting)
	if err != nil {
		return "Some errro occured in fetch page setting by ID function", err
	}
	return pageSetting.SettingData, nil
}

func DeleteSetting(page_setting_id int) int {
	db := orm.NewOrm()
	setting := HomePagesSettingTable{PageSettingId: page_setting_id}
	if _, err := db.Delete(&setting); err == nil {
		return 1
	}
	return 0

}

func HomePageSettingExistsDelete(u requestStruct.HomeSeetingDelete) int {
	page_setting_id := u.SettingId
	page_setting_data, err := FetchPageSettingByID(page_setting_id)
	if err != nil {
		return 0
	}
	file_name, file_directory := helpers.SplitFilePath(page_setting_data)
	log.Print(file_directory, file_name, "--------------------------7777777777777777777")
	helpers.RemoveFile(file_name, file_directory)
	DeleteSetting(page_setting_id)
	return 1

}

func FetchSetting() (interface{}, error) {
	db := orm.NewOrm()
	var homeResponse []struct {
		Section     string    `json:"section"`
		DataType    string    `json:"data_type"`
		SettingData string    `json:"setting_data"`
		CreatedDate time.Time `json:"created_date"`
		UpdatedDate time.Time `json:"updated_date"`
		CreatedBy   string    `json:"created_by"`
	}
	_, err := db.Raw(`SELECT section, data_type, setting_data,created_date, updated_date ,created_by FROM home_pages_setting_table `).QueryRows(&homeResponse)

	if err != nil {
		return nil, err
	}

	if len(homeResponse) == 0 {
		return "Not Found Cars", nil
	}
	return homeResponse, nil
}
