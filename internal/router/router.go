package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hieuprokk97/golang-eCommerce.git/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", controller.NewUserController().GetUserById)
	}
	return r
}
