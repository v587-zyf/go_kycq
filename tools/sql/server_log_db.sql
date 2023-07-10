/*
 Navicat MySQL Data Transfer

 Source Server         : 测试服
 Source Server Type    : MySQL
 Source Server Version : 50734
 Source Host           : 192.168.20.77:3306
 Source Schema         : cq_server_log_1

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 20/05/2022 11:33:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for log_item_flow
-- ----------------------------
DROP TABLE IF EXISTS `log_item_flow`;
CREATE TABLE `log_item_flow` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL COMMENT '玩家唯一编号',
  `serverId` int(11) NOT NULL DEFAULT '0' COMMENT '服务器Id',
  `dtEventTime` datetime NOT NULL COMMENT '事件时间',
  `openid` varchar(64) NOT NULL DEFAULT '' COMMENT '玩家账号',
  `iGoodsId` int(11) DEFAULT NULL COMMENT '道具ID',
  `Count` int(11) DEFAULT NULL COMMENT '数量',
  `AfterCount` int(11) DEFAULT NULL COMMENT '动作后的物品存量',
  `Reason` varchar(100) DEFAULT NULL COMMENT '道具流动一级原因',
  `Reason2` varchar(100) DEFAULT NULL COMMENT '道具流动二级原因',
  `AddOrReduce` int(11) DEFAULT NULL COMMENT '增加 0/减少 1',
  `UserLv` int(11) DEFAULT NULL COMMENT '玩家等级',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23824 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for log_player_login
-- ----------------------------
DROP TABLE IF EXISTS `log_player_login`;
CREATE TABLE `log_player_login`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dtEventTime` datetime NOT NULL COMMENT '事件时间',
  `UserId` int(11) NOT NULL DEFAULT 0 COMMENT '玩家唯一编号',
  `serverId` int(11) NOT NULL DEFAULT 0 COMMENT '服务器Id',
  `openid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '玩家账号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2420 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_player_logout
-- ----------------------------
DROP TABLE IF EXISTS `log_player_logout`;
CREATE TABLE `log_player_logout`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dtEventTime` datetime NOT NULL COMMENT '事件时间',
  `UserId` int(11) NOT NULL DEFAULT 0 COMMENT '玩家唯一编号',
  `serverId` int(11) NOT NULL DEFAULT 0 COMMENT '服务器Id',
  `openid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '玩家账号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2835 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_player_register
-- ----------------------------
DROP TABLE IF EXISTS `log_player_register`;
CREATE TABLE `log_player_register`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dtEventTime` datetime NOT NULL COMMENT '事件时间',
  `UserId` int(11) NOT NULL DEFAULT 0 COMMENT '玩家唯一编号',
  `serverId` int(11) NOT NULL DEFAULT 0 COMMENT '服务器Id',
  `openid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '玩家账号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 420 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
