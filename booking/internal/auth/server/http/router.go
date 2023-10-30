package http

import (
	"booking/internal/auth/repository"
	"github.com/gin-gonic/gin"
)

func InitRouter(rep *repository.Repo, r *gin.Engine) {
	v1 := r.Group("api/auth/v1")
	{
		v1.GET("/register", rep.CreateUserToken())
		v1.GET("/register/update", rep.UpdateUserToken())
	}
}
