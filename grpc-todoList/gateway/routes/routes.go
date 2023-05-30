package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"grpc-todoList/gateway/internal/handler"
	middleware "grpc-todoList/gateway/middlerware"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service))
	fmt.Println(service[0], "Zz", service[1])
	v1 := ginRouter.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("user/register", handler.UserRegister)
		v1.POST("user/login", handler.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JwT())
		{
			authed.POST("task/create", handler.TaskCreate)
			authed.POST("task/get/all", handler.GetAllTask)
			authed.POST("task/:tid", handler.SetTask)
			authed.POST("task/get/finish", handler.GetFinishTask)
			authed.POST("task/get/unfinished", handler.GetUnFinishTask)
			authed.DELETE("task/delete/:tid", handler.Delete)
			authed.DELETE("task/delete/finish", handler.DeleteFinishTask)
			authed.DELETE("task/delete/unfinished", handler.DeleteUnFinishTask)
		}
	}
	return ginRouter
}
