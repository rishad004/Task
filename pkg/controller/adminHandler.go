package controller

import (
	"fmt"
	"net/http"
	"temp/pkg/database"
	"temp/pkg/helper"
	"temp/pkg/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var msg string

const RoleAdmin = "admin"

func AdminLoginGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("")
		fmt.Println("Admin Login Page Showing.....................")
		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		session := sessions.Default(ctx)
		Check := session.Get(RoleAdmin)
		if Check == nil {
			ctx.HTML(200, "adminlog.html", msg)
			msg = ""
		} else {
			ctx.Redirect(http.StatusSeeOther, "/admin/panel")
		}
	}
}

func AdminLoginPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("")
		fmt.Println("Admin Entered Values.........................")

		var ad model.Admins
		database.DataBase.First(&ad, "Email=?", ctx.Request.PostFormValue("admail"))

		if ad.Email == ctx.Request.PostFormValue("admail") && ad.Password == ctx.Request.PostFormValue("adpass") {
			helper.JwtTokenStart(ctx, ad.Name, RoleAdmin)
			ctx.Redirect(http.StatusSeeOther, "/admin/panel")
		} else {
			msg = "You are not an Admin!!!"
			ctx.Redirect(http.StatusSeeOther, "/admin/login")
		}
	}
}

func AdminPanelGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("")
		fmt.Println("Admin Panel Showing....................")
		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		session := sessions.Default(ctx)
		Check := session.Get(RoleAdmin)
		if Check == nil {
			ctx.Redirect(http.StatusSeeOther, "/admin/login")
		} else {
			var detail []model.Users
			database.DataBase.Find(&detail)
			ctx.HTML(200, "adminpanel.html", gin.H{"detail": detail,
				"admin": Check,
				"msg":   msg})
			msg = ""
		}
	}
}
func AdminLogoutGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("")
		fmt.Println("Admin Logging Out.......................")
		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		ctx.SetCookie("admin", "", -1, "/", "localhost", false, true)
		session := sessions.Default(ctx)
		session.Delete(RoleAdmin)
		session.Save()
		msg = "Logged out successfully"
		ctx.Redirect(http.StatusSeeOther, "/admin/login")
	}
}
func BlockAndUnblockUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("")
		fmt.Println("Blocking User.........................")
		session := sessions.Default(ctx)
		Check := session.Get(RoleAdmin)
		if Check != nil {
			UserID := ctx.Param("id")
			var block model.Users
			database.DataBase.First(&block, UserID)
			if block.Bool == "Unblock" {
				database.DataBase.Model(&block).Update("Bool", "Block")
				msg = "Unblocked User"
			} else {
				database.DataBase.Model(&block).Update("Bool", "Unblock")
				msg = "Blocked User"
			}
			ctx.Redirect(http.StatusSeeOther, "/admin/panel")
		} else {
			ctx.Redirect(http.StatusSeeOther, "/admin/login")
		}
	}
}
func EditUserGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		UserID := ctx.Param("id")
		fmt.Println("")
		fmt.Println("User Editing Page Showing........................")
		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		session := sessions.Default(ctx)
		Check := session.Get(RoleAdmin)
		if Check == nil {
			ctx.Redirect(http.StatusSeeOther, "/admin/login")
		} else {
			ctx.HTML(200, "edit.html", UserID)
		}
	}
}
func EditUserPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("")
		fmt.Println("Editing User Post.....................")
		var edit model.Users
		UserID := ctx.Param("id")
		database.DataBase.First(&edit, UserID)

		database.DataBase.Model(&edit).Update("Name", ctx.Request.FormValue("cname"))
		err := database.DataBase.Model(&edit).Update("Email", ctx.Request.FormValue("cmail"))
		if err.Error != nil {
			msg = "Name changed but Email already exist"
			ctx.Redirect(http.StatusSeeOther, "/admin/panel")
		} else {
			msg = "Saved changes"
			ctx.Redirect(http.StatusSeeOther, "/admin/panel")
		}
	}
}
func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("")
		fmt.Println("Deleting User........................")
		session := sessions.Default(ctx)
		Check := session.Get(RoleAdmin)
		if Check != nil {
			var delete model.Users
			UserID := ctx.Param("id")
			database.DataBase.First(&delete, UserID)
			database.DataBase.Delete(&delete)
			msg = "Deleted!!!"
			ctx.Redirect(http.StatusSeeOther, "/admin/panel")
		} else {
			ctx.Redirect(http.StatusSeeOther, "/admin/login")
		}
	}
}
