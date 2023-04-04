package handler

import (
	"encoding/json"
	"net/http"

	_storyUsecase "github.com/VinncentWong/DiverBE/app/story/usecase"
	"github.com/VinncentWong/DiverBE/domain"
	_storyDto "github.com/VinncentWong/DiverBE/domain/dto/story"
	"github.com/VinncentWong/DiverBE/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoryHandler struct {
	u _storyUsecase.IStoryUsecase
}

func NewStoryHandler(u _storyUsecase.IStoryUsecase) *StoryHandler {
	return &StoryHandler{
		u: u,
	}
}

func (h *StoryHandler) CreateStory(c *gin.Context) {
	data := c.Request.FormValue("data")
	var container _storyDto.CreateStoryDto
	err := json.Unmarshal([]byte(data), &container)
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
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	result, err := h.u.CreateStory(container, file)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusCreated, "success create story", true, result)
}

func (h *StoryHandler) DeleteStory(c *gin.Context) {
	id := c.Param("id")
	err := h.u.DeleteStory(id)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success delete story", true, nil)
}

func (h *StoryHandler) UpdateStory(c *gin.Context) {
	id := c.Param("id")
	file, _, exist := c.Request.FormFile("file")
	var container domain.Story
	data := c.Request.FormValue("data")
	err := json.Unmarshal([]byte(data), &container)
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	if exist != nil {
		err = h.u.UpdateStory(id, nil, container)
	} else {
		err = h.u.UpdateStory(id, &file, container)
	}
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success update story", true, nil)
}

func (h *StoryHandler) GetStories(c *gin.Context) {
	result, err := h.u.GetStories()
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success get stories", true, result)
}

func (h *StoryHandler) GetStory(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	result, err := h.u.GetStory(objId)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success get story", true, result)
}
