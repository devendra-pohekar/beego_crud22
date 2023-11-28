package routers

import (
	"crudDemo/controllers"

	"github.com/astaxie/beego"
)

func RoutersFunction() {

	userController := &controllers.UserController{}
	carsController := &controllers.CarsControllers{}
	homeSettingController := &controllers.HomeSettingController{}
	// thirdpartyOTP := &controllers.ThirdPartyOtpSend{}

	user := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSRouter("/add_user", userController, "post:RegisterUser"),

			beego.NSRouter("/login", userController, "post:Login"),
			beego.NSRouter("/send_otp", userController, "post:SendMailForm"),
			beego.NSRouter("/verify_email", userController, "post:VerifyEmail"),
			beego.NSRouter("/send_otp_forgot", userController, "post:SendMailForForgotPassword"),
			beego.NSRouter("/verify_otp_forgot", userController, "post:ForgotPasswordUpdate"),
		),
		beego.NSNamespace("/cars",
			// beego.NSBefore(middelware.Auth),

			beego.NSRouter("/register_car", carsController, "post:RegisterCar"),
			beego.NSRouter("/update_cars", carsController, "post:UpdateCar"),
			beego.NSRouter("/delete_cars", carsController, "post:DeleteCar"),
			beego.NSRouter("/fetch_cars", carsController, "post:FetchCar"),
		),

		beego.NSNamespace("/homepage",
			// beego.NSBefore(middelware.Auth),
			beego.NSRouter("/register_settings", homeSettingController, "post:RegisterSettings"),
			beego.NSRouter("/update_settings", homeSettingController, "post:UpdateSettings"),
			beego.NSRouter("/fetch_settings", homeSettingController, "post:FetchSettings"),
			beego.NSRouter("/delete_settings", homeSettingController, "post:DeleteSetting"),
			// beego.NSRouter("/send_otp", homeSettingController, "post:SendMailForm"),

			// beego.NSRouter("/delete_cars", carsController, "post:DeleteCar"),
			// beego.NSRouter("/fetch_cars", carsController, "post:FetchCar"),
		),
	)

	beego.AddNamespace(user)
}
