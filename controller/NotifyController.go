package controller

import (
	"dwatch/bean"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotifyController struct {
}

//新增和保存
func (controller *NotifyController) Save(c *gin.Context) {
	var notify bean.Notify
	c.ShouldBind(&notify)
	notify.Save()
	c.JSON(http.StatusOK, notify)
}

//查询单个id
func (controller *NotifyController) Get(c *gin.Context) {
	var notify = &bean.Notify{}
	id := c.Query("id")
	notify.GetStr(id)
	c.JSON(http.StatusOK, notify)
}

//查询列表
func (controller *NotifyController) List(c *gin.Context) {
	// var notifys []bean.Notify
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	notify := &bean.Notify{}
	// notify.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "20"))
	// notify.Page, _ = strconv.Atoi(c.DefaultQuery("page", "0"))
	params := bean.CommonMap{
		"Template": c.Query("Template"),
		"Type":     c.Query("Type"),
		"State":    c.Query("State"),
		"Limit":    limit,
		"Page":     page,
	}
	notifys := notify.List(params)
	c.JSON(http.StatusOK, notifys)
}

//删除
func (controller *NotifyController) Del(c *gin.Context) {
	// var notify bean.Notify
	// c.ShouldBind(&notify)
	id := c.Query("id")
	(&bean.Notify{}).DelStr(id)
	c.JSON(http.StatusOK, gin.H{
		id: id,
	})
}
