package routes

import (
	"booking-api/middleware"
	"os"

	cusCont "booking-api/controllers/customer"
	cusRepo "booking-api/repository/customer"

	authCont "booking-api/controllers/auth"
	authRepo "booking-api/repository/auth"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	authRepo := authRepo.NewCustomerRepository()
	authCont := authCont.NewAuthController(authRepo)

	r.POST("/regis", authCont.Register)
	r.POST("/loggin", authCont.Loggin)

	customer := r.Group("/customer", middleware.ValidateToken())
	customer.Use(middleware.Authorization([]string{os.Getenv("CUSTOMER_ACCESSIBLE_ROLE")}))

	customerRepo := cusRepo.NewCustomerRepository()
	customerCont := cusCont.NewCustomerController(customerRepo)

	customer.GET("/cus", customerCont.GetCustomerById)

	customer.POST("/logout", middleware.LogoutHandler())

}
