package bean

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type TaskLog struct {
	Id         int       `json:"id"`
	TaskId     int       `json:"taskId"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	WarnWord   string    `json:"warnWord"`
	Status     int       `json:"status"`     // 网站状态 1:可用 2: 告警 9: 不可用
	NotifyType int       `json:"notifyType"` // 通知类型, 0 不通知
	NotifyId   int       `json:"notifyId"`   // 通知id
	Spec       string    `json:"spec"`
	Timeout    int       `json:"timeout"`
	Duration   int       `json:"duration"` //持续时间
	Remark     string    `json:"remark"`
	CreatedAt  time.Time `json:"createdAt"`
	IsDelete   int       `json:"isDelete"`
	BaseBean   `json:"-" gorm:"-:all"`
}

// // 新增
// func (one *TaskLog) Create() {
// 	db.Create(&one)
// }

func (one *TaskLog) GetNotify() (_notify Notify) {
	var notify Notify
	db.Where("state = ?", 1).First(&notify, one.NotifyId)
	return notify
}

func (taskLog *TaskLog) Save() {
	if taskLog.Id > 0 { // 更新
		fmt.Println("update taskLog...")
		db.Model(&taskLog).Updates(&taskLog)
	} else { // 新增
		fmt.Println("add taskLog...")
		db.Create(&taskLog)
	}
	// 更新task状态
	if taskLog.TaskId > 0 {
		task := &Task{
			Id:     taskLog.TaskId,
			Status: taskLog.Status,
			Remark: taskLog.Remark,
		}
		task.Save()
	}
	// Create
	fmt.Println(taskLog)
}

//查询单个id
func (taskLog *TaskLog) Get(id int) *TaskLog {
	println("get taskLog by id:", id)
	db.First(&taskLog, id)
	return taskLog
}

//查询单个id
func (taskLog *TaskLog) GetStr(id string) *TaskLog {
	println("get taskLog by id:", id)
	if id != "" {
		_id, _ := strconv.Atoi(id)
		db.First(&taskLog, _id)
	}
	return taskLog
}

//删除
func (taskLog *TaskLog) Del(id int) {
	println("delete taskLog by id:", id)
	if id > 0 {
		// taskLog.Id, _ = strconv.Atoi(id)
		db.Delete(&taskLog, id)
	}
}

//删除
func (taskLog *TaskLog) DelStr(id string) {
	println("delete taskLog by id:", id)
	if id != "" {
		_id, _ := strconv.Atoi(id)
		db.Delete(&taskLog, _id)
	}
}

func (taskLog *TaskLog) List(params CommonMap) []TaskLog {
	taskLog.parsePages(params)
	fmt.Println("list taskLogss...", taskLog, params)
	var taskLogs []TaskLog
	tx := db.Session(&gorm.Session{Initialized: true}) // Initialized: true
	taskLog.parseWhere(tx, params)
	tx.Order("id desc").Offset(taskLog.Offset()).Limit(taskLog.Limit).Find(&taskLogs)
	return taskLogs
}

// 解析where
func (taskLog *TaskLog) parseWhere(tx *gorm.DB, params CommonMap) {
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
	remark, ok := params["Remark"]
	if ok && remark.(string) != "" {
		tx.Where("remark LIKE ?", "%"+remark.(string)+"%") //fmt.Sprintf("%v", name)
	}
	status, ok := params["Status"]
	if ok && status.(string) != "" {
		tx.Where("status = ?", status)
	}
	startTime, ok := params["StartTime"]
	if ok && startTime.(string) != "" {
		tx.Where("created_at > ?", startTime)
	}
	endTime, ok := params["EndTime"]
	if ok && endTime.(string) != "" {
		tx.Where("created_at < ?", endTime)
	}

}
