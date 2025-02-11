package api

import (
	"demo1/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var UserRegister service.UserService
	if err := c.ShouldBind(&UserRegister); err == nil {
		res := UserRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func UserLogin(c *gin.Context) {
	var UserLogin service.UserService
	if err := c.ShouldBind(&UserLogin); err == nil {
		res := UserLogin.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
