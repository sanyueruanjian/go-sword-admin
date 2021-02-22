/*
 Navicat Premium Data Transfer

 Source Server         : lbl_mysql
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : 121.196.160.185:3306
 Source Schema         : go_sword

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 22/02/2021 23:28:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique_index`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (29, 'p', '1', '11', 'get', '', '', '');
INSERT INTO `casbin_rule` VALUES (11, 'p', '10', '11', 'get', '', '', '');
INSERT INTO `casbin_rule` VALUES (8, 'p', '12', '11', 'get', '', '', '');
INSERT INTO `casbin_rule` VALUES (10, 'p', '14', '11', 'get', '', '', '');
INSERT INTO `casbin_rule` VALUES (33, 'p', '9', '11', 'get', '', '', '');

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级部门（顶级部门为0，默认为0）',
  `sub_count` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '子部门数目',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `dept_sort` int(10) UNSIGNED NOT NULL DEFAULT 999 COMMENT '排序',
  `enabled` bit(1) NOT NULL DEFAULT b'1' COMMENT '状态：1启用（默认）、0禁用',
  `create_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者',
  `create_time` bigint(20) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '状态：1启用（默认）、0禁用',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_pid`(`pid`) USING BTREE COMMENT '普通索引——pid查询部门',
  INDEX `idx_enabled`(`enabled`) USING BTREE COMMENT '普通索引——enabled查询部门'
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (2, 7, 0, '研发部', 3, b'1', 1, 1, 20201213114104, 20201213114106, b'0');
INSERT INTO `sys_dept` VALUES (5, 7, 0, '运维部', 4, b'1', 1, 1, 20201213114136, 20201213114138, b'0');
INSERT INTO `sys_dept` VALUES (6, 8, 0, '测试部', 0, b'1', 1, 1, 20201213114217, 20201213114219, b'0');
INSERT INTO `sys_dept` VALUES (7, 0, 5, '华北分部', 5, b'1', 1, 1, 20201213114258, 1613912934520, b'0');
INSERT INTO `sys_dept` VALUES (8, 0, 2, '华南分部', 7, b'1', 1, 1, 20201213114340, 1613913600273, b'0');
INSERT INTO `sys_dept` VALUES (17, 7, 0, '哈哈哈', 999, b'1', 1, 1, 1613634683926, 1613634683926, b'0');
INSERT INTO `sys_dept` VALUES (18, 8, 0, '哈哈哈', 999, b'1', 1, 1, 1613637943931, 1613886300610, b'0');
INSERT INTO `sys_dept` VALUES (19, 7, 0, '运维部', 999, b'1', 1, 1, 1613640866701, 1613640866701, b'0');
INSERT INTO `sys_dept` VALUES (20, 8, 0, '测试删除', 999, b'1', 1, 1, 1613886421503, 1613887708779, b'1');
INSERT INTO `sys_dept` VALUES (21, 0, 0, '华北分部', 999, b'1', 1, 1, 1613896727952, 1613907290171, b'1');
INSERT INTO `sys_dept` VALUES (22, 0, 0, '哈贝', 999, b'1', 1, 1, 1613907616874, 1613907805818, b'1');
INSERT INTO `sys_dept` VALUES (23, 0, 0, '哈哈哈', 999, b'1', 1, 1, 1613908703795, 1613908836959, b'1');
INSERT INTO `sys_dept` VALUES (24, 0, 0, '华南', 999, b'1', 1, 1, 1613910111447, 1613910118496, b'1');
INSERT INTO `sys_dept` VALUES (25, 0, 0, '华北分部', 999, b'1', 1, 1, 1613912973123, 1613912993826, b'1');

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '岗位名称',
  `enabled` bit(1) NOT NULL DEFAULT b'1' COMMENT '状态：1启用（默认）、0禁用',
  `job_sort` int(10) UNSIGNED NOT NULL DEFAULT 999 COMMENT '排序',
  `create_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者id',
  `update_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者id',
  `create_time` bigint(20) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '软删除（默认值为0，1为删除）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unq_name`(`name`) USING BTREE COMMENT '岗位名唯一',
  INDEX `idx_enabled`(`enabled`) USING BTREE COMMENT '普通索引——enabled查询岗位'
) ENGINE = InnoDB AUTO_INCREMENT = 42 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '岗位' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_job
-- ----------------------------
INSERT INTO `sys_job` VALUES (1, '人事', b'1', 36, 1, 1, 1613390193132, 1613634533409, b'1');
INSERT INTO `sys_job` VALUES (2, '产品', b'1', 46, 1, 1, 1613390193132, 1613823200099, b'0');
INSERT INTO `sys_job` VALUES (21, 'CEO', b'1', 990, 1, 1, 1613637998070, 1613638006922, b'1');
INSERT INTO `sys_job` VALUES (28, 'CEO1', b'1', 999, 1, 1, 1613638067063, 1613638067063, b'1');
INSERT INTO `sys_job` VALUES (29, '漫威', b'1', 998, 1, 1, 1613638251200, 1613885304524, b'0');
INSERT INTO `sys_job` VALUES (31, 'sad', b'1', 999, 1, 1, 1613787679518, 1613787679518, b'1');
INSERT INTO `sys_job` VALUES (32, 'fg', b'1', 999, 1, 1, 1613790242141, 1613791281482, b'1');
INSERT INTO `sys_job` VALUES (33, 'Maven', b'1', 999, 1, 1, 1613823212597, 1613823217773, b'1');
INSERT INTO `sys_job` VALUES (34, '哈哈哈', b'1', 999, 1, 1, 1613896979011, 1613896983775, b'1');

-- ----------------------------
-- Table structure for sys_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_log`;
CREATE TABLE `sys_log`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志id',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作用户id',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `log_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '日志类型',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '方法名',
  `params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '参数',
  `request_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求ip',
  `request_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '请求耗时（毫秒值）',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址',
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `exception_detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '详细异常',
  `create_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人id',
  `update_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新人id',
  `create_time` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '软删除（默认值为0，1为删除）',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_create_time`(`create_time`) USING BTREE COMMENT '普通索引——根据创建时间查询日志',
  INDEX `idx_log_type`(`log_type`) USING BTREE COMMENT '普通索引——根据日志类型查询日志'
) ENGINE = InnoDB AUTO_INCREMENT = 159 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求路由',
  `action` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方式',
  `pid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级菜单ID',
  `sub_count` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '子菜单数目',
  `type` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单类型',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单标题',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '组件名称',
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '组件',
  `menu_sort` int(10) UNSIGNED NOT NULL DEFAULT 999 COMMENT '排序',
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '链接地址',
  `i_frame` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否外链',
  `cache` bit(1) NOT NULL DEFAULT b'0' COMMENT '缓存',
  `hidden` bit(1) NOT NULL DEFAULT b'0' COMMENT '隐藏',
  `permission` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限',
  `create_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  `create_time` bigint(20) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `is_deleted` bit(1) NOT NULL DEFAULT b'0',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_pid`(`pid`) USING BTREE COMMENT '普通索引——pid查询菜单'
) ENGINE = InnoDB AUTO_INCREMENT = 179 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, NULL, NULL, 0, 7, 0, '系统管理', '系统管理', 'Layout', 1, 'system', 'system', b'0', b'0', b'0', '', 1, 1, 20181218151129, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (2, '11', 'get', 1, 3, 1, '用户管理', 'User', 'system/user/index', 2, 'peoples', 'user', b'0', b'0', b'0', 'user:list', 1, 1, 20181218151444, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (3, NULL, NULL, 1, 3, 1, '角色管理', 'Role', 'system/role/index', 3, 'role', 'role', b'0', b'0', b'0', 'roles:list', 1, 1, 20181218151607, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (5, NULL, NULL, 1, 3, 1, '菜单管理', 'Menu', 'system/menu/index', 5, 'menu', 'menu', b'0', b'0', b'0', 'menu:list', 1, 1, 20181218151728, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (6, NULL, NULL, 0, 4, 0, '系统监控', '系统监控', 'Layout', 10, 'monitor', 'monitor', b'0', b'0', b'0', '', 1, 1, 20181218151748, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (7, NULL, NULL, 6, 0, 1, '操作日志', 'Log', 'monitor/log/index', 11, 'log', 'logs', b'0', b'0', b'0', '', 1, 1, 20181218151826, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (28, NULL, NULL, 1, 3, 1, '任务调度', 'Timing', 'system/timing/index', 999, 'timing', 'timing', b'0', b'0', b'0', 'timing:list', 1, 1, 20190107203440, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (32, NULL, NULL, 6, 0, 1, '异常日志', 'ErrorLog', 'monitor/log/errorLog', 12, 'error', 'errorLog', b'0', b'0', b'0', '', 1, 1, 20190113134903, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (35, NULL, NULL, 1, 3, 1, '部门管理', 'Dept', 'system/dept/index', 6, 'dept', 'dept', b'0', b'0', b'0', 'dept:list', 1, 1, 20190325094600, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (37, NULL, NULL, 1, 3, 1, '岗位管理', 'Job', 'system/job/index', 7, 'Steve-Jobs', 'job', b'0', b'0', b'0', 'job:list', 1, 1, 20190329135118, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (39, NULL, NULL, 1, 3, 1, '字典管理', 'Dict', 'system/dict/index', 8, 'dictionary', 'dict', b'0', b'0', b'0', 'dict:list', 1, 1, 20190410114904, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (41, NULL, NULL, 6, 0, 1, '在线用户', 'OnlineUser', 'monitor/online/index', 10, 'Steve-Jobs', 'online', b'0', b'0', b'0', '', 1, 1, 20191026220843, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (44, NULL, 'POST', 2, 0, 2, '用户新增', '', '', 2, '', '', b'0', b'0', b'0', 'user:add', 1, 1, 20191029105946, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (45, NULL, NULL, 2, 0, 2, '用户编辑', '', '', 3, '', '', b'0', b'0', b'0', 'user:edit', 1, 1, 20191029110008, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (46, NULL, NULL, 2, 0, 2, '用户删除', '', '', 4, '', '', b'0', b'0', b'0', 'user:del', 1, 1, 20191029110023, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (48, NULL, NULL, 3, 0, 2, '角色创建', '', '', 2, '', '', b'0', b'0', b'0', 'roles:add', 1, 1, 20191029124534, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (49, NULL, NULL, 3, 0, 2, '角色修改', '', '', 3, '', '', b'0', b'0', b'0', 'roles:edit', 1, 1, 20191029124616, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (50, NULL, NULL, 3, 0, 2, '角色删除', '', '', 4, '', '', b'0', b'0', b'0', 'roles:del', 1, 1, 20191029124651, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (52, NULL, NULL, 5, 0, 2, '菜单新增', '', '', 2, '', '', b'0', b'0', b'0', 'menu:add', 1, 1, 20191029125507, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (53, NULL, NULL, 5, 0, 2, '菜单编辑', '', '', 3, '', '', b'0', b'0', b'0', 'menu:edit', 1, 1, 20191029125540, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (54, NULL, NULL, 5, 0, 2, '菜单删除', '', '', 4, '', '', b'0', b'0', b'0', 'menu:del', 1, 1, 20191029125600, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (56, NULL, NULL, 35, 0, 2, '部门新增', '', '', 2, '', '', b'0', b'0', b'0', 'dept:add', 1, 1, 20191029125709, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (57, NULL, NULL, 35, 0, 2, '部门编辑', '', '', 3, '', '', b'0', b'0', b'0', 'dept:edit', 1, 1, 20191029125727, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (58, NULL, NULL, 35, 0, 2, '部门删除', '', '', 4, '', '', b'0', b'0', b'0', 'dept:del', 1, 1, 20191029125741, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (60, NULL, NULL, 37, 0, 2, '岗位新增', '', '', 2, '', '', b'0', b'0', b'0', 'job:add', 1, 1, 20191029125827, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (61, NULL, NULL, 37, 0, 2, '岗位编辑', '', '', 3, '', '', b'0', b'0', b'0', 'job:edit', 1, 1, 20191029125845, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (62, NULL, NULL, 37, 0, 2, '岗位删除', '', '', 4, '', '', b'0', b'0', b'0', 'job:del', 1, 1, 20191029125904, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (64, NULL, NULL, 39, 0, 2, '字典新增', '', '', 2, '', '', b'0', b'0', b'0', 'dict:add', 1, 1, 20191029130017, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (65, NULL, NULL, 39, 0, 2, '字典编辑', '', '', 3, '', '', b'0', b'0', b'0', 'dict:edit', 1, 1, 20191029130042, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (66, NULL, NULL, 39, 0, 2, '字典删除', '', '', 4, '', '', b'0', b'0', b'0', 'dict:del', 1, 1, 20191029130059, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (73, NULL, NULL, 28, 0, 2, '任务新增', '', '', 2, '', '', b'0', b'0', b'0', 'timing:add', 1, 1, 20191029130728, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (74, NULL, NULL, 28, 0, 2, '任务编辑', '', '', 3, '', '', b'0', b'0', b'0', 'timing:edit', 1, 1, 20191029130741, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (75, NULL, NULL, 28, 0, 2, '任务删除', '', '', 4, '', '', b'0', b'0', b'0', 'timing:del', 1, 1, 20191029130754, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (123, '', '', 0, 0, 0, '大多', '', 'Layout', 999, 'app', 'dd', b'0', b'0', b'0', '', 1, 1, 1613634690041, 1613634690041, b'1');
INSERT INTO `sys_menu` VALUES (125, '', '', 0, 0, 0, 'e31231', '312', 'Layout', 999, 'codeConsole', '312', b'0', b'0', b'0', '312', 1, 1, 1613639636372, 1613639636372, b'1');
INSERT INTO `sys_menu` VALUES (126, '', '', 0, 0, 0, '131', '', 'Layout', 999, 'Steve-Jobs', '3123', b'0', b'0', b'0', '', 1, 1, 1613641403168, 1613641403168, b'1');
INSERT INTO `sys_menu` VALUES (127, '', '', 0, 0, 0, '嘟嘟嘟嘟嘟', '', 'Layout', 999, 'Steve-Jobs', 'root', b'0', b'0', b'0', '', 1, 1, 1613641598407, 1613641598407, b'1');
INSERT INTO `sys_menu` VALUES (128, '', '', 0, 0, 0, '沙克斯', '', 'Layout', 999, 'Steve-Jobs', '/api/pictures', b'0', b'0', b'0', '', 1, 1, 1613822648011, 1613822648011, b'1');
INSERT INTO `sys_menu` VALUES (129, '', '', 128, 0, 0, '66666', '', 'Layout', 999, 'Steve-Jobs', '/api/group/search', b'0', b'0', b'0', '', 1, 1, 1613822667993, 1613822667993, b'1');
INSERT INTO `sys_menu` VALUES (130, '', '', 128, 0, 0, 'hello word', '', 'Layout', 999, 'alipay', '/api/pictures', b'0', b'0', b'0', '', 1, 1, 1613822742068, 1613822742068, b'1');
INSERT INTO `sys_menu` VALUES (131, '', '', 128, 0, 0, '大萨达所', '', 'Layout', 999, 'chart', 'root', b'0', b'0', b'0', '', 1, 1, 1613824933880, 1613824933880, b'1');
INSERT INTO `sys_menu` VALUES (132, '', '', 128, 0, 0, '大萨达所', '', 'Layout', 999, 'codeConsole', '大声道', b'0', b'0', b'0', '', 1, 1, 1613826622382, 1613826622382, b'1');
INSERT INTO `sys_menu` VALUES (133, '', '', 0, 1, 0, '就六角恐龙', '', '', 999, 'app', 'asdad', b'0', b'0', b'0', '', 1, 1, 1613826735668, 1613826735668, b'1');
INSERT INTO `sys_menu` VALUES (134, '', '', 128, 0, 0, '啊实打实多', '', 'Layout', 999, 'chart', '大萨达所', b'0', b'0', b'0', '', 1, 1, 1613826754637, 1613826754637, b'1');
INSERT INTO `sys_menu` VALUES (135, '', '', 133, 0, 0, '大法师的', '', 'Layout', 999, 'Steve-Jobs', '大幅度发', b'0', b'0', b'0', '', 1, 1, 1613827330050, 1613827330050, b'1');
INSERT INTO `sys_menu` VALUES (136, '', '', 0, 0, 0, '我是', '我是', 'Layout', 999, 'Steve-Jobs', 'root', b'0', b'0', b'0', '', 1, 1, 1613827554591, 1613827554591, b'1');
INSERT INTO `sys_menu` VALUES (137, '', '', 136, 0, 0, '大帅哥', '', 'Layout', 999, 'dashboard', 'spped', b'0', b'0', b'0', '', 1, 1, 1613827578752, 1613827578752, b'1');
INSERT INTO `sys_menu` VALUES (138, '', '', 136, 0, 0, '大帅哥', '', 'Layout', 999, 'anq', 'root', b'0', b'0', b'0', '', 1, 1, 1613827758794, 1613827758794, b'1');
INSERT INTO `sys_menu` VALUES (139, '', '', 136, 0, 0, '的境地', '', 'Layout', 999, 'Steve-Jobs', '的数据', b'0', b'0', b'0', '', 1, 1, 1613827778220, 1613827778220, b'1');
INSERT INTO `sys_menu` VALUES (140, '', '', 136, 0, 0, '打架了', '', 'Layout', 999, 'Steve-Jobs', 'root', b'0', b'0', b'0', '', 1, 1, 1613828092182, 1613828092182, b'1');
INSERT INTO `sys_menu` VALUES (141, '', '', 0, 0, 0, 'asdasd', '', 'Layout', 999, 'Steve-Jobs', 'dsad', b'0', b'0', b'0', '', 1, 1, 1613828185873, 1613828185873, b'1');
INSERT INTO `sys_menu` VALUES (142, '', '', 136, 0, 0, 'dasdasd', '', 'Layout', 999, 'anq', 'sadas', b'0', b'0', b'0', '', 1, 1, 1613828202498, 1613828202498, b'1');
INSERT INTO `sys_menu` VALUES (144, '', '', 136, 0, 0, 'sdad ', '', 'Layout', 999, '', 'asdasdas', b'0', b'0', b'0', '', 1, 1, 1613828703669, 1613828703669, b'1');
INSERT INTO `sys_menu` VALUES (145, '', '', 0, 1, 0, 'woshi ', '', '', 999, 'Steve-Jobs', 'sdaa', b'0', b'0', b'0', '', 1, 1, 1613828779103, 1613828779103, b'1');
INSERT INTO `sys_menu` VALUES (146, '', '', 145, 0, 0, 'adad', '', 'Layout', 999, 'backup', 'asdad', b'0', b'0', b'0', '', 1, 1, 1613828794030, 1613828794030, b'1');
INSERT INTO `sys_menu` VALUES (147, '', '', 145, 0, 0, 'dasd', '', 'Layout', 999, 'Steve-Jobs', 'dsadasd', b'0', b'0', b'0', '', 1, 1, 1613828866899, 1613828866899, b'1');
INSERT INTO `sys_menu` VALUES (148, '', '', 145, 0, 0, 'dasdas', '', 'Layout', 999, 'Steve-Jobs', 'dsadas', b'0', b'0', b'0', '', 1, 1, 1613829257381, 1613829257381, b'1');
INSERT INTO `sys_menu` VALUES (149, '', '', 0, 3, 0, '我是', '', 'Layout', 999, 'Steve-Jobs', 'root', b'0', b'0', b'0', '', 1, 1, 1613829354771, 1613829354771, b'0');
INSERT INTO `sys_menu` VALUES (150, '', '', 149, 0, 0, '啊实打实多', '', 'Layout', 999, 'Steve-Jobs', '啊实打实', b'0', b'0', b'0', '', 1, 1, 1613829426467, 1613829426467, b'1');
INSERT INTO `sys_menu` VALUES (151, '', '', 0, 0, 0, '啊实打实大所', '', 'Layout', 999, 'Steve-Jobs', ' asdasd', b'0', b'0', b'0', '', 1, 1, 1613829543929, 1613829543929, b'1');
INSERT INTO `sys_menu` VALUES (152, '', '', 151, 0, 0, '啊实打实大', '', 'Layout', 999, 'date', '啊实打实', b'0', b'0', b'0', '', 1, 1, 1613829563267, 1613829563267, b'1');
INSERT INTO `sys_menu` VALUES (153, '', '', 151, 0, 0, '大声道', '', 'Layout', 999, 'alipay', 'asdad', b'0', b'0', b'0', '', 1, 1, 1613829803990, 1613829803990, b'1');
INSERT INTO `sys_menu` VALUES (154, '', '', 151, 0, 0, 'dasd', '', 'Layout', 999, 'Steve-Jobs', 'asdasd', b'0', b'0', b'0', '', 1, 1, 1613829985777, 1613829985777, b'1');
INSERT INTO `sys_menu` VALUES (155, '', '', 149, 0, 0, 'dasd', '', 'Layout', 999, 'Steve-Jobs', 'dasd', b'0', b'0', b'0', '', 1, 1, 1613830157402, 1613830157402, b'1');
INSERT INTO `sys_menu` VALUES (156, '', '', 151, 0, 0, 'das', '', 'Layout', 999, 'Steve-Jobs', 'asda', b'0', b'0', b'0', '', 1, 1, 1613830379557, 1613830379557, b'1');
INSERT INTO `sys_menu` VALUES (157, '', '', 151, 0, 0, 'asdasd', '', 'Layout', 999, 'Steve-Jobs', 'asdas', b'0', b'0', b'0', '', 1, 1, 1613830573368, 1613830573368, b'1');
INSERT INTO `sys_menu` VALUES (158, '', '', 151, 0, 0, '打打', '', 'Layout', 999, 'Steve-Jobs', '大萨达所', b'0', b'0', b'0', '', 1, 1, 1613831125776, 1613831125776, b'1');
INSERT INTO `sys_menu` VALUES (159, '', '', 151, 0, 0, '大萨达多', '', 'Layout', 999, 'Steve-Jobs', '啊实打实', b'0', b'0', b'0', '', 1, 1, 1613831821137, 1613831821137, b'1');
INSERT INTO `sys_menu` VALUES (160, '', '', 0, 0, 0, '阿萨德饭', '', 'Layout', 999, 'Steve-Jobs', '阿道夫', b'0', b'0', b'0', '', 1, 1, 1613870362849, 1613870362849, b'0');
INSERT INTO `sys_menu` VALUES (161, '', '', 0, 0, 0, '啊实打实', '', 'Layout', 999, 'dashboard', 'sad', b'0', b'0', b'0', '', 1, 1, 1613877496131, 1613877496131, b'1');
INSERT INTO `sys_menu` VALUES (162, '', '', 0, 0, 1, '大声道', '阿萨德', 'Layout', 999, 'alipay', '阿萨德', b'0', b'0', b'0', '阿萨德', 1, 1, 1613877513440, 1613877513440, b'0');
INSERT INTO `sys_menu` VALUES (163, '', '', 0, 0, 0, '阿达', '', 'Layout', 999, 'app', 'root', b'0', b'0', b'0', '', 1, 1, 1613892977418, 1613892977418, b'0');
INSERT INTO `sys_menu` VALUES (164, '', '', 0, 0, 0, 'asdas', '', 'Layout', 999, 'Steve-Jobs', 'asdasd', b'0', b'0', b'0', '', 1, 1, 1613893543952, 1613893543952, b'0');
INSERT INTO `sys_menu` VALUES (165, '', '', 149, 0, 0, '我的哥哥', '', 'Layout', 999, 'Steve-Jobs', '嘟嘟嘟', b'0', b'0', b'0', '', 1, 1, 1613894383972, 1613894383972, b'0');
INSERT INTO `sys_menu` VALUES (166, '', '', 149, 0, 0, '啊实打实', '', 'Layout', 999, 'backup', '阿萨德', b'0', b'0', b'0', '', 1, 1, 1613894864659, 1613894864659, b'0');
INSERT INTO `sys_menu` VALUES (167, '', '', 149, 0, 0, '大帅哥', '', 'Layout', 999, 'app', 'pig', b'0', b'0', b'0', '', 1, 1, 1613895262187, 1613895262187, b'0');
INSERT INTO `sys_menu` VALUES (168, '', '', 0, 0, 0, '手打', '', 'Layout', 999, '', 'rootssss', b'0', b'0', b'0', '', 1, 1, 1613895288864, 1613895288864, b'1');
INSERT INTO `sys_menu` VALUES (169, '/api/roles', 'GET', 0, 0, 0, 'LJK', '', 'Layout', 999, 'Steve-Jobs', '/api/group/search', b'0', b'0', b'0', '', 1, 1, 1613896058197, 1613896058197, b'1');
INSERT INTO `sys_menu` VALUES (170, '\r\n/api/users/', 'GET', 0, 0, 1, '文章审核', '', 'Layout', 20, 'alipay', 'articleReview', b'0', b'0', b'0', '', 1, 1, 1613897767842, 1613897767842, b'1');
INSERT INTO `sys_menu` VALUES (171, '/api/roles', 'GET', 0, 0, 0, '1111111111', '', 'Layout', 999, 'app', '/api/pictures', b'0', b'0', b'0', '', 1, 1, 1613897912254, 1613897912254, b'0');
INSERT INTO `sys_menu` VALUES (172, '', '', 0, 0, 0, '文章审核', '', 'Layout', 999, 'Steve-Jobs', '嘟嘟嘟', b'0', b'0', b'0', '', 1, 1, 1613898029500, 1613898029500, b'1');
INSERT INTO `sys_menu` VALUES (173, '/api/role', 'POST', 0, 1, 0, '文章审核', '', 'Layout', 999, 'Steve-Jobs', '/api/pictures', b'0', b'0', b'0', '', 1, 1, 1613898271747, 1613898271747, b'0');
INSERT INTO `sys_menu` VALUES (174, '', '', 0, 0, 1, '文章', '111', 'jlj', 999, 'Steve-Jobs', 'lbl', b'0', b'0', b'0', '', 1, 1, 1613898372873, 1613898372873, b'1');
INSERT INTO `sys_menu` VALUES (175, '/api/role', 'GET', 173, 0, 1, '文章菜单', 'aaa', 'jlj', 999, 'develop', '/api/pictures', b'0', b'0', b'0', 'sad', 1, 1, 1613912704216, 1613912704216, b'0');
INSERT INTO `sys_menu` VALUES (176, '/api/role', 'GET', 0, 1, 0, 'qqqq', '', 'Layout', 999, 'app', 'qqqq', b'0', b'0', b'0', '', 1, 1, 1613965516081, 1613965516081, b'0');
INSERT INTO `sys_menu` VALUES (177, '/api/role', 'GET', 176, 1, 1, 'cccc', 'asadas', 'asdsad', 999, 'Steve-Jobs', '/api/pictures', b'0', b'0', b'0', 'asdas', 1, 1, 1613965566438, 1613965566438, b'0');
INSERT INTO `sys_menu` VALUES (178, '/api/role', 'GET', 177, 0, 1, 'tttt', 'asdsad', 'asdsad', 999, 'app', '/api/pictures', b'0', b'0', b'0', 'asd', 1, 1, 1613965642756, 1613965642756, b'0');

-- ----------------------------
-- Table structure for sys_quartz_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_quartz_job`;
CREATE TABLE `sys_quartz_job`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `bean_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'Spring Bean名称',
  `cron_expression` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'cron 表达式',
  `is_pause` bit(1) NOT NULL DEFAULT b'0' COMMENT '状态：0暂停、1启用',
  `job_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `method_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '方法名称',
  `params` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '参数',
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `person_in_charge` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '负责人',
  `email` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '报警邮箱',
  `sub_task` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '子任务ID',
  `pause_after_failure` bit(1) NOT NULL DEFAULT b'0' COMMENT '任务失败后是否暂停,0是暂停，1是不暂停',
  `create_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者',
  `create_time` bigint(20) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '逻辑删除：0启用（默认）、1删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `inx_is_pause`(`is_pause`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '定时任务' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_quartz_job
-- ----------------------------
INSERT INTO `sys_quartz_job` VALUES (2, 'testTask', '*/5 * * * * ?', b'1', 'test', 'run', '我想要休息一下', '测试有参数run', '李四', '', '', b'1', 1, 1, 20210122104851, 20210122104851, b'0');
INSERT INTO `sys_quartz_job` VALUES (4, 'testTask', '*/5 * * * * ?', b'1', 'test2', 'run', '', '测试子任务', '张三', '', '2', b'1', 1, 1, 20210122104851, 20210122104851, b'0');
INSERT INTO `sys_quartz_job` VALUES (5, 'testTask', '*/5 * * * * ?', b'1', 'rest', 'run', '', '测试无参数run', '李四', '', '', b'1', 1, 1, 20210122104851, 20210122104851, b'0');
INSERT INTO `sys_quartz_job` VALUES (9, 'testTask', '*/5 * * * * ?', b'1', 'test', 'job', '', '测试job任务', '张三', '', '', b'1', 1, 1, 20210122104851, 20210122104851, b'0');

-- ----------------------------
-- Table structure for sys_quartz_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_quartz_log`;
CREATE TABLE `sys_quartz_log`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `bean_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'bean对象名称',
  `cron_expression` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'cron表达式',
  `exception_detail` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '异常详情',
  `is_success` bit(1) NOT NULL DEFAULT b'0' COMMENT '状态（是否成功）1成功，0失败(默认)',
  `job_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `method_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '执行方法',
  `params` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '方法参数',
  `time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '执行时间(ms)',
  `create_time` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `create_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者',
  `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '逻辑删除：0启用（默认）、1删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 141 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '定时任务日志' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_quartz_log
-- ----------------------------
INSERT INTO `sys_quartz_log` VALUES (1, 'testTask', '*/5 * * * * ?', 'java.lang.NumberFormatException: For input string: \"\"\r\n	at java.base/java.lang.NumberFormatException.forInputString(NumberFormatException.java:65)\r\n	at java.base/java.lang.Long.parseLong(Long.java:702)\r\n	at java.base/java.lang.Long.parseLong(Long.java:817)\r\n	at marchsoft.modules.quartz.service.impl.QuartzJobServiceImpl.executionSubJob(QuartzJobServiceImpl.java:67)\r\n	at marchsoft.modules.quartz.service.impl.QuartzJobServiceImpl$$FastClassBySpringCGLIB$$c8d365fd.invoke(<generated>)\r\n	at org.springframework.cglib.proxy.MethodProxy.invoke(MethodProxy.java:218)\r\n	at org.springframework.aop.framework.CglibAopProxy$CglibMethodInvocation.invokeJoinpoint(CglibAopProxy.java:746)\r\n	at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:163)\r\n	at org.springframework.transaction.interceptor.TransactionAspectSupport.invokeWithinTransaction(TransactionAspectSupport.java:294)\r\n	at org.springframework.transaction.interceptor.TransactionInterceptor.invoke(TransactionInterceptor.java:98)\r\n	at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:186)\r\n	at org.springframework.aop.framework.CglibAopProxy$DynamicAdvisedInterceptor.intercept(CglibAopProxy.java:688)\r\n	at marchsoft.modules.quartz.service.impl.QuartzJobServiceImpl$$EnhancerBySpringCGLIB$$747379eb.executionSubJob(<generated>)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:74)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', '', 886, 20210122104910, 20210122104910, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (2, 'testTask', '*/5 * * * * ?', 'java.lang.NumberFormatException: For input string: \"\"\r\n	at java.base/java.lang.NumberFormatException.forInputString(NumberFormatException.java:65)\r\n	at java.base/java.lang.Long.parseLong(Long.java:702)\r\n	at java.base/java.lang.Long.parseLong(Long.java:817)\r\n	at marchsoft.modules.quartz.service.impl.QuartzJobServiceImpl.executionSubJob(QuartzJobServiceImpl.java:67)\r\n	at marchsoft.modules.quartz.service.impl.QuartzJobServiceImpl$$FastClassBySpringCGLIB$$c8d365fd.invoke(<generated>)\r\n	at org.springframework.cglib.proxy.MethodProxy.invoke(MethodProxy.java:218)\r\n	at org.springframework.aop.framework.CglibAopProxy$CglibMethodInvocation.invokeJoinpoint(CglibAopProxy.java:746)\r\n	at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:163)\r\n	at org.springframework.transaction.interceptor.TransactionAspectSupport.invokeWithinTransaction(TransactionAspectSupport.java:294)\r\n	at org.springframework.transaction.interceptor.TransactionInterceptor.invoke(TransactionInterceptor.java:98)\r\n	at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:186)\r\n	at org.springframework.aop.framework.CglibAopProxy$DynamicAdvisedInterceptor.intercept(CglibAopProxy.java:688)\r\n	at marchsoft.modules.quartz.service.impl.QuartzJobServiceImpl$$EnhancerBySpringCGLIB$$4c66ffb1.executionSubJob(<generated>)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:76)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', '', 315, 20210122105253, 20210122105253, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (3, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122105615, 20210122105615, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (4, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:28)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:56)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 1, 20210122111923, 20210122111923, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (5, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122112155, 20210122112155, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (6, '11', '*/5 * * * * ?', NULL, b'0', '2', '1', '', 0, 20210122112202, 20210122112202, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (7, '11', '*/5 * * * * ?', NULL, b'0', '2', '1', '', 0, 20210122112226, 20210122112226, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (8, '11', '*/5 * * * * ?', NULL, b'0', '2', '1', '', 0, 20210122112353, 20210122112353, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (9, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122112801, 20210122112801, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (10, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122112806, 20210122112806, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (11, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122112820, 20210122112820, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (12, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122112825, 20210122112825, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (13, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122112831, 20210122112831, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (14, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122112839, 20210122112839, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (15, '11', '*/5 * * * * ?', NULL, b'0', '2', '1', '', 0, 20210122113050, 20210122113050, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (16, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122113154, 20210122113154, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (17, '11', '*/5 * * * * ?', NULL, b'0', '2', '1', '', 0, 20210122113212, 20210122113212, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (18, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122114314, 20210122114314, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (19, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 2, 20210122114411, 20210122114411, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (20, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122114415, 20210122114415, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (21, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122114420, 20210122114420, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (22, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122114425, 20210122114425, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (23, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122114430, 20210122114430, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (24, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122114435, 20210122114435, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (25, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122114440, 20210122114440, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (26, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122114445, 20210122114445, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (27, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122114530, 20210122114530, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (28, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:28)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 2, 20210122140522, 20210122140522, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (29, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:28)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 3, 20210122140607, 20210122140607, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (30, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:28)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 1, 20210122140752, 20210122140752, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (31, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:28)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 3, 20210122140843, 20210122140843, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (32, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:28)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 0, 20210122140908, 20210122140908, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (33, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:28)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 2, 20210122141455, 20210122141455, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (34, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:28)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 3, 20210122141533, 20210122141533, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (35, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 2, 20210122142006, 20210122142006, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (36, '11', '*/5 * * * * ?', 'marchsoft.exception.BadRequestException: 未登录或当前登录状态过期\r\n	at marchsoft.utils.SecurityUtils.getCurrentUser(SecurityUtils.java:34)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:41)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 1, 20210122142011, 20210122142011, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (37, '11', '*/5 * * * * ?', 'marchsoft.exception.BadRequestException: 该bean对象或者方法不存在\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:40)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 3, 20210122142207, 20210122142207, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (38, '11', '*/5 * * * * ?', 'marchsoft.exception.BadRequestException: 该bean对象或者方法不存在\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:40)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 3, 20210122142437, 20210122142437, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (39, '11', '*/5 * * * * ?', 'marchsoft.exception.BadRequestException: 该bean对象或者方法不存在\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:40)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 3, 20210122142526, 20210122142526, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (40, '11', '*/5 * * * * ?', 'marchsoft.exception.BadRequestException: 该bean对象或者方法不存在\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:40)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 16591, 20210122142602, 20210122142602, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (41, '11', '*/5 * * * * ?', 'marchsoft.exception.BadRequestException: 该bean对象或者方法不存在\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:40)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 2705, 20210122143531, 20210122143531, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (42, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 2, 20210122143550, 20210122143550, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (43, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 2, 20210122143550, 20210122143550, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (44, '11', '*/5 * * * * ?', 'marchsoft.exception.BadRequestException: 该bean对象或者方法不存在\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:40)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 16710, 20210122143635, 20210122143635, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (45, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122143748, 20210122143748, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (46, '11', '*/5 * * * * ?', 'marchsoft.exception.BadRequestException: 该bean对象或者方法不存在\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:40)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 15954, 20210122144129, 20210122144130, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (47, '11', '*/5 * * * * ?', 'marchsoft.exception.BadRequestException: 该bean对象或者方法不存在\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:40)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 25369, 20210122144334, 20210122144334, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (48, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:27)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 3, 20210122144739, 20210122144739, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (49, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:27)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 0, 20210122145316, 20210122145316, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (50, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:27)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 1, 20210122145723, 20210122145723, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (51, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:27)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 1, 20210122145852, 20210122145852, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (52, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122145949, 20210122145949, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (53, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122145958, 20210122145958, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (54, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122150000, 20210122150000, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (55, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122150006, 20210122150006, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (56, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122150015, 20210122150015, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (57, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122150016, 20210122150016, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (58, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122150020, 20210122150020, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (59, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122150317, 20210122150317, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (60, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122150321, 20210122150321, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (61, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test2', 'run', '', 1, 20210122150321, 20210122150321, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (62, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test2', 'run', '', 1, 20210122150326, 20210122150326, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (63, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122150326, 20210122150326, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (64, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122150331, 20210122150331, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (65, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test2', 'run', '', 1, 20210122150331, 20210122150331, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (66, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test2', 'run', '', 0, 20210122150336, 20210122150336, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (67, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 3, 20210122151223, 20210122151223, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (68, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122151742, 20210122151742, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (69, 'testTask', '*/5 * * * * ?', 'marchsoft.exception.BadRequestException: 未登录或当前登录状态过期\r\n	at marchsoft.utils.SecurityUtils.getCurrentUser(SecurityUtils.java:34)\r\n	at marchsoft.modules.quartz.utils.QuartzManage.runJobNow(QuartzManage.java:198)\r\n	at marchsoft.modules.quartz.service.impl.QuartzJobServiceImpl.execution(QuartzJobServiceImpl.java:148)\r\n	at marchsoft.modules.quartz.service.impl.QuartzJobServiceImpl.executionSubJob(QuartzJobServiceImpl.java:71)\r\n	at marchsoft.modules.quartz.service.impl.QuartzJobServiceImpl$$FastClassBySpringCGLIB$$c8d365fd.invoke(<generated>)\r\n	at org.springframework.cglib.proxy.MethodProxy.invoke(MethodProxy.java:218)\r\n	at org.springframework.aop.framework.CglibAopProxy$CglibMethodInvocation.invokeJoinpoint(CglibAopProxy.java:746)\r\n	at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:163)\r\n	at org.springframework.transaction.interceptor.TransactionAspectSupport.invokeWithinTransaction(TransactionAspectSupport.java:294)\r\n	at org.springframework.transaction.interceptor.TransactionInterceptor.invoke(TransactionInterceptor.java:98)\r\n	at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:186)\r\n	at org.springframework.aop.framework.CglibAopProxy$DynamicAdvisedInterceptor.intercept(CglibAopProxy.java:688)\r\n	at marchsoft.modules.quartz.service.impl.QuartzJobServiceImpl$$EnhancerBySpringCGLIB$$7bda4b27.executionSubJob(<generated>)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:78)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test2', 'run', '', 629, 20210122151743, 20210122151743, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (70, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:27)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 3, 20210122151911, 20210122151911, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (71, '11', '*/5 * * * * ?', 'org.springframework.beans.factory.NoSuchBeanDefinitionException: No bean named \'11\' available\r\n	at org.springframework.beans.factory.support.DefaultListableBeanFactory.getBeanDefinition(DefaultListableBeanFactory.java:772)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getMergedLocalBeanDefinition(AbstractBeanFactory.java:1212)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.doGetBean(AbstractBeanFactory.java:294)\r\n	at org.springframework.beans.factory.support.AbstractBeanFactory.getBean(AbstractBeanFactory.java:199)\r\n	at org.springframework.context.support.AbstractApplicationContext.getBean(AbstractApplicationContext.java:1083)\r\n	at marchsoft.utils.SpringContextHolder.getBean(SpringContextHolder.java:31)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:27)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', '2', '1', '', 3, 20210122152309, 20210122152309, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (72, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', 'name', 3, 20210122195315, 20210122195315, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (73, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run()\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:35)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test2', 'run', '', 0, 20210122195327, 20210122195327, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (74, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 4, 20210122200027, 20210122200027, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (75, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 148, 20210122203016, 20210122203016, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (76, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 3, 20210122203021, 20210122203021, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (77, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 301, 20210122203022, 20210122203022, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (78, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 148, 20210122203026, 20210122203026, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (79, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122203026, 20210122203026, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (80, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122203031, 20210122203031, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (81, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 391, 20210122203032, 20210122203032, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (82, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122203036, 20210122203036, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (83, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 409, 20210122203036, 20210122203036, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (84, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 144, 20210122203041, 20210122203041, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (85, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122203041, 20210122203041, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (86, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122203046, 20210122203046, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (87, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 159, 20210122203046, 20210122203046, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (88, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122203051, 20210122203051, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (89, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 143, 20210122203051, 20210122203051, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (90, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 165, 20210122203056, 20210122203056, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (91, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 0, 20210122203056, 20210122203056, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (92, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122203101, 20210122203101, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (93, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 271, 20210122203101, 20210122203101, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (94, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122203106, 20210122203106, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (95, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 154, 20210122203106, 20210122203106, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (96, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 0, 20210122203111, 20210122203111, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (97, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 408, 20210122203112, 20210122203112, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (98, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 120, 20210122203116, 20210122203116, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (99, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 0, 20210122203117, 20210122203117, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (100, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 0, 20210122203121, 20210122203121, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (101, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 202, 20210122203122, 20210122203122, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (102, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 0, 20210122203126, 20210122203126, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (103, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 430, 20210122203126, 20210122203126, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (104, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 188, 20210122203131, 20210122203131, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (105, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 0, 20210122203131, 20210122203131, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (106, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122203136, 20210122203136, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (107, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 162, 20210122203137, 20210122203137, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (108, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 0, 20210122203141, 20210122203141, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (109, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 62, 20210122203152, 20210122203152, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (110, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 0, 20210122203212, 20210122203212, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (111, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 150, 20210122204953, 20210122204953, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (112, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test2', 'run', '', 3, 20210122204958, 20210122204958, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (113, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:33)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:57)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'name', 157, 20210122205034, 20210122205034, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (114, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122205040, 20210122205040, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (115, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122205208, 20210122205208, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (116, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122205217, 20210122205217, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (117, 'testTask', '*/5 * * * * ?', NULL, b'1', 'rest', 'run', '', 1, 20210122205218, 20210122205218, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (118, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 1, 20210122211447, 20210122211447, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (119, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122211451, 20210122211451, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (120, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test2', 'run', '', 1, 20210122211452, 20210122211452, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (121, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122211456, 20210122211456, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (122, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test2', 'run', '', 2, 20210122211456, 20210122211456, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (123, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test2', 'run', '', 1, 20210122211501, 20210122211501, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (124, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122211501, 20210122211501, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (125, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122211506, 20210122211506, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (126, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test2', 'run', '', 1, 20210122211506, 20210122211506, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (127, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test2', 'run', '', 0, 20210122211511, 20210122211511, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (128, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122211700, 20210122211700, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (129, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '', 0, 20210122211705, 20210122211705, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (130, 'testTask', '*/5 * * * * ?', 'java.lang.ClassNotFoundException: Integer\r\n	at java.base/jdk.internal.loader.BuiltinClassLoader.loadClass(BuiltinClassLoader.java:581)\r\n	at java.base/jdk.internal.loader.ClassLoaders$AppClassLoader.loadClass(ClassLoaders.java:178)\r\n	at java.base/java.lang.ClassLoader.loadClass(ClassLoader.java:521)\r\n	at java.base/java.lang.Class.forName0(Native Method)\r\n	at java.base/java.lang.Class.forName(Class.java:315)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:36)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:56)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'Integer,String', 3, 20210123084759, 20210123084759, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (131, 'testTask', '*/5 * * * * ?', 'java.lang.ClassNotFoundException: Integer\r\n	at java.base/jdk.internal.loader.BuiltinClassLoader.loadClass(BuiltinClassLoader.java:581)\r\n	at java.base/jdk.internal.loader.ClassLoaders$AppClassLoader.loadClass(ClassLoaders.java:178)\r\n	at java.base/java.lang.ClassLoader.loadClass(ClassLoader.java:521)\r\n	at java.base/java.lang.Class.forName0(Native Method)\r\n	at java.base/java.lang.Class.forName(Class.java:315)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:37)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:56)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'Integer,String', 3, 20210123085153, 20210123085153, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (132, 'testTask', '*/5 * * * * ?', 'java.lang.ClassNotFoundException: Integer\r\n	at java.base/jdk.internal.loader.BuiltinClassLoader.loadClass(BuiltinClassLoader.java:581)\r\n	at java.base/jdk.internal.loader.ClassLoaders$AppClassLoader.loadClass(ClassLoaders.java:178)\r\n	at java.base/java.lang.ClassLoader.loadClass(ClassLoader.java:521)\r\n	at java.base/java.lang.Class.forName0(Native Method)\r\n	at java.base/java.lang.Class.forName(Class.java:315)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:37)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:56)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'Integer,String', 0, 20210123085351, 20210123085351, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (133, 'testTask', '*/5 * * * * ?', 'java.lang.ClassNotFoundException: Integer\r\n	at java.base/jdk.internal.loader.BuiltinClassLoader.loadClass(BuiltinClassLoader.java:581)\r\n	at java.base/jdk.internal.loader.ClassLoaders$AppClassLoader.loadClass(ClassLoaders.java:178)\r\n	at java.base/java.lang.ClassLoader.loadClass(ClassLoader.java:521)\r\n	at java.base/java.lang.Class.forName0(Native Method)\r\n	at java.base/java.lang.Class.forName(Class.java:315)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:37)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:56)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'Integer,String', 1, 20210123085358, 20210123085358, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (134, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:40)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:56)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', 'Integer,String', 3, 20210123091851, 20210123091851, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (135, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', 'Integer,String', 3, 20210123091933, 20210123091933, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (136, 'testTask', '*/5 * * * * ?', NULL, b'1', 'test', 'run', '我想要休息一下', 3, 20210123092558, 20210123092558, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (137, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run1(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:40)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:56)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run1', '我想要休息一下', 1, 20210123103049, 20210123103049, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (138, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run1(java.lang.String)\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:40)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:56)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run1', '我想要休息一下', 0, 20210123185643, 20210123185643, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (139, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run()\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:42)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:56)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', '', 0, 20210123185732, 20210123185732, 0, 0, b'0');
INSERT INTO `sys_quartz_log` VALUES (140, 'testTask', '*/5 * * * * ?', 'java.lang.NoSuchMethodException: marchsoft.modules.quartz.task.TestTask.run()\r\n	at java.base/java.lang.Class.getDeclaredMethod(Class.java:2475)\r\n	at marchsoft.modules.quartz.utils.QuartzRunnable.<init>(QuartzRunnable.java:42)\r\n	at marchsoft.modules.quartz.utils.ExecutionJob.executeInternal(ExecutionJob.java:56)\r\n	at org.springframework.scheduling.quartz.QuartzJobBean.execute(QuartzJobBean.java:75)\r\n	at org.quartz.core.JobRunShell.run(JobRunShell.java:202)\r\n	at org.quartz.simpl.SimpleThreadPool$WorkerThread.run(SimpleThreadPool.java:573)\r\n', b'0', 'test', 'run', '', 0, 20210123185736, 20210123185736, 0, 0, b'0');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `level` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色级别（越小越大）',
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `data_scope` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '数据权限',
  `is_protection` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否受保护（内置角色，1为内置角色，默认值为0）',
  `create_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者id',
  `update_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者id',
  `create_time` bigint(20) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '软删除（默认值为0，1为删除）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unq_name`(`name`) USING BTREE COMMENT '角色名唯一',
  INDEX `idx_role_name`(`name`) USING BTREE COMMENT '普通索引——角色名查询角色信息'
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '超级管理员', 1, '-', '全部', b'1', 1, 1, 20201213113645, 20201213113647, b'0');
INSERT INTO `sys_role` VALUES (2, '普通用户1', 2, '-', '本级', b'0', 1, 1, 20201213113714, 1613639016559, b'0');
INSERT INTO `sys_role` VALUES (9, '123123', 3, '2131235', '本级', b'0', 1, 1, 1613390525872, 1613822244041, b'0');
INSERT INTO `sys_role` VALUES (10, 'das', 3, 'text1', '本级', b'0', 1, 1, 1613391123884, 1613823908741, b'0');
INSERT INTO `sys_role` VALUES (11, 'text2', 3, '232', '全部', b'1', 1, 1, 1613391149359, 1613391149359, b'0');
INSERT INTO `sys_role` VALUES (12, 'text3', 3, 'text=3', '本级', b'1', 1, 1, 1613391167396, 1613391167396, b'0');
INSERT INTO `sys_role` VALUES (13, 'text4', 3, 'text4', '全部', b'1', 1, 1, 1613391185290, 1613391185290, b'1');
INSERT INTO `sys_role` VALUES (14, 'text5', 3, 'text5', '全部', b'1', 1, 1, 1613391193447, 1613391193447, b'0');
INSERT INTO `sys_role` VALUES (15, 'tex20', 3, 'text6', '自定义', b'0', 1, 1, 1613391212971, 1613822261479, b'0');
INSERT INTO `sys_role` VALUES (17, 'text7', 3, '1112', '本级', b'0', 1, 1, 1613391246421, 1613639736094, b'1');
INSERT INTO `sys_role` VALUES (18, 'text8', 3, '1111', '全部', b'1', 1, 1, 1613391271238, 1613391271238, b'1');
INSERT INTO `sys_role` VALUES (19, 'text9', 3, 'text131', '本级', b'1', 1, 1, 1613391311326, 1613391311326, b'1');
INSERT INTO `sys_role` VALUES (20, 'text19', 3, '212121212', '全部', b'1', 1, 1, 1613395094708, 1613395094708, b'0');
INSERT INTO `sys_role` VALUES (21, 'text20', 3, '123123', '全部', b'1', 1, 1, 1613395166827, 1613395166827, b'0');
INSERT INTO `sys_role` VALUES (22, 'LJK', 3, '6666a', '全部', b'1', 1, 1, 1613701960106, 1613701960106, b'0');
INSERT INTO `sys_role` VALUES (25, '用户', 3, '1112233', '全部', b'1', 1, 1, 1613897125593, 1613897125593, b'0');

-- ----------------------------
-- Table structure for sys_roles_depts
-- ----------------------------
DROP TABLE IF EXISTS `sys_roles_depts`;
CREATE TABLE `sys_roles_depts`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `role_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色id',
  `dept_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_mid_dept_id`(`dept_id`) USING BTREE COMMENT '普通索引——根据dept_id查询',
  INDEX `idx_mid_role_id`(`role_id`) USING BTREE COMMENT '普通索引——根据role_id查询'
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色部门关联' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_roles_depts
-- ----------------------------
INSERT INTO `sys_roles_depts` VALUES (3, 15, 7);

-- ----------------------------
-- Table structure for sys_roles_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_roles_menus`;
CREATE TABLE `sys_roles_menus`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id\n',
  `menu_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单ID',
  `role_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_mid_role_id`(`role_id`) USING BTREE COMMENT '普通索引——根据role_id查询',
  INDEX `idx_mid_menu_id`(`menu_id`) USING BTREE COMMENT '普通索引——根据menu_id查询'
) ENGINE = InnoDB AUTO_INCREMENT = 622 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色菜单关联' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_roles_menus
-- ----------------------------
INSERT INTO `sys_roles_menus` VALUES (2, 1, 2);
INSERT INTO `sys_roles_menus` VALUES (4, 2, 2);
INSERT INTO `sys_roles_menus` VALUES (6, 3, 2);
INSERT INTO `sys_roles_menus` VALUES (8, 5, 2);
INSERT INTO `sys_roles_menus` VALUES (286, 6, 17);
INSERT INTO `sys_roles_menus` VALUES (287, 7, 17);
INSERT INTO `sys_roles_menus` VALUES (288, 32, 17);
INSERT INTO `sys_roles_menus` VALUES (289, 41, 17);
INSERT INTO `sys_roles_menus` VALUES (290, 1, 17);
INSERT INTO `sys_roles_menus` VALUES (291, 3, 17);
INSERT INTO `sys_roles_menus` VALUES (292, 5, 17);
INSERT INTO `sys_roles_menus` VALUES (293, 28, 17);
INSERT INTO `sys_roles_menus` VALUES (294, 35, 17);
INSERT INTO `sys_roles_menus` VALUES (295, 37, 17);
INSERT INTO `sys_roles_menus` VALUES (296, 39, 17);
INSERT INTO `sys_roles_menus` VALUES (297, 48, 17);
INSERT INTO `sys_roles_menus` VALUES (298, 49, 17);
INSERT INTO `sys_roles_menus` VALUES (299, 50, 17);
INSERT INTO `sys_roles_menus` VALUES (300, 52, 17);
INSERT INTO `sys_roles_menus` VALUES (301, 53, 17);
INSERT INTO `sys_roles_menus` VALUES (302, 54, 17);
INSERT INTO `sys_roles_menus` VALUES (303, 73, 17);
INSERT INTO `sys_roles_menus` VALUES (304, 74, 17);
INSERT INTO `sys_roles_menus` VALUES (305, 75, 17);
INSERT INTO `sys_roles_menus` VALUES (306, 56, 17);
INSERT INTO `sys_roles_menus` VALUES (307, 57, 17);
INSERT INTO `sys_roles_menus` VALUES (308, 58, 17);
INSERT INTO `sys_roles_menus` VALUES (309, 60, 17);
INSERT INTO `sys_roles_menus` VALUES (310, 61, 17);
INSERT INTO `sys_roles_menus` VALUES (311, 62, 17);
INSERT INTO `sys_roles_menus` VALUES (312, 64, 17);
INSERT INTO `sys_roles_menus` VALUES (313, 65, 17);
INSERT INTO `sys_roles_menus` VALUES (314, 66, 17);
INSERT INTO `sys_roles_menus` VALUES (321, 6, 12);
INSERT INTO `sys_roles_menus` VALUES (322, 7, 12);
INSERT INTO `sys_roles_menus` VALUES (323, 32, 12);
INSERT INTO `sys_roles_menus` VALUES (324, 41, 12);
INSERT INTO `sys_roles_menus` VALUES (331, 123, 14);
INSERT INTO `sys_roles_menus` VALUES (332, 6, 14);
INSERT INTO `sys_roles_menus` VALUES (333, 7, 14);
INSERT INTO `sys_roles_menus` VALUES (334, 32, 14);
INSERT INTO `sys_roles_menus` VALUES (335, 41, 14);
INSERT INTO `sys_roles_menus` VALUES (336, 6, 10);
INSERT INTO `sys_roles_menus` VALUES (337, 7, 10);
INSERT INTO `sys_roles_menus` VALUES (338, 32, 10);
INSERT INTO `sys_roles_menus` VALUES (339, 41, 10);
INSERT INTO `sys_roles_menus` VALUES (340, 125, 10);
INSERT INTO `sys_roles_menus` VALUES (341, 123, 10);
INSERT INTO `sys_roles_menus` VALUES (584, 1, 1);
INSERT INTO `sys_roles_menus` VALUES (585, 2, 1);
INSERT INTO `sys_roles_menus` VALUES (586, 3, 1);
INSERT INTO `sys_roles_menus` VALUES (587, 5, 1);
INSERT INTO `sys_roles_menus` VALUES (588, 28, 1);
INSERT INTO `sys_roles_menus` VALUES (589, 35, 1);
INSERT INTO `sys_roles_menus` VALUES (590, 37, 1);
INSERT INTO `sys_roles_menus` VALUES (591, 39, 1);
INSERT INTO `sys_roles_menus` VALUES (592, 6, 1);
INSERT INTO `sys_roles_menus` VALUES (593, 7, 1);
INSERT INTO `sys_roles_menus` VALUES (594, 32, 1);
INSERT INTO `sys_roles_menus` VALUES (595, 41, 1);
INSERT INTO `sys_roles_menus` VALUES (596, 169, 1);
INSERT INTO `sys_roles_menus` VALUES (597, 170, 1);
INSERT INTO `sys_roles_menus` VALUES (598, 171, 1);
INSERT INTO `sys_roles_menus` VALUES (599, 172, 1);
INSERT INTO `sys_roles_menus` VALUES (600, 173, 1);
INSERT INTO `sys_roles_menus` VALUES (601, 174, 1);
INSERT INTO `sys_roles_menus` VALUES (602, 175, 1);
INSERT INTO `sys_roles_menus` VALUES (603, 176, 1);
INSERT INTO `sys_roles_menus` VALUES (604, 177, 1);
INSERT INTO `sys_roles_menus` VALUES (605, 178, 1);
INSERT INTO `sys_roles_menus` VALUES (618, 6, 9);
INSERT INTO `sys_roles_menus` VALUES (619, 7, 9);
INSERT INTO `sys_roles_menus` VALUES (620, 32, 9);
INSERT INTO `sys_roles_menus` VALUES (621, 41, 9);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `dept_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门id',
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `gender` bit(1) NOT NULL DEFAULT b'0' COMMENT '性别（0为男默认，1为女）',
  `phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '手机号码',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `avatar_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '头像路径',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `is_admin` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否为admin账号',
  `enabled` bit(1) NOT NULL DEFAULT b'1' COMMENT '状态：1启用（默认）、0禁用',
  `create_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者id',
  `update_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者id',
  `pwd_reset_time` bigint(20) NULL DEFAULT NULL COMMENT '修改密码的时间',
  `create_time` bigint(20) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '软删除（默认值为0，1为删除）',
  `role_id` bigint(20) NULL DEFAULT NULL,
  `salt` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `avatar` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `sex` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `status` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `remark` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `post_id` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_dept_id`(`dept_id`) USING BTREE COMMENT '普通索引——根据dept_id查询用户',
  INDEX `idx_enabled`(`enabled`) USING BTREE COMMENT '普通索引——根据enabled查询'
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统用户' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 2, 'admin', '管理员', b'0', '18888888888', '201507802@qq.com', '44e0e97c-f09e-440c-879f-79014b8d8f13.png', '47474af9f2fa250735c68a2af27d66b0', b'1', b'1', 1, 1, 20201213112720, 20201213112652, 20201213112657, b'0', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (2, 5, 'test', '测试', b'1', '15689899898', '231@qq.com', '81c13fb2-e0cf-49a6-a4ed-3e83764133f2.png', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, NULL, 20201213113453, 20201213113459, b'0', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (3, 8, '成功', '1312', b'0', '15083138896', '', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'0', 1, 1, 1613639852910, 1613639852989, 1613639852989, b'0', 0, '', '', '', '', '', 0);
INSERT INTO `sys_user` VALUES (4, 7, '3123', '3123', b'0', '15083138896', '', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1613639939808, 1613639939874, 1613639939874, b'0', 0, '', '', '', '', '', 0);
INSERT INTO `sys_user` VALUES (5, 8, 'dududu', '安静', b'0', '15083138896', '', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1613640066925, 1613640066980, 1613640066980, b'0', 0, '', '', '', '', '', 0);
INSERT INTO `sys_user` VALUES (6, 8, '15083138896', '15083138896', b'0', '15083138896', '', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1613641040862, 1613641040908, 1613641040908, b'0', 0, '', '', '', '', '', 0);
INSERT INTO `sys_user` VALUES (7, 7, 'aaa', 'ss', b'0', '15757961798', '', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1613809331904, 1613809331965, 1613809331965, b'0', 0, '', '', '', '', '', 0);
INSERT INTO `sys_user` VALUES (8, 7, '打发斯蒂芬', 'dddd', b'0', '15737966928', '6934899864@qq.com', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1613822105901, 1613822105965, 1613822105965, b'0', 0, '', '', '', '', '', 0);
INSERT INTO `sys_user` VALUES (9, 7, '1111', '231312', b'1', '15798921688', '2534992866@qq.com', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1613896256620, 1613896256677, 1613896256677, b'0', 0, '', '', '', '', '', 0);
INSERT INTO `sys_user` VALUES (10, 7, 'user', '1111', b'1', '15798652168', '231@qq.com', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1613896395105, 1613896395169, 1613896395169, b'0', 0, '', '', '', '', '', 0);
INSERT INTO `sys_user` VALUES (11, 8, '11111', 'LJK', b'1', '15737987721', '231@qq.com', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1613896639825, 1613896639944, 1613896639944, b'0', 0, '', '', '', '', '', 0);
INSERT INTO `sys_user` VALUES (12, 8, '京125雪', '飞翔得子弹', b'1', '15236967283', '1720808104@qq.com', 'https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIomz1vPJX1xeY2pMaOvXQ6ticGJfQWaJw6wjoiaicYoIjwAOg2vFvhMdOianQ7A4OxicJ8Ml76N2an8Nw/132', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1613899039308, 1613899039538, 1613899039538, b'0', 0, '', '', '', '', '', 0);
INSERT INTO `sys_user` VALUES (13, 8, 'QXTest', 'LJK', b'1', '13837968289', '231@qq.com', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1613962895583, 1613962895771, 1613962895771, b'0', 0, '', '', '', '', '', 0);

-- ----------------------------
-- Table structure for sys_users_jobs
-- ----------------------------
DROP TABLE IF EXISTS `sys_users_jobs`;
CREATE TABLE `sys_users_jobs`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `job_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '岗位ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_mid_job_id`(`job_id`) USING BTREE COMMENT '普通索引——根据job_id查询',
  INDEX `idx_mid_user_id`(`user_id`) USING BTREE COMMENT '普通索引——根据user_id查询'
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users_jobs
-- ----------------------------
INSERT INTO `sys_users_jobs` VALUES (3, 2, 2);
INSERT INTO `sys_users_jobs` VALUES (4, 3, 0);
INSERT INTO `sys_users_jobs` VALUES (5, 3, 29);
INSERT INTO `sys_users_jobs` VALUES (6, 4, 2);
INSERT INTO `sys_users_jobs` VALUES (7, 4, 29);
INSERT INTO `sys_users_jobs` VALUES (8, 5, 29);
INSERT INTO `sys_users_jobs` VALUES (9, 6, 2);
INSERT INTO `sys_users_jobs` VALUES (10, 7, 29);
INSERT INTO `sys_users_jobs` VALUES (11, 7, 2);
INSERT INTO `sys_users_jobs` VALUES (14, 8, 0);
INSERT INTO `sys_users_jobs` VALUES (15, 8, 0);
INSERT INTO `sys_users_jobs` VALUES (16, 8, 2);
INSERT INTO `sys_users_jobs` VALUES (17, 9, 2);
INSERT INTO `sys_users_jobs` VALUES (18, 9, 29);
INSERT INTO `sys_users_jobs` VALUES (19, 10, 29);
INSERT INTO `sys_users_jobs` VALUES (20, 11, 2);
INSERT INTO `sys_users_jobs` VALUES (21, 11, 29);
INSERT INTO `sys_users_jobs` VALUES (22, 1, 29);
INSERT INTO `sys_users_jobs` VALUES (23, 1, 2);
INSERT INTO `sys_users_jobs` VALUES (26, 13, 29);
INSERT INTO `sys_users_jobs` VALUES (27, 12, 29);

-- ----------------------------
-- Table structure for sys_users_roles
-- ----------------------------
DROP TABLE IF EXISTS `sys_users_roles`;
CREATE TABLE `sys_users_roles`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id\n',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `role_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_mid_role_id`(`role_id`) USING BTREE COMMENT '普通索引——根据role_id查询',
  INDEX `idx_mid_user_id`(`user_id`) USING BTREE COMMENT '普通索引——根据user_id查询用户'
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户角色关联' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_users_roles
-- ----------------------------
INSERT INTO `sys_users_roles` VALUES (2, 2, 1);
INSERT INTO `sys_users_roles` VALUES (3, 3, 0);
INSERT INTO `sys_users_roles` VALUES (4, 3, 2);
INSERT INTO `sys_users_roles` VALUES (5, 4, 2);
INSERT INTO `sys_users_roles` VALUES (6, 5, 2);
INSERT INTO `sys_users_roles` VALUES (7, 6, 14);
INSERT INTO `sys_users_roles` VALUES (8, 7, 1);
INSERT INTO `sys_users_roles` VALUES (9, 7, 9);
INSERT INTO `sys_users_roles` VALUES (12, 8, 0);
INSERT INTO `sys_users_roles` VALUES (13, 8, 2);
INSERT INTO `sys_users_roles` VALUES (14, 9, 2);
INSERT INTO `sys_users_roles` VALUES (15, 9, 11);
INSERT INTO `sys_users_roles` VALUES (16, 10, 14);
INSERT INTO `sys_users_roles` VALUES (17, 11, 2);
INSERT INTO `sys_users_roles` VALUES (18, 11, 18);
INSERT INTO `sys_users_roles` VALUES (19, 1, 2);
INSERT INTO `sys_users_roles` VALUES (20, 1, 1);
INSERT INTO `sys_users_roles` VALUES (24, 13, 22);
INSERT INTO `sys_users_roles` VALUES (25, 12, 9);

SET FOREIGN_KEY_CHECKS = 1;
