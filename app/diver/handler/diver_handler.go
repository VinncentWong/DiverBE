package handler

import (
	"net/http"

	"github.com/VinncentWong/DiverBE/app/diver/usecase"
	"github.com/VinncentWong/DiverBE/domain"
	_diverDto "github.com/VinncentWong/DiverBE/domain/dto/diver"
	"github.com/VinncentWong/DiverBE/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DiverHandler struct {
	u usecase.IDiverUsecase
}

func NewDiverHandler(u usecase.IDiverUsecase) *DiverHandler {
	return &DiverHandler{
		u: u,
	}
}

func (h *DiverHandler) CreateUsernameIndex() error {
	err := h.u.CreateUsernameIndex()
	return err
}

func (h *DiverHandler) CreateDiver() (domain.Diver, error) {
	result, err := h.u.CreateDiver()
	if err != nil {
		if err.Error() != "diver already exist" {
			return domain.Diver{}, err
		}
	}
	return result, nil
}

func (h *DiverHandler) LoginDiver(c *gin.Context) {
	var container _diverDto.LoginDto
	err := c.ShouldBindJSON(&container)
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	validate := validator.New().Struct(&container)
	if validate != nil {
		validateErrors := validate.(validator.ValidationErrors)
		util.SendResponse(c, http.StatusBadRequest, validateErrors.Error(), false, nil)
		return
	}
	result, err := h.u.LoginDiver(container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	jwtToken, err := util.GenerateJwtToken()
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	response := make(map[string]any)
	response["message"] = "diver authenticated"
	response["success"] = true
	response["data"] = result
	response["jwt_token"] = jwtToken
	c.JSON(http.StatusOK, response)
}
