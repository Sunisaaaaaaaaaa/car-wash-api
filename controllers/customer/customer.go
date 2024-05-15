package customer

import (
	repo "booking-api/repository/customer"
	"fmt"
	"net/http"
	"os"
	"strconv"

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
		ctx.JSON(http.StatusUnauthorized, gin.H{
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
		return
	}

	res, err := cc.customerRepo.GetCustomerByIdRepository(id)

	if err != nil {
		if err == fmt.Errorf("user does not exists") {
			ctx.JSON(http.StatusNotFound, gin.H{
				"errorCode":   http.StatusNotFound,
				"errorDetail": err.Error(),
			})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errorCode":   http.StatusInternalServerError,
				"errorDetail": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, res)
}

func (cc *CustomerController) GetAllCustomer(ctx *gin.Context) {

	employee, ok := ctx.Get("role")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"errorCode":   http.StatusUnauthorized,
			"errorDetail": "role is not set",
		})
		return
	}

	roleInt, err := strconv.Atoi(os.Getenv("EMPLOYEE"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	roleVal := float64(roleInt)
	if employee == roleVal {
		allCus, err := cc.customerRepo.GetAllCustomerRepository()
		if err != nil {
			if err == fmt.Errorf("user does not exists") {
				ctx.JSON(http.StatusNotFound, gin.H{
					"errorCode":   http.StatusNotFound,
					"errorDetail": err.Error(),
				})
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"errorCode":   http.StatusInternalServerError,
					"errorDetail": err.Error(),
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, allCus)
		return

	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": "You are not authorized to access this path or Error to covert string",
		})
		return
	}
}

func (cc *CustomerController) UpdateCustomer(ctx *gin.Context) {
	var req repo.CustomerUpdateReq

	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"errorCode":   http.StatusUnauthorized,
			"errorDetail": "user id is not set",
		})
		return
	}

	var id uint
	if v, ok := userId.(float64); ok {
		id = uint(v)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": "user id invalid",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
	}

	req.ID = uint(id)

	res, err := cc.customerRepo.UpdateCustomerRepository(req)
	if err != nil {
		if err == fmt.Errorf("user does not exists") {
			ctx.JSON(http.StatusNotFound, gin.H{
				"errorCode":   http.StatusNotFound,
				"errorDetail": err.Error(),
			})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errorCode":   http.StatusInternalServerError,
				"errorDetail": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, res)
}

func (cc *CustomerController) DeleteCustomer(ctx *gin.Context) {

	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"errorCode":   http.StatusUnauthorized,
			"errorDetail": "user id is not set",
		})
		return
	}

	var id uint
	if v, ok := userId.(float64); ok {
		id = uint(v)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": "user id invalid",
		})
		return
	}

	_, err := cc.customerRepo.GetCustomerByIdRepository(id)
	if err != nil {
		if err == fmt.Errorf("user does not exists") {
			ctx.JSON(http.StatusNotFound, gin.H{
				"errorCode":   http.StatusNotFound,
				"errorDetail": err.Error(),
			})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errorCode":   http.StatusInternalServerError,
				"errorDetail": err.Error(),
			})
			return
		}
	}

	if err = cc.customerRepo.DeleteCustomerRepository(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"errorCode":   http.StatusOK,
		"errorDetail": "delete successful",
	})

}
