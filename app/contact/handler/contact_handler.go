package handler

import (
	"net/http"

	_contactUsecase "github.com/VinncentWong/DiverBE/app/contact/usecase"
	"github.com/VinncentWong/DiverBE/domain/dto/contact"
	"github.com/VinncentWong/DiverBE/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	validate := validator.New().Struct(&container)
	if validate != nil {
		errs := validate.(validator.ValidationErrors)
		util.SendResponse(c, http.StatusBadRequest, errs.Error(), false, nil)
		return
	}
	result, err := h.u.CreateContact(container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusCreated, "success create contact", true, result)
}

func (h *ContactHandler) GetContacts(c *gin.Context) {
	result, err := h.u.GetContacts()
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success get contacts", true, result)
}

func (h *ContactHandler) GetContact(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	result, err := h.u.GetContact(objId)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success get contact", true, result)
}

func (h *ContactHandler) DeleteContact(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	err = h.u.DeleteContact(objId)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success delete contact", true, nil)
}
