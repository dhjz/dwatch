package schedule

import (
	"dwatch/bean"
	"dwatch/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

var cn *cron.Cron = cron.New(cron.WithSeconds())

func StopAll() {
	fmt.Println("stop watch all...")
	IsStarted = 0
	cn.Stop()
}
func StartAll() {
	fmt.Println("start watch all...")
	IsStarted = 1
	cn.Start()
}

func WatchAll(isStop int) {
	fmt.Println("init watch all...")
	taskBean := new(bean.Task)
	tasks := taskBean.ListCronTask()

	if len(tasks) > 0 {
		for _, t := range tasks {
			AddTaskCron(t)
		}
	}
	if isStop != 1 {
		IsStarted = 1
		cn.Start()
	}

	// 初始化webhook
	var notify = &bean.Notify{}
	id := 1
	notify.Get(id)
	if notify.Id == 0 {
		notify.Url = "https://"
		notify.Template = `{"msgtype": "text","text": {"content":"*#{Name}的状态为#{Status}, 时间#{CreatedAt}, 超时设置#{Timeout}秒, 备注#{Remark}"}}`
		notify.State = 1
		notify.Save()
		fmt.Println("add init webhook....", notify)
	}

	// _id, _ := cn.AddFunc("*/10 * * * * *", func() {
	// 	fmt.Println("start watch cron task..." + time.Now().Format("2006-01-02 15:04:05"))
	// 	WatchSite("http://www.cqmu.edu.cn/")
	// 	WatchSite("http://github.com/")
	// 	WatchSite("http://cqmuweb.cqmu.edu.cn:8080/system/login.jsp")
	// 	WatchSite("http://localhost:3457/api/website/get?id=1")
	// })
	// fmt.Println(_id)

	// // c.AddFunc("@hourly",      func() { fmt.Println("Every hour") })
	// // c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	// cn.Start()
}

func AddTaskCron(task bean.Task) {
	if task.Id < 1 || task.Url == "" || task.Spec == "" || task.CronState != 1 {
		return
	}
	id, _ := cn.AddFunc(task.Spec, func() {
		WatchSite(task)
	})
	Add(strconv.Itoa(task.Id), task.Spec, int(id))
}

func UpdateTaskCron(task bean.Task) {
	if task.Id > 0 && task.Url != "" || task.Spec != "" {
		fmt.Println("更新UpdateTaskCron...del and add....", task)
		cmap, ok := Get(strconv.Itoa(task.Id))
		if ok {
			cn.Remove(cron.EntryID(cmap.Id))
			Delete(strconv.Itoa(task.Id))
		}
		if task.CronState == 1 {
			id, _ := cn.AddFunc(task.Spec, func() {
				WatchSite(task)
			})
			Add(strconv.Itoa(task.Id), task.Spec, int(id))
		}
	}
}

func DeleteTaskCron(task bean.Task) {
	if task.Id > 0 {
		cmap, ok := Get(strconv.Itoa(task.Id))
		if ok {
			cn.Remove(cron.EntryID(cmap.Id))
			Delete(strconv.Itoa(task.Id))
		}
	}
}

func WatchSite(task bean.Task) {
	fmt.Println("start watch site... <" + task.Url + ">  " + time.Now().Format("2006-01-02 15:04:05"))
	// 创建客户端
	client := &http.Client{}
	// 创建请求
	request, err := http.NewRequest("GET", task.Url, nil)
	if err != nil {
		fmt.Printf("创建请求失败！err:%+v", err)
		return
	}
	// 设置请求头
	// request.Header.Add("Cookie", "123")
	// request.Header.Add("Content-Type", "application/json;charset=utf-8")
	// request.Header.Add("Token", "456")
	client.Timeout = 5 * time.Second
	if task.Timeout > 0 && task.Timeout < 30 {
		client.Timeout = time.Duration(task.Timeout) * time.Second
	}

	tasklog := bean.TaskLog{
		TaskId:   task.Id,
		Name:     task.Name,
		Url:      task.Url,
		WarnWord: task.WarnWord,
		// Status:     task.Status,
		NotifyType: task.NotifyType,
		Spec:       task.Spec,
		Timeout:    task.Timeout,
		NotifyId:   task.NotifyId,
		// Duration:   task.Duration,
	}

	start := time.Now().UnixNano() / 1e6
	// 发起请求
	resp, err := client.Do(request)
	end := time.Now().UnixNano() / 1e6
	tasklog.Duration = int(end - start)
	if err != nil {
		fmt.Printf("请求失败: %+v", err)
		tasklog.Duration = task.Timeout * 1000
		tasklog.Status = 9
		tasklog.Remark = strings.ReplaceAll(err.Error(), "\"", "'")
		tasklog.Remark = strings.ReplaceAll(tasklog.Remark, ":", "：")
		tasklog.Remark = strings.ReplaceAll(tasklog.Remark, "{", "<")
		tasklog.Remark = strings.ReplaceAll(tasklog.Remark, "}", ">")
		tasklog.Save()
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("读取Body失败 error: %+v", err)
			return
		}
		bodystr := string(body)
		tasklog.Status = 1
		if strings.Trim(task.WarnWord, " ") != "" {
			wword := strings.Replace(task.WarnWord, "，", ",", -1)
			arr := strings.Split(wword, ",")
			for _, w := range arr {
				if strings.Trim(w, " ") != "" && strings.Index(bodystr, strings.Trim(w, " ")) > -1 {
					// tasklog.WarnWord = w
					tasklog.Remark = w
					tasklog.Status = 2
				}
			}
		}
		tasklog.Save()
		fmt.Println("search warnWords...", strings.Index(bodystr, "存储空间"))
		fmt.Println("-------------------------------------------------------", tasklog.Duration)
	}
	utils.DoNotify(&tasklog, nil)

}
