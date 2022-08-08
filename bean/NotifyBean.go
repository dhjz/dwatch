package bean

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

// Webhook Webhook
type Notify struct {
	Id       int                     `json:"id"`
	Type     int                     `json:"type" gorm:"default:1"` // 1: webhook
	Url      string                  `json:"url"`
	Template string                  `json:"template"`
	State    int                     `json:"state" gorm:"default:1"` // 1: 启用 2: 禁用
	Protocol string                  `json:"protocol" gorm:"default:smtp"`
	Host     string                  `json:"host"`
	Port     int                     `json:"port" gorm:"default:465"`
	Username string                  `json:"username"`
	Password string                  `json:"password"`
	BaseBean `json:"-" gorm:"-:all"` //migration
}

func (notify *Notify) Save() {
	if notify.Id > 0 { // 更新
		fmt.Println("update notify...")
		db.Model(&notify).Updates(&notify)
	} else { // 新增
		fmt.Println("add notify...")
		db.Create(&notify)
	}
	// Create
	fmt.Println(notify)
}

//查询单个id
func (notify *Notify) Get(id int) *Notify {
	println("get notify by id:", id)
	db.First(&notify, id)
	return notify
}

//查询单个id
func (notify *Notify) GetStr(id string) *Notify {
	println("get notify by id:", id)
	if id != "" {
		_id, _ := strconv.Atoi(id)
		db.First(&notify, _id)
	}
	return notify
}

//删除
func (notify *Notify) Del(id int) {
	println("delete notify by id:", id)
	if id > 0 {
		// notify.Id, _ = strconv.Atoi(id)
		db.Delete(&notify, id)
	}
}

//删除
func (notify *Notify) DelStr(id string) {
	println("delete notify by id:", id)
	if id != "" {
		_id, _ := strconv.Atoi(id)
		db.Delete(&notify, _id)
	}
}

func (notify *Notify) List(params CommonMap) []Notify {
	notify.parsePages(params)
	fmt.Println("list notifys...", notify, params)
	var notifys []Notify
	tx := db.Session(&gorm.Session{Initialized: true}) // Initialized: true
	notify.parseWhere(tx, params)
	tx.Order("id desc").Offset(notify.Offset()).Limit(notify.Limit).Find(&notifys)
	return notifys
}

// 解析where
func (notify *Notify) parseWhere(tx *gorm.DB, params CommonMap) {
	if len(params) == 0 {
		return
	}
	// Where("user LIKE ?", "%"+userstr+"%").Where("cont LIKE ?", "%"+contstr+"%").Where("time = ?", timestr)
	id, ok := params["Id"]
	if ok && id.(string) != "" {
		tx.Where("id = ?", id)
	}
	template, ok := params["Template"]
	if ok && template.(string) != "" {
		tx.Where("template LIKE ?", "%"+template.(string)+"%") //fmt.Sprintf("%v", name)
	}
	_type, ok := params["Type"]
	if ok && _type.(string) != "" {
		tx.Where("type = ?", _type)
	}
	state, ok := params["State"]
	if ok && state.(string) != "" {
		tx.Where("state = ?", state)
	}
}
