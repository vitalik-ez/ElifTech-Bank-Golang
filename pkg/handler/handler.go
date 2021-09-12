package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", h.showIndexPage)
	router.GET("/bank/view/:bank_id", h.getBank)
	router.POST("/bank/", h.createBank)

	return router
}
