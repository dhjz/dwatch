package bean

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	WarnWord   string    `json:"warnWord"`
	Status     int       `json:"status" gorm:"default:0"`     // 网站状态 1:可用 2: 告警 9: 不可用
	CronState  int       `json:"cronState" gorm:"default:1"`  // 1: 启用 2: 禁用
	NotifyType int       `json:"notifyType" gorm:"default:0"` // 通知类型, 0 不通知 1: webhook
	NotifyId   int       `json:"notifyId"`                    // 通知id
	Spec       string    `json:"spec"`
	Timeout    int       `json:"timeout" gorm:"default:5"`
	Remark     string    `json:"remark"`
	CreatedAt  time.Time `json:"createdAt"`
	IsDelete   int       `json:"isDelete" gorm:"default:0"`
	BaseBean   `json:"-" gorm:"-:all"`
}

func (task *Task) ListCronTask() []Task {
	// task.parsePages(params)
	// list := make([]Task, 0)
	var list []Task
	db.Where("is_delete = ? and cron_state = ? and spec IS NOT NULL", 0, 1).Find(&list)

	return list
}

func (task *Task) Save() {
	if task.Id > 0 { // 更新
		fmt.Println("update task...")
		db.Model(&task).Updates(&task)
	} else { // 新增
		fmt.Println("add task...")
		db.Create(&task)
	}
	// Create
	fmt.Println(task)
}

//查询单个id
func (task *Task) Get(id int) *Task {
	println("get task by id:", id)
	db.First(&task, id)
	return task
}

//查询单个id
func (task *Task) GetStr(id string) *Task {
	println("get task by id:", id)
	if id != "" {
		_id, _ := strconv.Atoi(id)
		db.First(&task, _id)
	}
	return task
}

//删除
func (task *Task) Del(id int) {
	println("delete task by id:", id)
	if id > 0 {
		// task.Id, _ = strconv.Atoi(id)
		db.Delete(&task, id)
	}
}

//删除
func (task *Task) DelStr(id string) {
	println("delete task by id:", id)
	if id != "" {
		_id, _ := strconv.Atoi(id)
		db.Delete(&task, _id)
	}
}

func (task *Task) List(params CommonMap) []Task {
	task.parsePages(params)
	fmt.Println("list taskss...", task, params)
	var tasks []Task
	tx := db.Session(&gorm.Session{Initialized: true}) // Initialized: true
	task.parseWhere(tx, params)
	tx.Order("id desc").Offset(task.Offset()).Limit(task.Limit).Find(&tasks)
	return tasks
}

// 解析where
func (task *Task) parseWhere(tx *gorm.DB, params CommonMap) {
	if len(params) == 0 {
		return
	}
	// Where("user LIKE ?", "%"+userstr+"%").Where("cont LIKE ?", "%"+contstr+"%").Where("time = ?", timestr)
	taskId, ok := params["TaskId"]
	if ok && taskId.(string) != "" {
		tx.Where("task_id = ?", taskId)
	}
	name, ok := params["Name"]
	if ok && name.(string) != "" {
		tx.Where("name LIKE ? OR url LIKE ?", "%"+name.(string)+"%", "%"+name.(string)+"%") //fmt.Sprintf("%v", name)
	}
	status, ok := params["Status"]
	if ok && status.(string) != "" {
		tx.Where("status = ?", status)
	}
	cronState, ok := params["CronState"]
	if ok && cronState.(string) != "" {
		tx.Where("cronState = ?", cronState)
	}
}
