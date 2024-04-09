package routers

import (
	"temp/pkg/controller"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.RouterGroup) {

	//GET
	r.GET("/login", controller.AdminLoginGet())
	r.GET("/panel", controller.AdminPanelGet())
	r.GET("/logout", controller.AdminLogoutGet())
	r.GET("/blockandunblock/:id", controller.BlockAndUnblockUser())
	r.GET("/delete/:id", controller.DeleteUser())
	r.GET("/edit/:id", controller.EditUserGet())
	
	//POST
	r.POST("/login", controller.AdminLoginPost())
	r.POST("/edit/:id", controller.EditUserPost())

}
