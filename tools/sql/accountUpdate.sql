CREATE TABLE `func_close` (
  `funcId` int(5) NOT NULL COMMENT '功能Id',
  `serverIds` varchar(255) NOT NULL DEFAULT '[]' COMMENT '关闭功能模块的服务器Id(空位全服全部)',
  PRIMARY KEY (`funcId`) USING BTREE,
  UNIQUE KEY `唯一` (`funcId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


ALTER TABLE challenge MODIFY COLUMN `fightUserInfo` text DEFAULT NULL COMMENT "战斗玩家数据";


ALTER TABLE `user` 
DROP COLUMN `level`,
DROP COLUMN `sex`,
ADD COLUMN `channelId` int(11) NULL COMMENT '渠道Id' AFTER `openId`,
ADD COLUMN `gold` int(11) NULL COMMENT '金币' AFTER `avatar`,
ADD COLUMN `ingot` int(11) NULL COMMENT '元宝' AFTER `gold`,
ADD COLUMN `taskId` int(8) NULL COMMENT '任务id' AFTER `vip`,
ADD COLUMN `recharge` int(11) NULL COMMENT '充值金额' AFTER `taskId`,
ADD COLUMN `offLineTime` datetime NULL COMMENT '下线时间' AFTER `createTime`,
ADD COLUMN `lastRechargeTime` datetime NULL COMMENT '最后充值时间' AFTER `offLineTime`,
ADD COLUMN `heros` varchar(500) NULL COMMENT '武将信息' AFTER `lastRechargeTime`;



DROP TABLE IF EXISTS `announcement`;
CREATE TABLE `announcement`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` int(11) NOT NULL DEFAULT 1,
  `title` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '公告名称',
  `startTime` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '开始时间',
  `endTime` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '结束时间',
  `channelIds` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `serverIds` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `announcement`longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;


DROP TABLE IF EXISTS `white_list`;
CREATE TABLE `white_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `gmId` int(11) NOT NULL DEFAULT '0' COMMENT 'gm平台Id',
  `valtype` int(1) NOT NULL COMMENT '白名单类型（1：ip,2:账号）',
  `value` varchar(200) NOT NULL COMMENT '白名单',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='白名单';



CREATE TABLE `announcement` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gmId` int(11) NOT NULL DEFAULT '0',
  `type` int(11) NOT NULL DEFAULT '1',
  `title` varchar(300) DEFAULT NULL COMMENT '公告名称',
  `startTime` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '开始时间',
  `endTime` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '结束时间',
  `channelIds` longtext,
  `serverIds` longtext,
  `Announcement` longtext,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

CREATE TABLE `paomadeng` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gmId` int(11) NOT NULL DEFAULT '0' COMMENT 'gm平台用id',
  `types` int(11) NOT NULL DEFAULT '0' COMMENT '0:循环播放 1:准点播放 2:间隔播放',
  `cycleTimes` int(11) NOT NULL DEFAULT '0' COMMENT '循环播放次数',
  `intervalTimes` int(11) NOT NULL DEFAULT '0' COMMENT '间隔播放时间',
  `startTime` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '开始时间',
  `endTime` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '结束时间',
  `channelIds` longtext,
  `serverIds` longtext,
  `content` longtext,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

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


ALTER TABLE `user` ADD COLUMN `exp` int(11) NULL COMMENT '经验' AFTER `recharge`;
ALTER TABLE `user` ADD COLUMN `loginTime` datetime NULL COMMENT '角色最后登录时间' AFTER `createTime`;


INSERT INTO `system_setting` VALUES (7, 'cross_activity_user_day', '3', '跨服开启取几天内活跃的人（-1 不限）');
INSERT INTO `system_setting` VALUES (8, 'max_activity_day', '7', '活跃人数存几天');

--
--支持多login修改
--
UPDATE `server_port` SET `name` = 'clientToLogin_1' WHERE `id` = 1;

--
--20220901 增加玩家代币充值记录
--
ALTER TABLE `user` ADD COLUMN `tokenRecharge` int(11) NOT NULL DEFAULT 0 COMMENT '代币充值' AFTER `recharge`;


--跨服擂台赛数据升级
ALTER TABLE `challenge` ADD COLUMN `guildName` varchar(100) DEFAULT NULL  COMMENT '公会名字';

--
--20221108 服务器版本
--
INSERT INTO `system_setting` VALUES (9, 'trial_server_version', '1.0.0', '提审服版本号');

--
--20221115 服务器提审标识
--
INSERT INTO `system_setting` VALUES (10, 'trial_server', '0', '是否提审服（0：不是，1：是）');
INSERT INTO `server_port`  VALUES (8, 'trialCrossCenterHttp', '192.168.91.15', 7500, '提审大区跨服中心http端口');

