/*
 Navicat Premium Data Transfer

 Source Server         : localhost_flyfox
 Source Server Type    : MySQL
 Source Server Version : 50624
 Source Host           : localhost:13306
 Source Schema         : gcs

 Target Server Type    : MySQL
 Target Server Version : 50624
 File Encoding         : 65001

 Date: 08/10/2019 12:43:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `key` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '键',
  `value` varchar(4000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '值',
  `code` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '编码',
  `data_type` int(2) NULL DEFAULT NULL COMMENT '数据类型//radio/1,KV,2,字典,3,数组',
  `parent_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '类型',
  `parent_key` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` int(11) NOT NULL DEFAULT 10 COMMENT '排序号',
  `project_id` bigint(20) NULL DEFAULT 1 COMMENT '项目ID',
  `copy_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '拷贝状态 1 拷贝  2  不拷贝',
  `change_status` tinyint(1) NOT NULL DEFAULT 2 COMMENT '1 不可更改 2 可以更改',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 68 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统配置表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (24, '系统参数', 'system', '', '', NULL, 0, NULL, 15, 1, 1, 2, 1, '2017-09-15 17:02:30', 4, '2017-09-15 16:54:52', 4);
INSERT INTO `sys_config` VALUES (46, '日志控制配置', 'system.debug', 'false', '', NULL, 24, 'system', 15, 1, 1, 1, 1, '2019-02-24 00:00:08', 0, '2017-09-15 17:06:21', 4);
INSERT INTO `sys_config` VALUES (47, '短信配置', 'sms', '', '', NULL, 0, '', 15, 1, 1, 2, 1, '2019-02-20 22:45:41', 1, '2017-09-15 17:06:56', 4);
INSERT INTO `sys_config` VALUES (50, '短信账号', 'sms.username', 'test', '', NULL, 47, 'sms', 10, 1, 1, 2, 1, '2019-02-20 22:26:29', 1, '2019-02-18 01:07:47', 1);
INSERT INTO `sys_config` VALUES (51, '短信密码', 'sms.passwd', '111111', '', NULL, 47, 'sms', 10, 1, 1, 2, 1, '2019-07-10 13:27:01', 1, '2019-02-18 01:08:16', 1);
INSERT INTO `sys_config` VALUES (52, '短信类型', 'sms.type', '阿里云', '', NULL, 47, 'sms', 10, 1, 1, 2, 1, '2019-02-20 22:26:21', 1, '2019-02-20 22:26:21', 1);
INSERT INTO `sys_config` VALUES (53, '性别', 'sex', '', '', NULL, 0, NULL, 90, 1, 1, 2, 1, '2019-02-20 23:35:18', 1, '2019-02-20 23:35:18', 1);
INSERT INTO `sys_config` VALUES (54, '性别男', 'sex.male', '男', '1', NULL, 53, 'sex', 91, 1, 1, 2, 1, '2019-02-20 23:40:19', 1, '2019-02-20 23:35:45', 1);
INSERT INTO `sys_config` VALUES (55, '性别女', 'sex.female', '女', '2', NULL, 53, 'sex', 92, 1, 1, 2, 1, '2019-02-20 23:40:24', 1, '2019-02-20 23:36:12', 1);
INSERT INTO `sys_config` VALUES (56, '性别未知', 'sex.unknown', '未知', '3', NULL, 53, 'sex', 93, 1, 1, 2, 1, '2019-02-20 23:40:29', 1, '2019-02-20 23:36:46', 1);

-- ----------------------------
-- Table structure for sys_department
-- ----------------------------
DROP TABLE IF EXISTS `sys_department`;
CREATE TABLE `sys_department`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int(11) NULL DEFAULT 0 COMMENT '上级机构',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '部门/11111',
  `code` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '机构编码',
  `sort` int(11) NULL DEFAULT 0 COMMENT '序号',
  `linkman` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系人',
  `linkman_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系人电话',
  `remark` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '机构描述',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_depart_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10018 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '组织机构' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_department
-- ----------------------------
INSERT INTO `sys_department` VALUES (10001, 0, 'FLY的狐狸', 'ABC000', 100, '', '', '', 1, '2017-04-28 01:16:43', 1, '2016-07-31 18:12:30', 1);
INSERT INTO `sys_department` VALUES (10002, 10001, '开发组', 'ABC001', 101, NULL, NULL, NULL, 1, '2019-07-09 11:40:40', 1, '2016-07-31 18:15:29', 1);
INSERT INTO `sys_department` VALUES (10003, 10001, '产品组', 'ABC003', 103, '', '', '', 1, '2017-04-28 00:58:41', 1, '2016-07-31 18:16:06', 1);
INSERT INTO `sys_department` VALUES (10004, 10001, '运营组', 'ABC004', 104, NULL, NULL, NULL, 1, '2016-07-31 18:16:30', 1, '2016-07-31 18:16:30', 1);
INSERT INTO `sys_department` VALUES (10005, 10001, '测试组', '12323', 10, '', '', '', 0, '2019-06-30 22:33:44', 1, '2017-10-18 18:13:09', 1);

-- ----------------------------
-- Table structure for sys_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_log`;
CREATE TABLE `sys_log`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `log_type` int(11) NOT NULL COMMENT '类型',
  `oper_object` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作对象',
  `oper_table` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '操作表',
  `oper_id` int(11) NULL DEFAULT 0 COMMENT '操作主键',
  `oper_type` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作类型',
  `oper_remark` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作备注',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12091 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '日志' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_log
-- ----------------------------
INSERT INTO `sys_log` VALUES (11813, 1, NULL, 'sys_user', 1, '登录', NULL, 1, '2019-07-08 23:14:28', 1, '2019-07-08 23:14:28', 1);
INSERT INTO `sys_log` VALUES (11814, 1, NULL, 'sys_user', 1, '登出', NULL, 1, '2019-07-08 23:14:32', 1, '2019-07-08 23:14:32', 1);
INSERT INTO `sys_log` VALUES (11815, 1, NULL, 'sys_user', 1, '登录', NULL, 1, '2019-07-08 23:15:11', 1, '2019-07-08 23:15:11', 1);
INSERT INTO `sys_log` VALUES (11816, 1, NULL, 'sys_user', 1, '登录', NULL, 1, '2019-07-08 23:16:52', 1, '2019-07-08 23:16:52', 1);
INSERT INTO `sys_log` VALUES (11817, 1, NULL, 'sys_user', 1, '登录', NULL, 1, '2019-07-09 09:28:55', 1, '2019-07-09 09:28:55', 1);
INSERT INTO `sys_log` VALUES (11818, 1, NULL, 'sys_user', 1, '登录', NULL, 1, '2019-07-09 09:34:17', 1, '2019-07-09 09:34:17', 1);
INSERT INTO `sys_log` VALUES (11819, 1, NULL, 'sys_user', 1, '登录', NULL, 1, '2019-07-09 09:39:39', 1, '2019-07-09 09:39:39', 1);
INSERT INTO `sys_log` VALUES (11820, 1, NULL, 'sys_user', 1, '登录', NULL, 1, '2019-07-09 17:06:27', 1, '2019-07-09 17:06:27', 1);
INSERT INTO `sys_log` VALUES (11821, 1, NULL, 'sys_user', 1, '登录', NULL, 1, '2019-07-09 17:22:12', 1, '2019-07-09 17:22:12', 1);
INSERT INTO `sys_log` VALUES (11822, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-07-09 22:07:31', 1, '2019-07-09 22:07:31', 1);
INSERT INTO `sys_log` VALUES (11823, 1, '用户', 'sys_user', 1, '登出', NULL, 1, '2019-07-09 22:07:55', 1, '2019-07-09 22:07:55', 1);
INSERT INTO `sys_log` VALUES (11824, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-07-09 22:07:56', 1, '2019-07-09 22:07:56', 1);
INSERT INTO `sys_log` VALUES (11825, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-07-09 22:38:34', 1, '2019-07-09 22:38:34', 1);
INSERT INTO `sys_log` VALUES (11826, 2, '系统配置表', 'sys_config', 58, '更新', '{\"id\":58,\"name\":\"3333\",\"key\":\"3333\",\"value\":\"\",\"code\":\"\",\"dataType\":0,\"parentId\":0,\"parentKey\":\"\",\"sort\":10,\"projectId\":1,\"copyStatus\":\"1\",\"changeStatus\":\"2\",\"enable\":0,\"updateTime\":\"2019-07-09 22:38:41\",\"updateId\":1,\"createTime\":\"\",\"createId\":0}', 1, '2019-07-09 22:38:41', 1, '2019-07-09 22:38:41', 1);
INSERT INTO `sys_log` VALUES (11827, 1, NULL, 'sys_user', 1, '登录', NULL, 1, '2019-07-10 11:45:41', 1, '2019-07-10 11:45:41', 1);
INSERT INTO `sys_log` VALUES (11828, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-07-10 11:48:05', 1, '2019-07-10 11:48:05', 1);
INSERT INTO `sys_log` VALUES (11829, 1, '用户', 'sys_user', 1, '登出', NULL, 1, '2019-07-10 11:48:19', 1, '2019-07-10 11:48:19', 1);
INSERT INTO `sys_log` VALUES (11830, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-07-10 11:48:34', 1, '2019-07-10 11:48:34', 1);
INSERT INTO `sys_log` VALUES (11831, 2, '系统配置表', 'sys_config', 62, '更新', '{\"id\":62,\"name\":\"2898\",\"key\":\"123\",\"value\":\"123\",\"code\":\"123\",\"dataType\":0,\"parentId\":58,\"parentKey\":\"\",\"sort\":10,\"projectId\":1,\"copyStatus\":\"1\",\"changeStatus\":\"2\",\"enable\":0,\"updateTime\":\"2019-07-10 11:48:43\",\"updateId\":1,\"createTime\":\"\",\"createId\":0}', 1, '2019-07-10 11:48:43', 1, '2019-07-10 11:48:43', 1);
INSERT INTO `sys_log` VALUES (11832, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-07-10 11:51:03', 1, '2019-07-10 11:51:03', 1);
INSERT INTO `sys_log` VALUES (11833, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-07-10 12:10:11', 1, '2019-07-10 12:10:11', 1);
INSERT INTO `sys_log` VALUES (11834, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-07-10 13:24:29', 1, '2019-07-10 13:24:29', 1);
INSERT INTO `sys_log` VALUES (11835, 2, '系统配置表', 'sys_config', 67, '插入', '{\"id\":67,\"name\":\"123\",\"key\":\"123\",\"value\":\"\",\"code\":\"\",\"dataType\":0,\"parentId\":0,\"parentKey\":\"\",\"sort\":10,\"projectId\":0,\"copyStatus\":\"1\",\"changeStatus\":\"2\",\"enable\":0,\"updateTime\":\"2019-07-10 13:24:38\",\"updateId\":1,\"createTime\":\"2019-07-10 13:24:38\",\"createId\":1}', 1, '2019-07-10 13:24:38', 1, '2019-07-10 13:24:38', 1);
INSERT INTO `sys_log` VALUES (11836, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-07-10 13:26:23', 1, '2019-07-10 13:26:23', 1);
INSERT INTO `sys_log` VALUES (11837, 2, '系统配置表', 'sys_config', 67, '更新', '{\"id\":67,\"name\":\"123\",\"key\":\"123\",\"value\":\"\",\"code\":\"\",\"dataType\":0,\"parentId\":0,\"parentKey\":\"\",\"sort\":10,\"projectId\":1,\"copyStatus\":\"1\",\"changeStatus\":\"2\",\"enable\":0,\"updateTime\":\"2019-07-10 13:26:48\",\"updateId\":1,\"createTime\":\"\",\"createId\":0}', 1, '2019-07-10 13:26:48', 1, '2019-07-10 13:26:48', 1);
INSERT INTO `sys_log` VALUES (11838, 2, '系统配置表', 'sys_config', 51, '更新', '{\"id\":51,\"name\":\"短信密码\",\"key\":\"sms.passwd\",\"value\":\"111111\",\"code\":\"\",\"dataType\":0,\"parentId\":47,\"parentKey\":\"sms\",\"sort\":10,\"projectId\":1,\"copyStatus\":\"1\",\"changeStatus\":\"2\",\"enable\":0,\"updateTime\":\"2019-07-10 13:27:01\",\"updateId\":1,\"createTime\":\"\",\"createId\":0}', 1, '2019-07-10 13:27:01', 1, '2019-07-10 13:27:01', 1);
INSERT INTO `sys_log` VALUES (11839, 2, '系统配置表', 'sys_config', 67, '删除', '{\"id\":67,\"name\":\"\",\"key\":\"\",\"value\":\"\",\"code\":\"\",\"dataType\":0,\"parentId\":0,\"parentKey\":\"\",\"sort\":0,\"projectId\":0,\"copyStatus\":\"\",\"changeStatus\":\"\",\"enable\":0,\"updateTime\":\"2019-07-10 13:27:12\",\"updateId\":1,\"createTime\":\"\",\"createId\":0}', 1, '2019-07-10 13:27:12', 1, '2019-07-10 13:27:12', 1);
INSERT INTO `sys_log` VALUES (11840, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-07-10 19:00:32', 1, '2019-07-10 19:00:32', 1);
INSERT INTO `sys_log` VALUES (11841, 2, '组织机构', 'sys_department', 10016, '插入', '{\"id\":10016,\"parentId\":10001,\"name\":\"123\",\"code\":\"123\",\"sort\":10,\"linkman\":\"123\",\"linkmanNo\":\"123\",\"remark\":\"\",\"enable\":0,\"updateTime\":\"2019-07-10 19:00:44\",\"updateId\":1,\"createTime\":\"2019-07-10 19:00:44\",\"createId\":1}', 1, '2019-07-10 19:00:44', 1, '2019-07-10 19:00:44', 1);
INSERT INTO `sys_log` VALUES (11842, 2, '组织机构', 'sys_department', 10016, '更新', '{\"id\":10016,\"parentId\":10001,\"name\":\"1234\",\"code\":\"1234\",\"sort\":104,\"linkman\":\"1234\",\"linkmanNo\":\"1234\",\"remark\":\"\",\"enable\":0,\"updateTime\":\"2019-07-10 19:00:52\",\"updateId\":1,\"createTime\":\"\",\"createId\":0}', 1, '2019-07-10 19:00:52', 1, '2019-07-10 19:00:52', 1);
INSERT INTO `sys_log` VALUES (11843, 2, '组织机构', 'sys_department', 10016, '删除', '{\"id\":10016,\"parentId\":0,\"name\":\"\",\"code\":\"\",\"sort\":0,\"linkman\":\"\",\"linkmanNo\":\"\",\"remark\":\"\",\"enable\":0,\"updateTime\":\"2019-07-10 19:00:56\",\"updateId\":1,\"createTime\":\"\",\"createId\":0}', 1, '2019-07-10 19:00:56', 1, '2019-07-10 19:00:56', 1);
INSERT INTO `sys_log` VALUES (11844, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-07-10 21:21:19', 1, '2019-07-10 21:21:19', 1);
INSERT INTO `sys_log` VALUES (11845, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-07-10 21:24:49', 1, '2019-07-10 21:24:49', 1);
INSERT INTO `sys_log` VALUES (11846, 2, '用户', 'sys_user', 8, '插入', '{\"id\":8,\"uuid\":\"81lEDM63TuH06KT6QA88z89009TAp1Uy\",\"username\":\"123\",\"password\":\"4aa3332c38ae75866ae594abcc3433f5\",\"salt\":\"wOO1hQ\",\"realName\":\"123\",\"departId\":10002,\"userType\":2,\"status\":0,\"thirdid\":\"\",\"endtime\":\"\",\"email\":\"123\",\"tel\":\"123\",\"address\":\"\",\"titleUrl\":\"\",\"remark\":\"123\",\"theme\":\"\",\"backSiteId\":0,\"createSiteId\":0,\"projectId\":0,\"projectName\":\"\",\"departName\":\"\",\"enable\":0,\"updateTime\":\"2019-07-10 21:25:22\",\"updateId\":1,\"createTime\":\"2019-07-10 21:25:22\",\"createId\":1}', 1, '2019-07-10 21:25:22', 1, '2019-07-10 21:25:22', 1);
INSERT INTO `sys_log` VALUES (11847, 2, '用户', 'sys_user', 8, '更新', '{\"id\":8,\"uuid\":\"81lEDM63TuH06KT6QA88z89009TAp1Uy\",\"username\":\"1234\",\"password\":\"4aa3332c38ae75866ae594abcc3433f5\",\"salt\":\"wOO1hQ\",\"realName\":\"1234\",\"departId\":10002,\"userType\":2,\"status\":10,\"thirdid\":\"\",\"endtime\":\"\",\"email\":\"1234\",\"tel\":\"1234\",\"address\":\"\",\"titleUrl\":\"\",\"remark\":\"1234\",\"theme\":\"default\",\"backSiteId\":0,\"createSiteId\":1,\"projectId\":0,\"projectName\":\"\",\"departName\":\"\",\"enable\":0,\"updateTime\":\"2019-07-10 21:25:31\",\"updateId\":1,\"createTime\":\"\",\"createId\":0}', 1, '2019-07-10 21:25:31', 1, '2019-07-10 21:25:31', 1);
INSERT INTO `sys_log` VALUES (11848, 2, '用户和角色关联', 'sys_user_role', 0, '删除', '{\"id\":0,\"userId\":8,\"roleId\":0,\"enable\":0,\"updateTime\":\"\",\"updateId\":0,\"createTime\":\"\",\"createId\":0}', 1, NULL, 0, NULL, 0);
INSERT INTO `sys_log` VALUES (12076, 1, '用户', 'sys_user', 11, '登出', NULL, 1, '2019-09-29 17:40:11', 11, '2019-09-29 17:40:11', 11);
INSERT INTO `sys_log` VALUES (12077, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-10-08 11:01:13', 1, '2019-10-08 11:01:13', 1);
INSERT INTO `sys_log` VALUES (12078, 2, '组织机构', 'sys_department', 10017, '插入', '{\"id\":10017,\"parentId\":10002,\"name\":\"123\",\"code\":\"123\",\"sort\":123,\"linkman\":\"123\",\"linkmanNo\":\"123\",\"remark\":\"123\",\"parentName\":\"\",\"enable\":0,\"updateTime\":\"2019-10-08 11:01:29\",\"updateId\":1,\"createTime\":\"2019-10-08 11:01:29\",\"createId\":1}', 1, '2019-10-08 11:01:29', 1, '2019-10-08 11:01:29', 1);
INSERT INTO `sys_log` VALUES (12079, 2, '组织机构', 'sys_department', 10017, '更新', '{\"id\":10017,\"parentId\":10002,\"name\":\"1234\",\"code\":\"1234\",\"sort\":1234,\"linkman\":\"1234\",\"linkmanNo\":\"1234\",\"remark\":\"1234\",\"parentName\":\"\",\"enable\":0,\"updateTime\":\"2019-10-08 11:01:38\",\"updateId\":1,\"createTime\":\"\",\"createId\":0}', 1, '2019-10-08 11:01:38', 1, '2019-10-08 11:01:38', 1);
INSERT INTO `sys_log` VALUES (12080, 2, '组织机构', 'sys_department', 10017, '删除', '{\"id\":10017,\"parentId\":0,\"name\":\"\",\"code\":\"\",\"sort\":0,\"linkman\":\"\",\"linkmanNo\":\"\",\"remark\":\"\",\"parentName\":\"\",\"enable\":0,\"updateTime\":\"2019-10-08 11:01:46\",\"updateId\":1,\"createTime\":\"\",\"createId\":0}', 1, '2019-10-08 11:01:46', 1, '2019-10-08 11:01:46', 1);
INSERT INTO `sys_log` VALUES (12081, 2, '用户', 'sys_user', 7, '删除', '{\"id\":7,\"uuid\":\"\",\"username\":\"\",\"password\":\"\",\"salt\":\"\",\"realName\":\"\",\"departId\":0,\"userType\":0,\"status\":0,\"thirdid\":\"\",\"endtime\":\"\",\"email\":\"\",\"tel\":\"\",\"address\":\"\",\"titleUrl\":\"\",\"remark\":\"\",\"theme\":\"\",\"backSiteId\":0,\"createSiteId\":0,\"projectId\":0,\"projectName\":\"\",\"departName\":\"\",\"enable\":0,\"updateTime\":\"2019-10-08 11:02:12\",\"updateId\":1,\"createTime\":\"\",\"createId\":0}', 1, '2019-10-08 11:02:12', 1, '2019-10-08 11:02:12', 1);
INSERT INTO `sys_log` VALUES (12082, 2, '用户', 'sys_user', 9, '删除', '{\"id\":9,\"uuid\":\"\",\"username\":\"\",\"password\":\"\",\"salt\":\"\",\"realName\":\"\",\"departId\":0,\"userType\":0,\"status\":0,\"thirdid\":\"\",\"endtime\":\"\",\"email\":\"\",\"tel\":\"\",\"address\":\"\",\"titleUrl\":\"\",\"remark\":\"\",\"theme\":\"\",\"backSiteId\":0,\"createSiteId\":0,\"projectId\":0,\"projectName\":\"\",\"departName\":\"\",\"enable\":0,\"updateTime\":\"2019-10-08 11:02:30\",\"updateId\":1,\"createTime\":\"\",\"createId\":0}', 1, '2019-10-08 11:02:30', 1, '2019-10-08 11:02:30', 1);
INSERT INTO `sys_log` VALUES (12083, 2, '项目表', 'tb_project', 11, '更新', '{\"id\":11,\"name\":\"邮件系统\",\"secretKey\":\"8bb0b1082de459dd92bcfbd0ae659d37\",\"type\":\"1\",\"sort\":10,\"remark\":\"发送邮件平台\",\"enable\":1,\"updateTime\":\"2019-10-08 11:11:59\",\"updateId\":1,\"createTime\":\"2019-08-27 09:40:25\",\"createId\":1}', 1, '2019-10-08 11:11:59', 1, '2019-10-08 11:11:59', 1);
INSERT INTO `sys_log` VALUES (12084, 2, '项目表', 'tb_project', 10, '更新', '{\"id\":10,\"name\":\"crm\",\"secretKey\":\"4554595bb5ec950f214b73c29588cd1c\",\"type\":\"1\",\"sort\":10,\"remark\":\"客户关系管理\",\"enable\":2,\"updateTime\":\"2019-10-08 11:12:28\",\"updateId\":1,\"createTime\":\"2019-07-24 16:39:57\",\"createId\":1}', 1, '2019-10-08 11:12:28', 1, '2019-10-08 11:12:28', 1);
INSERT INTO `sys_log` VALUES (12085, 2, '项目表', 'tb_project', 11, '更新', '{\"id\":11,\"name\":\"mail\",\"secretKey\":\"8bb0b1082de459dd92bcfbd0ae659d37\",\"type\":\"1\",\"sort\":10,\"remark\":\"发送邮件平台\",\"enable\":1,\"updateTime\":\"2019-10-08 11:12:36\",\"updateId\":1,\"createTime\":\"2019-08-27 09:40:25\",\"createId\":1}', 1, '2019-10-08 11:12:36', 1, '2019-10-08 11:12:36', 1);
INSERT INTO `sys_log` VALUES (12086, 2, '项目表', 'tb_project', 1, '更新', '{\"id\":1,\"name\":\"test\",\"secretKey\":\"12345678\",\"type\":\"1\",\"sort\":1,\"remark\":\"测试\",\"enable\":1,\"updateTime\":\"2019-10-08 11:12:43\",\"updateId\":1,\"createTime\":\"2017-07-23 00:11:46\",\"createId\":1}', 1, '2019-10-08 11:12:43', 1, '2019-10-08 11:12:43', 1);
INSERT INTO `sys_log` VALUES (12087, 2, '项目表', 'tb_project', 11, '更新', '{\"id\":11,\"name\":\"mail\",\"secretKey\":\"8bb0b1082de459dd92bcfbd0ae659d37\",\"type\":\"1\",\"sort\":10,\"remark\":\"邮件系统\",\"enable\":1,\"updateTime\":\"2019-10-08 11:12:55\",\"updateId\":1,\"createTime\":\"2019-08-27 09:40:25\",\"createId\":1}', 1, '2019-10-08 11:12:55', 1, '2019-10-08 11:12:55', 1);
INSERT INTO `sys_log` VALUES (12088, 2, '配置发布内容表', 'tb_config_public', 395, '插入', '{\"id\":395,\"projectId\":1,\"projectName\":\"test\",\"version\":\"20191008111411398\",\"content\":\"[{\\\"code\\\":\\\"\\\",\\\"key\\\":\\\"sms.username\\\",\\\"name\\\":\\\"短信账号\\\",\\\"parentKey\\\":\\\"sms\\\",\\\"value\\\":\\\"test\\\"},{\\\"code\\\":\\\"\\\",\\\"key\\\":\\\"sms.type\\\",\\\"name\\\":\\\"短信类型\\\",\\\"parentKey\\\":\\\"sms\\\",\\\"value\\\":\\\"阿里云\\\"},{\\\"code\\\":\\\"\\\",\\\"key\\\":\\\"sms.passwd\\\",\\\"name\\\":\\\"短信密码\\\",\\\"parentKey\\\":\\\"sms\\\",\\\"value\\\":\\\"111111\\\"},{\\\"code\\\":\\\"\\\",\\\"key\\\":\\\"system.debug\\\",\\\"name\\\":\\\"日志控制配置\\\",\\\"parentKey\\\":\\\"system\\\",\\\"value\\\":\\\"false\\\"},{\\\"code\\\":\\\"\\\",\\\"key\\\":\\\"system\\\",\\\"name\\\":\\\"系统参数\\\",\\\"parentKey\\\":\\\"\\\",\\\"value\\\":\\\"\\\"},{\\\"code\\\":\\\"\\\",\\\"key\\\":\\\"sms\\\",\\\"name\\\":\\\"短信配置\\\",\\\"parentKey\\\":\\\"\\\",\\\"value\\\":\\\"\\\"},{\\\"code\\\":\\\"\\\",\\\"key\\\":\\\"sex\\\",\\\"name\\\":\\\"性别\\\",\\\"parentKey\\\":\\\"\\\",\\\"value\\\":\\\"\\\"},{\\\"code\\\":\\\"1\\\",\\\"key\\\":\\\"sex.male\\\",\\\"name\\\":\\\"性别男\\\",\\\"parentKey\\\":\\\"sex\\\",\\\"value\\\":\\\"男\\\"},{\\\"code\\\":\\\"2\\\",\\\"key\\\":\\\"sex.female\\\",\\\"name\\\":\\\"性别女\\\",\\\"parentKey\\\":\\\"sex\\\",\\\"value\\\":\\\"女\\\"},{\\\"code\\\":\\\"3\\\",\\\"key\\\":\\\"sex.unknown\\\",\\\"name\\\":\\\"性别未知\\\",\\\"parentKey\\\":\\\"sex\\\",\\\"value\\\":\\\"未知\\\"}]\",\"beforeContent\":\"\",\"enable\":0,\"updateTime\":\"2019-10-08 11:14:11\",\"updateId\":1,\"createTime\":\"2019-10-08 11:14:11\",\"createId\":1}', 1, '2019-10-08 11:14:11', 1, '2019-10-08 11:14:11', 1);
INSERT INTO `sys_log` VALUES (12089, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-10-08 11:21:13', 1, '2019-10-08 11:21:13', 1);
INSERT INTO `sys_log` VALUES (12090, 1, '用户', 'sys_user', 1, '登录', NULL, 1, '2019-10-08 11:29:07', 1, '2019-10-08 11:29:07', 1);

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int(11) NOT NULL DEFAULT 0 COMMENT '父id',
  `name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称/11111',
  `icon` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `urlkey` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单key',
  `url` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '链接地址',
  `perms` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '授权(多个用逗号分隔，如：user:list,user:create)',
  `status` int(11) NULL DEFAULT 1 COMMENT '状态//radio/2,隐藏,1,显示',
  `type` int(11) NULL DEFAULT 1 COMMENT '类型//select/1,目录,2,菜单,3,按钮',
  `sort` int(11) NULL DEFAULT 1 COMMENT '排序',
  `level` int(11) NULL DEFAULT 1 COMMENT '级别',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, 20, '系统首页', 'fa fa-home', 'home', '/admin/welcome.html', '', 1, 2, 10, 2, 1, '2019-02-17 23:24:28', 1, '2015-04-27 17:28:06', 1);
INSERT INTO `sys_menu` VALUES (2, 0, '系统管理', 'fa fa-institution', 'system_root', NULL, NULL, 1, 1, 190, 1, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:28:06', 1);
INSERT INTO `sys_menu` VALUES (3, 2, '组织机构', 'fa fa-users', 'department', '/system/department/index', NULL, 1, 2, 191, 2, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:28:25', 1);
INSERT INTO `sys_menu` VALUES (4, 2, '用户管理', 'fa fa-user-o', 'user', '/system/user/index', NULL, 1, 2, 192, 2, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:28:46', 1);
INSERT INTO `sys_menu` VALUES (5, 2, '角色管理', 'fa fa-address-book-o', 'role', '/system/role/index', NULL, 1, 2, 194, 2, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:29:13', 1);
INSERT INTO `sys_menu` VALUES (6, 2, '菜单管理', 'fa fa-bars', 'menu', '/system/menu/index', NULL, 1, 2, 196, 2, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:29:43', 1);
INSERT INTO `sys_menu` VALUES (8, 20, '参数配置', 'fa fa-file-text-o', 'config', '/system/config/index', '', 1, 2, 198, 2, 1, '2019-07-30 15:33:55', 1, '2016-12-17 23:34:13', 1);
INSERT INTO `sys_menu` VALUES (9, 2, '日志管理', 'fa fa-line-chart', 'log', '/system/log/index', NULL, 1, 2, 199, 2, 1, '2015-04-27 17:28:06', 1, '2016-01-03 18:09:18', 1);
INSERT INTO `sys_menu` VALUES (15, 20, '项目管理', 'fa fa-th-list', '', '/admin/project/index', '', 1, 2, 52, 2, 1, '2017-07-22 23:59:37', 1, '2017-07-22 23:58:32', 1);
INSERT INTO `sys_menu` VALUES (16, 20, '配置发布', 'fa fa-edit', '', '/admin/configpublic/index', '', 1, 2, 200, 2, 1, '2017-09-15 14:53:59', 1, '2017-07-22 23:59:27', 1);
INSERT INTO `sys_menu` VALUES (17, 20, '配置拷贝', 'fa fa-exchange', '', '/admin/project/copyindex', '', 2, 2, 210, 2, 1, '2019-08-27 11:44:52', 1, '2017-09-27 15:18:06', 5);
INSERT INTO `sys_menu` VALUES (20, 0, '业务处理', 'fa fa-home', 'home', '', '', 1, 1, 10, 1, 1, '2019-02-17 23:24:08', 1, '2019-02-17 23:24:08', 1);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称/11111/',
  `status` int(11) NULL DEFAULT 1 COMMENT '状态//radio/2,隐藏,1,显示',
  `sort` int(11) NULL DEFAULT 1 COMMENT '排序',
  `remark` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '说明//textarea',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '项目管理角色', 1, 10, '', 1, '2019-07-30 16:16:06', 1, '2017-09-15 14:54:26', 1);
INSERT INTO `sys_role` VALUES (2, '配置管理角色', 1, 10, NULL, 1, '2019-07-30 16:16:31', 1, '2019-07-03 00:31:25', 1);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 88 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色和菜单关联' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (59, 3, 20);
INSERT INTO `sys_role_menu` VALUES (60, 3, 1);
INSERT INTO `sys_role_menu` VALUES (61, 3, 2);
INSERT INTO `sys_role_menu` VALUES (62, 3, 3);
INSERT INTO `sys_role_menu` VALUES (63, 3, 4);
INSERT INTO `sys_role_menu` VALUES (64, 3, 5);
INSERT INTO `sys_role_menu` VALUES (65, 3, 6);
INSERT INTO `sys_role_menu` VALUES (66, 3, 8);
INSERT INTO `sys_role_menu` VALUES (67, 3, 9);
INSERT INTO `sys_role_menu` VALUES (79, 1, 20);
INSERT INTO `sys_role_menu` VALUES (80, 1, 1);
INSERT INTO `sys_role_menu` VALUES (81, 1, 15);
INSERT INTO `sys_role_menu` VALUES (82, 1, 8);
INSERT INTO `sys_role_menu` VALUES (83, 1, 16);
INSERT INTO `sys_role_menu` VALUES (84, 1, 17);
INSERT INTO `sys_role_menu` VALUES (85, 2, 20);
INSERT INTO `sys_role_menu` VALUES (86, 2, 8);
INSERT INTO `sys_role_menu` VALUES (87, 2, 16);

-- ----------------------------
-- Table structure for sys_role_project
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_project`;
CREATE TABLE `sys_role_project`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  `project_id` bigint(20) NOT NULL COMMENT '项目id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色和项目关联' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_role_project
-- ----------------------------
INSERT INTO `sys_role_project` VALUES (2, 1, 1);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'UUID',
  `username` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录名/11111',
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `salt` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '1111' COMMENT '密码盐',
  `real_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `depart_id` int(11) NULL DEFAULT 0 COMMENT '部门/11111/dict',
  `user_type` int(11) NULL DEFAULT 2 COMMENT '类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户',
  `status` int(11) NULL DEFAULT 10 COMMENT '状态',
  `thirdid` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '第三方ID',
  `endtime` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '结束时间',
  `email` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'email',
  `tel` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `address` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地址',
  `title_url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像地址',
  `remark` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '说明',
  `theme` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'default' COMMENT '主题',
  `back_site_id` int(11) NULL DEFAULT 0 COMMENT '后台选择站点ID',
  `create_site_id` int(11) NULL DEFAULT 1 COMMENT '创建站点ID',
  `project_id` bigint(20) NULL DEFAULT 0 COMMENT '项目ID',
  `project_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '项目名称',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_user_username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, '94091b1fa6ac4a27a06c0b92155aea6a', 'admin', '9fb3dc842c899aa63d6944a55080b795', '1111', '系统管理员', 10001, 1, 10, '', '', 'zcool321@sina.com', '123', '', '', '时间是最好的老师，但遗憾的是&mdash;&mdash;最后他把所有的学生都弄死了', 'flat', 5, 1, 1, 'test', 1, '2019-08-27 09:40:12', 1, '2017-03-19 20:41:25', 1);
INSERT INTO `sys_user` VALUES (10, 'lcH3HObj93ELuT121r84A7ZQGMkW42j4', 'testLogin', '86a3a78c459e1e7dabe219dc921eaa30', '1a187n', '登录测试', 10002, 2, 10, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'default', 0, 1, 0, NULL, 1, '2019-09-29 17:12:17', 1, '2019-09-29 17:12:17', 1);
INSERT INTO `sys_user` VALUES (11, '51N697z90R60he2x63d6Cm96009XT2G7', 'testLogout', 'dd441ac58f01361d502a1b984143cd7d', '1801rb', '登出测试', 10002, 2, 10, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'default', 0, 1, 0, NULL, 1, '2019-09-29 17:12:44', 1, '2019-09-29 17:12:34', 1);
INSERT INTO `sys_user` VALUES (12, '560P333pV2yB3akw22Y2yDwx4z6a2UO9', 'flyfox', '1869d9d7eadddadc2c2d6e86e870a677', '6P53Ni', '测试账户', 10002, 2, 10, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'default', 0, 1, 0, NULL, 1, '2019-09-29 17:13:02', 1, '2019-09-29 17:13:02', 1);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户和角色关联' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (6, 5, 1);
INSERT INTO `sys_user_role` VALUES (7, 5, 2);
INSERT INTO `sys_user_role` VALUES (11, 4, 2);
INSERT INTO `sys_user_role` VALUES (12, 6, 1);
INSERT INTO `sys_user_role` VALUES (13, 6, 2);
INSERT INTO `sys_user_role` VALUES (14, 7, 1);
INSERT INTO `sys_user_role` VALUES (15, 7, 2);
INSERT INTO `sys_user_role` VALUES (21, 9, 1);

-- ----------------------------
-- Table structure for tb_config_public
-- ----------------------------
DROP TABLE IF EXISTS `tb_config_public`;
CREATE TABLE `tb_config_public`  (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `project_id` int(11) NULL DEFAULT NULL COMMENT '项目ID',
  `project_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '项目名称',
  `version` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '版本',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '内容',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` int(11) NULL DEFAULT NULL COMMENT '更新人',
  `create_time` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` int(11) NULL DEFAULT NULL COMMENT '创建人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 396 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '配置发布内容表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of tb_config_public
-- ----------------------------
INSERT INTO `tb_config_public` VALUES (393, 1, 'test', '20190827142008716', '[{\"id\":50,\"name\":\"短信账号\",\"key\":\"sms.username\",\"value\":\"test\",\"code\":\"\",\"dataType\":0,\"parentId\":47,\"parentKey\":\"sms\",\"sort\":10,\"projectId\":1,\"copyStatus\":0,\"changeStatus\":0,\"enable\":0,\"updateTime\":\"\",\"updateId\":0,\"createTime\":\"\",\"createId\":0},{\"id\":52,\"name\":\"短信类型\",\"key\":\"sms.type\",\"value\":\"阿里云\",\"code\":\"\",\"dataType\":0,\"parentId\":47,\"parentKey\":\"sms\",\"sort\":10,\"projectId\":1,\"copyStatus\":0,\"changeStatus\":0,\"enable\":0,\"updateTime\":\"\",\"updateId\":0,\"createTime\":\"\",\"createId\":0},{\"id\":51,\"name\":\"短信密码\",\"key\":\"sms.passwd\",\"value\":\"111111\",\"code\":\"\",\"dataType\":0,\"parentId\":47,\"parentKey\":\"sms\",\"sort\":10,\"projectId\":1,\"copyStatus\":0,\"changeStatus\":0,\"enable\":0,\"updateTime\":\"\",\"updateId\":0,\"createTime\":\"\",\"createId\":0},{\"id\":46,\"name\":\"日志控制配置\",\"key\":\"system.debug\",\"value\":\"false\",\"code\":\"\",\"dataType\":0,\"parentId\":24,\"parentKey\":\"system\",\"sort\":15,\"projectId\":1,\"copyStatus\":0,\"changeStatus\":0,\"enable\":0,\"updateTime\":\"\",\"updateId\":0,\"createTime\":\"\",\"createId\":0},{\"id\":24,\"name\":\"系统参数\",\"key\":\"system\",\"value\":\"\",\"code\":\"\",\"dataType\":0,\"parentId\":0,\"parentKey\":\"\",\"sort\":15,\"projectId\":1,\"copyStatus\":0,\"changeStatus\":0,\"enable\":0,\"updateTime\":\"\",\"updateId\":0,\"createTime\":\"\",\"createId\":0},{\"id\":47,\"name\":\"短信配置\",\"key\":\"sms\",\"value\":\"\",\"code\":\"\",\"dataType\":0,\"parentId\":0,\"parentKey\":\"\",\"sort\":15,\"projectId\":1,\"copyStatus\":0,\"changeStatus\":0,\"enable\":0,\"updateTime\":\"\",\"updateId\":0,\"createTime\":\"\",\"createId\":0},{\"id\":53,\"name\":\"性别\",\"key\":\"sex\",\"value\":\"\",\"code\":\"\",\"dataType\":0,\"parentId\":0,\"parentKey\":\"\",\"sort\":90,\"projectId\":1,\"copyStatus\":0,\"changeStatus\":0,\"enable\":0,\"updateTime\":\"\",\"updateId\":0,\"createTime\":\"\",\"createId\":0},{\"id\":54,\"name\":\"性别男\",\"key\":\"sex.male\",\"value\":\"男\",\"code\":\"1\",\"dataType\":0,\"parentId\":53,\"parentKey\":\"sex\",\"sort\":91,\"projectId\":1,\"copyStatus\":0,\"changeStatus\":0,\"enable\":0,\"updateTime\":\"\",\"updateId\":0,\"createTime\":\"\",\"createId\":0},{\"id\":55,\"name\":\"性别女\",\"key\":\"sex.female\",\"value\":\"女\",\"code\":\"2\",\"dataType\":0,\"parentId\":53,\"parentKey\":\"sex\",\"sort\":92,\"projectId\":1,\"copyStatus\":0,\"changeStatus\":0,\"enable\":0,\"updateTime\":\"\",\"updateId\":0,\"createTime\":\"\",\"createId\":0},{\"id\":56,\"name\":\"性别未知\",\"key\":\"sex.unknown\",\"value\":\"未知\",\"code\":\"3\",\"dataType\":0,\"parentId\":53,\"parentKey\":\"sex\",\"sort\":93,\"projectId\":1,\"copyStatus\":0,\"changeStatus\":0,\"enable\":0,\"updateTime\":\"\",\"updateId\":0,\"createTime\":\"\",\"createId\":0}]', 1, '2019-08-27 14:20:08', 1, '2019-08-27 14:20:08', 1);
INSERT INTO `tb_config_public` VALUES (394, 1, 'test', '20190827143220029', '[{\"code\":\"\",\"key\":\"sms.username\",\"name\":\"短信账号\",\"parentKey\":\"sms\",\"value\":\"test\"},{\"code\":\"\",\"key\":\"sms.type\",\"name\":\"短信类型\",\"parentKey\":\"sms\",\"value\":\"阿里云\"},{\"code\":\"\",\"key\":\"sms.passwd\",\"name\":\"短信密码\",\"parentKey\":\"sms\",\"value\":\"111111\"},{\"code\":\"\",\"key\":\"system.debug\",\"name\":\"日志控制配置\",\"parentKey\":\"system\",\"value\":\"false\"},{\"code\":\"\",\"key\":\"system\",\"name\":\"系统参数\",\"parentKey\":\"\",\"value\":\"\"},{\"code\":\"\",\"key\":\"sms\",\"name\":\"短信配置\",\"parentKey\":\"\",\"value\":\"\"},{\"code\":\"\",\"key\":\"sex\",\"name\":\"性别\",\"parentKey\":\"\",\"value\":\"\"},{\"code\":\"1\",\"key\":\"sex.male\",\"name\":\"性别男\",\"parentKey\":\"sex\",\"value\":\"男\"},{\"code\":\"2\",\"key\":\"sex.female\",\"name\":\"性别女\",\"parentKey\":\"sex\",\"value\":\"女\"},{\"code\":\"3\",\"key\":\"sex.unknown\",\"name\":\"性别未知\",\"parentKey\":\"sex\",\"value\":\"未知\"}]', 1, '2019-08-27 14:32:20', 1, '2019-08-27 14:32:20', 1);
INSERT INTO `tb_config_public` VALUES (395, 1, 'test', '20191008111411398', '[{\"code\":\"\",\"key\":\"sms.username\",\"name\":\"短信账号\",\"parentKey\":\"sms\",\"value\":\"test\"},{\"code\":\"\",\"key\":\"sms.type\",\"name\":\"短信类型\",\"parentKey\":\"sms\",\"value\":\"阿里云\"},{\"code\":\"\",\"key\":\"sms.passwd\",\"name\":\"短信密码\",\"parentKey\":\"sms\",\"value\":\"111111\"},{\"code\":\"\",\"key\":\"system.debug\",\"name\":\"日志控制配置\",\"parentKey\":\"system\",\"value\":\"false\"},{\"code\":\"\",\"key\":\"system\",\"name\":\"系统参数\",\"parentKey\":\"\",\"value\":\"\"},{\"code\":\"\",\"key\":\"sms\",\"name\":\"短信配置\",\"parentKey\":\"\",\"value\":\"\"},{\"code\":\"\",\"key\":\"sex\",\"name\":\"性别\",\"parentKey\":\"\",\"value\":\"\"},{\"code\":\"1\",\"key\":\"sex.male\",\"name\":\"性别男\",\"parentKey\":\"sex\",\"value\":\"男\"},{\"code\":\"2\",\"key\":\"sex.female\",\"name\":\"性别女\",\"parentKey\":\"sex\",\"value\":\"女\"},{\"code\":\"3\",\"key\":\"sex.unknown\",\"name\":\"性别未知\",\"parentKey\":\"sex\",\"value\":\"未知\"}]', 1, '2019-10-08 11:14:11', 1, '2019-10-08 11:14:11', 1);

-- ----------------------------
-- Table structure for tb_project
-- ----------------------------
DROP TABLE IF EXISTS `tb_project`;
CREATE TABLE `tb_project`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '项目名称',
  `secret_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '12345678' COMMENT '秘钥',
  `type` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '类型',
  `sort` int(11) NULL DEFAULT 100 COMMENT '排序',
  `remark` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` int(11) NULL DEFAULT NULL COMMENT '更新人',
  `create_time` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` int(11) NULL DEFAULT NULL COMMENT '创建人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_project_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of tb_project
-- ----------------------------
INSERT INTO `tb_project` VALUES (1, 'test', '12345678', '1', 1, '测试', 1, '2019-10-08 11:12:43', 1, '2017-07-23 00:11:46', 1);
INSERT INTO `tb_project` VALUES (10, 'crm', '4554595bb5ec950f214b73c29588cd1c', '1', 10, '客户关系管理', 2, '2019-10-08 11:12:28', 1, '2019-07-24 16:39:57', 1);
INSERT INTO `tb_project` VALUES (11, 'mail', '8bb0b1082de459dd92bcfbd0ae659d37', '1', 10, '邮件系统', 1, '2019-10-08 11:12:55', 1, '2019-08-27 09:40:25', 1);

SET FOREIGN_KEY_CHECKS = 1;
