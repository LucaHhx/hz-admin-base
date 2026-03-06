/*
 Navicat Premium Dump SQL

 Source Server         : dev2
 Source Server Type    : MySQL
 Source Server Version : 80042 (8.0.42)
 Source Host           : 16.163.7.125:3306
 Source Schema         : hab

 Target Server Type    : MySQL
 Target Server Version : 80042 (8.0.42)
 File Encoding         : 65001

 Date: 04/03/2026 17:05:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
COMMIT;


-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_apis_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=479 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
BEGIN;
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (1, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwtManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (2, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/user/deleteUser', '删除用户', 'userManagement', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (3, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/user/admin_register', '用户注册', 'userManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (4, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/user/getUserList', '获取用户列表', 'userManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (5, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/user/setUserInfo', '设置用户信息', 'userManagement', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (6, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/user/setSelfInfo', '设置自身信息(必选)', 'userManagement', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (7, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/user/getUserInfo', '获取自身信息(必选)', 'userManagement', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (8, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/user/setUserAuthorities', '设置权限组', 'userManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (9, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/user/changePassword', '修改密码（建议选择)', 'userManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (10, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/user/setUserAuthority', '修改用户角色(必选)', 'userManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (11, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/user/resetPassword', '重置用户密码', 'userManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (12, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/user/setSelfSetting', '用户界面配置', 'userManagement', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (13, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/api/createApi', '创建api', 'apiManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (14, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/api/deleteApi', '删除Api', 'apiManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (15, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/api/updateApi', '更新Api', 'apiManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (16, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/api/getApiList', '获取api列表', 'apiManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (17, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/api/getAllApis', '获取所有api', 'apiManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (18, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/api/getApiById', '获取api详细信息', 'apiManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (19, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/api/deleteApisByIds', '批量删除api', 'apiManagement', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (20, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/api/syncApi', '获取待同步API', 'apiManagement', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (21, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/api/getApiGroups', '获取路由组', 'apiManagement', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (22, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/api/enterSyncApi', '确认同步API', 'apiManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (23, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/api/ignoreApi', '忽略API', 'apiManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (24, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/authority/copyAuthority', '拷贝角色', 'emailPlugin', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (25, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/authority/createAuthority', '创建角色', 'emailPlugin', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (26, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/authority/deleteAuthority', '删除角色', 'emailPlugin', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (27, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/authority/updateAuthority', '更新角色信息', 'emailPlugin', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (28, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/authority/getAuthorityList', '获取角色列表', 'emailPlugin', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (29, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/authority/setDataAuthority', '设置角色资源权限', 'emailPlugin', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (30, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbinManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (31, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbinManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (32, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/menu/addBaseMenu', '新增菜单', 'menuManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (33, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/menu/getMenu', '获取菜单树(必选)', 'menuManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (34, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/menu/deleteBaseMenu', '删除菜单', 'menuManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (35, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/menu/updateBaseMenu', '更新菜单', 'menuManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (36, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/menu/getBaseMenuById', '根据id获取菜单', 'menuManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (37, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/menu/getMenuList', '分页获取基础menu列表', 'menuManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (38, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', 'menuManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (39, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/menu/getMenuAuthority', '获取指定角色menu', 'menuManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (40, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', 'menuManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (50, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/system/getServerInfo', '获取服务器信息', 'systemConfig', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (51, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/system/getSystemConfig', '获取配置文件内容', 'systemConfig', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (52, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/system/setSystemConfig', '设置配置文件内容', 'systemConfig', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (58, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/getDB', '获取所有数据库', 'codeGenerator', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (59, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/getTables', '获取数据库表', 'codeGenerator', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (60, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/createTemp', '自动化代码', 'codeGenerator', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (61, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/preview', '预览自动化代码', 'codeGenerator', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (62, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/getColumn', '获取所选table的所有字段', 'codeGenerator', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (63, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/installPlugin', '安装插件', 'codeGenerator', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (64, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/pubPlug', '打包插件', 'codeGenerator', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (65, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/createPackage', '配置模板', 'templateConfig', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (66, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/getTemplates', '获取模板文件', 'templateConfig', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (67, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/getPackage', '获取所有模板', 'templateConfig', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (68, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/delPackage', '删除模板', 'templateConfig', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (69, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/getMeta', '获取meta信息', 'autocodeDetail', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (70, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/rollback', '回滚自动生成代码', 'autocodeDetail', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (71, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/getSysHistory', '查询回滚记录', 'autocodeDetail', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (72, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/delSysHistory', '删除回滚记录', 'autocodeDetail', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (73, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/autoCode/addFunc', '增加模板方法', 'autocodeDetail', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (74, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', 'dictionaryManagement', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (75, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', 'dictionaryManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (76, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', 'dictionaryManagement', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (77, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', 'dictionaryManagement', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (78, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', 'dictionaryManagement', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (79, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysDictionary/createSysDictionary', '新增字典', 'dictionaryManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (80, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', 'dictionaryManagement', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (81, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysDictionary/updateSysDictionary', '更新字典', 'dictionaryManagement', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (82, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典（建议选择）', 'dictionaryManagement', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (83, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', 'dictionaryManagement', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (84, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', 'operationHistory', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (85, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', 'operationHistory', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (86, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', 'operationHistory', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (87, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', 'operationHistory', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (88, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', 'operationHistory', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (94, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/authorityBtn/setAuthorityBtn', '设置按钮权限', 'roleManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (95, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', 'roleManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (96, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', 'roleManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (97, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysExportTemplate/createSysExportTemplate', '新增导出模板', 'exportTemplate', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (98, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysExportTemplate/deleteSysExportTemplate', '删除导出模板', 'exportTemplate', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (99, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysExportTemplate/deleteSysExportTemplateByIds', '批量删除导出模板', 'exportTemplate', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (100, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysExportTemplate/updateSysExportTemplate', '更新导出模板', 'exportTemplate', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (101, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysExportTemplate/findSysExportTemplate', '根据ID获取导出模板', 'exportTemplate', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (102, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysExportTemplate/getSysExportTemplateList', '获取导出模板列表', 'exportTemplate', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (103, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysExportTemplate/exportExcel', '导出Excel', 'exportTemplate', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (104, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysExportTemplate/exportTemplate', '下载模板', 'exportTemplate', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (105, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysExportTemplate/importExcel', '导入Excel', 'exportTemplate', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (112, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysParams/createSysParams', '新建参数', 'parameterManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (113, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysParams/deleteSysParams', '删除参数', 'parameterManagement', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (114, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysParams/deleteSysParamsByIds', '批量删除参数', 'parameterManagement', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (115, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysParams/updateSysParams', '更新参数', 'parameterManagement', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (116, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysParams/findSysParams', '根据ID获取参数', 'parameterManagement', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (117, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysParams/getSysParamsList', '获取参数列表', 'parameterManagement', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (118, '2025-03-12 11:33:34.723', '2025-03-12 11:33:34.723', NULL, '/sysParams/getSysParam', '获取参数列表', 'parameterManagement', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (164, '2025-03-13 14:18:35.407', '2025-03-13 14:18:35.407', NULL, '/sysTableColumns/createSysTableColumns', 'add', 'sysTableColumns', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (165, '2025-03-13 14:18:35.414', '2025-03-13 14:18:35.414', NULL, '/sysTableColumns/deleteSysTableColumns', 'delete', 'sysTableColumns', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (166, '2025-03-13 14:18:35.417', '2025-03-13 14:18:35.417', NULL, '/sysTableColumns/deleteSysTableColumnsByIds', 'batchDelete', 'sysTableColumns', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (167, '2025-03-13 14:18:35.419', '2025-03-13 14:18:35.419', NULL, '/sysTableColumns/updateSysTableColumns', 'update', 'sysTableColumns', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (168, '2025-03-13 14:18:35.422', '2025-03-13 14:18:35.422', NULL, '/sysTableColumns/findSysTableColumns', 'find', 'sysTableColumns', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (169, '2025-03-13 14:18:35.424', '2025-03-13 14:18:35.424', NULL, '/sysTableColumns/getSysTableColumnsList', 'getList', 'sysTableColumns', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (170, '2025-03-13 14:18:35.426', '2025-03-13 14:18:35.426', NULL, '/sysTableColumns/importSysTableColumns', 'import', 'sysTableColumns', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (171, '2025-03-13 14:38:26.320', '2025-03-13 14:38:26.320', NULL, '/sysTableColumns/getStructNameColumns', 'byStruct', 'sysTableColumns', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (201, '2025-03-14 03:50:23.315', '2025-03-14 03:50:23.315', NULL, '/authorityCol/getAuthorityCol', 'getAuthorityCol', 'authorityCol', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (202, '2025-03-14 03:50:30.913', '2025-03-14 03:50:30.913', NULL, '/authorityCol/canRemoveAuthorityCol', 'canRemoveAuthorityCol', 'authorityCol', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (203, '2025-03-14 03:50:38.705', '2025-03-14 03:50:38.705', NULL, '/authorityCol/setAuthorityCol', 'setAuthorityCol', 'authorityCol', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (204, '2025-03-14 03:50:52.203', '2025-03-14 03:50:52.203', NULL, '/sysTableColumns/importSysTableColumns', 'importSysTableColumns', 'sysTableColumns', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (205, '2025-03-14 03:51:31.750', '2025-03-14 03:51:31.750', NULL, '/authorityCol/getMenuColumns', 'getMenuColumns', 'authorityCol', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (255, '2025-04-11 03:33:47.703', '2025-04-11 03:33:47.703', NULL, '/sysDataFilter/createSysDataFilter', 'add', 'sysDataFilter', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (256, '2025-04-11 03:33:47.874', '2025-04-11 03:33:47.874', NULL, '/sysDataFilter/deleteSysDataFilter', 'delete', 'sysDataFilter', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (257, '2025-04-11 03:33:48.037', '2025-04-11 03:33:48.037', NULL, '/sysDataFilter/deleteSysDataFilterByIds', 'batchDelete', 'sysDataFilter', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (258, '2025-04-11 03:33:48.203', '2025-04-11 03:33:48.203', NULL, '/sysDataFilter/updateSysDataFilter', 'update', 'sysDataFilter', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (259, '2025-04-11 03:33:48.357', '2025-04-11 03:33:48.357', NULL, '/sysDataFilter/findSysDataFilter', 'find', 'sysDataFilter', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (260, '2025-04-11 03:33:48.509', '2025-04-11 03:33:48.509', NULL, '/sysDataFilter/getSysDataFilterList', 'getList', 'sysDataFilter', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (261, '2025-04-11 03:33:48.671', '2025-04-11 03:33:48.671', NULL, '/sysDataFilter/importSysDataFilter', 'import', 'sysDataFilter', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (437, '2025-07-15 16:57:35.016', '2025-07-15 16:57:35.016', NULL, '/user/getGoogleAuthQR', '获取谷歌验证码', 'userManagement', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (438, '2025-07-15 17:04:49.011', '2025-07-15 17:04:49.011', NULL, '/translation/tree', '加载文件树', 'translation', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (439, '2025-07-15 17:05:03.415', '2025-07-15 17:05:03.415', NULL, '/translation/file', '查看翻译文件', 'translation', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (440, '2025-07-15 17:05:27.263', '2025-07-15 17:05:27.263', NULL, '/translation/update', '更新', 'translation', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (441, '2025-07-15 17:45:43.364', '2025-07-15 17:45:43.364', NULL, '/sysTableColumns/updateSysTableColumnsInfo', '更新系统表列信息', 'sysTableColumns', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (442, '2025-07-15 18:00:45.964', '2025-07-15 18:00:45.964', NULL, '/sysDataFilter/filterData', '过滤条件', 'sysDataFilter', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (466, '2025-07-25 19:21:40.766', '2025-07-25 19:21:40.766', NULL, '/user/bindGoogleAuth', '绑定谷歌验证码', 'userManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (467, '2025-07-25 19:21:54.697', '2025-07-25 19:21:54.697', NULL, '/user/resetGoogleAuth', '重置谷歌验证码', 'userManagement', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (471, '2025-07-25 19:23:35.673', '2025-07-25 19:23:35.673', NULL, '/translation/delete-language', '删除语言', 'translation', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (472, '2025-07-25 19:23:51.471', '2025-07-25 19:23:51.471', NULL, '/translation/export-language', '导出翻译', 'translation', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (473, '2025-07-25 19:24:04.232', '2025-07-25 19:24:04.232', NULL, '/translation/batch', '批量', 'translation', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (474, '2025-07-25 19:24:08.807', '2025-07-25 19:24:08.807', NULL, '/translation/copy', '复制', 'translation', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (475, '2025-07-25 19:24:22.268', '2025-07-25 19:24:22.268', NULL, '/translation/upload-language', '上传翻译', 'translation', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (478, '2025-07-25 19:25:07.138', '2025-07-25 19:25:07.138', NULL, '/translation/download-language', '下载翻译', 'translation', 'GET');
COMMIT;

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'userInfo' COMMENT '默认菜单',
  `type` bigint DEFAULT '0' COMMENT '用户类型 1普通用户 2商户',
  `parameter` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '用户参数',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE KEY `uni_sys_authorities_authority_id` (`authority_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
BEGIN;
INSERT INTO `sys_authorities` (`created_at`, `updated_at`, `deleted_at`, `authority_id`, `authority_name`, `parent_id`, `default_router`, `type`, `parameter`) VALUES ('2025-03-12 11:33:34.739', '2025-12-29 17:17:35.442', NULL, 1, 'Admin', 0, 'userInfo', 0, NULL);
INSERT INTO `sys_authorities` (`created_at`, `updated_at`, `deleted_at`, `authority_id`, `authority_name`, `parent_id`, `default_router`, `type`, `parameter`) VALUES ('2025-07-15 16:55:29.046', '2025-12-09 13:37:24.785', NULL, 55, '超级管理员', 1, 'userInfo', 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_authority_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_btns`;
CREATE TABLE `sys_authority_btns` (
  `authority_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_authority_btns
-- ----------------------------
BEGIN;
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (1, 37, 146);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (1, 37, 148);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (1, 37, 147);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (1, 37, 149);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (1, 37, 151);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (1, 37, 152);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (1, 37, 153);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (1, 37, 154);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (55, 37, 146);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (55, 37, 148);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (55, 37, 147);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (55, 37, 149);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (55, 37, 151);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (55, 37, 152);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (55, 37, 153);
INSERT INTO `sys_authority_btns` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`) VALUES (55, 37, 154);
COMMIT;

-- ----------------------------
-- Table structure for sys_authority_cols
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_cols`;
CREATE TABLE `sys_authority_cols` (
  `authority_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint unsigned DEFAULT NULL COMMENT '菜单按钮ID',
  `sys_table_columns_id` bigint unsigned DEFAULT NULL COMMENT '列ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_authority_cols
-- ----------------------------
BEGIN;
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 141);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 142);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 143);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 144);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 145);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 146);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 147);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 148);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 149);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 150);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 151);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 152);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 153);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 154);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 155);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 156);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 157);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 158);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 159);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 36, NULL, 160);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 161);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 162);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 163);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 164);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 165);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 166);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 167);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 168);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 169);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 170);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 171);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 172);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 190);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 191);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (1, 37, NULL, 192);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 141);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 142);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 143);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 144);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 145);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 146);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 147);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 148);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 149);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 150);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 151);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 152);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 153);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 154);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 155);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 156);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 157);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 158);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 159);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 36, NULL, 160);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 161);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 162);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 163);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 164);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 165);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 166);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 167);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 168);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 169);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 170);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 171);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 172);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 190);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 191);
INSERT INTO `sys_authority_cols` (`authority_id`, `sys_menu_id`, `sys_base_menu_btn_id`, `sys_table_columns_id`) VALUES (55, 37, NULL, 192);
COMMIT;

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (1, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (2, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (3, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (3, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (4, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (4, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (5, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (5, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (6, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (6, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (7, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (7, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (8, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (8, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (9, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (9, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (10, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (10, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (15, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (16, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (17, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (18, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (19, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (20, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (21, 1);

INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (23, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (24, 1);

INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (26, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (27, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (29, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (31, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (31, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (33, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (33, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (34, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (34, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (34, 56);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (34, 57);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (35, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (35, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (37, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (37, 55);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (73, 1);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (73, 55);
COMMIT;

-- ----------------------------
-- Table structure for sys_auto_code_histories
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_histories`;
CREATE TABLE `sys_auto_code_histories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '表名',
  `package` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '模块名/插件名',
  `request` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '前端传入的结构化信息',
  `struct_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '结构体名称',
  `abbreviation` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '结构体名称缩写',
  `business_db` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '业务库',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Struct中文名称',
  `templates` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '模板信息',
  `Injections` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '注入路径',
  `flag` bigint DEFAULT NULL COMMENT '[0:创建,1:回滚]',
  `api_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api表注册内容',
  `menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `export_template_id` bigint unsigned DEFAULT NULL COMMENT '导出模板ID',
  `package_id` bigint unsigned DEFAULT NULL COMMENT '包ID',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_auto_code_histories
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_auto_code_packages
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_packages`;
CREATE TABLE `sys_auto_code_packages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '描述',
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '展示名',
  `template` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '模版',
  `package_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '包名',
  `module` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_auto_code_packages_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_auto_code_packages
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_btns`;
CREATE TABLE `sys_base_menu_btns` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=509 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_base_menu_btns
-- ----------------------------
BEGIN;
INSERT INTO `sys_base_menu_btns` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES (146, '2025-04-11 03:33:49.011', '2025-04-11 03:33:49.011', NULL, 'add', 'add', 37);
INSERT INTO `sys_base_menu_btns` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES (147, '2025-04-11 03:33:49.011', '2025-04-11 03:33:49.011', NULL, 'batchDelete', 'batchDelete', 37);
INSERT INTO `sys_base_menu_btns` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES (148, '2025-04-11 03:33:49.011', '2025-04-11 03:33:49.011', NULL, 'delete', 'delete', 37);
INSERT INTO `sys_base_menu_btns` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES (149, '2025-04-11 03:33:49.011', '2025-04-11 03:33:49.011', NULL, 'edit', 'edit', 37);
INSERT INTO `sys_base_menu_btns` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES (150, '2025-04-11 03:33:49.011', '2025-04-11 03:33:49.011', NULL, 'info', 'info', 37);
INSERT INTO `sys_base_menu_btns` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES (151, '2025-04-11 03:33:49.011', '2025-04-11 03:33:49.011', NULL, 'exportTemplate', 'exportTemplate', 37);
INSERT INTO `sys_base_menu_btns` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES (152, '2025-04-11 03:33:49.011', '2025-04-11 03:33:49.011', NULL, 'exportExcel', 'exportExcel', 37);
INSERT INTO `sys_base_menu_btns` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES (153, '2025-04-11 03:33:49.011', '2025-04-11 03:33:49.011', NULL, 'importExcel', 'importExcel', 37);
INSERT INTO `sys_base_menu_btns` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES (154, '2025-04-11 03:33:49.011', '2025-04-11 03:33:49.011', NULL, 'columnConfig', 'columnConfig', 37);
INSERT INTO `sys_base_menu_btns` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES (508, '2025-07-31 03:59:05.120', '2025-07-31 03:59:05.120', NULL, 'copy', '复制按钮', 37);
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint unsigned DEFAULT NULL,
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `transition_type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (1, '2025-03-12 11:33:34.783', '2025-07-15 17:16:03.726', NULL, 0, 0, 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 99, '', 0, 0, 'dashboard', 'odometer', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (2, '2025-03-12 11:33:34.783', '2025-07-15 17:15:17.045', NULL, 0, 0, 'about', 'about', 0, 'view/about/index.vue', 99, '', 0, 0, 'aboutUs', 'info-filled', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (3, '2025-03-12 11:33:34.783', '2025-07-15 18:15:18.437', NULL, 0, 0, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 1, '', 0, 0, 'superAdmin', 'user-filled', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (4, '2025-03-12 11:33:34.783', '2025-07-15 18:27:24.611', NULL, 0, 3, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, '', 0, 0, 'roleManagement', 'setting', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (5, '2025-03-12 11:33:34.783', '2025-07-15 18:26:39.913', NULL, 0, 3, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 4, '', 1, 0, 'menuManagement', 'tickets', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (6, '2025-03-12 11:33:34.783', '2025-07-15 18:27:01.993', NULL, 0, 3, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, '', 1, 0, 'apiManagement', 'rank', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (7, '2025-03-12 11:33:34.783', '2025-07-15 18:27:14.523', NULL, 0, 3, 'user', 'user', 0, 'view/superAdmin/user/user.vue', 2, '', 0, 0, 'userManagement', 'user', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (8, '2025-03-12 11:33:34.783', '2025-03-12 11:33:34.783', NULL, 0, 3, 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, '', 0, 0, 'dictionaryManagement', 'notebook', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (9, '2025-03-12 11:33:34.783', '2025-03-12 11:33:34.783', NULL, 0, 3, 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, '', 0, 0, 'operationHistory', 'pie-chart', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (10, '2025-03-12 11:33:34.783', '2025-07-15 17:15:54.328', NULL, 0, 0, 'person', 'person', 1, 'view/person/person.vue', 99, '', 0, 0, 'personalInfo', 'message', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (15, '2025-03-12 11:33:34.783', '2025-07-15 17:15:44.539', NULL, 0, 0, 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 99, '', 0, 0, 'systemTools', 'tools', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (16, '2025-03-12 11:33:34.783', '2025-03-12 11:33:34.783', NULL, 0, 15, 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, '', 1, 0, 'codeGenerator', 'cpu', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (17, '2025-03-12 11:33:34.783', '2025-03-12 11:33:34.783', NULL, 0, 15, 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 3, '', 1, 0, 'formGenerator', 'magic-stick', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (18, '2025-03-12 11:33:34.783', '2025-03-12 11:33:34.783', NULL, 0, 15, 'system', 'system', 0, 'view/systemTools/system/system.vue', 4, '', 0, 0, 'systemConfig', 'operation', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (19, '2025-03-12 11:33:34.783', '2025-03-12 11:33:34.783', NULL, 0, 15, 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 2, '', 0, 0, 'autocodeAdmin', 'magic-stick', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (20, '2025-03-12 11:33:34.783', '2025-03-12 11:33:34.783', NULL, 0, 15, 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, '', 0, 0, 'autocodeDetail', 'magic-stick', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (21, '2025-03-12 11:33:34.783', '2025-03-12 11:33:34.783', NULL, 0, 15, 'autoPkg', 'autoPkg', 0, 'view/systemTools/autoPkg/autoPkg.vue', 0, '', 0, 0, 'templateConfig', 'folder', 0, NULL);

INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (23, '2025-03-12 11:33:34.783', '2025-07-15 17:15:23.544', NULL, 0, 0, 'state', 'state', 0, 'view/system/state.vue', 99, '', 0, 0, 'serverStatus', 'cloudy', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (24, '2025-03-12 11:33:34.783', '2025-07-15 17:15:37.271', NULL, 0, 0, 'plugin', 'plugin', 0, 'view/routerHolder.vue', 99, '', 0, 0, 'pluginSystem', 'cherry', 0, NULL);

INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (26, '2025-03-12 11:33:34.783', '2025-03-12 11:33:34.783', NULL, 0, 24, 'installPlugin', 'installPlugin', 0, 'view/systemTools/installPlugin/index.vue', 1, '', 0, 0, 'pluginInstall', 'box', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (27, '2025-03-12 11:33:34.783', '2025-03-12 11:33:34.783', NULL, 0, 24, 'pubPlug', 'pubPlug', 0, 'view/systemTools/pubPlug/pubPlug.vue', 3, '', 0, 0, 'pluginPacking', 'files', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (29, '2025-03-12 11:33:34.783', '2025-03-12 11:33:34.783', NULL, 0, 15, 'exportTemplate', 'exportTemplate', 0, 'view/systemTools/exportTemplate/exportTemplate.vue', 5, '', 0, 0, 'exportTemplate', 'reading', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (31, '2025-03-12 11:33:34.783', '2025-03-12 11:33:34.783', NULL, 0, 3, 'sysParams', 'sysParams', 0, 'view/superAdmin/params/sysParams.vue', 7, '', 0, 0, 'parameterManagement', 'compass', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (33, '2025-03-13 14:18:35.433', '2025-07-15 18:26:06.188', NULL, 0, 3, 'sysTableColumns', 'sysTableColumns', 0, 'view/system/sysTableColumns/sysTableColumns.vue', 10, '', 0, 0, 'sysTableColumns', 'list', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (34, '2025-03-18 03:49:20.437', '2025-07-15 18:14:58.894', NULL, 0, 0, 'userInfo', 'userInfo', 0, 'view/person/person.vue', 0, '', 0, 0, 'userInfo', 'home-filled', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (35, '2025-03-19 07:49:33.111', '2025-07-15 18:25:04.985', NULL, 0, 3, 'index', 'index', 0, 'view/systemTools/translation/index.vue', 15, '', 0, 0, 'translation', 'refresh', 0, NULL);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (37, '2025-04-11 03:33:48.933', '2025-08-01 11:38:21.205', NULL, 0, 73, 'sysDataFilter', 'sysDataFilter', 0, 'view/system/sysDataFilter/sysDataFilter.vue', 11, '', 0, 0, 'sysDataFilter', '', 0, '');
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`) VALUES (73, '2025-07-25 19:52:06.806', '2025-07-31 11:46:08.666', NULL, 0, 0, 'sys', 'sys', 0, 'view/routerHolder.vue', 2, '', 1, 0, 'sys', 'compass', 0, '');
COMMIT;

-- ----------------------------
-- Table structure for sys_bind_sessions
-- ----------------------------
DROP TABLE IF EXISTS `sys_bind_sessions`;
CREATE TABLE `sys_bind_sessions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `session_token` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '会话token',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户名',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户昵称',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_sys_bind_sessions_session_token` (`session_token`),
  KEY `idx_sys_bind_sessions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_bind_sessions
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
BEGIN;
INSERT INTO `sys_data_authority_id` (`sys_authority_authority_id`, `data_authority_id_authority_id`) VALUES (1, 1);
INSERT INTO `sys_data_authority_id` (`sys_authority_authority_id`, `data_authority_id_authority_id`) VALUES (55, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_data_filter
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_filter`;
CREATE TABLE `sys_data_filter` (
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sql` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `columns` json DEFAULT NULL,
  `order_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `note` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint unsigned DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_data_filter_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_data_filter
-- ----------------------------
BEGIN;
INSERT INTO `sys_data_filter` (`name`, `sql`, `columns`, `order_by`, `note`, `key`, `value`, `label`, `id`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`, `deleted_by`) VALUES ('用户', 'SELECT id,`name`,note FROM sys_data_filter', '[{\"sort\": 0, \"label\": \"编号\", \"filter\": true, \"isShow\": true, \"columnName\": \"id\"}, {\"sort\": 0, \"label\": \"名称\", \"filter\": true, \"isShow\": true, \"columnName\": \"name\"}, {\"sort\": 0, \"label\": \"备注\", \"filter\": true, \"isShow\": true, \"columnName\": \"note\"}]', '', '测试', 'id', 'id', 'name', 1, '2025-04-11 07:24:50.000', '2025-12-24 16:44:12.981', NULL, 1, 12, 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
BEGIN;
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (1, '2025-03-12 11:33:34.753', '2025-03-12 11:33:34.758', NULL, '性别', 'gender', 1, '性别字典');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (2, '2025-03-12 11:33:34.753', '2025-03-12 11:33:34.762', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (3, '2025-03-12 11:33:34.753', '2025-03-12 11:33:34.767', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (4, '2025-03-12 11:33:34.753', '2025-03-12 11:33:34.771', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (5, '2025-03-12 11:33:34.753', '2025-03-12 11:33:34.775', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (6, '2025-03-12 11:33:34.753', '2025-03-12 11:33:34.779', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (7, '2025-03-18 03:20:44.403', '2025-03-18 03:20:44.403', NULL, '语言', 'Language', 1, '语言');
COMMIT;

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '展示值',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
BEGIN;
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (1, '2025-03-12 11:33:34.759', '2025-03-12 11:33:34.759', NULL, '男', '1', '', 1, 1, 1);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (2, '2025-03-12 11:33:34.759', '2025-03-12 11:33:34.759', NULL, '女', '2', '', 1, 2, 1);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (3, '2025-03-12 11:33:34.763', '2025-03-12 11:33:34.763', NULL, 'smallint', '1', 'mysql', 1, 1, 2);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (4, '2025-03-12 11:33:34.763', '2025-03-12 11:33:34.763', NULL, 'mediumint', '2', 'mysql', 1, 2, 2);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (5, '2025-03-12 11:33:34.763', '2025-03-12 11:33:34.763', NULL, 'int', '3', 'mysql', 1, 3, 2);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (6, '2025-03-12 11:33:34.763', '2025-03-12 11:33:34.763', NULL, 'bigint', '4', 'mysql', 1, 4, 2);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (7, '2025-03-12 11:33:34.763', '2025-03-12 11:33:34.763', NULL, 'int2', '5', 'pgsql', 1, 5, 2);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (8, '2025-03-12 11:33:34.763', '2025-03-12 11:33:34.763', NULL, 'int4', '6', 'pgsql', 1, 6, 2);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (9, '2025-03-12 11:33:34.763', '2025-03-12 11:33:34.763', NULL, 'int6', '7', 'pgsql', 1, 7, 2);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (10, '2025-03-12 11:33:34.763', '2025-03-12 11:33:34.763', NULL, 'int8', '8', 'pgsql', 1, 8, 2);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (11, '2025-03-12 11:33:34.767', '2025-03-12 11:33:34.767', NULL, 'date', '', '', 1, 0, 3);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (12, '2025-03-12 11:33:34.767', '2025-03-12 11:33:34.767', NULL, 'time', '1', 'mysql', 1, 1, 3);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (13, '2025-03-12 11:33:34.767', '2025-03-12 11:33:34.767', NULL, 'year', '2', 'mysql', 1, 2, 3);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (14, '2025-03-12 11:33:34.767', '2025-03-12 11:33:34.767', NULL, 'datetime', '3', 'mysql', 1, 3, 3);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (15, '2025-03-12 11:33:34.767', '2025-03-12 11:33:34.767', NULL, 'timestamp', '5', 'mysql', 1, 5, 3);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (16, '2025-03-12 11:33:34.767', '2025-03-12 11:33:34.767', NULL, 'timestamptz', '6', 'pgsql', 1, 5, 3);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (17, '2025-03-12 11:33:34.772', '2025-03-12 11:33:34.772', NULL, 'float', '', '', 1, 0, 4);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (18, '2025-03-12 11:33:34.772', '2025-03-12 11:33:34.772', NULL, 'double', '1', 'mysql', 1, 1, 4);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (19, '2025-03-12 11:33:34.772', '2025-03-12 11:33:34.772', NULL, 'decimal', '2', 'mysql', 1, 2, 4);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (20, '2025-03-12 11:33:34.772', '2025-03-12 11:33:34.772', NULL, 'numeric', '3', 'pgsql', 1, 3, 4);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (21, '2025-03-12 11:33:34.772', '2025-03-12 11:33:34.772', NULL, 'smallserial', '4', 'pgsql', 1, 4, 4);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (22, '2025-03-12 11:33:34.776', '2025-03-12 11:33:34.776', NULL, 'char', '', '', 1, 0, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (23, '2025-03-12 11:33:34.776', '2025-03-12 11:33:34.776', NULL, 'varchar', '1', 'mysql', 1, 1, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (24, '2025-03-12 11:33:34.776', '2025-03-12 11:33:34.776', NULL, 'tinyblob', '2', 'mysql', 1, 2, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (25, '2025-03-12 11:33:34.776', '2025-03-12 11:33:34.776', NULL, 'tinytext', '3', 'mysql', 1, 3, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (26, '2025-03-12 11:33:34.776', '2025-03-12 11:33:34.776', NULL, 'text', '4', 'mysql', 1, 4, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (27, '2025-03-12 11:33:34.776', '2025-03-12 11:33:34.776', NULL, 'blob', '5', 'mysql', 1, 5, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (28, '2025-03-12 11:33:34.776', '2025-03-12 11:33:34.776', NULL, 'mediumblob', '6', 'mysql', 1, 6, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (29, '2025-03-12 11:33:34.776', '2025-03-12 11:33:34.776', NULL, 'mediumtext', '7', 'mysql', 1, 7, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (30, '2025-03-12 11:33:34.776', '2025-03-12 11:33:34.776', NULL, 'longblob', '8', 'mysql', 1, 8, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (31, '2025-03-12 11:33:34.776', '2025-03-12 11:33:34.776', NULL, 'longtext', '9', 'mysql', 1, 9, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (32, '2025-03-12 11:33:34.780', '2025-03-12 11:33:34.780', NULL, 'tinyint', '1', 'mysql', 1, 0, 6);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (33, '2025-03-12 11:33:34.780', '2025-03-12 11:33:34.780', NULL, 'bool', '2', 'pgsql', 1, 0, 6);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (34, '2025-03-18 03:22:00.725', '2025-03-18 03:30:44.538', NULL, 'CN', 'zh-CN', '', 1, 1, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES (35, '2025-03-18 03:22:16.618', '2025-03-18 03:30:50.911', NULL, 'EN', 'en-US', '', 1, 2, 7);
COMMIT;

-- ----------------------------
-- Table structure for sys_export_template_condition
-- ----------------------------
DROP TABLE IF EXISTS `sys_export_template_condition`;
CREATE TABLE `sys_export_template_condition` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '模板标识',
  `from` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '条件取的key',
  `column` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '作为查询条件的字段',
  `operator` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作符',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_export_template_condition_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_export_template_condition
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_export_template_join
-- ----------------------------
DROP TABLE IF EXISTS `sys_export_template_join`;
CREATE TABLE `sys_export_template_join` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '模板标识',
  `joins` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '关联',
  `table` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '关联表',
  `on` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '关联条件',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_export_template_join_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_export_template_join
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_export_templates
-- ----------------------------
DROP TABLE IF EXISTS `sys_export_templates`;
CREATE TABLE `sys_export_templates` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `db_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '数据库名称',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '模板名称',
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '表名称',
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '模板标识',
  `template_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `limit` bigint DEFAULT NULL COMMENT '导出限制',
  `order` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_export_templates_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_export_templates
-- ----------------------------
BEGIN;
INSERT INTO `sys_export_templates` (`id`, `created_at`, `updated_at`, `deleted_at`, `db_name`, `name`, `table_name`, `template_id`, `template_info`, `limit`, `order`) VALUES (1, '2025-03-12 11:33:34.945', '2025-03-12 11:33:34.945', NULL, '', 'api', 'sys_apis', 'api', '{\n\"path\":\"路径\",\n\"method\":\"方法（大写）\",\n\"description\":\"方法介绍\",\n\"api_group\":\"方法分组\"\n}', NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for sys_ignore_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_ignore_apis`;
CREATE TABLE `sys_ignore_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api路径',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_ignore_apis_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_ignore_apis
-- ----------------------------
BEGIN;
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (1, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/swagger/*any', 'GET');
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (2, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/api/freshCasbin', 'GET');
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (3, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/uploads/file/*filepath', 'GET');
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (4, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/health', 'GET');
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (5, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/uploads/file/*filepath', 'HEAD');
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (6, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/autoCode/llmAuto', 'POST');
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (7, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/system/reloadSystem', 'POST');
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (8, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/base/login', 'POST');
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (9, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/base/captcha', 'POST');
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (10, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/init/initdb', 'POST');
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (11, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/init/checkdb', 'POST');
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (12, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/info/getInfoDataSource', 'GET');
INSERT INTO `sys_ignore_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `method`) VALUES (13, '2025-03-12 11:33:34.729', '2025-03-12 11:33:34.729', NULL, '/info/getInfoPublic', 'GET');
COMMIT;

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求路径',
  `status` bigint DEFAULT NULL COMMENT '请求状态',
  `latency` bigint DEFAULT NULL COMMENT '延迟',
  `agent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '响应Body',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_params
-- ----------------------------
DROP TABLE IF EXISTS `sys_params`;
CREATE TABLE `sys_params` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '参数名称',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '参数键',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '参数值',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '参数说明',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_params_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_params
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_table_columns
-- ----------------------------
DROP TABLE IF EXISTS `sys_table_columns`;
CREATE TABLE `sys_table_columns` (
  `struct_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '结构名称',
  `json_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'json名称',
  `column_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '数据库名称',
  `with` int DEFAULT NULL COMMENT '宽度',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '类型',
  `sortable` tinyint(1) DEFAULT NULL COMMENT '排序',
  `filter` tinyint(1) DEFAULT NULL COMMENT '筛选',
  `default_filter` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '默认筛选',
  `filter_list` json DEFAULT NULL COMMENT '筛选列表',
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `sort` bigint DEFAULT NULL COMMENT '排序',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `note` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '备注',
  `is_show` tinyint DEFAULT NULL COMMENT '是否显示',
  `enum` json DEFAULT NULL COMMENT '枚举',
  `menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `tb_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '表名称',
  `fixed` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '固定',
  `format` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '格式',
  `is_add_search` tinyint DEFAULT NULL COMMENT '是否添加搜索',
  `search_width` int DEFAULT NULL COMMENT '搜索宽度',
  `is_additional` tinyint DEFAULT NULL COMMENT '是否附加',
  `filter_id` bigint unsigned DEFAULT NULL COMMENT '筛选ID',
  `form_disabled` tinyint DEFAULT NULL COMMENT '表单禁用',
  `form_hidden` tinyint DEFAULT NULL COMMENT '表单隐藏',
  `form_with` bigint DEFAULT NULL COMMENT '表单宽度',
  `form_order` bigint DEFAULT NULL COMMENT '表单排序',
  `is_having` tinyint DEFAULT NULL COMMENT '是否having',
  `form_must` tinyint DEFAULT NULL COMMENT '表单必填',
  `editor_filter_id` bigint DEFAULT NULL COMMENT '编辑器筛选ID',
  `is_sum` tinyint DEFAULT NULL COMMENT '是否求和',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_table_col` (`struct_name`,`json_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=193 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_table_columns
-- ----------------------------
BEGIN;
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'name', 'name', 160, 'string', 1, 1, 'like', '[\"=\", \"!=\", \"like\", \"in\", \"not in\"]', 161, 1, NULL, NULL, 0, 1, 0, '', 1, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'sql', 'sql', 450, 'richtext', 1, 1, 'like', '[\"=\", \"!=\", \"like\", \"in\", \"not in\"]', 162, 4, NULL, NULL, 0, 1, 0, '', 1, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'columns', 'columns', 200, 'json', 1, 1, 'like', '[\"=\", \"!=\", \"like\"]', 163, 5, NULL, NULL, 0, 1, 0, '', 0, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'orderBy', 'order_by', 150, 'string', 1, 1, 'like', '[\"=\", \"!=\", \"like\", \"in\", \"not in\"]', 164, 6, NULL, NULL, 0, 1, 0, '', 1, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'note', 'note', 450, 'richtext', 1, 1, 'like', '[\"=\", \"!=\", \"like\", \"in\", \"not in\"]', 165, 7, NULL, NULL, 0, 1, 0, '', 1, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'ID', 'id', 120, 'int64', 1, 1, '=', '[\"=\", \"!=\", \">\", \">=\", \"<\", \"<=\", \"in\", \"not in\", \"between\", \"not between\"]', 166, 0, NULL, NULL, 0, 1, 0, '主键ID', 1, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'CreatedAt', 'created_at', 150, 'datetime', 1, 1, 'between', '[\"=\", \"!=\", \">\", \">=\", \"<\", \"<=\", \"between\", \"not between\"]', 167, 999, NULL, NULL, 0, 1, 0, '创建时间', 0, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'UpdatedAt', 'updated_at', 170, 'datetime', 1, 1, 'between', '[\"=\", \"!=\", \">\", \">=\", \"<\", \"<=\", \"between\", \"not between\"]', 168, 99, NULL, NULL, 0, 1, 0, '修改时间', 1, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'DeletedAt', 'deleted_at', 150, 'datetime', 1, 1, 'between', '[\"=\", \"!=\", \">\", \">=\", \"<\", \"<=\", \"between\", \"not between\"]', 169, 999, NULL, NULL, 0, 1, 0, '删除时间', 0, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'CreatedBy', 'created_by', 120, 'int64', 1, 1, '=', '[\"=\", \"!=\", \">\", \">=\", \"<\", \"<=\", \"between\", \"not between\"]', 170, 999, NULL, NULL, 0, 1, 0, '创建人', 0, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'UpdatedBy', 'updated_by', 120, 'int64', 1, 1, '=', '[\"=\", \"!=\", \">\", \">=\", \"<\", \"<=\", \"between\", \"not between\"]', 171, 999, NULL, NULL, 0, 1, 0, '修改人', 0, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'DeletedBy', 'deleted_by', 120, 'int64', 1, 1, '=', '[\"=\", \"!=\", \">\", \">=\", \"<\", \"<=\", \"between\", \"not between\"]', 172, 999, NULL, NULL, 0, 1, 0, '删除人', 0, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'key', 'key', 150, 'string', 1, 1, '=', '[\"=\", \"!=\", \"like\", \"in\", \"not in\"]', 190, 2, NULL, NULL, 1, 1, 0, '', 1, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'value', 'value', 160, 'string', 1, 1, '=', '[\"=\", \"!=\", \"like\", \"in\", \"not in\"]', 191, 1, NULL, NULL, 1, 1, 0, '', 1, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
INSERT INTO `sys_table_columns` (`struct_name`, `json_name`, `column_name`, `with`, `type`, `sortable`, `filter`, `default_filter`, `filter_list`, `id`, `sort`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_by`, `note`, `is_show`, `enum`, `menu_id`, `tb_name`, `fixed`, `format`, `is_add_search`, `search_width`, `is_additional`, `filter_id`, `form_disabled`, `form_hidden`, `form_with`, `form_order`, `is_having`, `form_must`, `editor_filter_id`, `is_sum`) VALUES ('sysDataFilter', 'label', 'label', 160, 'string', 1, 1, '=', '[\"=\", \"!=\", \"like\", \"in\", \"not in\"]', 192, 3, NULL, NULL, 1, 1, 0, '', 1, '[]', 37, 'sys_data_filter', '', '', 0, 120, 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES (1, 888);
COMMIT;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'https://qmplusimg.henrongyi.top/HAB_header.jpg' COMMENT '用户头像',
  `authority_id` bigint unsigned DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  `origin_setting` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '配置',
  `type` bigint DEFAULT '0' COMMENT '用户类型 1普通用户 2商户',
  `parameter` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '用户参数',
  `language` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'zh-CN' COMMENT '用户语言',
  `side_mode` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'dark' COMMENT '用户角色ID',
  `active_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '#1890ff' COMMENT '用户角色ID',
  `base_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '#fff' COMMENT '用户角色ID',
  `google_auth_enabled` tinyint(1) DEFAULT '0' COMMENT '是否启用谷歌验证器',
  `google_auth_secret` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '谷歌验证器密钥',
  `passkey_enabled` tinyint(1) DEFAULT '0' COMMENT '是否启用通行密钥',
  `passkey_credentials` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '通行密钥凭证JSON',
  `passkey` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户Passkey',
  `test_passkey` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户测试Passkey',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_users_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_sys_users_uuid` (`uuid`) USING BTREE,
  KEY `idx_sys_users_username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
BEGIN;
INSERT INTO `sys_users` (`id`, `created_at`, `updated_at`, `deleted_at`, `uuid`, `username`, `password`, `nick_name`, `header_img`, `authority_id`, `phone`, `email`, `enable`, `origin_setting`, `type`, `parameter`, `language`, `side_mode`, `active_color`, `base_color`, `google_auth_enabled`, `google_auth_secret`, `passkey_enabled`, `passkey_credentials`, `passkey`, `test_passkey`) VALUES (1, '2025-03-12 11:33:34.933', '2025-12-16 14:11:06.743', NULL, 'c32f5f84-fd16-4362-8af6-b32f658ab07b', 'admin', '$2a$10$HKHDUgjJzKuFziBcf9KkKe7Myu2uGtwQTp.Yg6F875NbolfemHh5W', 'Admin', '', 1, '17611111111', '333333333@qq.com', 1, '{\"darkMode\":\"auto\",\"grey\":false,\"language\":\"zh-CN\",\"layout_side_collapsed_width\":80,\"layout_side_item_height\":48,\"layout_side_width\":256,\"messages\":{},\"primaryColor\":\"#3b82f6\",\"showTabs\":true,\"show_watermark\":false,\"side_mode\":\"normal\",\"table_hight\":800,\"table_hight_enable\":true,\"transition_type\":\"slide\",\"weakness\":false}', 0, '2', 'zh-CN', 'dark', '#1890ff', '#fff', 0, '', 0, NULL, '0THX6Ysf8CgkoCjrReNm4g', 'gmzWGxRu_Ps16LQl8pbIbw');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
