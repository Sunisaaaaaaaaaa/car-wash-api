package invoice

import (
	repo "booking-api/repository/invoice"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InvoiceController struct {
	invoiceRepo repo.InvoiceRepository
}

func NewInvoiceController(invoiceRepo repo.InvoiceRepository) *InvoiceController {
	return &InvoiceController{
		invoiceRepo: invoiceRepo,
	}
}

type InvoiceIdReq struct {
	InvoiceId uint `form:"invoiceId" binding:"required"`
}

func (ic *InvoiceController) GetInvoiceById(ctx *gin.Context) {
	var req InvoiceIdReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := ic.invoiceRepo.GetInvoiceByIdRepo(req.InvoiceId)
	if err != nil {
		if err.Error() == ("invoice does not exist") {
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

func (ic *InvoiceController) GetInvoiceByDate(ctx *gin.Context) {
	var req repo.InvoiceByDateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := ic.invoiceRepo.GetInvoiceByDateRepo(req)
	if err != nil {
		if err.Error() == ("invoice does not exist") {
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

func (ic *InvoiceController) GetPaidInvoiceByDate(ctx *gin.Context) {
	var req repo.InvoiceByDateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := ic.invoiceRepo.GetPaidInvoiceByDateRepo(req)
	if err != nil {
		if err.Error() == ("invoice does not exist") {
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

func (ic *InvoiceController) GetCancledInvoiceByDate(ctx *gin.Context) {
	var req repo.InvoiceByDateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := ic.invoiceRepo.GetCancledInvoiceByDateRepo(req)
	if err != nil {
		if err.Error() == ("invoice does not exist") {
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

func (ic *InvoiceController) GetInvoiceByEmpIdAndDate(ctx *gin.Context) {
	var req repo.InvoiceByDateReq

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

	res, err := ic.invoiceRepo.GetInvoiceByEmpIdAndDateRepo(id, req)
	if err != nil {
		if err.Error() == ("invoice does not exist") {
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

func (ic *InvoiceController) GetInvoiceByCusId(ctx *gin.Context) {

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

	res, err := ic.invoiceRepo.GetInvoiceByCusIdRepo(id)
	if err != nil {
		if err.Error() == ("invoice does not exist") {
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

type UpdateToPayInvoiceReq struct {
	InvoiceId uint `form:"invoiceId" binding:"required"`
}

func (ic *InvoiceController) UpdateToPayInvoice(ctx *gin.Context) {
	var req UpdateToPayInvoiceReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   http.StatusBadRequest,
			"errorDetail": err.Error(),
		})
		return
	}

	res, err := ic.invoiceRepo.UpdatePaidInvoiceRepo(req.InvoiceId)
	if err != nil {
		if err.Error() == ("invoice does not exist") {
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
