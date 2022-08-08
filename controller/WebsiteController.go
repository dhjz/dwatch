package controller

import (
	"dwatch/bean"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type WebSiteController struct {
}

//新增和保存
func (controller *WebSiteController) Save(c *gin.Context) {
	var task bean.Task
	c.ShouldBind(&task)
	fmt.Println(task)
	// if task.Time == "" {
	// 	task.Time = time.Now().Format("2006-01-02")
	// }
	// id, _ := strconv.Atoi(c.DefaultPostForm("id", "0"))
	// task := bean.Task{
	// 	Id:   id,
	// 	User: c.PostForm("user"),
	// 	Time: c.DefaultPostForm("id", time.Now().Format("2006-01-02")), //2006-01-02 15:04:05
	// 	Cont: c.PostForm("cont"),
	// }
	// name, exist := c.GetPostForm("name")
	// if !exist || name == "" {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "请输入用户名:name",
	// 	})
	// 	return
	// }
	db, err := gorm.Open(sqlite.Open("dwatch.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if task.Id > 0 { // 更新
		fmt.Println("update task...")
		db.Model(&task).Updates(&task)
	} else { // 新增
		fmt.Println("add task...")
		db.Create(&task)
	}
	// Create

	fmt.Println(task)
	c.JSON(http.StatusOK, task)
}

//查询单个id
func (controller *WebSiteController) Get(c *gin.Context) {
	var task bean.Task
	id := c.Query("id")
	println("get by id:", id)
	db, err := gorm.Open(sqlite.Open("dwatch.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.First(&task, id)
	c.JSON(http.StatusOK, task)
}

//查询列表
func (controller *WebSiteController) List(c *gin.Context) {
	var tasks []bean.Task
	timestr := c.Query("time")
	userstr := c.Query("user")
	contstr := c.Query("cont")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	println("list by time:", timestr)
	db, err := gorm.Open(sqlite.Open("dwatch.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if timestr == "" {
		// timestr = time.Now().Format("2006-01-02")
		db.Order("time desc").Limit(limit).Where("user LIKE ?", "%"+userstr+"%").Where("cont LIKE ?", "%"+contstr+"%").Find(&tasks)
	} else {
		db.Order("time desc").Limit(limit).Where("user LIKE ?", "%"+userstr+"%").Where("cont LIKE ?", "%"+contstr+"%").Where("time = ?", timestr).Find(&tasks)
	}
	c.JSON(http.StatusOK, tasks)
}

//删除
func (controller *WebSiteController) Del(c *gin.Context) {
	var task bean.Task
	// c.ShouldBind(&task)
	id := c.Query("id")
	println("delete by id:", id)
	db, err := gorm.Open(sqlite.Open("dwatch.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if id != "" {
		task.Id, _ = strconv.Atoi(id)
		db.Delete(&task)
	}
	c.JSON(http.StatusOK, task)
}
