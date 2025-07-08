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

type UserHandler struct {
    service services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) CreateUserHandler(ctx *gin.Context) {
    var req request.CreateUserRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        response.SendError(ctx, http.StatusBadRequest, "invalid request")
        return
    }

    if err := req.Validate(); err != nil {
        response.SendError(ctx, http.StatusBadRequest, err.Error())
        return
    }

    user, err := h.service.CreateUser(req)
    if err != nil {
        response.SendError(ctx, http.StatusInternalServerError, "error creating user on database")
        return
    }

    response.SendSuccess(ctx, "create-user", user)
}

func (h *UserHandler) DeleteUserHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    if id == "" {
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "parameter").Error())
        return
    }

    err := h.service.DeleteUser(id)
    if err != nil {
        if err == services.ErrUserNotFound {
            response.SendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id: %s not found", id))
            return
        }
        response.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting user with id: %s", id))
        return
    }

    response.SendSuccess(ctx, "delete-user", nil)
}

func (h *UserHandler) ListUserHandler(ctx *gin.Context) {
    users, err := h.service.ListUsers()
    if err != nil {
        response.SendError(ctx, http.StatusInternalServerError, "error listing users")
        return
    }
    response.SendSuccess(ctx, "list-users", users)
}

func (h *UserHandler) ShowUserHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    if id == "" {
        response.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "parameter").Error())
        return
    }

    user, err := h.service.ShowUser(id)
    if err != nil {
        if err == services.ErrUserNotFound {
            response.SendError(ctx, http.StatusNotFound, "user not found")
            return
        }
        response.SendError(ctx, http.StatusInternalServerError, "error fetching user")
        return
    }

    response.SendSuccess(ctx, "show-user", user)
}

func (h *UserHandler) UpdateUserHandler(ctx *gin.Context) {
    var req request.UpdatedUserRequest

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

    user, err := h.service.UpdateUser(id, req)
    if err != nil {
        if err == services.ErrUserNotFound {
           	response.SendError(ctx, http.StatusNotFound, "user not found")
            return
        }
        response.SendError(ctx, http.StatusInternalServerError, "error updating user on database")
        return
    }

    response.SendSuccess(ctx, "update-user", user)
}
