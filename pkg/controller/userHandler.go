package controller

import (
	"fmt"
	"net/http"
	"strings"
	"temp/pkg/database"
	"temp/pkg/helper"
	"temp/pkg/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var message string
var fetch model.Users

const RoleUser = "user"

func SignupPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("")
		fmt.Println("Signup Values Entered.....................")
		pass := helper.HashPass(c.Request.PostFormValue("signpass"))
		err := database.DataBase.Create(&model.Users{
			Name:     c.Request.PostFormValue("signname"),
			Email:    strings.ToLower(c.Request.PostFormValue("signmail")),
			Password: pass,
			Bool:     "Block"})
		if err.Error != nil {
			message = "Email already exist"
			c.Redirect(http.StatusSeeOther, "/signup")
		} else {
			message = "Successfully signed up"
			c.Redirect(http.StatusSeeOther, "/")
		}
	}
}

func SignupGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("")
		fmt.Println("Signup Page Showing.....................")
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

		session := sessions.Default(c)
		Check := session.Get(RoleUser)
		if Check == nil {
			c.HTML(http.StatusOK, "signup.html", message)
			message = ""
		} else {
			c.Redirect(http.StatusSeeOther, "/home")
		}

	}
}

func LoginGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("")
		fmt.Println("Login Page Showing.....................")
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		session := sessions.Default(c)
		Check := session.Get(RoleUser)
		if Check == nil {
			c.HTML(http.StatusOK, "login.html", message)
			message = ""
		} else {
			c.Redirect(http.StatusSeeOther, "/home")
		}
	}
}

func LoginPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("")
		fmt.Println("Login Values Entered.......................")
		database.DataBase.First(&fetch, "email=?", strings.ToLower(c.Request.PostFormValue("usermail")))
		err := bcrypt.CompareHashAndPassword([]byte(fetch.Password), []byte(c.Request.PostFormValue("userpass")))
		if err != nil {
			message = "Invalid Email or Password"
			c.Redirect(http.StatusSeeOther, "/")
		} else {
			if fetch.Bool == "Unblock" {
				message = "User blocked by admin"
				c.Redirect(http.StatusSeeOther, "/")
			} else {
				helper.JwtTokenStart(c, fetch.Name, RoleUser)
				c.Redirect(http.StatusSeeOther, "/home")
			}
		}

	}
}
func HomeGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("")
		fmt.Println("Home Page Showing...................")
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		session := sessions.Default(c)
		Check := session.Get(RoleUser)
		if Check == nil {
			c.Redirect(http.StatusSeeOther, "/")
		} else {
			c.HTML(200, "home.html", gin.H{
				"name": Check,
			})
		}
	}
}
func LogoutGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("")
		fmt.Println("Loging Out.......................")
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.SetCookie("user", "", -1, "/", "localhost", false, true)
		session := sessions.Default(c)
		session.Delete(RoleUser)
		session.Save()
		fetch = model.Users{}
		message = "Logged out successfully"
		c.Redirect(http.StatusSeeOther, "/")
	}
}
