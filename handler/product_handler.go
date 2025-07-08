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

type ProductHandler struct {
    service services.ProductService
}

func NewProductHandler(s services.ProductService) *ProductHandler {
    return &ProductHandler{service: s}
}

func (h *ProductHandler) CreateProductHandler(ctx *gin.Context) {
    currentUserRole := ctx.Query("role")
    if currentUserRole == ""{
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("role", "queryParameter").Error())
        return
    }

    var req request.CreateProductRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        response.SendError(ctx, http.StatusBadRequest, "invalid request")
        return
    }

    if err := req.Validate(); err != nil {
        response.SendError(ctx, http.StatusBadRequest, err.Error())
        return
    }

    product, err := h.service.CreateProduct(req, currentUserRole)
    if err != nil {
        response.SendError(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response.SendSuccess(ctx, "create-product", product)
}

func (h *ProductHandler) DeleteProductHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    if id == "" {
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "parameter").Error())
        return
    }

    currentUserRole := ctx.Query("role")
    if currentUserRole == ""{
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("role", "queryParameter").Error())
        return
    }

    err := h.service.DeleteProduct(id, currentUserRole)
    if err != nil {
        if err == services.ErrProductNotFound {
            response.SendError(ctx, http.StatusNotFound, fmt.Sprintf("product with id: %s not found", id))
            return
        }
        response.SendError(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response.SendSuccess(ctx, "delete-product", nil)
}

func (h *ProductHandler) ListProductHandler(ctx *gin.Context) {
    products, err := h.service.ListProducts()
    if err != nil {
        response.SendError(ctx, http.StatusInternalServerError, "error listing products")
        return
    }
    response.SendSuccess(ctx, "list-products", products)
}

func (h *ProductHandler) ShowProductHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    if id == "" {
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "parameter").Error())
        return
    }

    product, err := h.service.ShowProduct(id)
    if err != nil {
        if err == services.ErrProductNotFound {
            response.SendError(ctx, http.StatusNotFound, "product not found")
            return
        }
        response.SendError(ctx, http.StatusInternalServerError, "error fetching product")
        return
    }

    response.SendSuccess(ctx, "show-product", product)
}

func (h *ProductHandler) UpdateProductHandler(ctx *gin.Context) {
    currentUserRole := ctx.Query("role")
    if currentUserRole == ""{
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("role", "queryParameter").Error())
        return
    }
    var req request.UpdatedProductRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        response.SendError(ctx, http.StatusBadRequest, "invalid request")
        return
    }

    if err := req.Validate(); err != nil {
        response.SendError(ctx, http.StatusBadRequest, err.Error())
        return
    }

    id := ctx.Param("id")
    if id == "" {
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "parameter").Error())
        return
    }

    product, err := h.service.UpdateProduct(id, req, currentUserRole)
    if err != nil {
        if err == services.ErrProductNotFound {
            response.SendError(ctx, http.StatusNotFound, "product not found")
            return
        }
        response.SendError(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response.SendSuccess(ctx, "update-product", product)
}
