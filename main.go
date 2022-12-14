package main

import (
	"RestApi/dbgorm"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"io"
	"os"
)

func main() {
	fmt.Println("DB Project")
	dbGorm := new(dbgorm.DB)
	dbGorm.Connect()
	dbGorm.Migrate()

	connDB := new(DBObject)
	connDB.InitialConnectDB()
	connDB.InitialLogFile()

	defer connDB.CloseConnection()

	gin.DefaultWriter = io.MultiWriter(connDB.log, os.Stdout)
	r := gin.Default()

	r.LoadHTMLGlob("C:\\Users\\admin\\GolandProjects\\RestApi\\templates\\*")
	r.Static("/assets", "C:\\Users\\admin\\GolandProjects\\RestApi\\assets")

	auth := r.Group("/")
	{
		auth.GET("/", connDB.Authorization)
		//auth.GET("/authorization", connDB.Authorization)
		auth.POST("/register", connDB.RegisterUser)
		auth.POST("/token", connDB.GenerateToken)
		auth.GET("/exit", connDB.Exit)
		api := auth.Group("/api").Use(connDB.Auth())
		{
			api.GET("/infoObject", connDB.CheckInfoObjects)
			api.GET("/users", connDB.GetAccountsUser)
			api.GET("/site", connDB.GetListSite)
			api.GET("/pc-site/:id", connDB.GetPCToSite)
			api.GET("/pc-info/:id", connDB.GetInfoComputer)
			api.GET("/test", connDB.CreateTableSite)
		}
	}

	r.Run("localhost:8080")
}
