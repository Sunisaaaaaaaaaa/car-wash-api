package controller

import (
	repo "booking-api/repository/booking"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	bookingRepo repo.BookingRepository
}

func NewBookingController(bookingRepo repo.BookingRepository) *BookingController {
	return &BookingController{
		bookingRepo: bookingRepo,
	}
}

func (bc *BookingController) CreateBooking(ctx *gin.Context) {
	var req repo.BookingCreateReq

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
	res, err := bc.bookingRepo.CreateBookingRepo(req)
	if err != nil {
		if err.Error() == ("user does not exist") || err.Error() == ("vehicle does not exist") {
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

type GetBookingReq struct {
	BookingId uint `form:"bookingId" binding:"required"`
}

func (bc *BookingController) GetBookingById(ctx *gin.Context) {
	var req GetBookingReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := bc.bookingRepo.GetBookingByIdRepo(req.BookingId)
	if err != nil {
		if err.Error() == ("booking does not exist") {
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

func (bc *BookingController) GetBookingByCusId(ctx *gin.Context) {
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

	res, err := bc.bookingRepo.GetBookingByCusIdRepo(id)
	if err != nil {
		if err.Error() == ("booking does not exist") {
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

func (bc *BookingController) GetBookingByEmpId(ctx *gin.Context) {
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

	res, err := bc.bookingRepo.GetBookingByEmpIdRepo(id)
	if err != nil {
		if err.Error() == ("booking does not exist") {
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

func (bc *BookingController) GetBookingByDate(ctx *gin.Context) {
	var req repo.BookingByDateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := bc.bookingRepo.GetBookingByDateRepo(req)
	if err != nil {
		if err.Error() == ("booking does not exist") {
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

func (bc *BookingController) GetBookingThatNotTakenByDate(ctx *gin.Context) {
	var req repo.BookingByDateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := bc.bookingRepo.GetBookingThatNotTakenByDateRepo(req)
	if err != nil {
		if err.Error() == ("booking does not exist") {
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

func (bc *BookingController) UpdateToTakeBooking(ctx *gin.Context) {
	var req repo.UpdateTakenBookingReq

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

	req.EmployeeId = &id

	res, err := bc.bookingRepo.UpdateToTakeBookingRepo(req)
	if err != nil {
		if err.Error() == ("booking does not exist") {
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

func (bc *BookingController) UpdateToFininhBooking(ctx *gin.Context) {
	var req repo.UpdateFinishedBookingReq

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

	req.EmployeeId = &id

	res, err := bc.bookingRepo.UpdateToFinishBookingRepo(req)
	if err != nil {
		if err.Error() == ("booking does not exist") {
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

func (bc *BookingController) CancleBooking(ctx *gin.Context) {
	var req repo.CancleBookingReq

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

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	req.CustomerId = id
	err := bc.bookingRepo.DeleteBookingRepo(req)
	if err != nil {
		if err.Error() == ("booking does not exist") {
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
		"message": "cancle booking already",
	})

}
