package requestStruct

type InsertUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Mobile    string `json:"mobile"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MobileNumber struct {
	MobileNumber string `json:"mobile_number"`
}
