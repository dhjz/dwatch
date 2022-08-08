package controller

import (
	"dwatch/bean"
	"dwatch/schedule"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
}

func (controller *TaskController) GetWatch(c *gin.Context) {
	c.JSON(http.StatusOK, schedule.IsStarted)
}

func (controller *TaskController) SetWatch(c *gin.Context) {
	start := c.Query("start")
	if start == "1" {
		schedule.StartAll()
	} else if start == "0" {
		schedule.StopAll()
	}
	c.JSON(http.StatusOK, 1)
}

//手动执行单个任务
func (controller *TaskController) Watch(c *gin.Context) {
	var task = &bean.Task{}
	id := c.Query("id")
	task.GetStr(id)
	schedule.WatchSite(*task)
	c.JSON(http.StatusOK, task)
}

//新增和保存
func (controller *TaskController) Save(c *gin.Context) {
	var task bean.Task
	c.ShouldBind(&task)
	task.Save()
	task.Get(task.Id)
	schedule.UpdateTaskCron(task)
	c.JSON(http.StatusOK, task)
}

//查询单个id
func (controller *TaskController) Get(c *gin.Context) {
	var task = &bean.Task{}
	id := c.Query("id")
	task.GetStr(id)
	c.JSON(http.StatusOK, task)
}

//查询列表
func (controller *TaskController) List(c *gin.Context) {
	// var tasks []bean.Task
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	task := &bean.Task{}
	// task.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "20"))
	// task.Page, _ = strconv.Atoi(c.DefaultQuery("page", "0"))
	params := bean.CommonMap{
		"Name":      c.Query("name"),
		"Url":       c.Query("url"),
		"Status":    c.Query("status"),
		"CronState": c.Query("cronState"),
		"Limit":     limit,
		"Page":      page,
	}
	tasks := task.List(params)
	c.JSON(http.StatusOK, tasks)
}

//删除
func (controller *TaskController) Del(c *gin.Context) {
	// var task bean.Task
	// c.ShouldBind(&task)
	id := c.Query("id")
	(&bean.Task{}).DelStr(id)
	c.JSON(http.StatusOK, gin.H{
		id: id,
	})
}
