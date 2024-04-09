package routers

import (
	"temp/pkg/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {

	// GET
	r.GET("/signup", controller.SignupGet())
	r.GET("/", controller.LoginGet())
	r.GET("/home", controller.HomeGet())
	r.GET("/logout", controller.LogoutGet())

	// POST
	r.POST("/signup", controller.SignupPost())
	r.POST("/", controller.LoginPost())

}
