package schedule

import (
	"fmt"
)

var CronMap map[string]CronItem // key 对应实体主键
var IsStarted = 0

type CronItem struct {
	Id   int
	Spec string
}

// 初始化数据库
func init() {
	fmt.Println("init CronMap....")
	CronMap = map[string]CronItem{
		// "8868": {
		// 	Id:   8,
		// 	Spec: "*/8 * * * * *",
		// },
	}
	// cn = cron.New(cron.WithSeconds())
}

func Add(mainId, spec string, id int) {
	CronMap[mainId] = CronItem{
		Id:   id,
		Spec: spec,
	}
}

func Delete(mainId string) {
	delete(CronMap, mainId)
}

func Get(mainId string) (CronItem, bool) {
	_map, ok := CronMap[mainId]
	return _map, ok
}
