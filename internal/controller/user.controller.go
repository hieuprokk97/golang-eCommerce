package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hieuprokk97/golang-eCommerce.git/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController { //con trỏ
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": service.NewUserService().GetInfoUserService(),
	})
}
