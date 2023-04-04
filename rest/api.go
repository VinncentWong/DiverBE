package rest

import (
	"net/http"

	_contactHandler "github.com/VinncentWong/DiverBE/app/contact/handler"
	_diverHandler "github.com/VinncentWong/DiverBE/app/diver/handler"
	_storyHandler "github.com/VinncentWong/DiverBE/app/story/handler"
	"github.com/VinncentWong/DiverBE/middleware"
	"github.com/gin-gonic/gin"
)

type Routing struct {
	router *gin.Engine
}

func NewRouting(r *gin.Engine) *Routing {
	return &Routing{
		router: r,
	}
}

func (r *Routing) InitializeCheckHealthRouting() {
	r.router.GET("/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}

func (r *Routing) InitializeDiverRouting(handler *_diverHandler.DiverHandler) error {
	err := handler.CreateUsernameIndex()
	if err != nil {
		return err
	}
	_, err = handler.CreateDiver()
	if err != nil {
		return err
	}
	group := r.router.Group("/diver")
	group.POST("/login", handler.LoginDiver)
	return nil
}

func (r *Routing) InitializeStoryRouting(handler *_storyHandler.StoryHandler) {
	group := r.router.Group("/story")
	group.POST("/create", middleware.ValidateToken(), handler.CreateStory)
	group.DELETE("/delete/:id", middleware.ValidateToken(), handler.DeleteStory)
	group.PATCH("/update/:id", middleware.ValidateToken(), handler.UpdateStory)
	group.GET("/gets", handler.GetStories)
	group.GET("/get/:id", handler.GetStory)
}

func (r *Routing) InitializeContactRouting(handler *_contactHandler.ContactHandler) {
	group := r.router.Group("/contact")
	group.POST("/create", handler.CreateContact)
}
