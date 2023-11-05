package http

import (
	"booking/internal/booking/booking"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EndpointHandler struct {
	bookingService booking.Service
}

func NewEndpointHandler(bookingService booking.Service) *EndpointHandler {
	return &EndpointHandler{
		bookingService: bookingService,
	}
}

func (eh *EndpointHandler) BookApartment() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (eh *EndpointHandler) GiveStaticFiles() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := ctx.Param("filename")
		filesPath := "booking/internal/files/" + filename
		file := http.Dir(filesPath)
		ctx.FileFromFS(filesPath, file)
	}
}
