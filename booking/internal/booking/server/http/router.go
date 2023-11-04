package http

import (
	"booking/internal/booking/repository"
	"booking/internal/booking/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, rep *repository.Repo, eh *EndpointHandler) {
	v1 := r.Group("api/booking/v1")
	{
		v1.POST("/booking/create", rep.CreateApartment())
		v1.GET("/booking/read/all", rep.GetApartments())
		v1.PUT("/booking/update/:id", rep.UpdateApartment())
		v1.DELETE("/booking/delete/:id", rep.DeleteApartment())
		v1.GET("/booking/read/:id", rep.GetApartmentById())
		v1.GET("/booking/book/:id", middleware.JWTVerify(), eh.BookApartment())
	}
}
