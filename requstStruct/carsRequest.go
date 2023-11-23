package requestStruct

type CarsInsert struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	CarType     string `json:"car_types" form:"car_types"`
	File        string `json:"file" form:"file"`
}

type UpdateCars struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	CarType     string `json:"car_types" form:"car_types"`
	CarsId      int    `json:"car_id" form:"car_id"`
	File        string `json:"file" form:"file"`
}

type DeleteCar struct {
	CarsId int `json:"car_id"`
}

type SearchigData struct {
	FilterString string `json:"filter_string" form:"filter_string"`
	CurrentPage  int    `json:"current_page"`
	TotalPages   int    `json:"total_pages"`
}
