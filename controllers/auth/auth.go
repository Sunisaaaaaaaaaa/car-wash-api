package auth

import (
	repo "booking-api/repository/auth"
	"fmt"
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
		if err == fmt.Errorf("user exists") {
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

func (ac *AuthController) Loggin(ctx *gin.Context) {
	var req repo.CustomerLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := ac.authRepo.LogginRepository(req)
	if err != nil {
		if err == fmt.Errorf("user does not exists") {
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

	ctx.JSON(http.StatusOK, gin.H{
		"token": res,
	})
}
