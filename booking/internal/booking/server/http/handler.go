package http

import (
	"booking/internal/booking/booking"
	"github.com/gin-gonic/gin"
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
