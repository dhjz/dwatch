package utils

import (
	"dwatch/bean"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// updateStatusType 更新状态
type updateStatusType string

// const (
// 	// UpdatedNothing 未改变
// 	UpdatedNothing updateStatusType = "未改变"
// 	// UpdatedFailed 更新失败
// 	UpdatedFailed = "失败"
// 	// UpdatedSuccess 更新成功
// 	UpdatedSuccess = "成功"
// )

var notifyInit bean.Notify

func init() {
	notifyInit = bean.Notify{
		Url:      "https://oapi.dingtalk.com/robot/send?access_token=f6738866ed758e20e03e5ca94d21986665408b5062b37818ef6e434d196e2d60",
		Template: `{"msgtype": "text","text": {"content":"*#{Name}的状态为#{Status}, 时间#{CreatedAt}, 超时设置#{Timeout}秒, 备注#{Remark}"}}`,
	}
}

// DoNotify
func DoNotify(taskLog *bean.TaskLog, notify *bean.Notify) {
	if taskLog != nil && taskLog.Status > 1 {
		if taskLog.NotifyType == 1 && taskLog.NotifyId > 0 {
			DoWebhook(taskLog, notify)
		}
	}

}

func DoWebhook(taskLog *bean.TaskLog, notify *bean.Notify) {
	if notify == nil {
		// notify = &notifyInit
		if taskLog.TaskId > 0 {
			_wh := taskLog.GetNotify()
			notify = &_wh
			fmt.Println("util.webhook. get notify by tasklog...", notify.Id)
		}
	}
	if notify != nil && notify.State == 1 && notify.Url != "" {
		// 成功和失败都要触发notify
		method := "GET"
		postPara := ""
		contentType := "application/x-www-form-urlencoded"
		if notify.Template != "" {
			method = "POST"
			postPara = replacePara(taskLog, notify.Template)
			if json.Valid([]byte(postPara)) {
				contentType = "application/json"
			}
		}
		requestURL := replacePara(taskLog, notify.Url)
		u, err := url.Parse(requestURL)
		if err != nil {
			log.Println("notify配置中的URL不正确")
			return
		}
		req, err := http.NewRequest(method, fmt.Sprintf("%s://%s%s?%s", u.Scheme, u.Host, u.Path, u.Query().Encode()), strings.NewReader(postPara))
		if err != nil {
			log.Println("创建notify请求异常, Err:", err)
			return
		}
		req.Header.Add("content-type", contentType)

		clt := CreateHTTPClient()
		resp, err := clt.Do(req)
		body, err := GetHTTPResponseOrg(resp, requestURL, err)
		if err == nil {
			log.Println(fmt.Sprintf("notify调用成功, 返回数据: %s", string(body)), taskLog.Status, taskLog)
		} else {
			log.Println(fmt.Sprintf("notify调用失败，Err：%s", err), taskLog.Status, taskLog)
		}
	}
}

// replacePara 替换参数
func replacePara(taskLog *bean.TaskLog, orgPara string) (newPara string) {
	orgPara = strings.ReplaceAll(orgPara, "#{TaskId}", strconv.Itoa(taskLog.TaskId))
	orgPara = strings.ReplaceAll(orgPara, "#{Name}", taskLog.Name)
	orgPara = strings.ReplaceAll(orgPara, "#{Url}", taskLog.Url)
	orgPara = strings.ReplaceAll(orgPara, "#{WarnWord}", taskLog.WarnWord)
	orgPara = strings.ReplaceAll(orgPara, "#{Status}", getStatus(taskLog.Status))
	orgPara = strings.ReplaceAll(orgPara, "#{Duration}", strconv.Itoa(taskLog.Duration))
	orgPara = strings.ReplaceAll(orgPara, "#{Timeout}", strconv.Itoa(taskLog.Timeout))
	orgPara = strings.ReplaceAll(orgPara, "#{Spec}", taskLog.Spec)
	orgPara = strings.ReplaceAll(orgPara, "#{Remark}", taskLog.Remark)
	orgPara = strings.ReplaceAll(orgPara, "#{CreatedAt}", taskLog.CreatedAt.Format("2006-01-02 15:04:05"))

	return orgPara
}

func getStatus(status int) (result string) {
	if status == 2 {
		return "告警中"
	}
	if status == 9 {
		return "无法访问"
	}
	return "正常"
}
