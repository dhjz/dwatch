package controller

import (
	"dwatch/bean"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskLogController struct {
}

//新增和保存
func (controller *TaskLogController) Save(c *gin.Context) {
	var taskLog bean.TaskLog
	c.ShouldBind(&taskLog)
	taskLog.Save()
	c.JSON(http.StatusOK, taskLog)
}

//查询单个id
func (controller *TaskLogController) Get(c *gin.Context) {
	var taskLog = &bean.TaskLog{}
	id := c.Query("id")
	taskLog.GetStr(id)
	c.JSON(http.StatusOK, taskLog)
}

//查询列表
func (controller *TaskLogController) List(c *gin.Context) {
	// var taskLogs []bean.TaskLog
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	taskLog := &bean.TaskLog{}
	params := bean.CommonMap{
		"TaskId":    c.Query("taskId"),
		"Name":      c.Query("name"),
		"Url":       c.Query("url"),
		"Status":    c.Query("status"),
		"Remark":    c.Query("remark"),
		"StartTime": c.Query("startTime"),
		"EndTime":   c.Query("endTime"),
		"Limit":     limit,
		"Page":      page,
	}
	taskLogs := taskLog.List(params)
	c.JSON(http.StatusOK, taskLogs)
}

//删除
func (controller *TaskLogController) Del(c *gin.Context) {
	// var taskLog bean.TaskLog
	// c.ShouldBind(&taskLog)
	id := c.Query("id")
	(&bean.TaskLog{}).DelStr(id)
	c.JSON(http.StatusOK, gin.H{
		id: id,
	})
}
