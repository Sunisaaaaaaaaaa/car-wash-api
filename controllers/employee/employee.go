package employee

import (
	repo "booking-api/repository/employee"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeRepo repo.EmployeeRepository
}

func NewCustomerController(employeeRepo repo.EmployeeRepository) *EmployeeController {
	return &EmployeeController{
		employeeRepo: employeeRepo,
	}
}

func (ec *EmployeeController) GetEmployeeById(ctx *gin.Context) {
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

	res, err := ec.employeeRepo.GetEmployeeByIdRepo(id)

	if err != nil {
		if err == fmt.Errorf("employee does not exist") {
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

func (ec *EmployeeController) GetAllEmployee(ctx *gin.Context) {

	res, err := ec.employeeRepo.GetAllEmployeeRepo()
	if err != nil {
		if err == fmt.Errorf("employee does not exist") {
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

func (ec *EmployeeController) GetActivedEmployee(ctx *gin.Context) {

	res, err := ec.employeeRepo.GetActivedEmployeeRepo()
	if err != nil {
		if err == fmt.Errorf("employee does not exist") {
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

func (ec *EmployeeController) GetInActivedEmployee(ctx *gin.Context) {

	res, err := ec.employeeRepo.GetInActivedEmployeeRepo()
	if err != nil {
		if err == fmt.Errorf("employee does not exist") {
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

func (ec *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	var req repo.EmployeeReq

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
		return
	}

	req.EmployeeId = id

	res, err := ec.employeeRepo.UpdateEmployeeRepo(req)

	if err != nil {
		if err == fmt.Errorf("employee does not exist") {
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

type DeactivateEmployeereq struct {
	EmployeeId uint `json:"employeeid"`
}

func (ec *EmployeeController) DeactivateEmployee(ctx *gin.Context) {
	var req DeactivateEmployeereq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := ec.employeeRepo.DeleteEmployeeRepo(req.EmployeeId)
	if err != nil {
		if err == fmt.Errorf("employee does not exist") {
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
