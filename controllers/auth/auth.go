package auth

import (
	repo "booking-api/repository/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authRepo repo.AuthenRepository
}

func NewAuthController(authRepo repo.AuthenRepository) *AuthController {
	return &AuthController{
		authRepo: authRepo,
	}
}

func (ac *AuthController) Register(ctx *gin.Context) {
	var req repo.CustomerRegRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := ac.authRepo.RegisterRepository(req)
	if err != nil {
		if err.Error() == ("user exists") {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errorCode":   http.StatusInternalServerError,
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

func (ac *AuthController) Login(ctx *gin.Context) {
	var req repo.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := ac.authRepo.LoginRepository(req)
	if err != nil {
		if err.Error() == ("user does not exist") {
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

	ctx.JSON(http.StatusOK, gin.H{
		"token": res,
	})
}

func (ac *AuthController) RegisterForEmployee(ctx *gin.Context) {
	var req repo.EmployeeRegRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := ac.authRepo.RegisterForEmployeeRepository(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (ac *AuthController) LoginForEmployee(ctx *gin.Context) {
	var req repo.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := ac.authRepo.LoginForEmployeeRepository(req)
	if err != nil {
		if err.Error() == ("user does not exist") {
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

	ctx.JSON(http.StatusOK, gin.H{
		"token": res,
	})
}
