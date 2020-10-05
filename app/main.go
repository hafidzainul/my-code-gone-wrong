package main

import (
	"assesmentbulk/app/controllers"
	"assesmentbulk/app/models"
	"connection/app/config"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// database connection
	db := config.Connect()
	strDB := controllers.StrDB{DB: db}

	// migrations
	models.Migrations(db)

	router := gin.Default()

	office := router.Group("office")
	{
		office.GET("/", strDB.ListOffice)
		office.GET("/detail/:id", strDB.DetailOffice)
		office.GET("/search", strDB.SearchOffice)
		office.POST("/", strDB.CreateOffice)
		office.PUT("/:id", strDB.UpdateOffice)
		office.DELETE("/:id", strDB.DeleteOffice)
	}

	user := router.Group("user")
	{
		user.GET("/", strDB.ListUser)
		user.GET("/detail/:id", strDB.DetailUser)
		user.GET("/search", strDB.SearchUser)
		user.POST("/", strDB.CreateUser)
		user.PUT("/:id", strDB.UpdateUser)
		user.DELETE("/:id", strDB.DeleteUser)
	}

	todos := router.Group("todos")
	{
		todos.GET("/", strDB.ListTodos)
		todos.GET("/detail/:id", strDB.DetailTodos)
		todos.GET("/search", strDB.SearchTodos)
		todos.POST("/", strDB.CreateTodos)
		todos.PUT("/:id", strDB.UpdateTodos)
		todos.DELETE("/:id", strDB.DeleteTodos)
	}

	custom := router.Group("custom")
	{
		custom.GET("/user-office", strDB.UserOnOffice)
		custom.GET("/user-office-redis", strDB.UserOnOfficeWithRedis)
		custom.GET("/user-jobs", strDB.UserJobs)
		custom.GET("/office-jobs", strDB.OfficeJobs)
		custom.GET("/user-jobs/:id", strDB.UserJobs)
		custom.GET("/office-jobs/:id", strDB.OfficeJobs)
		custom.GET("/office-user-job", strDB.OfficeUserJobs)
		custom.GET("/office-by-user/:id", strDB.OfficeByUser)
		custom.GET("/user-by-job/:id", strDB.UserByJob)
		custom.GET("office-by-job/:id", strDB.OfficeByJob)
	}

	router.Run()
}
