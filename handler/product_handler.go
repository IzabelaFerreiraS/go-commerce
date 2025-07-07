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
    var req request.CreateProductRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        response.SendError(ctx, http.StatusBadRequest, "invalid request")
        return
    }

    if err := req.Validate(); err != nil {
        response.SendError(ctx, http.StatusBadRequest, err.Error())
        return
    }

    product, err := h.service.CreateProduct(req)
    if err != nil {
        response.SendError(ctx, http.StatusInternalServerError, "error creating product on database")
        return
    }

    response.SendSuccess(ctx, "create-product", product)
}

func (h *ProductHandler) DeleteProductHandler(ctx *gin.Context) {
    id := ctx.Query("id")
    if id == "" {
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "queryParameter").Error())
        return
    }

    var req request.DeletedProductRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        response.SendError(ctx, http.StatusForbidden, err.Error())
        return
    }

    err := h.service.DeleteProduct(id, req.Role)
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
    id := ctx.Query("id")
    if id == "" {
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "queryParameter").Error())
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
    var req request.UpdatedProductRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        response.SendError(ctx, http.StatusBadRequest, "invalid request")
        return
    }

    if err := req.Validate(); err != nil {
        response.SendError(ctx, http.StatusBadRequest, err.Error())
        return
    }

    id := ctx.Query("id")
    if id == "" {
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "queryParameter").Error())
        return
    }

    product, err := h.service.UpdateProduct(id, req)
    if err != nil {
        if err == services.ErrProductNotFound {
            response.SendError(ctx, http.StatusNotFound, "product not found")
            return
        }
        response.SendError(ctx, http.StatusInternalServerError, "error updating product on database")
        return
    }

    response.SendSuccess(ctx, "update-product", product)
}
