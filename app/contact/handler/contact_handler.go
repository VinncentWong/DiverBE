package handler

import (
	"net/http"

	_contactUsecase "github.com/VinncentWong/DiverBE/app/contact/usecase"
	"github.com/VinncentWong/DiverBE/domain/dto/contact"
	"github.com/VinncentWong/DiverBE/util"
	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	u _contactUsecase.IContactUsecase
}

func NewContactHandler(u _contactUsecase.IContactUsecase) *ContactHandler {
	return &ContactHandler{
		u: u,
	}
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
	var container contact.CreateContactDto
	err := c.ShouldBindJSON(&container)
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	result, err := h.u.CreateContact(container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusCreated, "success create contact", true, result)
}
