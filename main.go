package main

import (
	"temp/pkg/database"
	"temp/pkg/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	//Initializing GIN
	router := gin.Default()

	//Initializing Session and Cookie
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	//Loaded all html files
	router.LoadHTMLGlob("templates/*")

	//Initialized database and sent the models into database
	database.InitDatabse()

	//Grouped user's route
	user := router.Group("/")
	routers.UserRouter(user)

	//Grouped Admin's route
	admin := router.Group("/admin/")
	routers.AdminRouter(admin)

	//Localhost port set as :8080
	router.Run(":8081")
}
