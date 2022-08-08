package routers

import (
	"dwatch/controller"

	"github.com/gin-gonic/gin"
)

//路由设置
func RegisterRouter(c *gin.Engine) {
	routerTask(c)
	routerTaskLog(c)
	routerNotify(c)
	c.Static("/web", "./webapp") //.StaticFS("/more_static", http.Dir("my_file_system"))

}

func routerTask(c *gin.Engine) {
	var group = c.Group("/api/task")
	{
		con := &controller.TaskController{}
		group.POST("/save", con.Save)
		group.GET("/get", con.Get)
		group.GET("/list", con.List)
		group.DELETE("/del", con.Del)
		group.GET("/watch", con.Watch)
		group.GET("/setwatch", con.SetWatch)
		group.GET("/getwatch", con.GetWatch)
	}
}

func routerTaskLog(c *gin.Engine) {
	var group = c.Group("/api/tasklog")
	{
		con := &controller.TaskLogController{}
		group.POST("/save", con.Save)
		group.GET("/get", con.Get)
		group.GET("/list", con.List)
		group.DELETE("/del", con.Del)
	}
}

func routerNotify(c *gin.Engine) {
	var group = c.Group("/api/notify")
	{
		con := &controller.NotifyController{}
		group.POST("/save", con.Save)
		group.GET("/get", con.Get)
		group.GET("/list", con.List)
		group.DELETE("/del", con.Del)
	}
}
