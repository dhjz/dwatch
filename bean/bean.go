package bean

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Status int8
type CommonMap map[string]interface{}

var db *gorm.DB

const (
	Page     = 1    // 当前页数
	Limit    = 20   // 每页多少条数据
	MaxLimit = 1000 // 每次最多取多少条
)

const DefaultTimeFormat = "2006-01-02 15:04:05"

const (
	dbPingInterval = 90 * time.Second
	dbMaxLiftTime  = 2 * time.Hour
)

type BaseBean struct {
	Page  int `json:"page" gorm:"-:all"`
	Limit int `json:"limit" gorm:"-:all"`
}

func (bean *BaseBean) parsePages(params CommonMap) {
	page, ok := params["Page"]
	if ok {
		bean.Page = page.(int)
	}
	pageSize, ok := params["Limit"]
	if ok {
		bean.Limit = pageSize.(int)
	}
	if bean.Page <= 0 {
		bean.Page = Page
	}
	if bean.Limit <= 0 {
		bean.Limit = MaxLimit
	}
}

func (bean *BaseBean) Offset() int {
	return (bean.Page - 1) * bean.Limit
}

// 解析where
// func parseWhere(db *gorm.DB, params CommonMap) {
// 	if len(params) == 0 {
// 		return
// 	}
// 	id, ok := params["Id"]
// 	if ok && id.(int) > 0 {
// 		session.And("id = ?", id)
// 	}
// 	name, ok := params["Name"]
// 	if ok && name.(string) != "" {
// 		session.And("name = ?", name)
// 	}
// }
func init() {
	db = CreateDb()
}

// 创建Db
func CreateDb() *gorm.DB {
	_db, err := gorm.Open(sqlite.Open("dwatch.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connect database...")

	// 迁移 schema
	_db.AutoMigrate(&Task{}, &TaskLog{}, &Notify{})

	return _db
}

// func keepDbAlived(engine *xorm.Engine) {
// 	t := time.Tick(dbPingInterval)
// 	var err error
// 	for {
// 		<-t
// 		err = engine.Ping()
// 		if err != nil {
// 			logger.Infof("database ping: %s", err)
// 		}
// 	}
// }
