### 网站监控程序
#### 开发语言
- 前端: vue2
- 后端: go
- 数据库: sqlite (方便本地运行, 无需额外环境)
#### 功能介绍
- 添加网站监控任务, 通过http请求来监控网站状态\
- 基于cron表达式定时监控
- 支持配置响应超时时间, 并对响应时间数据做统计, 使用echart做看板分析
- 支持任务的总体启动暂停和单个任务的启动暂停
- 支持设置告警词汇, 检测页面中包含关键字则设置为告警状态
- 支持查看监控日志明细, 以及分析图
- 支持自定义webhook发送告警或者无法访问的通知, 例如可以配合钉钉机器人进行消息推送
#### 打包安装
- 双击build.bat 可以打包window和linux运行文件, 双击生成的dist/dwatch.exe, 即可运行
- 打开页面`http://localhost:3457/web/` 访问页面
- dwatch.exe -p 3457 可以指定端口
#### 其他说明
- 目前前端页面用的cdn资源, 如需离线运行可以下载对应的js和css文件, 修改index.html里面的路径即可
- 页面效果图见`appimg`目录
- ![alt 首页图](https://raw.githubusercontent.com/dhjz/dwatch/master/appimg/app.jpg)
- ![alt 明细图](https://raw.githubusercontent.com/dhjz/dwatch/master/appimg/detail.jpg)
- ![alt 数据看板图](https://raw.githubusercontent.com/dhjz/dwatch/master/appimg/data.jpg)
