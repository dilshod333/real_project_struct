package api

import (
	h "gin/api/handler"
	"gin/storage"

	"github.com/gin-gonic/gin"
)

type Option struct {
	Storage storage.IStorage
}

func New(option Option) *gin.Engine {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	

	handler := h.New(&h.HandlerConfig{
		Storage: option.Storage,
	})

	router.POST("/course", handler.CreateCourse)

	return router
}
