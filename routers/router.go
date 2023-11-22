package routers

import (
	"crudDemo/controllers"
	"crudDemo/middelware"

	"github.com/astaxie/beego"
)

func RoutersFunction() {

	userController := &controllers.UserController{}
	carsController := &controllers.CarsControllers{}
	thirdpartyOTP := &controllers.ThirdPartyOtpSend{}
	fileUpload := &controllers.FileUploadController{}

	user := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSRouter("/login", userController, "post:Login"),
			beego.NSRouter("/send_otp", thirdpartyOTP, "post:SendOTP"),
			beego.NSRouter("/add_user", userController, "post:RegisterUser"),
			beego.NSRouter("/upload_file", fileUpload, "post:Upload"),
			// beego.NSRouter("/delete_user", userController, "post:Delete"),
			// beego.NSRouter("/login_user", userController, "post:LoginUser"),
		),
		beego.NSNamespace("/cars",
			beego.NSBefore(middelware.Auth),

			beego.NSRouter("/register_car", carsController, "post:RegisterCar"),
			beego.NSRouter("/update_cars", carsController, "post:UpdateCar"),
			beego.NSRouter("/delete_cars", carsController, "post:DeleteCar"),
			beego.NSRouter("/fetch_cars", carsController, "post:FetchCar"),
		),
	)

	beego.AddNamespace(user)
}
