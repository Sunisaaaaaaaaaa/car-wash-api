package routes

import (
	"booking-api/middleware"
	"os"

	cusCont "booking-api/controllers/customer"
	cusRepo "booking-api/repository/customer"

	authCont "booking-api/controllers/auth"
	authRepo "booking-api/repository/auth"

	bookCont "booking-api/controllers/booking"
	bookRepo "booking-api/repository/booking"

	invoiceCont "booking-api/controllers/invoice"
	invoiceRepo "booking-api/repository/invoice"

	vehicleCont "booking-api/controllers/vehicle"
	vehicleRepo "booking-api/repository/vehicle"

	employeeCont "booking-api/controllers/employee"
	employeeRepo "booking-api/repository/employee"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	authRepo := authRepo.NewCustomerRepository()
	authCont := authCont.NewAuthController(authRepo)
	customerRepo := cusRepo.NewCustomerRepository()
	customerCont := cusCont.NewCustomerController(customerRepo)
	bookRepo := bookRepo.NewBookingRepository()
	bookCont := bookCont.NewBookingController(bookRepo)
	invoiceRepo := invoiceRepo.NewInvoiceRepository()
	invoiceCont := invoiceCont.NewInvoiceController(invoiceRepo)
	vehicleRepo := vehicleRepo.NewVehicleRepository()
	vehicleCont := vehicleCont.NewVehicleController(vehicleRepo)
	employeeRepo := employeeRepo.NewEmployeeRepository()
	employeeCont := employeeCont.NewCustomerController(employeeRepo)

	r.POST("/regis", authCont.Register)
	r.POST("/login", authCont.Login)
	r.POST("/regis-employee", authCont.RegisterForEmployee)
	r.POST("/login-employee", authCont.LoginForEmployee)

	customer := r.Group("/customer", middleware.ValidateToken())
	customer.Use(middleware.Authorization([]string{os.Getenv("CUSTOMER")}))
	customer.GET("/detail", customerCont.GetCustomerById)
	customer.PUT("/detail", customerCont.UpdateCustomer)
	customer.POST("/logout", middleware.LogoutHandler())
	customer.DELETE("/delete", customerCont.DeleteCustomer)
	customer.POST("/booking", bookCont.CreateBooking)
	customer.GET("/booking-by-id", bookCont.GetBookingById)
	customer.GET("/booking", bookCont.GetBookingByCusId)
	customer.GET("/booking-date", bookCont.GetBookingByDate)
	customer.DELETE("/booking", bookCont.CancleBooking)
	customer.GET("/invoice", invoiceCont.GetInvoiceByCusId)
	customer.POST("/vehicle", vehicleCont.CreateVehicle)
	customer.GET("/vehicle", vehicleCont.GetVehicleByCusId)
	customer.PUT("/vehicle", vehicleCont.UpdateVehicle)
	customer.DELETE("/vehicle", vehicleCont.DeleteVehicle)

	employee := r.Group("/employee", middleware.ValidateToken())
	employee.Use(middleware.Authorization([]string{os.Getenv("EMPLOYEE")}))
	employee.GET("/all-cus", customerCont.GetAllCustomer)
	employee.GET("/booking-by-id", bookCont.GetBookingById)
	employee.GET("/booking", bookCont.GetBookingByEmpId)
	employee.GET("/booking-date", bookCont.GetBookingByDate)
	employee.GET("/not-taken-booking", bookCont.GetBookingThatNotTakenByDate)
	employee.PUT("/take-booking", bookCont.UpdateToTakeBooking)
	employee.PUT("/finish-booking", bookCont.UpdateToFininhBooking)
	employee.GET("/invoice", invoiceCont.GetInvoiceById)
	employee.GET("/invoice-date", invoiceCont.GetInvoiceByDate)
	employee.GET("/paid-invoice", invoiceCont.GetPaidInvoiceByDate)
	employee.GET("/cancled-invoice", invoiceCont.GetCancledInvoiceByDate)
	employee.GET("/invoice-by-emId", invoiceCont.GetInvoiceByEmpIdAndDate)
	employee.PUT("/pay-invoice", invoiceCont.UpdateToPayInvoice)
	employee.GET("/vehicle", vehicleCont.GetVehicleById)
	employee.GET("/all-vehicle", vehicleCont.GetAllVehicle)
	employee.GET("/detail", employeeCont.GetEmployeeById)
	employee.GET("/all-employee", employeeCont.GetAllEmployee)
	employee.GET("/actived-employee", employeeCont.GetActivedEmployee)
	employee.GET("/inactived-employee", employeeCont.GetInActivedEmployee)
	employee.PUT("/detail", employeeCont.UpdateEmployee)
	employee.PUT("/deactivate", employeeCont.DeactivateEmployee)

}
