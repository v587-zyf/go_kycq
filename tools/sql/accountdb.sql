/*
 Navicat MySQL Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50712
 Source Host           : localhost:3306
 Source Schema         : cq_accountdb

 Target Server Type    : MySQL
 Target Server Version : 50712
 File Encoding         : 65001

 Date: 27/07/2022 10:32:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `openId` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账号名',
  `channelId` int(11) NULL DEFAULT 0 COMMENT '渠道id',
  `loginCount` int(11) NULL DEFAULT 0,
  `status` int(2) NULL DEFAULT 0 COMMENT '账号状态  0正常 ，1普通封号，2设备封号',
  `freeze` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '封号备注',
  `lastLoginServerId` int(11) NULL DEFAULT -1 COMMENT '最后登录服务器ID',
  `lastLoginTime` datetime NULL DEFAULT '0000-00-00 00:00:00',
  `createIP` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '创建ip',
  `createTime` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `InviterOpenId` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '英雄贴-邀请者openId',
  `banInfo` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '封禁信息',
  INDEX `idx_id`(`id`) USING BTREE,
  INDEX `account_openId`(`openId`) USING BTREE,
  INDEX `idx_createTime`(`createTime`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1182 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'shard_key \"openId\"' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of account
-- ----------------------------


-- ----------------------------
-- Table structure for announcement
-- ----------------------------
DROP TABLE IF EXISTS `announcement`;
CREATE TABLE `announcement`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gmId` int(11) NOT NULL DEFAULT 0,
  `type` int(11) NOT NULL DEFAULT 1,
  `title` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '公告名称',
  `startTime` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '开始时间',
  `endTime` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '结束时间',
  `channelIds` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `serverIds` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `Announcement` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of announcement
-- ----------------------------

-- ----------------------------
-- Table structure for challenge
-- ----------------------------
DROP TABLE IF EXISTS `challenge`;
CREATE TABLE `challenge`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `season` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `userId` int(11) NOT NULL DEFAULT 0,
  `openId` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `nickName` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `crossFsId` int(11) NOT NULL DEFAULT 0,
  `serverId` int(11) NOT NULL DEFAULT 0,
  `combat` bigint(20) NOT NULL DEFAULT 0,
  `isLose` int(11) NOT NULL DEFAULT 0,
  `loseRound` int(11) NOT NULL DEFAULT 0,
  `winUserId` int(11) NOT NULL DEFAULT 0,
  `fightUserInfo` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '战斗玩家数据',
  `round` int(11) NOT NULL DEFAULT 0,
  `expireTime` bigint(20) NOT NULL DEFAULT 0,
  `challenge`  guildName varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of challenge
-- ----------------------------

-- ----------------------------
-- Table structure for challenge_data
-- ----------------------------
DROP TABLE IF EXISTS `challenge_data`;
CREATE TABLE `challenge_data`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `season` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `crossFsId` int(11) NOT NULL DEFAULT 0,
  `round` int(11) NOT NULL DEFAULT 0,
  `userIds` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `expireTime` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of challenge_data
-- ----------------------------

-- ----------------------------
-- Table structure for cross_redis
-- ----------------------------
DROP TABLE IF EXISTS `cross_redis`;
CREATE TABLE `cross_redis`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `network` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `address` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `password` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `db` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '跨服redis服务器配置,创建者:liuxsh' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of cross_redis
-- ----------------------------
INSERT INTO `cross_redis` VALUES (1, 'tcp', '127.0.0.1:6379', '123456', 0);

-- ----------------------------
-- Table structure for crossfight_server_info
-- ----------------------------
DROP TABLE IF EXISTS `crossfight_server_info`;
CREATE TABLE `crossfight_server_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `host` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '127.0.0.1' COMMENT '内网服务器ip',
  `gatePort` int(5) NOT NULL DEFAULT 0 COMMENT '开放给GateServer的端口',
  `gsPort` int(5) NOT NULL DEFAULT 0 COMMENT '开放给GameServer的端口',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 100001 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '跨服战斗服关系表,创建者:liuxsh' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of crossfight_server_info
-- ----------------------------
INSERT INTO `crossfight_server_info` VALUES (100000, '127.0.0.1', 19501, 19502);

-- ----------------------------
-- Table structure for func_close
-- ----------------------------
DROP TABLE IF EXISTS `func_close`;
CREATE TABLE `func_close`  (
  `funcId` int(5) NOT NULL COMMENT '功能Id',
  `serverIds` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '[]' COMMENT '关闭功能模块的服务器Id(空位全服全部)',
  PRIMARY KEY (`funcId`) USING BTREE,
  UNIQUE INDEX `唯一`(`funcId`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of func_close
-- ----------------------------

-- ----------------------------
-- Table structure for ids
-- ----------------------------
DROP TABLE IF EXISTS `ids`;
CREATE TABLE `ids`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `number` int(11) UNSIGNED NOT NULL COMMENT '组号',
  `step` int(10) UNSIGNED NOT NULL COMMENT '每组数量',
  `modifiedtime` datetime NOT NULL,
  `createdtime` datetime NOT NULL,
  UNIQUE INDEX `idxname`(`name`) USING BTREE,
  INDEX `idx_id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ids
-- ----------------------------
INSERT INTO `ids` VALUES (2, 'equip', 100, 1000, '2021-08-05 17:19:20', '2021-08-05 17:19:26');
INSERT INTO `ids` VALUES (3, 'guild', 1, 100, '2022-03-14 20:47:54', '2022-03-14 20:48:02');
INSERT INTO `ids` VALUES (1, 'user', 100, 1000, '2021-07-14 10:49:48', '2021-07-14 10:49:48');

-- ----------------------------
-- Table structure for paomadeng
-- ----------------------------
DROP TABLE IF EXISTS `paomadeng`;
CREATE TABLE `paomadeng`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gmId` int(11) NOT NULL DEFAULT 0 COMMENT 'gm平台用id',
  `types` int(11) NOT NULL DEFAULT 0 COMMENT '0:循环播放 1:准点播放 2:间隔播放',
  `cycleTimes` int(11) NOT NULL DEFAULT 0 COMMENT '循环播放次数',
  `intervalTimes` int(11) NOT NULL DEFAULT 0 COMMENT '间隔播放时间',
  `startTime` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '开始时间',
  `endTime` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '结束时间',
  `channelIds` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `serverIds` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of paomadeng
-- ----------------------------

-- ----------------------------
-- Table structure for server_info
-- ----------------------------
DROP TABLE IF EXISTS `server_info`;
CREATE TABLE `server_info`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `serverIndex` int(11) NOT NULL DEFAULT 0,
  `serverId` int(11) NOT NULL DEFAULT 0,
  `appId` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `mergeServerId` int(11) NOT NULL DEFAULT 0 COMMENT '合服后serverId',
  `mergeTime` datetime NULL DEFAULT '1970-03-28 15:04:05' COMMENT '合服时间',
  `crossFsId` int(5) NULL DEFAULT 0 COMMENT '战斗跨服Id',
  `crossFirst` int(1) NULL DEFAULT 0 COMMENT '是否第一次跨服（0：是，1不是）',
  `gates` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT 'gate地址及端口号',
  `gsHostIn` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '127.0.0.1' COMMENT '服务器内网ip',
  `gsHostWww` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '服务器外网ip',
  `gsPort` int(5) NULL DEFAULT 0 COMMENT 'GameServer的端口，gate链接',
  `gsfsPort` int(5) NULL DEFAULT NULL COMMENT '本地战斗服端口，game链接',
  `gatefsPort` int(5) NULL DEFAULT NULL COMMENT '本地战斗服端口，gate链接',
  `httpPort` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '充值回调的端口',
  `isNew` int(1) NULL DEFAULT 1 COMMENT '是否是新区',
  `status` int(1) NULL DEFAULT 1 COMMENT '状态:, 1:良好，2:正常，3:爆满',
  `openTime` datetime NULL DEFAULT '2069-03-28 15:04:05' COMMENT '开服时间',
  `isClose` int(2) NULL DEFAULT 0 COMMENT '是否维护',
  `closeExplain` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '维护说明',
  `prefix` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 's' COMMENT '玩家名字前缀',
  `version` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1.0.1' COMMENT '版本号',
  `clientVersion` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0.8.5.0' COMMENT '客户端资源版本号',
  `isTrialVersion` int(1) NULL DEFAULT 0 COMMENT '是否是体验服',
  `redisAddr` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '服务器redis地址ip:port:pwd:db',
  `dblink` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '数据库连接地址',
  `dblinkLog` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '日志数据库连接地址',
  `ipFilter` int(1) NULL DEFAULT 0 COMMENT '是否开启白名单登录0否，1是',
  UNIQUE INDEX `idxServerIndex`(`serverIndex`) USING BTREE,
  INDEX `idx_id`(`id`) USING BTREE,
  INDEX `idxServerId`(`serverId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of server_info
-- ----------------------------
INSERT INTO `server_info` VALUES (1, '传奇1', 1, 1, 'ky', 1, '1970-03-28 15:04:05', 100000, 1, '192.168.91.15:7001', '192.168.91.15', '192.168.91.15', 10001, 10012, 10013, '10014', 1, 3, '2021-07-15 15:04:05', 0, '系统维护', 's', '1.0.1', '0.8.5.0', 0, '127.0.0.1:6379:123456:7', 'server=127.0.0.1;port=3306;database=cq_server_game_1;uid=root;pwd=123456;charset=utf8', '127.0.0.1:3306/cq_server_log_1', 0);
INSERT INTO `server_info` VALUES (99999, '先遣服', 99999, 99999, 'ky', 99999, '1970-03-28 15:04:05', 0, 0, '192.168.91.15:7999', '192.168.91.15', '192.168.91.15', 19991, 19992, 19993, '10024', 1, 3, '2021-07-15 15:04:05', 0, '系统维护', 's', '1.0.1', '0.8.5.0', 1, '127.0.0.1:6379:123456:1', 'server=127.0.0.1;port=3306;database=cq_server_game_99999;uid=root;pwd=123456;charset=utf8', '127.0.0.1:3306/cq_server_log_99999', 0);

-- ----------------------------
-- Table structure for server_port
-- ----------------------------
DROP TABLE IF EXISTS `server_port`;
CREATE TABLE `server_port`  (
  `id` int(5) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '配置名字',
  `host` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '服务器ip(内网)',
  `port` int(5) NOT NULL COMMENT '服务器端口号',
  `explain` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '说明',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of server_port
-- ----------------------------
INSERT INTO `server_port` VALUES (1, 'clientToLogin_1', '127.0.0.1', 7100, 'client->loginServer');
INSERT INTO `server_port` VALUES (3, 'gameToCrosscenter', '127.0.0.1', 7300, 'gameServer->crossCenterServer');
INSERT INTO `server_port` VALUES (4, 'gateToFightcenter', '127.0.0.1', 7400, 'gateServer->fightCenterServer');
INSERT INTO `server_port` VALUES (5, 'gameToFightcenter', '127.0.0.1', 7401, 'gameServer->fightCenterServer');
INSERT INTO `server_port` VALUES (6, 'crossCenterHttpPort', '127.0.0.1', 7500, '跨服中心http端口');
INSERT INTO `server_port`  VALUES (8, 'trialCrossCenterHttp', '192.168.91.15', 7500, '提审大区跨服中心http端口');

-- ----------------------------
-- Table structure for system_setting
-- ----------------------------
DROP TABLE IF EXISTS `system_setting`;
CREATE TABLE `system_setting`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `setting` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '配置名称',
  `value` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '配置值',
  `mark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '注释',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `setting_unique`(`setting`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_setting
-- ----------------------------
INSERT INTO `system_setting` VALUES (1, 'area_name', '2', '大区(恺英平台Id)，弃用');
INSERT INTO `system_setting` VALUES (2, 'cross_open_day', '5', '跨服开启开服天数（-1 不限）');
INSERT INTO `system_setting` VALUES (3, 'cross_open_active_player', '20', '跨服开启低于活跃人数限制（-1 不限）');
INSERT INTO `system_setting` VALUES (4, 'cross_open_server_day_recharge', '500', '跨服开启服务器日充值低于数量（-1不限）');
INSERT INTO `system_setting` VALUES (5, 'cross_max_server', '5', '跨服最大服务器数量');
INSERT INTO `system_setting` VALUES (6, 'cross_min_server', '2', '跨服最小服务器数量');
INSERT INTO `system_setting` VALUES (7, 'cross_activity_user_day', '3', '跨服开启取几天内活跃的人（-1 不限）');
INSERT INTO `system_setting` VALUES (8, 'max_activity_day', '7', '活跃人数存几天');
INSERT INTO `system_setting` VALUES (9, 'trial_server_version', '1.0.0', '提审服版本号');
INSERT INTO `system_setting` VALUES (10, 'trial_server', '0', '是否提审服（0：不是，1：是）');



-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `userId` bigint(20) NOT NULL COMMENT '玩家唯一ID',
  `openId` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '玩家账号',
  `channelId` int(11) NOT NULL DEFAULT 0 COMMENT '渠道Id',
  `serverIndex` int(11) NULL DEFAULT 0 COMMENT '服务器索引',
  `serverId` int(11) NULL DEFAULT 0 COMMENT '服务器ID',
  `nickname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '玩家昵称',
  `avatar` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '玩家头像链接',
  `gold` int(11) NULL DEFAULT NULL COMMENT '金币',
  `ingot` int(11) NOT NULL DEFAULT 0 COMMENT '元宝',
  `vip` int(11) NOT NULL DEFAULT 0 COMMENT 'VIP等级',
  `taskId` int(8) NOT NULL DEFAULT 0 COMMENT '任务id',
  `recharge` int(11) NOT NULL DEFAULT 0 COMMENT '充值金额',
  `tokenRecharge` int(11) NOT NULL DEFAULT 0 COMMENT '代币充值',
  `combat` int(20) NOT NULL DEFAULT 0 COMMENT '战力',
  `updateTime` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后更新时间',
  `offLineTime` datetime NULL DEFAULT NULL COMMENT '下线时间',
  `createTime` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创角时间',
  `lastRechargeTime` datetime NULL DEFAULT NULL COMMENT '最后充值时间',
  `heros` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '武将信息',
  PRIMARY KEY (`userId`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '玩家信息表,创建者:liuxsh' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------


-- ----------------------------
-- Table structure for white_list
-- ----------------------------
DROP TABLE IF EXISTS `white_list`;
CREATE TABLE `white_list`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `gmId` int(11) NOT NULL DEFAULT 0 COMMENT 'gm平台Id',
  `valtype` int(1) NOT NULL COMMENT '白名单类型（1：ip,2:账号）',
  `value` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '白名单',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '白名单' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of white_list
-- ----------------------------

-- ----------------------------
-- Table structure for gift_code
-- ----------------------------
DROP TABLE IF EXISTS `gift_code`;
CREATE TABLE `gift_code` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(50) NOT NULL COMMENT '礼包码',
  `batchId` int(11) NOT NULL DEFAULT '0' COMMENT '批次',
  `batchName` varchar(100) DEFAULT NULL COMMENT '批次名',
  `batchNum` int(11) NOT NULL DEFAULT '1' COMMENT '本批次可使用个数',
  `reward` longtext NOT NULL COMMENT '奖励信息',
  `serverId` int(11) DEFAULT '0' COMMENT '可使用服务器id',
  `channel` int(11) DEFAULT '0' COMMENT '可领取渠道',
  `startTime` datetime NOT NULL COMMENT '开始时间',
  `endTime` datetime NOT NULL COMMENT '结束时间',
  `codeType` int(11) NOT NULL DEFAULT '1' COMMENT '礼包码类型',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `gift_code_receive`;
CREATE TABLE `gift_code_receive` (
  `codeId` int(11) NOT NULL COMMENT '礼包码id',
  `userId` int(11) NOT NULL COMMENT '用户id',
  `batchId` int(11) NOT NULL COMMENT '批次',
  `receiveTime` datetime NOT NULL COMMENT '领取时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
