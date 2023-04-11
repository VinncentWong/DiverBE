package main

import (
	"log"

	_contactHandler "github.com/VinncentWong/DiverBE/app/contact/handler"
	_contactRepository "github.com/VinncentWong/DiverBE/app/contact/repository"
	_contactUsecase "github.com/VinncentWong/DiverBE/app/contact/usecase"
	_diverHandler "github.com/VinncentWong/DiverBE/app/diver/handler"
	_diverRepository "github.com/VinncentWong/DiverBE/app/diver/repository"
	_diverUsecase "github.com/VinncentWong/DiverBE/app/diver/usecase"
	_photoRepository "github.com/VinncentWong/DiverBE/app/photo/repository"
	_storyHandler "github.com/VinncentWong/DiverBE/app/story/handler"
	_storyRepository "github.com/VinncentWong/DiverBE/app/story/repository"
	_storyUsecase "github.com/VinncentWong/DiverBE/app/story/usecase"
	"github.com/VinncentWong/DiverBE/infrastructure"
	"github.com/VinncentWong/DiverBE/rest"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	err := infrastructure.ConnectToDatabase()
	if err != nil {
		log.Fatal("error occured ", err.Error())
	}
	infrastructure.ConnectToStorage()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	diverRepository := _diverRepository.NewDiverRepository()
	diverUsecase := _diverUsecase.NewDiverUsecase(diverRepository)
	diverHandler := _diverHandler.NewDiverHandler(diverUsecase)

	storyRepository := _storyRepository.NewStoryRepository()
	photoRepository := _photoRepository.NewPhotoRepository()
	storyUsecase := _storyUsecase.NewStoryUsecase(storyRepository, photoRepository)
	storyHandler := _storyHandler.NewStoryHandler(storyUsecase)

	contactRepository := _contactRepository.NewContactRepository()
	contactUsecase := _contactUsecase.NewContactUsecase(contactRepository)
	contactHandler := _contactHandler.NewContactHandler(contactUsecase)

	routing := rest.NewRouting(r)
	routing.InitializeCheckHealthRouting()
	err = routing.InitializeDiverRouting(diverHandler)
	if err != nil {
		log.Fatal("error occured ", err.Error())
	}
	routing.InitializeStoryRouting(storyHandler)
	routing.InitializeContactRouting(contactHandler)
	r.Run()
}
