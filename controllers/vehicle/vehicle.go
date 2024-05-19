package vehicle

import (
	repo "booking-api/repository/vehicle"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	vehicleRepo repo.VehicleRepository
}

func NewVehicleController(vehicleRepo repo.VehicleRepository) *VehicleController {
	return &VehicleController{
		vehicleRepo: vehicleRepo,
	}
}

func (vc *VehicleController) CreateVehicle(ctx *gin.Context) {
	var req repo.VehicleReq

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

	req.CustomerId = id
	res, err := vc.vehicleRepo.CreateVehicleRepo(req)
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
	ctx.JSON(http.StatusOK, res)
}

type VehicleIdReq struct {
	VehicleId uint `form:"vehicleId" binding:"required"`
}

func (vc *VehicleController) GetVehicleById(ctx *gin.Context) {
	var req VehicleIdReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := vc.vehicleRepo.GetVehicleByIdRepo(req.VehicleId)
	if err != nil {
		if err.Error() == ("vehicle does not exist") {
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

func (vc *VehicleController) GetAllVehicle(ctx *gin.Context) {

	res, err := vc.vehicleRepo.GetAllVehicleRepo()
	if err != nil {
		if err.Error() == ("vehicle does not exist") {
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

func (vc *VehicleController) GetVehicleByCusId(ctx *gin.Context) {

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

	res, err := vc.vehicleRepo.GetVehicleByCusIdRepo(id)
	if err != nil {
		if err.Error() == ("vehicle does not exist") {
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

func (vc *VehicleController) UpdateVehicle(ctx *gin.Context) {
	var req repo.UpdateVehicleReq

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

	req.CustomerId = id
	res, err := vc.vehicleRepo.UpdateVehicleRepo(req)
	if err != nil {
		if err.Error() == ("vehicle does not exist") {
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

func (vc *VehicleController) DeleteVehicle(ctx *gin.Context) {
	var req repo.DeleteVehicleReq

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

	req.CustomerId = id
	err := vc.vehicleRepo.DeleteVehicleRepo(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete vehicle succesful",
	})
}
