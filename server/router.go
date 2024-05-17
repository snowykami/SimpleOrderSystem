package server

import (
	"github.com/gin-gonic/gin"
	"main/api"
)

func initRouter() {

	router := gin.Default()
	router.POST("/user/register", api.Register)
	router.POST("/user/login", api.Login)
	router.GET("/auth/test", api.AuthTest)

	router.GET("/item/query", api.QueryItem)
	router.POST("/item/add", api.AddItem)
	router.POST("/item/update", api.UpdateItem)

	router.POST("/activity/add", api.AddActivity)
	router.POST("/activity/update", api.UpdateActivity)
	router.POST("/activity/delete", api.DeleteActivity)
	router.POST("/activity/join", api.JoinAct)

	router.POST("/action/grab", GrabItem)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
