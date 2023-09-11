package handler

import (
	"net/http"

	"github.com/ThailanTec/rascunho/internal/core/domain"
	"github.com/ThailanTec/rascunho/internal/core/services"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	svc services.UserService
}

func NewHttpHandler(svc services.UserService) *HTTPHandler {
	return &HTTPHandler{
		svc: svc,
	}
}

func (h *HTTPHandler) CreateUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})

		return
	}

	usr, err := h.svc.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, usr)
}

func (h *HTTPHandler) GetUser(ctx *gin.Context) {
	name := ctx.Param("name")
	user, err := h.svc.GetUser(name)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
