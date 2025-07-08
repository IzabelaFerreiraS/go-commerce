package handler

import (
	"fmt"
	"net/http"

	"go-commerce/dtos/request"
	"go-commerce/services"
	"go-commerce/utils"
	"go-commerce/utils/response"

	"github.com/gin-gonic/gin"
)

type SaleHandler struct {
    service services.SaleService
}

func NewSaleHandler(s services.SaleService) *SaleHandler {
    return &SaleHandler{service: s}
}

func (h *SaleHandler) CreateSaleHandler(ctx *gin.Context) {
    var req request.CreateSaleRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        response.SendError(ctx, http.StatusBadRequest, "invalid request")
        return
    }

    if err := req.Validate(); err != nil {
        response.SendError(ctx, http.StatusBadRequest, err.Error())
        return
    }

    sale, err := h.service.CreateSale(req)
    if err != nil {
        response.SendError(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response.SendSuccess(ctx, "create-sale", sale)
}

func (h *SaleHandler) DeleteSaleHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    if id == "" {
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "parameter").Error())
        return
    }

    err := h.service.DeleteSale(id)
    if err != nil {
        if err == services.ErrSaleNotFound {
            response.SendError(ctx, http.StatusNotFound, fmt.Sprintf("sale with id: %s not found", id))
            return
        }
        response.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting sale with id: %s", id))
        return
    }

    response.SendSuccess(ctx, "delete-sale", nil)
}

func (h *SaleHandler) ListSalesHandler(ctx *gin.Context) {
    sales, err := h.service.ListSales()
    if err != nil {
        response.SendError(ctx, http.StatusInternalServerError, "error listing sales")
        return
    }

    response.SendSuccess(ctx, "list-sales", sales)
}

func (h *SaleHandler) ShowSaleHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    if id == "" {
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "parameter").Error())
        return
    }

    sale, err := h.service.ShowSale(id)
    if err != nil {
        response.SendError(ctx, http.StatusNotFound, "sale not found")
        return
    }

    response.SendSuccess(ctx, "show-sale", sale)
}
