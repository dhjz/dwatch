/*
 Navicat Premium Data Transfer

 Source Server         : dwatch
 Source Server Type    : SQLite
 Source Server Version : 3017000
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3017000
 File Encoding         : 65001

 Date: 08/08/2022 12:06:46
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for notifies
-- ----------------------------
DROP TABLE IF EXISTS "notifies";
CREATE TABLE "notifies" (
  "id" integer,
  "type" integer DEFAULT 1,
  "url" text,
  "template" text,
  "state" integer DEFAULT 1,
  "protocol" text DEFAULT "smtp",
  "host" text,
  "port" integer DEFAULT 465,
  "username" text,
  "password" text,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of "notifies"
-- ----------------------------
INSERT INTO "notifies" VALUES (1, 1, 'https://oapi.dingtalk.com/robot/send?access_token=f6738866ed758e20e03e5ca94d21986665408b5062b37818ef6e434d196e2d60', '{"msgtype": "text","text": {"content":"*#{Name}的状态为#{Status}, 时间#{CreatedAt}, 超时设置#{Timeout}秒, 备注#{Remark}"}}', 1, 'smtp', '', 465, '', '');

-- ----------------------------
-- Table structure for task_logs
-- ----------------------------
DROP TABLE IF EXISTS "task_logs";
CREATE TABLE "task_logs" (
  "id" integer,
  "task_id" integer,
  "name" text,
  "url" text,
  "warn_word" text,
  "status" integer,
  "notify_type" integer,
  "notify_id" integer,
  "spec" text,
  "timeout" integer,
  "duration" integer,
  "remark" text,
  "created_at" datetime,
  "is_delete" integer,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of "task_logs"
-- ----------------------------
INSERT INTO "task_logs" VALUES (1, 2, '二师官网', 'http://www.cque.edu.cn', '', 1, 1, 1, '5 */20 * * * * ', 5, 454, '', '2022-08-08 11:00:05.5203231+08:00', 0);
INSERT INTO "task_logs" VALUES (2, 1, '川外官网', 'http://www.sisu.edu.cn', '', 1, 1, 1, '10 */20 * * * * ', 5, 98, '', '2022-08-08 11:00:10.191628+08:00', 0);
INSERT INTO "task_logs" VALUES (3, 2, '二师官网', 'https://www.cque.edu.cn', '', 1, 1, 1, '5 */20 * * * * ', 5, 283, '', '2022-08-08 11:20:05.2845405+08:00', 0);
INSERT INTO "task_logs" VALUES (4, 1, '川外官网', 'http://www.sisu.edu.cn', '', 1, 1, 1, '10 */20 * * * * ', 5, 93, '', '2022-08-08 11:20:10.6138122+08:00', 0);
INSERT INTO "task_logs" VALUES (5, 3, '重医官网', 'https://www.cqmu.edu.cn', '', 1, 1, 1, '15 */20 * * * * ', 5, 65, '', '2022-08-08 11:20:15.0786449+08:00', 0);
INSERT INTO "task_logs" VALUES (6, 4, '工职院官网', 'http://www.cqipc.edu.cn/', '', 1, 1, 1, '20 */20 * * * * ', 5, 19, '', '2022-08-08 11:20:20.0276528+08:00', 0);
INSERT INTO "task_logs" VALUES (7, 5, '重师官网', 'https://www.cqnu.edu.cn/', '', 1, 1, 1, '25 */20 * * * * ', 5, 24, '', '2022-08-08 11:20:25.024343+08:00', 0);
INSERT INTO "task_logs" VALUES (8, 6, '西南医院官网', 'http://www.xnyy.cn', '', 1, 1, 1, '30 */20 * * * * ', 5, 24, '', '2022-08-08 11:20:30.0260389+08:00', 0);
INSERT INTO "task_logs" VALUES (9, 7, '工贸学院官网', 'https://www.cqgmy.edu.cn/', '', 1, 1, 1, '40 */20 * * * * ', 5, 212, '', '2022-08-08 11:20:40.2126959+08:00', 0);
INSERT INTO "task_logs" VALUES (10, 8, '港物流官网', 'http://www.cqg.com.cn', '', 1, 1, 1, '50 */20 * * * * ', 5, 24, '', '2022-08-08 11:20:50.0283867+08:00', 0);
INSERT INTO "task_logs" VALUES (11, 2, '二师官网', 'https://www.cque.edu.cn', '', 1, 1, 1, '5 */20 * * * * ', 5, 281, '', '2022-08-08 11:40:05.344865+08:00', 0);
INSERT INTO "task_logs" VALUES (12, 1, '川外官网', 'http://www.sisu.edu.cn', '', 1, 1, 1, '10 */20 * * * * ', 5, 109, '', '2022-08-08 11:40:10.1986949+08:00', 0);
INSERT INTO "task_logs" VALUES (13, 3, '重医官网', 'https://www.cqmu.edu.cn', '', 1, 1, 1, '15 */20 * * * * ', 5, 556, '', '2022-08-08 11:40:15.5713678+08:00', 0);
INSERT INTO "task_logs" VALUES (14, 4, '工职院官网', 'http://www.cqipc.edu.cn/', '', 1, 1, 1, '20 */20 * * * * ', 5, 29, '', '2022-08-08 11:40:20.036789+08:00', 0);
INSERT INTO "task_logs" VALUES (15, 5, '重师官网', 'https://www.cqnu.edu.cn/', '', 1, 1, 1, '25 */20 * * * * ', 5, 22, '', '2022-08-08 11:40:25.0238453+08:00', 0);
INSERT INTO "task_logs" VALUES (16, 6, '西南医院官网', 'http://www.xnyy.cn', '', 1, 1, 1, '30 */20 * * * * ', 5, 42, '', '2022-08-08 11:40:30.0445542+08:00', 0);
INSERT INTO "task_logs" VALUES (17, 7, '工贸学院官网', 'https://www.cqgmy.edu.cn/', '', 1, 1, 1, '40 */20 * * * * ', 5, 235, '', '2022-08-08 11:40:40.2350533+08:00', 0);
INSERT INTO "task_logs" VALUES (18, 8, '港物流官网', 'http://www.cqg.com.cn', '', 1, 1, 1, '50 */20 * * * * ', 5, 19, '', '2022-08-08 11:40:50.0236781+08:00', 0);
INSERT INTO "task_logs" VALUES (19, 10, '二师后台', 'http://183.230.3.19:40085/system/login.jsp', '存储空间,数据库', 1, 1, 1, '0 */30 * * * * ', 5, 76, '', '2022-08-08 12:00:00.0768623+08:00', 0);
INSERT INTO "task_logs" VALUES (20, 2, '二师官网', 'https://www.cque.edu.cn', '', 1, 1, 1, '5 */20 * * * * ', 5, 276, '', '2022-08-08 12:00:05.340636+08:00', 0);
INSERT INTO "task_logs" VALUES (21, 11, '重医后台', 'http://cqmuweb.cqmu.edu.cn:8080/system/login.jsp', '存储空间,数据库', 1, 1, 1, '5 */30 * * * * ', 5, 907, '', '2022-08-08 12:00:05.9073769+08:00', 0);
INSERT INTO "task_logs" VALUES (22, 12, '工贸后台', 'http://www1.cqgmy.edu.cn:8080/system/login.jsp', '存储空间,数据库', 1, 1, 1, '10 */30 * * * * ', 5, 213, '', '2022-08-08 12:00:10.2144011+08:00', 0);
INSERT INTO "task_logs" VALUES (23, 1, '川外官网', 'http://www.sisu.edu.cn', '', 1, 1, 1, '10 */20 * * * * ', 5, 116, '', '2022-08-08 12:00:10.2234078+08:00', 0);
INSERT INTO "task_logs" VALUES (24, 3, '重医官网', 'https://www.cqmu.edu.cn', '', 1, 1, 1, '15 */20 * * * * ', 5, 560, '', '2022-08-08 12:00:15.5732416+08:00', 0);
INSERT INTO "task_logs" VALUES (25, 4, '工职院官网', 'http://www.cqipc.edu.cn/', '', 1, 1, 1, '20 */20 * * * * ', 5, 38, '', '2022-08-08 12:00:20.0462894+08:00', 0);
INSERT INTO "task_logs" VALUES (26, 13, '港务后台', 'http://219.153.136.214:8080/system/login.jsp', '', 1, 1, 1, '20 */30 * * * * ', 5, 115, '', '2022-08-08 12:00:20.1153372+08:00', 0);
INSERT INTO "task_logs" VALUES (27, 5, '重师官网', 'https://www.cqnu.edu.cn/', '', 1, 1, 1, '25 */20 * * * * ', 5, 20, '', '2022-08-08 12:00:25.0202547+08:00', 0);
INSERT INTO "task_logs" VALUES (28, 6, '西南医院官网', 'http://www.xnyy.cn', '', 1, 1, 1, '30 */20 * * * * ', 5, 30, '', '2022-08-08 12:00:30.0311366+08:00', 0);
INSERT INTO "task_logs" VALUES (29, 7, '工贸学院官网', 'https://www.cqgmy.edu.cn/', '', 1, 1, 1, '40 */20 * * * * ', 5, 250, '', '2022-08-08 12:00:40.2506622+08:00', 0);
INSERT INTO "task_logs" VALUES (30, 8, '港物流官网', 'http://www.cqg.com.cn', '', 1, 1, 1, '50 */20 * * * * ', 5, 19, '', '2022-08-08 12:00:50.0249237+08:00', 0);
INSERT INTO "task_logs" VALUES (31, 9, '融智学院官网', 'http://www.cfec.edu.cn/', '', 1, 1, 1, '55 */20 * * * * ', 5, 96, '', '2022-08-08 12:00:55.0982429+08:00', 0);

-- ----------------------------
-- Table structure for tasks
-- ----------------------------
DROP TABLE IF EXISTS "tasks";
CREATE TABLE "tasks" (
  "id" integer,
  "name" text,
  "url" text,
  "warn_word" text,
  "status" integer DEFAULT 0,
  "cron_state" integer DEFAULT 1,
  "notify_type" integer DEFAULT 0,
  "notify_id" integer,
  "spec" text,
  "timeout" integer DEFAULT 5,
  "remark" text,
  "created_at" datetime,
  "is_delete" integer DEFAULT 0,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of "tasks"
-- ----------------------------
INSERT INTO "tasks" VALUES (1, '川外官网', 'http://www.sisu.edu.cn', '', 1, 1, 1, 1, '10 */20 * * * * ', 5, '', '2022-08-08 10:52:09.5027174+08:00', 0);
INSERT INTO "tasks" VALUES (2, '二师官网', 'https://www.cque.edu.cn', '', 1, 1, 1, 1, '5 */20 * * * * ', 5, '', '2022-08-08 10:54:35.4563083+08:00', 0);
INSERT INTO "tasks" VALUES (3, '重医官网', 'https://www.cqmu.edu.cn', '', 1, 1, 1, 1, '15 */20 * * * * ', 5, '', '2022-08-08 11:01:43.7438512+08:00', 0);
INSERT INTO "tasks" VALUES (4, '工职院官网', 'http://www.cqipc.edu.cn/', '', 1, 1, 1, 1, '20 */20 * * * * ', 5, '', '2022-08-08 11:02:44.1038529+08:00', 0);
INSERT INTO "tasks" VALUES (5, '重师官网', 'https://www.cqnu.edu.cn/', '', 1, 1, 1, 1, '25 */20 * * * * ', 5, '', '2022-08-08 11:03:36.5988943+08:00', 0);
INSERT INTO "tasks" VALUES (6, '西南医院官网', 'http://www.xnyy.cn', '', 1, 1, 1, 1, '30 */20 * * * * ', 5, '', '2022-08-08 11:04:10.9346799+08:00', 0);
INSERT INTO "tasks" VALUES (7, '工贸学院官网', 'https://www.cqgmy.edu.cn/', '', 1, 1, 1, 1, '40 */20 * * * * ', 5, '', '2022-08-08 11:04:52.0967817+08:00', 0);
INSERT INTO "tasks" VALUES (8, '港物流官网', 'http://www.cqg.com.cn', '', 1, 1, 1, 1, '50 */20 * * * * ', 5, '', '2022-08-08 11:06:18.2866619+08:00', 0);
INSERT INTO "tasks" VALUES (9, '融智学院官网', 'http://www.cfec.edu.cn/', '', 1, 1, 1, 1, '55 */20 * * * * ', 5, '', '2022-08-08 11:52:51.2178306+08:00', 0);
INSERT INTO "tasks" VALUES (10, '二师后台', 'http://183.230.3.19:40085/system/login.jsp', '存储空间,数据库', 1, 1, 1, 1, '0 */30 * * * * ', 5, '', '2022-08-08 11:54:17.0617318+08:00', 0);
INSERT INTO "tasks" VALUES (11, '重医后台', 'http://cqmuweb.cqmu.edu.cn:8080/system/login.jsp', '存储空间,数据库', 1, 1, 1, 1, '5 */30 * * * * ', 5, '', '2022-08-08 11:54:57.5192866+08:00', 0);
INSERT INTO "tasks" VALUES (12, '工贸后台', 'http://www1.cqgmy.edu.cn:8080/system/login.jsp', '存储空间,数据库', 1, 1, 1, 1, '10 */30 * * * * ', 5, '', '2022-08-08 11:55:26.577208+08:00', 0);
INSERT INTO "tasks" VALUES (13, '港务后台', 'http://219.153.136.214:8080/system/login.jsp', '存储空间,数据库', 1, 1, 1, 1, '20 */30 * * * * ', 5, '', '2022-08-08 11:57:43.2194815+08:00', 0);

PRAGMA foreign_keys = true;
