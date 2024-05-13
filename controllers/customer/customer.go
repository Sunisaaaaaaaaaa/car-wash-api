package customer

import (
	repo "booking-api/repository/customer"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerRepo repo.CustomerRepository
}

func NewCustomerController(customerRepo repo.CustomerRepository) *CustomerController {
	return &CustomerController{
		customerRepo: customerRepo,
	}
}

func (cc *CustomerController) GetCustomerById(ctx *gin.Context) {

	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errorCode":   http.StatusUnauthorized,
			"errorDetail": "user id is not set",
		})
	}

	var id uint
	if v, ok := userId.(float64); ok {
		id = uint(v)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": "user id invalid",
		})
	}

	res, err := cc.customerRepo.GetCustomerByIdRepository(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": err.Error(),
		})
		return
	}

	if res == (repo.CustomerRes{}) {
		errNotFound := fmt.Errorf("customer id not found")
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorCode":   http.StatusNotFound,
			"errorDetail": errNotFound.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}


func (cc *CustomerController) GetAllCustomer(ctx *gin.Context) {

	employee, ok := ctx.Get("role")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errorCode":   http.StatusUnauthorized,
			"errorDetail": "role is not set",
		})
	}

	if employee == 

}
