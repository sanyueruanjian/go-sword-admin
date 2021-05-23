/*
 Navicat Premium Data Transfer

 Source Server         : 121.196.160.185
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : 121.196.160.185:3306
 Source Schema         : go_sword

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 20/05/2021 17:53:42
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
) ENGINE = InnoDB AUTO_INCREMENT = 995 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (990, 'p', '1', '/api/dept/', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (977, 'p', '1', '/api/dept/', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (988, 'p', '1', '/api/dept/', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (989, 'p', '1', '/api/dept/', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (993, 'p', '1', '/api/job', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (978, 'p', '1', '/api/job', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (991, 'p', '1', '/api/job', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (992, 'p', '1', '/api/job', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (987, 'p', '1', '/api/menus/', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (985, 'p', '1', '/api/menus/', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (986, 'p', '1', '/api/menus/', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (984, 'p', '1', '/api/roles', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (976, 'p', '1', '/api/roles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (982, 'p', '1', '/api/roles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (983, 'p', '1', '/api/roles', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (994, 'p', '1', '/api/student', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (981, 'p', '1', '/api/users/', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (975, 'p', '1', '/api/users/', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (979, 'p', '1', '/api/users/', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (980, 'p', '1', '/api/users/', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (655, 'p', '2', '/api/users/', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (652, 'p', '2', '/api/users/', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (653, 'p', '2', '/api/users/', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (654, 'p', '2', '/api/users/', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (572, 'p', '9', '/api/dept/', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (559, 'p', '9', '/api/dept/', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (570, 'p', '9', '/api/dept/', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (571, 'p', '9', '/api/dept/', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (575, 'p', '9', '/api/job', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (560, 'p', '9', '/api/job', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (573, 'p', '9', '/api/job', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (574, 'p', '9', '/api/job', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (569, 'p', '9', '/api/menus/', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (567, 'p', '9', '/api/menus/', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (568, 'p', '9', '/api/menus/', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (566, 'p', '9', '/api/roles', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (558, 'p', '9', '/api/roles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (564, 'p', '9', '/api/roles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (565, 'p', '9', '/api/roles', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (563, 'p', '9', '/api/users/', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (557, 'p', '9', '/api/users/', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (561, 'p', '9', '/api/users/', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (562, 'p', '9', '/api/users/', 'PUT', '', '', '');

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
  `create_by` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者',
  `create_time` bigint(20) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `is_deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '状态：1启用（默认）、0禁用',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_pid`(`pid`) USING BTREE COMMENT '普通索引——pid查询部门',
  INDEX `idx_enabled`(`enabled`) USING BTREE COMMENT '普通索引——enabled查询部门'
) ENGINE = InnoDB AUTO_INCREMENT = 65 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (2, 7, 0, '研发部', 3, b'1', 1, 1, 20201213114104, 20201213114106, b'0');
INSERT INTO `sys_dept` VALUES (5, 7, 0, '运维部', 4, b'1', 1, 1, 20201213114136, 20201213114138, b'0');
INSERT INTO `sys_dept` VALUES (6, 8, 0, '测试部', 0, b'1', 1, 1, 20201213114217, 20201213114219, b'0');
INSERT INTO `sys_dept` VALUES (7, 0, 32, '华北分部', 5, b'1', 1, 1, 20201213114258, 1613912934520, b'0');
INSERT INTO `sys_dept` VALUES (8, 0, 2, '华南分部', 7, b'1', 1, 1, 20201213114340, 1613913600273, b'0');
INSERT INTO `sys_dept` VALUES (17, 7, 0, '哈哈哈', 999, b'1', 1, 1, 1613634683926, 1620135672948, b'1');
INSERT INTO `sys_dept` VALUES (18, 8, 0, '哈哈哈', 999, b'1', 1, 1, 1613637943931, 1613886300610, b'0');
INSERT INTO `sys_dept` VALUES (19, 7, 0, '运维部', 999, b'1', 1, 1, 1613640866701, 1620135668846, b'1');
INSERT INTO `sys_dept` VALUES (20, 8, 0, '测试删除', 999, b'1', 1, 1, 1613886421503, 1613887708779, b'1');
INSERT INTO `sys_dept` VALUES (21, 0, 0, '华北分部', 999, b'1', 1, 1, 1613896727952, 1613907290171, b'1');
INSERT INTO `sys_dept` VALUES (22, 0, 0, '哈贝', 999, b'1', 1, 1, 1613907616874, 1613907805818, b'1');
INSERT INTO `sys_dept` VALUES (23, 0, 0, '哈哈哈', 999, b'1', 1, 1, 1613908703795, 1613908836959, b'1');
INSERT INTO `sys_dept` VALUES (24, 0, 0, '华南', 999, b'1', 1, 1, 1613910111447, 1613910118496, b'1');
INSERT INTO `sys_dept` VALUES (25, 0, 0, '华北分部', 999, b'1', 1, 1, 1613912973123, 1613912993826, b'1');
INSERT INTO `sys_dept` VALUES (26, 0, 0, '哈哈哈', 999, b'1', 1, 1, 1616375147080, 1616375151270, b'1');
INSERT INTO `sys_dept` VALUES (27, 0, 0, '零零零零', 999, b'1', 1, 1, 1620112928765, 1620112934351, b'1');
INSERT INTO `sys_dept` VALUES (28, 7, 0, '零零零零', 999, b'1', 1, 1, 1620112944512, 1620135663950, b'1');
INSERT INTO `sys_dept` VALUES (29, 7, 0, '酷酷酷酷酷酷', 999, b'1', 1, 1, 1620112954886, 1620135658488, b'1');
INSERT INTO `sys_dept` VALUES (30, 7, 0, '将军即将', 999, b'1', 1, 1, 1620112967852, 1620135653694, b'1');
INSERT INTO `sys_dept` VALUES (31, 7, 0, '哈哈哈哈哈哈哈', 999, b'1', 1, 1, 1620112978385, 1620135648764, b'1');
INSERT INTO `sys_dept` VALUES (32, 7, 0, '灌灌灌灌灌', 999, b'1', 1, 1, 1620112988539, 1620135642321, b'1');
INSERT INTO `sys_dept` VALUES (33, 7, 0, '顶顶顶顶顶顶顶', 999, b'1', 1, 1, 1620112999591, 1620135637964, b'1');
INSERT INTO `sys_dept` VALUES (34, 7, 0, '嘟嘟嘟嘟嘟嘟', 999, b'1', 1, 1, 1620113009103, 1620135632710, b'1');
INSERT INTO `sys_dept` VALUES (35, 7, 0, '水水水水水水', 999, b'1', 1, 1, 1620113016141, 1620135628436, b'1');
INSERT INTO `sys_dept` VALUES (36, 7, 0, '啊啊啊啊啊啊啊啊啊啊', 999, b'1', 1, 1, 1620113022803, 1620135623966, b'1');
INSERT INTO `sys_dept` VALUES (37, 7, 0, '钱钱钱钱钱钱', 999, b'1', 1, 1, 1620113032221, 1620135619391, b'1');
INSERT INTO `sys_dept` VALUES (38, 7, 0, '哇哇哇哇哇哇哇哇哇哇哇', 999, b'1', 1, 1, 1620113039070, 1620135614280, b'1');
INSERT INTO `sys_dept` VALUES (39, 7, 0, '呃呃呃呃呃呃呃呃', 999, b'1', 1, 1, 1620113045614, 1620135248453, b'1');
INSERT INTO `sys_dept` VALUES (40, 7, 0, '日日日日日日日日日日', 999, b'1', 1, 1, 1620113056907, 1620135243237, b'1');
INSERT INTO `sys_dept` VALUES (41, 7, 0, '她她她她她她她她她', 999, b'1', 1, 1, 1620113064370, 1620135239310, b'1');
INSERT INTO `sys_dept` VALUES (42, 7, 0, '有有有有有有有有有有', 999, b'1', 1, 1, 1620113071621, 1620135234340, b'1');
INSERT INTO `sys_dept` VALUES (43, 7, 0, 'uuuuuu', 999, b'1', 1, 1, 1620113077787, 1620135229287, b'1');
INSERT INTO `sys_dept` VALUES (44, 7, 0, 'iiiiiiiiiiiiiiiiiiii', 999, b'1', 1, 1, 1620113258136, 1620135223845, b'1');
INSERT INTO `sys_dept` VALUES (45, 0, 0, '大委', 999, b'1', 1, 1, 1620121322834, 1620128010399, b'1');
INSERT INTO `sys_dept` VALUES (46, 0, 0, '大伟大·', 999, b'1', 1, 1, 1620121329434, 1620128013608, b'1');
INSERT INTO `sys_dept` VALUES (47, 0, 0, '大伟大··11', 999, b'1', 1, 1, 1620121336735, 1620128015112, b'1');
INSERT INTO `sys_dept` VALUES (48, 0, 0, '大卫奥', 999, b'1', 1, 1, 1620121340915, 1620128016945, b'1');
INSERT INTO `sys_dept` VALUES (49, 0, 0, '让我让我·', 999, b'1', 1, 1, 1620121345333, 1620128018523, b'1');
INSERT INTO `sys_dept` VALUES (50, 0, 0, '的娃儿啊我日', 999, b'1', 1, 1, 1620121350189, 1620128020755, b'1');
INSERT INTO `sys_dept` VALUES (51, 0, 0, '达瓦达瓦', 999, b'1', 1, 1, 1620121354516, 1620128022441, b'1');
INSERT INTO `sys_dept` VALUES (52, 0, 0, '达瓦达瓦他·', 999, b'1', 1, 1, 1620121359309, 1620128024177, b'1');
INSERT INTO `sys_dept` VALUES (53, 0, 0, '爱他他问题·', 999, b'1', 1, 1, 1620121365601, 1620128026127, b'1');
INSERT INTO `sys_dept` VALUES (54, 0, 0, '让外人哇·', 999, b'1', 1, 1, 1620121371776, 1620128028185, b'1');
INSERT INTO `sys_dept` VALUES (55, 7, 0, '6666', 999, b'1', 1, 1, 1620131037083, 1620135218989, b'1');
INSERT INTO `sys_dept` VALUES (56, 7, 0, '啦啦啦啦啦', 999, b'1', 1, 1, 1620184527062, 1620204445826, b'1');
INSERT INTO `sys_dept` VALUES (57, 7, 0, '啦啦啦啦啦', 999, b'1', 1, 1, 1620204533678, 1620206243049, b'1');
INSERT INTO `sys_dept` VALUES (58, 7, 0, '看看看看', 999, b'1', 1, 1, 1620204587953, 1620204781231, b'1');
INSERT INTO `sys_dept` VALUES (59, 7, 0, '斤斤计较急急急急急急', 999, b'1', 1, 1, 1620206281950, 1620206290075, b'1');
INSERT INTO `sys_dept` VALUES (60, 7, 0, '大伟大', 999, b'1', 1, 1, 1620206366252, 1620206420307, b'1');
INSERT INTO `sys_dept` VALUES (61, 7, 0, '大伟大', 999, b'1', 1, 1, 1620206433133, 1620206553374, b'1');
INSERT INTO `sys_dept` VALUES (62, 7, 0, '大伟大', 999, b'1', 1, 1, 1620206611853, 1620206611853, b'0');
INSERT INTO `sys_dept` VALUES (63, 7, 0, '将军即将', 999, b'1', 1, 1, 1620218538180, 1620218538180, b'0');
INSERT INTO `sys_dept` VALUES (64, 7, 0, 'llll', 999, b'1', 1, 1, 1620465266782, 1620465266782, b'0');

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
INSERT INTO `sys_job` VALUES (29, '漫威', b'1', 999, 1, 1, 1613638251200, 1615776677804, b'0');
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
) ENGINE = InnoDB AUTO_INCREMENT = 195 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, NULL, NULL, 0, 7, 0, '系统管理', '系统管理', 'Layout', 1, 'system', 'system', b'0', b'0', b'0', '', 1, 1, 20181218151129, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (2, '/api/users/', 'GET', 1, 3, 1, '用户管理', 'User', 'system/user/index', 2, 'peoples', 'user', b'0', b'0', b'0', 'user:list', 1, 1, 20181218151444, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (3, '/api/roles', 'GET', 1, 3, 1, '角色管理', 'Role', 'system/role/index', 3, 'role', 'role', b'0', b'0', b'0', 'roles:list', 1, 1, 20181218151607, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (5, NULL, NULL, 1, 3, 1, '菜单管理', 'Menu', 'system/menu/index', 5, 'menu', 'menu', b'0', b'0', b'0', 'menu:list', 1, 1, 20181218151728, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (6, NULL, NULL, 0, 4, 0, '系统监控', '系统监控', 'Layout', 10, 'monitor', 'monitor', b'0', b'0', b'0', '', 1, 1, 20181218151748, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (7, NULL, NULL, 6, 0, 1, '操作日志', 'Log', 'monitor/log/index', 11, 'log', 'logs', b'0', b'0', b'0', '', 1, 1, 20181218151826, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (28, NULL, NULL, 1, 3, 1, '任务调度', 'Timing', 'system/timing/index', 999, 'timing', 'timing', b'0', b'0', b'0', 'timing:list', 1, 1, 20190107203440, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (32, NULL, NULL, 6, 0, 1, '异常日志', 'ErrorLog', 'monitor/log/errorLog', 12, 'error', 'errorLog', b'0', b'0', b'0', '', 1, 1, 20190113134903, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (35, '/api/dept/', 'GET', 1, 3, 1, '部门管理', 'Dept', 'system/dept/index', 6, 'dept', 'dept', b'0', b'0', b'0', 'dept:list', 1, 1, 20190325094600, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (37, '/api/job', 'GET', 1, 3, 1, '岗位管理', 'Job', 'system/job/index', 7, 'Steve-Jobs', 'job', b'0', b'0', b'0', 'job:list', 1, 1, 20190329135118, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (39, NULL, NULL, 1, 3, 1, '字典管理', 'Dict', 'system/dict/index', 8, 'dictionary', 'dict', b'0', b'0', b'0', 'dict:list', 1, 1, 20190410114904, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (41, NULL, NULL, 6, 0, 1, '在线用户', 'OnlineUser', 'monitor/online/index', 10, 'Steve-Jobs', 'online', b'0', b'0', b'0', '', 1, 1, 20191026220843, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (44, '/api/users/', 'POST', 2, 0, 2, '用户新增', '', '', 2, '', '', b'0', b'0', b'0', 'user:add', 1, 1, 20191029105946, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (45, '/api/users/', 'PUT', 2, 0, 2, '用户编辑', '', '', 3, '', '', b'0', b'0', b'0', 'user:edit', 1, 1, 20191029110008, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (46, '/api/users/', 'DELETE', 2, 0, 2, '用户删除', '', '', 4, '', '', b'0', b'0', b'0', 'user:del', 1, 1, 20191029110023, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (48, '/api/roles', 'POST', 3, 0, 2, '角色创建', '', '', 2, '', '', b'0', b'0', b'0', 'roles:add', 1, 1, 20191029124534, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (49, '/api/roles', 'PUT', 3, 0, 2, '角色修改', '', '', 3, '', '', b'0', b'0', b'0', 'roles:edit', 1, 1, 20191029124616, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (50, '/api/roles', 'DELETE', 3, 0, 2, '角色删除', '', '', 4, '', '', b'0', b'0', b'0', 'roles:del', 1, 1, 20191029124651, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (52, '/api/menus/', 'POST', 5, 0, 2, '菜单新增', '', '', 2, '', '', b'0', b'0', b'0', 'menu:add', 1, 1, 20191029125507, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (53, '/api/menus/', 'PUT', 5, 0, 2, '菜单编辑', '', '', 3, '', '', b'0', b'0', b'0', 'menu:edit', 1, 1, 20191029125540, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (54, '/api/menus/', 'DELETE', 5, 0, 2, '菜单删除', '', '', 4, '', '', b'0', b'0', b'0', 'menu:del', 1, 1, 20191029125600, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (56, '/api/dept/', 'POST', 35, 0, 2, '部门新增', '', '', 2, '', '', b'0', b'0', b'0', 'dept:add', 1, 1, 20191029125709, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (57, '/api/dept/', 'PUT', 35, 0, 2, '部门编辑', '', '', 3, '', '', b'0', b'0', b'0', 'dept:edit', 1, 1, 20191029125727, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (58, '/api/dept/', 'DELETE', 35, 0, 2, '部门删除', '', '', 4, '', '', b'0', b'0', b'0', 'dept:del', 1, 1, 20191029125741, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (60, '/api/job', 'POST', 37, 0, 2, '岗位新增', '', '', 2, '', '', b'0', b'0', b'0', 'job:add', 1, 1, 20191029125827, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (61, '/api/job', 'PUT', 37, 0, 2, '岗位编辑', '', '', 3, '', '', b'0', b'0', b'0', 'job:edit', 1, 1, 20191029125845, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (62, '/api/job', 'DELETE', 37, 0, 2, '岗位删除', '', '', 4, '', '', b'0', b'0', b'0', 'job:del', 1, 1, 20191029125904, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (64, NULL, NULL, 39, 0, 2, '字典新增', '', '', 2, '', '', b'0', b'0', b'0', 'dict:add', 1, 1, 20191029130017, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (65, NULL, NULL, 39, 0, 2, '字典编辑', '', '', 3, '', '', b'0', b'0', b'0', 'dict:edit', 1, 1, 20191029130042, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (66, NULL, NULL, 39, 0, 2, '字典删除', '', '', 4, '', '', b'0', b'0', b'0', 'dict:del', 1, 1, 20191029130059, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (73, NULL, NULL, 28, 0, 2, '任务新增', '', '', 2, '', '', b'0', b'0', b'0', 'timing:add', 1, 1, 20191029130728, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (74, NULL, NULL, 28, 0, 2, '任务编辑', '', '', 3, '', '', b'0', b'0', b'0', 'timing:edit', 1, 1, 20191029130741, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (75, NULL, NULL, 28, 0, 2, '任务删除', '', '', 4, '', '', b'0', b'0', b'0', 'timing:del', 1, 1, 20191029130754, 20201213162507, b'0');
INSERT INTO `sys_menu` VALUES (179, '', '', 0, 2, 0, '学生管理', '学生管理', 'Layout', 13, 'peoples', 'peoples', b'0', b'0', b'0', '', 1, 1, 1615540666096, 1615540666096, b'1');
INSERT INTO `sys_menu` VALUES (180, '/api/job', 'GET', 179, 0, 0, 'ssssss', '', 'Layout', 999, 'Steve-Jobs', '/api/power', b'0', b'0', b'0', '', 1, 1, 1615813724284, 1615813724284, b'0');
INSERT INTO `sys_menu` VALUES (181, '/api/student', 'GET', 179, 0, 0, 'fff', '', 'Layout', 999, 'app', 'gsdsd', b'0', b'0', b'0', '', 1, 1, 1616226163460, 1616226163460, b'0');
INSERT INTO `sys_menu` VALUES (182, '/aa', 'GET', 0, 0, 0, 'asd', '', 'Layout', 999, 'alipay', 'ad', b'0', b'0', b'0', '', 1, 1, 1616231985279, 1616231985279, b'1');
INSERT INTO `sys_menu` VALUES (183, '/api/ping', 'GET', 0, 0, 0, '测试', '', 'Layout', 999, 'Steve-Jobs', '/text', b'0', b'0', b'0', '', 1, 1, 1616374632947, 1616374632947, b'1');
INSERT INTO `sys_menu` VALUES (184, 'add', 'add', 0, 0, 1, 'ddd', 'dd', 'dd', 999, 'Steve-Jobs', 'dd', b'0', b'0', b'0', 'dd', 1, 1, 1616374721696, 1616374721696, b'1');
INSERT INTO `sys_menu` VALUES (185, '', '', 0, 1, 0, '大萨达', '', 'Layout', 999, 'app', '大大', b'0', b'0', b'0', '', 1, 1, 1618546553401, 1618546553401, b'1');
INSERT INTO `sys_menu` VALUES (186, '', '', 185, 0, 0, '打算打算', '', 'Layout', 999, 'chart', '大萨达多', b'0', b'0', b'0', '', 1, 1, 1618546568874, 1618546568874, b'0');
INSERT INTO `sys_menu` VALUES (187, '', '', 0, 0, 0, '11', '', 'Layout', 999, 'Steve-Jobs', '11', b'0', b'0', b'0', '', 1, 1, 1618547526615, 1618547526615, b'1');
INSERT INTO `sys_menu` VALUES (188, '', '', 188, 1, 1, 'ss', 'Ss', '', 999, 'app', 'ss', b'0', b'0', b'0', 'ss:list', 1, 1, 1618557035101, 1618557035101, b'0');
INSERT INTO `sys_menu` VALUES (189, '', '', 188, 0, 1, 'ss2', '', '', 999, 'app', 'ss2', b'0', b'0', b'0', '', 1, 1, 1618557050855, 1618557050855, b'0');
INSERT INTO `sys_menu` VALUES (190, '', '', 0, 0, 0, 'ss', '', 'Layout', 999, 'chart', 'ss', b'0', b'0', b'0', '', 1, 1, 1618557163205, 1618557163205, b'1');
INSERT INTO `sys_menu` VALUES (191, '', '', 0, 3, 0, '考评管理', '', 'Layout', 999, 'app', 'evaluation', b'0', b'0', b'0', '', 1, 1, 1618559792246, 1618559792246, b'0');
INSERT INTO `sys_menu` VALUES (192, '', '', 191, 0, 1, '考评记录', 'Evalog', 'evaluation/log/index', 999, 'alipay', 'evalog', b'0', b'0', b'0', 'evalog:list', 1, 1, 1618559877170, 1618559877170, b'0');
INSERT INTO `sys_menu` VALUES (193, '', '', 191, 0, 1, '考评类型', 'Evatype', 'evaluation/type/index', 999, 'anq', 'evatype', b'0', b'0', b'0', 'evatype:list', 1, 1, 1618559949578, 1618559949578, b'0');
INSERT INTO `sys_menu` VALUES (194, '', '', 191, 0, 1, '全勤计算', 'Evaoperation', 'evaluation/operation/index', 999, 'chart', 'evaoperation', b'0', b'0', b'0', 'evaoperation:list', 1, 1, 1618560036494, 1618560036494, b'0');

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
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = Compact;

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
INSERT INTO `sys_role` VALUES (26, 'hello', 3, 'ssss', '全部', b'1', 1, 1, 1615777739552, 1615777739552, b'0');

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
) ENGINE = InnoDB AUTO_INCREMENT = 2526 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色菜单关联' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_roles_menus
-- ----------------------------
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
INSERT INTO `sys_roles_menus` VALUES (1774, 1, 9);
INSERT INTO `sys_roles_menus` VALUES (1775, 2, 9);
INSERT INTO `sys_roles_menus` VALUES (1776, 3, 9);
INSERT INTO `sys_roles_menus` VALUES (1777, 5, 9);
INSERT INTO `sys_roles_menus` VALUES (1778, 6, 9);
INSERT INTO `sys_roles_menus` VALUES (1779, 7, 9);
INSERT INTO `sys_roles_menus` VALUES (1780, 28, 9);
INSERT INTO `sys_roles_menus` VALUES (1781, 32, 9);
INSERT INTO `sys_roles_menus` VALUES (1782, 35, 9);
INSERT INTO `sys_roles_menus` VALUES (1783, 37, 9);
INSERT INTO `sys_roles_menus` VALUES (1784, 39, 9);
INSERT INTO `sys_roles_menus` VALUES (1785, 41, 9);
INSERT INTO `sys_roles_menus` VALUES (1786, 44, 9);
INSERT INTO `sys_roles_menus` VALUES (1787, 45, 9);
INSERT INTO `sys_roles_menus` VALUES (1788, 46, 9);
INSERT INTO `sys_roles_menus` VALUES (1789, 48, 9);
INSERT INTO `sys_roles_menus` VALUES (1790, 49, 9);
INSERT INTO `sys_roles_menus` VALUES (1791, 50, 9);
INSERT INTO `sys_roles_menus` VALUES (1792, 52, 9);
INSERT INTO `sys_roles_menus` VALUES (1793, 53, 9);
INSERT INTO `sys_roles_menus` VALUES (1794, 54, 9);
INSERT INTO `sys_roles_menus` VALUES (1795, 56, 9);
INSERT INTO `sys_roles_menus` VALUES (1796, 57, 9);
INSERT INTO `sys_roles_menus` VALUES (1797, 58, 9);
INSERT INTO `sys_roles_menus` VALUES (1798, 60, 9);
INSERT INTO `sys_roles_menus` VALUES (1799, 61, 9);
INSERT INTO `sys_roles_menus` VALUES (1800, 62, 9);
INSERT INTO `sys_roles_menus` VALUES (1801, 64, 9);
INSERT INTO `sys_roles_menus` VALUES (1802, 65, 9);
INSERT INTO `sys_roles_menus` VALUES (1803, 66, 9);
INSERT INTO `sys_roles_menus` VALUES (1804, 73, 9);
INSERT INTO `sys_roles_menus` VALUES (1805, 74, 9);
INSERT INTO `sys_roles_menus` VALUES (1806, 75, 9);
INSERT INTO `sys_roles_menus` VALUES (1807, 179, 9);
INSERT INTO `sys_roles_menus` VALUES (1808, 6, 10);
INSERT INTO `sys_roles_menus` VALUES (1809, 7, 10);
INSERT INTO `sys_roles_menus` VALUES (1810, 32, 10);
INSERT INTO `sys_roles_menus` VALUES (1811, 41, 10);
INSERT INTO `sys_roles_menus` VALUES (1812, 179, 10);
INSERT INTO `sys_roles_menus` VALUES (1813, 6, 12);
INSERT INTO `sys_roles_menus` VALUES (1814, 7, 12);
INSERT INTO `sys_roles_menus` VALUES (1815, 32, 12);
INSERT INTO `sys_roles_menus` VALUES (1816, 41, 12);
INSERT INTO `sys_roles_menus` VALUES (1817, 6, 14);
INSERT INTO `sys_roles_menus` VALUES (1818, 7, 14);
INSERT INTO `sys_roles_menus` VALUES (1819, 32, 14);
INSERT INTO `sys_roles_menus` VALUES (1820, 41, 14);
INSERT INTO `sys_roles_menus` VALUES (1941, 1, 2);
INSERT INTO `sys_roles_menus` VALUES (1942, 2, 2);
INSERT INTO `sys_roles_menus` VALUES (1943, 44, 2);
INSERT INTO `sys_roles_menus` VALUES (1944, 45, 2);
INSERT INTO `sys_roles_menus` VALUES (1945, 46, 2);
INSERT INTO `sys_roles_menus` VALUES (2487, 1, 1);
INSERT INTO `sys_roles_menus` VALUES (2488, 2, 1);
INSERT INTO `sys_roles_menus` VALUES (2489, 3, 1);
INSERT INTO `sys_roles_menus` VALUES (2490, 5, 1);
INSERT INTO `sys_roles_menus` VALUES (2491, 6, 1);
INSERT INTO `sys_roles_menus` VALUES (2492, 7, 1);
INSERT INTO `sys_roles_menus` VALUES (2493, 28, 1);
INSERT INTO `sys_roles_menus` VALUES (2494, 32, 1);
INSERT INTO `sys_roles_menus` VALUES (2495, 35, 1);
INSERT INTO `sys_roles_menus` VALUES (2496, 37, 1);
INSERT INTO `sys_roles_menus` VALUES (2497, 41, 1);
INSERT INTO `sys_roles_menus` VALUES (2498, 44, 1);
INSERT INTO `sys_roles_menus` VALUES (2499, 45, 1);
INSERT INTO `sys_roles_menus` VALUES (2500, 46, 1);
INSERT INTO `sys_roles_menus` VALUES (2501, 48, 1);
INSERT INTO `sys_roles_menus` VALUES (2502, 49, 1);
INSERT INTO `sys_roles_menus` VALUES (2503, 50, 1);
INSERT INTO `sys_roles_menus` VALUES (2504, 52, 1);
INSERT INTO `sys_roles_menus` VALUES (2505, 53, 1);
INSERT INTO `sys_roles_menus` VALUES (2506, 54, 1);
INSERT INTO `sys_roles_menus` VALUES (2507, 56, 1);
INSERT INTO `sys_roles_menus` VALUES (2508, 57, 1);
INSERT INTO `sys_roles_menus` VALUES (2509, 58, 1);
INSERT INTO `sys_roles_menus` VALUES (2510, 60, 1);
INSERT INTO `sys_roles_menus` VALUES (2511, 61, 1);
INSERT INTO `sys_roles_menus` VALUES (2512, 62, 1);
INSERT INTO `sys_roles_menus` VALUES (2513, 73, 1);
INSERT INTO `sys_roles_menus` VALUES (2514, 74, 1);
INSERT INTO `sys_roles_menus` VALUES (2515, 75, 1);
INSERT INTO `sys_roles_menus` VALUES (2516, 179, 1);
INSERT INTO `sys_roles_menus` VALUES (2517, 181, 1);
INSERT INTO `sys_roles_menus` VALUES (2518, 185, 1);
INSERT INTO `sys_roles_menus` VALUES (2519, 186, 1);
INSERT INTO `sys_roles_menus` VALUES (2520, 188, 1);
INSERT INTO `sys_roles_menus` VALUES (2521, 189, 1);
INSERT INTO `sys_roles_menus` VALUES (2522, 191, 1);
INSERT INTO `sys_roles_menus` VALUES (2523, 192, 1);
INSERT INTO `sys_roles_menus` VALUES (2524, 193, 1);
INSERT INTO `sys_roles_menus` VALUES (2525, 194, 1);

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
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统用户' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 7, 'admin', '超级管理员', b'0', '18888888888', '201507802@qq.com', '32778d47-9839-4af9-886d-6cddfb0ba832.png', '47474af9f2fa250735c68a2af27d66b0', b'1', b'1', 1, 1, 1616231882930, 20201213112652, 20201213112657, b'0', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
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
INSERT INTO `sys_user` VALUES (14, 7, 'cgl', '修改一下下', b'0', '15083138896', '1720808104@qq.com', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1614074908040, 1614074908212, 1614074908212, b'0', 0, '', '', '', '', '', 0);
INSERT INTO `sys_user` VALUES (15, 8, 'AA', '法师', b'0', '15083138896', '1720808104@qq.com', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1614136659888, 1614136660083, 1614136660083, b'0', 0, '', '', '', '', '', 0);
INSERT INTO `sys_user` VALUES (16, 6, 'qwe', 'qwer', b'0', '15945632874', '1946117378@qq.com', '', '47474af9f2fa250735c68a2af27d66b0', b'0', b'1', 1, 1, 1615472314360, 1615472314567, 1615472314567, b'0', 0, '', '', '', '', '', 0);

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
) ENGINE = InnoDB AUTO_INCREMENT = 56 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

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
INSERT INTO `sys_users_jobs` VALUES (20, 11, 2);
INSERT INTO `sys_users_jobs` VALUES (21, 11, 29);
INSERT INTO `sys_users_jobs` VALUES (26, 13, 29);
INSERT INTO `sys_users_jobs` VALUES (34, 12, 29);
INSERT INTO `sys_users_jobs` VALUES (37, 16, 2);
INSERT INTO `sys_users_jobs` VALUES (42, 14, 29);
INSERT INTO `sys_users_jobs` VALUES (49, 15, 29);
INSERT INTO `sys_users_jobs` VALUES (50, 15, 2);
INSERT INTO `sys_users_jobs` VALUES (53, 1, 29);
INSERT INTO `sys_users_jobs` VALUES (54, 1, 2);
INSERT INTO `sys_users_jobs` VALUES (55, 10, 29);

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
) ENGINE = InnoDB AUTO_INCREMENT = 47 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户角色关联' ROW_FORMAT = Compact;

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
INSERT INTO `sys_users_roles` VALUES (17, 11, 2);
INSERT INTO `sys_users_roles` VALUES (18, 11, 18);
INSERT INTO `sys_users_roles` VALUES (24, 13, 22);
INSERT INTO `sys_users_roles` VALUES (31, 12, 9);
INSERT INTO `sys_users_roles` VALUES (33, 16, 9);
INSERT INTO `sys_users_roles` VALUES (38, 14, 9);
INSERT INTO `sys_users_roles` VALUES (42, 15, 20);
INSERT INTO `sys_users_roles` VALUES (43, 15, 11);
INSERT INTO `sys_users_roles` VALUES (45, 1, 1);
INSERT INTO `sys_users_roles` VALUES (46, 10, 2);

SET FOREIGN_KEY_CHECKS = 1;
