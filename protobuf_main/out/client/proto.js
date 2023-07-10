// auto generated, do not edit

var pb = dcodeIO.ProtoBuf.newBuilder({"populateAccessors": false})['import']({
    "package": "pb",
    "syntax": "proto3",
    "messages": [
        {
            "name": "AchievementLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "AchievementLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "AchievementInfo",
                    "name": "achievementInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "allPoint",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "Medal",
                    "id": 3,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "AchievementGetAwardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "id",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "AchievementGetAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "AchievementInfo",
                    "name": "achievementInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "allPoint",
                    "id": 2
                }
            ]
        },
        {
            "name": "ActiveMedalReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "ActiveMedalAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "allPoint",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "Medal",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "AchievementInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "conditionType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "canGetId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "process",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isGetAllAward",
                    "id": 4
                }
            ]
        },
        {
            "name": "AchievementTaskInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "taskId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "process",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "conditionType",
                    "id": 3
                }
            ]
        },
        {
            "name": "ErrorAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "code",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "message",
                    "id": 2
                }
            ]
        },
        {
            "name": "UserLoginInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "nickName",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "avatar",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "vipLevel",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "vipScore",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "gold",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ingot",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "createTime",
                    "id": 10
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 11
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageWave",
                    "id": 12
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "combat",
                    "id": 14
                },
                {
                    "rule": "repeated",
                    "type": "HeroInfo",
                    "name": "heros",
                    "id": 15
                },
                {
                    "rule": "optional",
                    "type": "Rein",
                    "name": "rein",
                    "id": 16
                },
                {
                    "rule": "repeated",
                    "type": "ReinCost",
                    "name": "reinCost",
                    "id": 17
                },
                {
                    "rule": "repeated",
                    "type": "Fabao",
                    "name": "fabao",
                    "id": 18
                },
                {
                    "rule": "optional",
                    "type": "FieldBossInfo",
                    "name": "fieldBossInfo",
                    "id": 19
                },
                {
                    "rule": "optional",
                    "type": "WorldBossInfoNtf",
                    "name": "worldBossInfo",
                    "id": 20
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ArenaFightNum",
                    "id": 21
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fightModel",
                    "id": 22
                },
                {
                    "rule": "optional",
                    "type": "TaskInfoNtf",
                    "name": "task",
                    "id": 23
                },
                {
                    "rule": "map",
                    "type": "ShopInfo",
                    "keytype": "int32",
                    "name": "shopInfo",
                    "id": 24
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "chuanqiBi",
                    "id": 25
                },
                {
                    "rule": "map",
                    "type": "MaterialStage",
                    "keytype": "int32",
                    "name": "materialStage",
                    "id": 26
                },
                {
                    "rule": "map",
                    "type": "PanaceaInfo",
                    "keytype": "int32",
                    "name": "panaceas",
                    "id": 27
                },
                {
                    "rule": "optional",
                    "type": "SignInfo",
                    "name": "signInfo",
                    "id": 28
                },
                {
                    "rule": "optional",
                    "type": "DayStateInfo",
                    "name": "dayStateInfo",
                    "id": 29
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "official",
                    "id": 30
                },
                {
                    "rule": "repeated",
                    "type": "Holy",
                    "name": "holy",
                    "id": 31
                },
                {
                    "rule": "repeated",
                    "type": "Atlas",
                    "name": "atlases",
                    "id": 32
                },
                {
                    "rule": "repeated",
                    "type": "AtlasGather",
                    "name": "atlasGathers",
                    "id": 33
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "miningWorkTime",
                    "id": 34
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "miner",
                    "id": 35
                },
                {
                    "rule": "optional",
                    "type": "ExpStage",
                    "name": "expStage",
                    "id": 36
                },
                {
                    "rule": "map",
                    "type": "PetInfo",
                    "keytype": "int32",
                    "name": "pets",
                    "id": 37
                },
                {
                    "rule": "repeated",
                    "type": "Juexue",
                    "name": "juexues",
                    "id": 38
                },
                {
                    "rule": "optional",
                    "type": "UserWear",
                    "name": "userWear",
                    "id": 39
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isHaveGetDailyCompetitveReward",
                    "id": 40
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "honour",
                    "id": 41
                },
                {
                    "rule": "optional",
                    "type": "DarkPalaceInfo",
                    "name": "darkPalaceInfo",
                    "id": 42
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "personBoss",
                    "id": 43
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "vipBoss",
                    "id": 44
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "vipGift",
                    "id": 45,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "Fit",
                    "name": "fit",
                    "id": 46
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rechargeAll",
                    "id": 47
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "accumulativeAllGetIds",
                    "id": 48,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "map",
                    "type": "MonthCardUnit",
                    "keytype": "int32",
                    "name": "monthCard",
                    "id": 49
                },
                {
                    "rule": "optional",
                    "type": "FirstRecharge",
                    "name": "firstRecharge",
                    "id": 50
                },
                {
                    "rule": "optional",
                    "type": "SpendRebates",
                    "name": "spendRebates",
                    "id": 51
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "dailyPack",
                    "id": 52
                },
                {
                    "rule": "optional",
                    "type": "GrowFund",
                    "name": "growFund",
                    "id": 53
                },
                {
                    "rule": "optional",
                    "type": "WarOrder",
                    "name": "warOrder",
                    "id": 54
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "redPacketGetNum",
                    "id": 55
                },
                {
                    "rule": "optional",
                    "type": "Elf",
                    "name": "elf",
                    "id": 56
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cutTreasureLv",
                    "id": 57
                },
                {
                    "rule": "optional",
                    "type": "FitHolyEquip",
                    "name": "fitHolyEquip",
                    "id": 58
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "recharge",
                    "id": 59,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "shaBakeIsEnd",
                    "id": 60
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "crossShabakeIsEnd",
                    "id": 61
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isFriendApply",
                    "id": 62
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bindingIngot",
                    "id": 63
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "hookupTime",
                    "id": 64
                },
                {
                    "rule": "repeated",
                    "type": "itemUnit",
                    "name": "hookupBag",
                    "id": 65
                },
                {
                    "rule": "optional",
                    "type": "ContRecharge",
                    "name": "contRecharge",
                    "id": 66
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "openGift",
                    "id": 67
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "crossChallengeIsApply",
                    "id": 68
                },
                {
                    "rule": "repeated",
                    "type": "AnnouncementInfo",
                    "name": "AnnouncementInfos",
                    "id": 69
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "vipCustomer",
                    "id": 70
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "chatBanTime",
                    "id": 71
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "haveUseRecharge",
                    "id": 72
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "goldIngot",
                    "id": 73
                },
                {
                    "rule": "optional",
                    "type": "AncientBossInfo",
                    "name": "ancientBossInfo",
                    "id": 74
                },
                {
                    "rule": "repeated",
                    "type": "Title",
                    "name": "titleList",
                    "id": 75
                },
                {
                    "rule": "repeated",
                    "type": "MiJiInfo",
                    "name": "miJiInfos",
                    "id": 76
                },
                {
                    "rule": "map",
                    "type": "AncientTreasureInfo",
                    "keytype": "int32",
                    "name": "ancientTreasureInfo",
                    "id": 77
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "petAppendage",
                    "id": 78
                },
                {
                    "rule": "optional",
                    "type": "HellBossInfo",
                    "name": "hellBossInfo",
                    "id": 79
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "redPacketNum",
                    "id": 80
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "daBaoEquip",
                    "id": 81
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "daBaoMysteryEnergy",
                    "id": 82
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "appletsEnergy",
                    "id": 83
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "appletsResumeTime",
                    "id": 84
                },
                {
                    "rule": "map",
                    "type": "appletsInfo",
                    "keytype": "int32",
                    "name": "appletsInfos",
                    "id": 85
                },
                {
                    "rule": "optional",
                    "type": "Label",
                    "name": "label",
                    "id": 86
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "subscribe",
                    "id": 87,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "privilege",
                    "id": 88,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "useRedPacketNum",
                    "id": 89
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dailyRecharge",
                    "id": 90
                }
            ]
        },
        {
            "name": "appletsInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientTreasureInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "zhuLinLv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "starLv",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "jueXinLv",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "types",
                    "id": 4
                }
            ]
        },
        {
            "name": "MiJiInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 2
                }
            ]
        },
        {
            "name": "AnnouncementInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "title",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "announcement",
                    "id": 3
                }
            ]
        },
        {
            "name": "HeroInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "index",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sex",
                    "id": 3
                },
                {
                    "rule": "map",
                    "type": "EquipUnit",
                    "keytype": "int32",
                    "name": "equips",
                    "id": 4
                },
                {
                    "rule": "repeated",
                    "type": "EquipGrid",
                    "name": "equipGrids",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "HeroProp",
                    "name": "heroProp",
                    "id": 6
                },
                {
                    "rule": "repeated",
                    "type": "Wing",
                    "name": "wing",
                    "id": 7
                },
                {
                    "rule": "map",
                    "type": "SpecialEquipUnit",
                    "keytype": "int32",
                    "name": "zodiacs",
                    "id": 8
                },
                {
                    "rule": "map",
                    "type": "SpecialEquipUnit",
                    "keytype": "int32",
                    "name": "kingarms",
                    "id": 9
                },
                {
                    "rule": "repeated",
                    "type": "DictateInfo",
                    "name": "dictates",
                    "id": 10
                },
                {
                    "rule": "repeated",
                    "type": "WingSpecialNtf",
                    "name": "wingSpecial",
                    "id": 11
                },
                {
                    "rule": "map",
                    "type": "JewelInfo",
                    "keytype": "int32",
                    "name": "jewels",
                    "id": 12
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 13
                },
                {
                    "rule": "optional",
                    "type": "InsideInfo",
                    "name": "insideInfo",
                    "id": 14
                },
                {
                    "rule": "map",
                    "type": "Fashion",
                    "keytype": "int32",
                    "name": "Fashions",
                    "id": 15
                },
                {
                    "rule": "optional",
                    "type": "Wears",
                    "name": "wears",
                    "id": 16
                },
                {
                    "rule": "map",
                    "type": "Ring",
                    "keytype": "int32",
                    "name": "rings",
                    "id": 17
                },
                {
                    "rule": "repeated",
                    "type": "SkillUnit",
                    "name": "skills",
                    "id": 18
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "skillBag",
                    "id": 19
                },
                {
                    "rule": "repeated",
                    "type": "SkillUnit",
                    "name": "uniqueSkills",
                    "id": 20
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "uniqueSkillBag",
                    "id": 21
                },
                {
                    "rule": "map",
                    "type": "GodEquip",
                    "keytype": "int32",
                    "name": "godEquips",
                    "id": 22
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "area",
                    "id": 23
                },
                {
                    "rule": "map",
                    "type": "EquipClearArr",
                    "keytype": "int32",
                    "name": "equipClears",
                    "id": 24
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "expLvl",
                    "id": 25
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "dragonEquip",
                    "id": 26
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "MagicCircle",
                    "id": 27
                },
                {
                    "rule": "optional",
                    "type": "TalentInfo",
                    "name": "talents",
                    "id": 28
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "chuanShiEquip",
                    "id": 29
                },
                {
                    "rule": "optional",
                    "type": "AncientSkill",
                    "name": "ancientSkill",
                    "id": 30
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "chuanShiStrengthen",
                    "id": 31
                }
            ]
        },
        {
            "name": "UserWear",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "petid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitFashionId",
                    "id": 2
                }
            ]
        },
        {
            "name": "Wears",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "FashionWeaponId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "FashionClothId",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "atlasWear",
                    "id": 3,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "wingId",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "magicCircleLvId",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "titleId",
                    "id": 6
                }
            ]
        },
        {
            "name": "BriefUserInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sex",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lvl",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "vip",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "combat",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "avatar",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "maxLv",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverId",
                    "id": 10
                },
                {
                    "rule": "map",
                    "type": "Display",
                    "keytype": "int32",
                    "name": "display",
                    "id": 11
                }
            ]
        },
        {
            "name": "BriefUserInfoWithDisplay",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "userInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "DisplayNtf",
                    "name": "display",
                    "id": 14
                }
            ]
        },
        {
            "name": "TopDataChangeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "TopDataChange",
                    "name": "changeInfos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 2
                }
            ]
        },
        {
            "name": "TopDataChange",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "change",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "nowNum",
                    "id": 3
                }
            ]
        },
        {
            "name": "BagDataChangeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "ItemChange",
                    "name": "changeInfos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 2
                }
            ]
        },
        {
            "name": "ItemChange",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "position",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "change",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "nowNum",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "GetSource",
                    "name": "getSource",
                    "id": 5
                }
            ]
        },
        {
            "name": "BagEquipDataChangeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "EquipChange",
                    "name": "changeInfos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 2
                }
            ]
        },
        {
            "name": "EquipChange",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "position",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "change",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "nowNum",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "EquipUnit",
                    "name": "equip",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "GetSource",
                    "name": "getSource",
                    "id": 6
                }
            ]
        },
        {
            "name": "Item",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "count",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "position",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "EquipUnit",
                    "name": "equip",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "GetSource",
                    "name": "getSource",
                    "id": 5
                }
            ]
        },
        {
            "name": "GoodsChangeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "itemUnit",
                    "name": "items",
                    "id": 1
                }
            ]
        },
        {
            "name": "itemUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "count",
                    "id": 2
                }
            ]
        },
        {
            "name": "EquipUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "EquipRandProp",
                    "name": "randProps",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "lock",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipIndex",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lucky",
                    "id": 5
                }
            ]
        },
        {
            "name": "EquipRandProp",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "propId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "color",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "value",
                    "id": 3
                }
            ]
        },
        {
            "name": "EquipClearArr",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "EquipClearInfo",
                    "name": "equipClearInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "EquipClearInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "grade",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "color",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "propId",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "value",
                    "id": 4
                }
            ]
        },
        {
            "name": "HeroProp",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int64",
                    "keytype": "int32",
                    "name": "props",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "int64",
                    "keytype": "int32",
                    "name": "modulesCombat",
                    "id": 2
                }
            ]
        },
        {
            "name": "TaskInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "taskId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "process",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "markProcess",
                    "id": 3
                }
            ]
        },
        {
            "name": "DisplayNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "Display",
                    "keytype": "int32",
                    "name": "display",
                    "id": 1
                }
            ]
        },
        {
            "name": "EventNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ts",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "string",
                    "name": "args",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sourceId",
                    "id": 4
                }
            ]
        },
        {
            "name": "DailyConditionNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "dailyConditions",
                    "id": 1
                }
            ]
        },
        {
            "name": "PropInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "key",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "value",
                    "id": 2
                }
            ]
        },
        {
            "name": "EquipGrid",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "strength",
                    "id": 2
                }
            ]
        },
        {
            "name": "Fabao",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 3
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "skills",
                    "id": 4,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "GodEquip",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "blood",
                    "id": 3
                }
            ]
        },
        {
            "name": "Juexue",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 2
                }
            ]
        },
        {
            "name": "Fashion",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 2
                }
            ]
        },
        {
            "name": "Wing",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 2
                }
            ]
        },
        {
            "name": "WingSpecialNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "specialType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 2
                }
            ]
        },
        {
            "name": "Rein",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "exp",
                    "id": 2
                }
            ]
        },
        {
            "name": "ReinCost",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "num",
                    "id": 2
                }
            ]
        },
        {
            "name": "Atlas",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "star",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isActive",
                    "id": 3
                }
            ]
        },
        {
            "name": "AtlasGather",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "star",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isActive",
                    "id": 3
                }
            ]
        },
        {
            "name": "Preference",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "key",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "value",
                    "id": 2
                }
            ]
        },
        {
            "name": "WorldBossInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "prepareTime",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "openTime",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "closeTime",
                    "id": 4
                }
            ]
        },
        {
            "name": "VipBoss",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "ExpStage",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int64",
                    "keytype": "int32",
                    "name": "expStages",
                    "id": 3
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "appraise",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "layer",
                    "id": 5
                }
            ]
        },
        {
            "name": "MaterialStage",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "nowLayer",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lastLayer",
                    "id": 4
                }
            ]
        },
        {
            "name": "Display",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "clothItemId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "clothType",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "weaponItemId",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "weaponType",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "wingId",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "magicCircleLvId",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "titleId",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "labelId",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "labelJob",
                    "id": 9
                }
            ]
        },
        {
            "name": "SpecialEquipUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GetSource",
                    "name": "getSource",
                    "id": 2
                }
            ]
        },
        {
            "name": "GetSource",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "map",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "monster",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "skillUser",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "skillDate",
                    "id": 4
                }
            ]
        },
        {
            "name": "RankInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rank",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "score",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "userInfo",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "Display",
                    "name": "display",
                    "id": 4
                }
            ]
        },
        {
            "name": "SkillUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "startTime",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 4
                }
            ]
        },
        {
            "name": "ShopInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "shopItem",
                    "id": 1
                }
            ]
        },
        {
            "name": "DictateInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 2
                }
            ]
        },
        {
            "name": "PanaceaInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "numbers",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "number",
                    "id": 2
                }
            ]
        },
        {
            "name": "JewelInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "left",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "right",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "down",
                    "id": 3
                }
            ]
        },
        {
            "name": "DayStateInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rankWorship",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "monthCardReceive",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "SignInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "signCount",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "signDay",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "cumulativeDay",
                    "id": 3
                }
            ]
        },
        {
            "name": "InsideInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "acupoint",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "InsideSkill",
                    "keytype": "int32",
                    "name": "insideSkill",
                    "id": 2
                }
            ]
        },
        {
            "name": "InsideSkill",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 2
                }
            ]
        },
        {
            "name": "Holy",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 3
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "skills",
                    "id": 4
                }
            ]
        },
        {
            "name": "Ring",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "strengthen",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pid",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "talent",
                    "id": 4
                },
                {
                    "rule": "map",
                    "type": "RingPhantom",
                    "keytype": "int32",
                    "name": "phantom",
                    "id": 5
                }
            ]
        },
        {
            "name": "RingPhantom",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "talent",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "phantom",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "skill",
                    "id": 3
                }
            ]
        },
        {
            "name": "PetInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "grade",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "break",
                    "id": 4
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "skill",
                    "id": 5,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "ResetNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "newDayTime",
                    "id": 2
                }
            ]
        },
        {
            "name": "FieldFightRivalUserInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rivalUserId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rivalDifficult",
                    "id": 2
                }
            ]
        },
        {
            "name": "DarkPalaceInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "helpNum",
                    "id": 3
                }
            ]
        },
        {
            "name": "HellBossInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "helpNum",
                    "id": 3
                }
            ]
        },
        {
            "name": "FieldBossInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "firstReceive",
                    "id": 3
                }
            ]
        },
        {
            "name": "AncientBossInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "TalentInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "getPoints",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "surplusPoints",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "TalentUnit",
                    "keytype": "int32",
                    "name": "talents",
                    "id": 3
                }
            ]
        },
        {
            "name": "TalentUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "usePoints",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "talents",
                    "id": 2
                }
            ]
        },
        {
            "name": "Fit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "cdStart",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "cdEnd",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "fashion",
                    "id": 3
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "skillBag",
                    "id": 4
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "lv",
                    "id": 5
                },
                {
                    "rule": "map",
                    "type": "FitSkill",
                    "keytype": "int32",
                    "name": "skills",
                    "id": 6
                }
            ]
        },
        {
            "name": "FitSkill",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "star",
                    "id": 2
                }
            ]
        },
        {
            "name": "MonthCardUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "startTime",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isExpire",
                    "id": 3
                }
            ]
        },
        {
            "name": "FirstRecharge",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isRecharge",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "days",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "openDay",
                    "id": 3
                }
            ]
        },
        {
            "name": "SpendRebates",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "countIngot",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ingot",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "reward",
                    "id": 3,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cycle",
                    "id": 4
                }
            ]
        },
        {
            "name": "GrowFund",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isBuy",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "ids",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "WarOrderTaskUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "one",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "two",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "three",
                    "id": 3
                }
            ]
        },
        {
            "name": "WarOrderTaskInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "WarOrderTaskUnit",
                    "name": "val",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "finish",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "reward",
                    "id": 3
                }
            ]
        },
        {
            "name": "WarOrderTask",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "WarOrderTaskInfo",
                    "keytype": "int32",
                    "name": "task",
                    "id": 1
                }
            ]
        },
        {
            "name": "WarOrderReward",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "elite",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "luxury",
                    "id": 2
                }
            ]
        },
        {
            "name": "WarOrder",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "season",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "startTime",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isLuxury",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "WarOrderTask",
                    "name": "task",
                    "id": 7
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "exchange",
                    "id": 8
                },
                {
                    "rule": "map",
                    "type": "WarOrderTask",
                    "keytype": "int32",
                    "name": "weekTask",
                    "id": 9
                },
                {
                    "rule": "map",
                    "type": "WarOrderReward",
                    "keytype": "int32",
                    "name": "reward",
                    "id": 10
                }
            ]
        },
        {
            "name": "Elf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "skills",
                    "id": 3
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "skillBag",
                    "id": 4
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "receiveLimit",
                    "id": 5
                }
            ]
        },
        {
            "name": "FriendUserInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "UserId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "combat",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "FriendHeroInfo",
                    "keytype": "int32",
                    "name": "friendHeroInfo",
                    "id": 3
                }
            ]
        },
        {
            "name": "FriendHeroInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "EquipUnit",
                    "keytype": "int32",
                    "name": "equips",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "int64",
                    "keytype": "int32",
                    "name": "props",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sex",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "Display",
                    "name": "display",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "combat",
                    "id": 8
                }
            ]
        },
        {
            "name": "FitHolyEquipUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "equip",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitHolyEquip",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "suitId",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "FitHolyEquipUnit",
                    "keytype": "int32",
                    "name": "equips",
                    "id": 2
                }
            ]
        },
        {
            "name": "ContRecharge",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cycle",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "recharge",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "receive",
                    "id": 3,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "todayPay",
                    "id": 4
                }
            ]
        },
        {
            "name": "PaoMaDengInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cycleTimes",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "content",
                    "id": 3
                }
            ]
        },
        {
            "name": "PaoMaDengNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "PaoMaDengInfo",
                    "name": "PaoMaDengInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientSkill",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "grade",
                    "id": 3
                }
            ]
        },
        {
            "name": "Title",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "titleId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "startTime",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isLook",
                    "id": 4
                }
            ]
        },
        {
            "name": "BriefServerInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "serverName",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "crossFsId",
                    "id": 3
                }
            ]
        },
        {
            "name": "Label",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "labelId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "transfer",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "dayReward",
                    "id": 4
                }
            ]
        },
        {
            "name": "AncientBossLoadReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "area",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientBossLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "AncientBossNtf",
                    "name": "ancientBoss",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientBossBuyNumReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "use",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "AncientBossBuyNumAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterAncientBossFightReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterAncientBossFightAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "AncientBossFightResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "winner",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 5
                }
            ]
        },
        {
            "name": "AncientBossOwnerReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientBossOwnerAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "AncientBossOwnerInfo",
                    "name": "list",
                    "id": 2
                }
            ]
        },
        {
            "name": "AncientBossOwnerInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "time",
                    "id": 2
                }
            ]
        },
        {
            "name": "AncientBossNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "float",
                    "name": "blood",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "reliveTime",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "area",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userCount",
                    "id": 5
                }
            ]
        },
        {
            "name": "AncientSkillActiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientSkillActiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "AncientSkill",
                    "name": "ancientSkill",
                    "id": 2
                }
            ]
        },
        {
            "name": "AncientSkillUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientSkillUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 2
                }
            ]
        },
        {
            "name": "AncientSkillUpGradeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientSkillUpGradeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "grade",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 2
                }
            ]
        },
        {
            "name": "AncientTreasuresLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "AncientTreasuresLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "AncientTreasuresInfo",
                    "keytype": "int32",
                    "name": "AncientTreasuresInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientTreasuresInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "zhuLinLv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "starLv",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isAwakening",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "types",
                    "id": 4
                }
            ]
        },
        {
            "name": "AncientTreasuresActivateReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureId",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientTreasuresActivateAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureId",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientTreasuresZhuLinReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureId",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientTreasuresZhuLinAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "zhuLinLv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureId",
                    "id": 2
                }
            ]
        },
        {
            "name": "AncientTreasuresUpStarReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureId",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientTreasuresUpStarAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "starLv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureId",
                    "id": 2
                }
            ]
        },
        {
            "name": "AncientTreasuresJueXingReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureId",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "chooseItemInfos",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "index",
                    "id": 3
                }
            ]
        },
        {
            "name": "AncientTreasuresJueXingAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureId",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientTreasuresResertReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureId",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientTreasuresResertAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureId",
                    "id": 1
                }
            ]
        },
        {
            "name": "AncientTreasuresCondotionInfosReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "AncientTreasuresCondotionInfosAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "ancientTreasureConditionInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterAppletsReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "appletsType",
                    "id": 1
                }
            ]
        },
        {
            "name": "AppletsEnergyNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "energy",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "ResumeTime",
                    "id": 2
                }
            ]
        },
        {
            "name": "AppletsReceiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "receiveId",
                    "id": 1
                }
            ]
        },
        {
            "name": "AppletsReceiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "receiveId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "CronGetAwardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "index",
                    "id": 2
                }
            ]
        },
        {
            "name": "CronGetAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "EndResultReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "appletsType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 2
                }
            ]
        },
        {
            "name": "EndResultAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "appletsType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "energy",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "AreaUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 2
                }
            ]
        },
        {
            "name": "AreaUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 3
                }
            ]
        },
        {
            "name": "ArenaOpenReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ArenaOpenAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "RankInfo",
                    "name": "three",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "ArenaRank",
                    "name": "arenaRank",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyDareNum",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyDareNums",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ranking",
                    "id": 6
                }
            ]
        },
        {
            "name": "EnterArenaFightReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "challengeUid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "challengeRanking",
                    "id": 2
                }
            ]
        },
        {
            "name": "ArenaFightNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "myRank",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "seasonScore",
                    "id": 4
                }
            ]
        },
        {
            "name": "BuyArenaFightNumReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "BuyArenaFightNumAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyDareNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyDareNums",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "RefArenaRankReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "RefArenaRankAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "ArenaRank",
                    "name": "arenaRank",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "RankInfo",
                    "name": "three",
                    "id": 2
                }
            ]
        },
        {
            "name": "ArenaRank",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ranking",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "userinfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "AtlasActiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "AtlasActiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Atlas",
                    "name": "atlas",
                    "id": 1
                }
            ]
        },
        {
            "name": "AtlasUpStarReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "AtlasUpStarAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Atlas",
                    "name": "atlas",
                    "id": 1
                }
            ]
        },
        {
            "name": "AtlasGatherActiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "AtlasGatherActiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "AtlasGather",
                    "name": "atlasGather",
                    "id": 1
                }
            ]
        },
        {
            "name": "AtlasGatherUpStarReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "AtlasGatherUpStarAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "AtlasGather",
                    "name": "atlasGather",
                    "id": 1
                }
            ]
        },
        {
            "name": "AtlasWearChangeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 2
                }
            ]
        },
        {
            "name": "AtlasWearChangeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "removeId",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "atlasWear",
                    "id": 3,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "AuctionInfoReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionType",
                    "id": 1
                }
            ]
        },
        {
            "name": "AuctionInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionType",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "AuctionItemInfo",
                    "name": "auctionInfos",
                    "id": 2
                }
            ]
        },
        {
            "name": "AuctionItemInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "auctionId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "auctionTime",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionDuration",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "nowBidPrice",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "nowBidUserId",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionType",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "nowBidderNickname",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "nowBidderAvatar",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionSrc",
                    "id": 10
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "bidGuildId",
                    "id": 12,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dropState",
                    "id": 13
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "haveBidUsers",
                    "id": 14,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "finBidTimes",
                    "id": 15
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemCount",
                    "id": 16
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "putAwayPrice",
                    "id": 17
                }
            ]
        },
        {
            "name": "BidInfoReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionType",
                    "id": 2
                }
            ]
        },
        {
            "name": "BidInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "AuctionItemInfo",
                    "name": "auctionInfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "BidReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "auctionId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isBuyNow",
                    "id": 3
                }
            ]
        },
        {
            "name": "BidNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "AuctionItemInfo",
                    "name": "auctionInfo",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "code",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isBuyNow",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionId",
                    "id": 5
                }
            ]
        },
        {
            "name": "MyBidReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MyBidNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "AuctionItemInfo",
                    "name": "myBidInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "BidItemUpdateNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "AuctionItemInfo",
                    "name": "newInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemStatus",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lastBidUserId",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionId",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "auctionType",
                    "id": 5
                }
            ]
        },
        {
            "name": "AuctionPutawayItemReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "count",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "position",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "price",
                    "id": 4
                }
            ]
        },
        {
            "name": "AuctionPutawayItemNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "AuctionItemInfo",
                    "name": "auctionInfos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "code",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                }
            ]
        },
        {
            "name": "BidSuccessInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "nickname",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "bidItemName",
                    "id": 2
                }
            ]
        },
        {
            "name": "BidSuccessNoticeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "BidSuccessInfo",
                    "name": "noticeInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "AuctionBuyTimesReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                }
            ]
        },
        {
            "name": "AuctionBuyTimesAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "CanBuyTimes",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "MyPutAwayItemInfoReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MyPutAwayItemInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "AuctionItemInfo",
                    "name": "myBidInfosNotInBid",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "AuctionItemInfo",
                    "name": "myBidInfosInBid",
                    "id": 2
                }
            ]
        },
        {
            "name": "MyBidInfoItemReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MyBidInfoItemAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "AuctionBidInfo",
                    "name": "myBidInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "AuctionBidInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "Id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "UserId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "AuctionId",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "AuctionType",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ItemId",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "FirstBidTime",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "FinallyBidTime",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "State",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "ExpireTime",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ItemCount",
                    "id": 10
                }
            ]
        },
        {
            "name": "RedPointStateNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isBright",
                    "id": 2
                }
            ]
        },
        {
            "name": "ConversionGoldIngotReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "num",
                    "id": 1
                }
            ]
        },
        {
            "name": "ConversionGoldIngotAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "haveUseRecharge",
                    "id": 1
                }
            ]
        },
        {
            "name": "AwakenLoadReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "AwakenLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "AwakenUnit",
                    "keytype": "int32",
                    "name": "awakens",
                    "id": 2
                }
            ]
        },
        {
            "name": "AwakenReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                }
            ]
        },
        {
            "name": "AwakenAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "AwakenUnit",
                    "name": "awakenInfo",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "AwakenUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 2
                }
            ]
        },
        {
            "name": "BagInfoReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "BagInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagMax",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "Item",
                    "name": "items",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "haveBuyTimes",
                    "id": 3
                }
            ]
        },
        {
            "name": "BagSpaceAddReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "BagSpaceAddAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagMax",
                    "id": 1
                }
            ]
        },
        {
            "name": "BagSortReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "BagSortAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "Item",
                    "name": "items",
                    "id": 1
                }
            ]
        },
        {
            "name": "GiftUseReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "position",
                    "id": 2
                }
            ]
        },
        {
            "name": "GiftUseAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 1
                }
            ]
        },
        {
            "name": "EquipRecoverReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "positions",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "EquipRecoverAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                }
            ]
        },
        {
            "name": "ItemUseReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 3
                }
            ]
        },
        {
            "name": "ItemUseAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                }
            ]
        },
        {
            "name": "EquipDestroyReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "positions",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "count",
                    "id": 2
                }
            ]
        },
        {
            "name": "EquipDestroyAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                }
            ]
        },
        {
            "name": "WarehouseInfoReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "WarehouseInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagMax",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "Item",
                    "name": "items",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "haveBuyTimes",
                    "id": 3
                }
            ]
        },
        {
            "name": "WareHouseSpaceAddReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "WareHouseSpaceAddAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagMax",
                    "id": 1
                }
            ]
        },
        {
            "name": "WarehouseAddReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "positions",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "WarehouseAddAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "Item",
                    "name": "items",
                    "id": 1
                }
            ]
        },
        {
            "name": "WarehouseShiftOutReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "positions",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "WarehouseShiftOutAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "Item",
                    "name": "items",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "WarehouseSortReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "WarehouseSortAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "Item",
                    "name": "items",
                    "id": 1
                }
            ]
        },
        {
            "name": "GetBossFamilyInfoReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bossFamilyType",
                    "id": 1
                }
            ]
        },
        {
            "name": "GetBossFamilyInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "bossFamilyInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterBossFamilyReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "CardActivityApplyGetReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "times",
                    "id": 1
                }
            ]
        },
        {
            "name": "CardActivityApplyGetAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "cards",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cardTime",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 4
                },
                {
                    "rule": "repeated",
                    "type": "CardInfoUnit",
                    "name": "myDrawInfo",
                    "id": 5
                },
                {
                    "rule": "repeated",
                    "type": "CardInfoUnit",
                    "name": "serverDrawInfo",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "integral",
                    "id": 7
                }
            ]
        },
        {
            "name": "CardActivityInfosReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "CardActivityInfosAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "integral",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "totalDrawCardTimes",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "CardInfoUnit",
                    "name": "myDrawInfo",
                    "id": 3
                },
                {
                    "rule": "repeated",
                    "type": "CardInfoUnit",
                    "name": "serverDrawInfo",
                    "id": 4
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "haveGetIndex",
                    "id": 5,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "nowSeason",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "mergeMark",
                    "id": 7
                }
            ]
        },
        {
            "name": "CardInfoUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "count",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "time",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "userName",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 5
                }
            ]
        },
        {
            "name": "CardInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "time",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "userName",
                    "id": 3
                }
            ]
        },
        {
            "name": "GetIntegralAwardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "index",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "times",
                    "id": 2
                }
            ]
        },
        {
            "name": "GetIntegralAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "haveGetIndex",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "integral",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                }
            ]
        },
        {
            "name": "CardCloseNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isClose",
                    "id": 1
                }
            ]
        },
        {
            "name": "ChallengeInfoReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ChallengeInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "peopleInfos",
                    "keytype": "int32",
                    "name": "challengePeopleInfo",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "PeopleInfo",
                    "name": "BottomUserInfo",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "PeopleInfo",
                    "name": "ApplyUserInfo",
                    "id": 3
                },
                {
                    "rule": "repeated",
                    "type": "PeopleInfo",
                    "name": "FirstPlayer",
                    "id": 4
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "JoinServer",
                    "id": 5,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "ApplyChallengeReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ApplyChallengeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "PeopleInfo",
                    "name": "ApplyUserInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "ChallengeEachRoundPeopleReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ChallengeEachRoundPeopleAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "nowRound",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "PeopleInfo",
                    "name": "challengePeopleInfo",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "PeopleInfo",
                    "name": "BottomUserInfo",
                    "id": 3
                }
            ]
        },
        {
            "name": "peopleInfos",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "PeopleInfo",
                    "name": "peopleInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "BottomPourReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                }
            ]
        },
        {
            "name": "BottomPourAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "state",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "PeopleInfo",
                    "name": "BottomUserInfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "PeopleInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "avatar",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverId",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "combat",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "guildName",
                    "id": 6
                }
            ]
        },
        {
            "name": "ChallengeOpenNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isOpen",
                    "id": 1
                }
            ]
        },
        {
            "name": "ChallengeRoundEndNtf",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ChallengeApplyUserInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "PeopleInfo",
                    "name": "ApplyUserInfo",
                    "id": 3
                }
            ]
        },
        {
            "name": "ChatMessageNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "sender",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "msg",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ts",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "toId",
                    "id": 5
                }
            ]
        },
        {
            "name": "ChatMessageListReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ChatMessageListAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "ChatMessageNtf",
                    "name": "msgs",
                    "id": 1
                }
            ]
        },
        {
            "name": "ChatSendReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "msg",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "toId",
                    "id": 3
                }
            ]
        },
        {
            "name": "ChatSendAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isBanSpeak",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "banSpeakTime",
                    "id": 2
                }
            ]
        },
        {
            "name": "ChatBanNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "endTime",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "reason",
                    "id": 2
                }
            ]
        },
        {
            "name": "ChatBanRemoveNtf",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ChuanShiWearReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagPos",
                    "id": 2
                }
            ]
        },
        {
            "name": "ChuanShiWearAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipId",
                    "id": 3
                }
            ]
        },
        {
            "name": "ChuanShiRemoveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                }
            ]
        },
        {
            "name": "ChuanShiRemoveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                }
            ]
        },
        {
            "name": "ChuanShiDeComposeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagPos",
                    "id": 1
                }
            ]
        },
        {
            "name": "ChuanShiDeComposeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                }
            ]
        },
        {
            "name": "ChuanshiStrengthenReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stone",
                    "id": 3
                }
            ]
        },
        {
            "name": "ChuanshiStrengthenAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isUp",
                    "id": 4
                }
            ]
        },
        {
            "name": "ClearReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "propIndex",
                    "id": 3
                }
            ]
        },
        {
            "name": "ClearAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "EquipClearArr",
                    "name": "equipClear",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "CompetitveLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "CompetitveLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "seasonTimes",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sessionWinTimes",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "remainChallengeTimes",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "todayCanBuyTimes",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userScore",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "yestardayReward",
                    "id": 6
                },
                {
                    "rule": "repeated",
                    "type": "CompetitveRankInfo",
                    "name": "seasonRank",
                    "id": 7
                },
                {
                    "rule": "repeated",
                    "type": "CompetitveRankInfo",
                    "name": "lastSeasonRank",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "beginTimes",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lastSeasonUserRank",
                    "id": 10
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lastSeasonUserRankScore",
                    "id": 11
                }
            ]
        },
        {
            "name": "EnterCompetitveFightReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "challengeUid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "challengeRanking",
                    "id": 2
                }
            ]
        },
        {
            "name": "CompetitveFightNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "myRank",
                    "id": 3
                }
            ]
        },
        {
            "name": "BuyCompetitveChallengeTimesReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "BuyCompetitveChallengeTimesAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "residueTimes",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "todayCanBuyTimes",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "RefCompetitveRankReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "RefCompetitveRankAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "userInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "score",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "HeroInfo",
                    "name": "heros",
                    "id": 3
                }
            ]
        },
        {
            "name": "CompetitveRankInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ranking",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "avatar",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "score",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "nickName",
                    "id": 4
                }
            ]
        },
        {
            "name": "GetCompetitveDailyRewardReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "GetCompetitveDailyRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "haveGetRewardState",
                    "id": 2
                }
            ]
        },
        {
            "name": "CompetitveMultipleClaimReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "num",
                    "id": 1
                }
            ]
        },
        {
            "name": "CompetitveMultipleClaimAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "seasonScore",
                    "id": 1
                }
            ]
        },
        {
            "name": "ComposeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "subId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "composeNum",
                    "id": 3
                }
            ]
        },
        {
            "name": "ComposeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "ComposeEquipReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "composeEquipSubId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isLuckyStone",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "bigLuckyStone",
                    "id": 3
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "bagPos",
                    "id": 4,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 5,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "items",
                    "id": 6,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "ComposeEquipAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "composeEquipSubId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isLuckyStone",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "bigLuckyStone",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "ComposeChuanShiEquipReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "composeSubId",
                    "id": 1
                }
            ]
        },
        {
            "name": "ComposeChuanShiEquipAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "composeSubId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "CutTreasureUpLvReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "CutTreasureUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cutTreasureLv",
                    "id": 1
                }
            ]
        },
        {
            "name": "CutTreasureUseReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "CutTreasureUseAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "useTime",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cdEndTime",
                    "id": 2
                }
            ]
        },
        {
            "name": "DaBaoEquipUpReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipType",
                    "id": 1
                }
            ]
        },
        {
            "name": "DaBaoEquipUpAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 2
                }
            ]
        },
        {
            "name": "EnterDaBaoMysteryReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "DaBaoMysteryResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                }
            ]
        },
        {
            "name": "DaBaoMysteryEnergyItemBuyReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "shopId",
                    "id": 1
                }
            ]
        },
        {
            "name": "DaBaoMysteryEnergyAddReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                }
            ]
        },
        {
            "name": "DaBaoMysteryEnergyAddAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                }
            ]
        },
        {
            "name": "DaBaoMysteryEnergyNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "energy",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterDailyActivityReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "activityId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 2
                }
            ]
        },
        {
            "name": "DailyActivityResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "activityId",
                    "id": 1
                }
            ]
        },
        {
            "name": "DailyActivityListReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "DailyActivityListAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "DailyActivityInfo",
                    "name": "list",
                    "id": 1
                }
            ]
        },
        {
            "name": "DailyActivityInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "activityId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "startTime",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 3
                }
            ]
        },
        {
            "name": "DailyPackBuyReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "DailyPackBuyAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                }
            ]
        },
        {
            "name": "DailyRankLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "DailyRankLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "RankInfo",
                    "name": "ranks",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "self",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "selfScore",
                    "id": 4
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "haveGetIds",
                    "id": 5,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "buyGiftInfos",
                    "id": 6
                }
            ]
        },
        {
            "name": "DailyRankGetMarkRewardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "DailyRankGetMarkRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "haveGetIds",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "DailyRankBuyGiftReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "DailyRankBuyGiftAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "buyGiftInfos",
                    "id": 2
                }
            ]
        },
        {
            "name": "DailyTaskLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "DailyTaskLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dayExp",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "weekExp",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dayResourcesBackExp",
                    "id": 3
                },
                {
                    "rule": "repeated",
                    "type": "HaveChallengeTime",
                    "name": "haveChallengeTimes",
                    "id": 4
                },
                {
                    "rule": "repeated",
                    "type": "ResourcesBackInfo",
                    "name": "ResourcesBackInfos",
                    "id": 5
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "GetDayRewardIds",
                    "id": 6,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "GetWeekRewardIds",
                    "id": 7,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "BuyChallengeTimeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "activityId",
                    "id": 1
                }
            ]
        },
        {
            "name": "BuyChallengeTimeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "activityId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "haveChallengeTime",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyChallengTimes",
                    "id": 3
                }
            ]
        },
        {
            "name": "GetExpReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "activityId",
                    "id": 1
                }
            ]
        },
        {
            "name": "GetExpAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dayExp",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "weekExp",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dayResourcesBackExp",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isCanGetAward",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "activityId",
                    "id": 5
                }
            ]
        },
        {
            "name": "GetAwardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 2
                }
            ]
        },
        {
            "name": "GetAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "GetDayRewardIds",
                    "id": 6,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "GetWeekRewardIds",
                    "id": 7,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "ResourcesBackGetRewardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "activityId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "backTimes",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "useIngot",
                    "id": 3
                }
            ]
        },
        {
            "name": "ResourcesBackGetRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "ResourcesBackInfo",
                    "name": "ResourcesBackInfos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dayResourcesBackExp",
                    "id": 2
                }
            ]
        },
        {
            "name": "ResourcesBackGetAllRewardReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ResourcesBackGetAllRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "ResourcesBackInfo",
                    "name": "ResourcesBackInfos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dayResourcesBackExp",
                    "id": 2
                }
            ]
        },
        {
            "name": "HaveChallengeTime",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "activityId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "haveChallengeTime",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isGetAward",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyChallengTimes",
                    "id": 4
                }
            ]
        },
        {
            "name": "ResourcesBackInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "activityId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "residueChallengeTimes",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "haveChallengeTimes",
                    "id": 3
                }
            ]
        },
        {
            "name": "DarkPalaceLoadReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "floor",
                    "id": 1
                }
            ]
        },
        {
            "name": "DarkPalaceLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "DarkPalaceBossNtf",
                    "name": "darkPalaceBoss",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterDarkPalaceFightReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "DarkPalaceFightResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "winner",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isHelper",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "helpNum",
                    "id": 7
                }
            ]
        },
        {
            "name": "DarkPalaceBuyNumReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "use",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "DarkPalaceBuyNumAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "DarkPalaceBossNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "float",
                    "name": "blood",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "reliveTime",
                    "id": 3
                }
            ]
        },
        {
            "name": "EnterDarkPalaceHelpFightReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "DarkPalaceHelpFightResultNtf",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "DarkPalaceDareNumNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "DictateUpReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "body",
                    "id": 2
                }
            ]
        },
        {
            "name": "DictateUpAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "DictateInfo",
                    "name": "dictateInfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "DragonEquipUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 2
                }
            ]
        },
        {
            "name": "DragonEquipUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 3
                }
            ]
        },
        {
            "name": "ElfFeedReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "positions",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "ElfFeedAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "receiveLimit",
                    "id": 4
                }
            ]
        },
        {
            "name": "ElfSkillUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 1
                }
            ]
        },
        {
            "name": "ElfSkillUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillLv",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "skillBag",
                    "id": 3
                }
            ]
        },
        {
            "name": "ElfSkillChangePosReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                }
            ]
        },
        {
            "name": "ElfSkillChangePosAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "skillBag",
                    "id": 1
                }
            ]
        },
        {
            "name": "EquipChangeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagPos",
                    "id": 3
                }
            ]
        },
        {
            "name": "EquipChangeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "EquipUnit",
                    "keytype": "int32",
                    "name": "equips",
                    "id": 2
                }
            ]
        },
        {
            "name": "EquipLockReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                }
            ]
        },
        {
            "name": "EquipLockAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "lock",
                    "id": 3
                }
            ]
        },
        {
            "name": "EquipStrengthenReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                }
            ]
        },
        {
            "name": "EquipStrengthenAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "EquipGrid",
                    "name": "equipGrids",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isUp",
                    "id": 3
                }
            ]
        },
        {
            "name": "EquipRemoveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                }
            ]
        },
        {
            "name": "EquipRemoveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                }
            ]
        },
        {
            "name": "EquipBlessNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lucky",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "res",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "EquipUnit",
                    "name": "equip",
                    "id": 4
                }
            ]
        },
        {
            "name": "EquipStrengthenAutoReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isBreak",
                    "id": 2
                }
            ]
        },
        {
            "name": "EquipStrengthenAutoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "EquipGrid",
                    "name": "equipGrids",
                    "id": 2
                }
            ]
        },
        {
            "name": "ExpPoolLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ExpPoolLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "worlLvl",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "expPool",
                    "id": 3
                }
            ]
        },
        {
            "name": "ExpPoolUpGradeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "index",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "times",
                    "id": 2
                }
            ]
        },
        {
            "name": "ExpPoolUpGradeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "Lvl",
                    "id": 2
                }
            ]
        },
        {
            "name": "ExpStageFightReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "ExpStageDareNumNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "ExpStageFightResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "exp",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "monsterNum",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "grade",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "layer",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isFree",
                    "id": 6
                }
            ]
        },
        {
            "name": "ExpStageDoubleReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "ExpStageDoubleAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "exp",
                    "id": 2
                }
            ]
        },
        {
            "name": "ExpStageRefNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isRef",
                    "id": 1
                }
            ]
        },
        {
            "name": "ExpStageBuyNumNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "ExpStageSweepReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "ExpStageSweepAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "exp",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "monsterNum",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "grade",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 5
                }
            ]
        },
        {
            "name": "ExpStageBuyNumReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "use",
                    "id": 1
                }
            ]
        },
        {
            "name": "ExpStageBuyNumAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "FabaoActiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "FabaoActiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Fabao",
                    "name": "fabao",
                    "id": 1
                }
            ]
        },
        {
            "name": "FabaoUpLevelReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "FabaoUpLevelAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Fabao",
                    "name": "fabao",
                    "id": 1
                }
            ]
        },
        {
            "name": "FabaoSkillActiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 2
                }
            ]
        },
        {
            "name": "FabaoSkillActiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Fabao",
                    "name": "fabao",
                    "id": 1
                }
            ]
        },
        {
            "name": "FashionUpLevelReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fashionId",
                    "id": 2
                }
            ]
        },
        {
            "name": "FashionUpLevelAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "Fashion",
                    "name": "fashion",
                    "id": 2
                }
            ]
        },
        {
            "name": "FashionWearReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fashionId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isWear",
                    "id": 3
                }
            ]
        },
        {
            "name": "FashionWearAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "wearFashionId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isWear",
                    "id": 3
                }
            ]
        },
        {
            "name": "FieldBossLoadReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "area",
                    "id": 1
                }
            ]
        },
        {
            "name": "FieldBossLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "FieldBossNtf",
                    "name": "fieldBoss",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterFieldBossFightReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterFieldBossFightAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "FieldBossFightResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "winner",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 5
                }
            ]
        },
        {
            "name": "FieldBossBuyNumReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "use",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "FieldBossBuyNumAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "FieldBossFirstReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FieldBossFirstAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "firstReceive",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "FieldBossNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "float",
                    "name": "blood",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "reliveTime",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "area",
                    "id": 4
                }
            ]
        },
        {
            "name": "FieldFightLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FieldFightLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "FieldFightListInfo",
                    "name": "listInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "myCombat",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "remainChallengeTimes",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "todayCanBuyTimes",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "changeRivalCd",
                    "id": 5
                },
                {
                    "rule": "repeated",
                    "type": "FieldFightBeatBackUserInfo",
                    "name": "BeatBackOwnUserInfo",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isCheckNoPromptState",
                    "id": 7
                }
            ]
        },
        {
            "name": "EnterFieldFightReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "challengeUid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isBeatBack",
                    "id": 2
                }
            ]
        },
        {
            "name": "FieldFightNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "remainChallengeTimes",
                    "id": 3
                },
                {
                    "rule": "repeated",
                    "type": "FieldFightListInfo",
                    "name": "listInfo",
                    "id": 4
                },
                {
                    "rule": "repeated",
                    "type": "FieldFightBeatBackUserInfo",
                    "name": "BeatBackOwnUserInfo",
                    "id": 5
                }
            ]
        },
        {
            "name": "BuyFieldFightChallengeTimesReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "BuyFieldFightChallengeTimesAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "residueTimes",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "todayCanBuyTimes",
                    "id": 2
                }
            ]
        },
        {
            "name": "RefFieldFightRivalUserReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "RefFieldFightRivalUserAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "FieldFightListInfo",
                    "name": "listInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "changeRivalCd",
                    "id": 2
                }
            ]
        },
        {
            "name": "FieldFightListInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "difficultyLevel",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "avatar",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "nickName",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "combat",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userLv",
                    "id": 5
                },
                {
                    "rule": "repeated",
                    "type": "itemUnit",
                    "name": "rewardInfos",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sex",
                    "id": 9
                }
            ]
        },
        {
            "name": "FieldFightBeatBackUserInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "nickName",
                    "id": 2
                }
            ]
        },
        {
            "name": "BeatBackInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "FieldFightBeatBackUserInfo",
                    "name": "BeatBackOwnUserInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterPublicCopyReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "condition",
                    "id": 2
                }
            ]
        },
        {
            "name": "EnterPublicCopyAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "failReason",
                    "id": 1
                }
            ]
        },
        {
            "name": "FightItemUseReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FightItemUseAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "useTimes",
                    "id": 2
                }
            ]
        },
        {
            "name": "FightUserReliveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "safeRelive",
                    "id": 1
                }
            ]
        },
        {
            "name": "FightUserReliveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "reliveTimes",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "reliveByIngotTimes",
                    "id": 2
                }
            ]
        },
        {
            "name": "FightPickUpReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "dropItemIds",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "FightPickUpAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "itemUnit",
                    "keytype": "int32",
                    "name": "items",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isOneKey",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "inMail",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "ErrorAck",
                    "name": "err",
                    "id": 4
                }
            ]
        },
        {
            "name": "FightGetCheerNumReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FightGetCheerNumNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cheerNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildCheerNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "FightCheerReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FightCheerAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cheerNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildCheerNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "FightCheerNumChangeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildCheerNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "FightPotionReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FightPotionAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "coolDown",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverTime",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "endTime",
                    "id": 3
                }
            ]
        },
        {
            "name": "FightPotionCdReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FightPotionCdAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "coolDown",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverTime",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "endTime",
                    "id": 3
                }
            ]
        },
        {
            "name": "FightCollectionReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FightCollectionAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "startTime",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 3
                }
            ]
        },
        {
            "name": "FightCollectionNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                }
            ]
        },
        {
            "name": "FightCollectionCancelReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FightCollectionCancelAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "result",
                    "id": 1
                }
            ]
        },
        {
            "name": "FightApplyForHelpReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "helpUserId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "source",
                    "id": 2
                }
            ]
        },
        {
            "name": "FightApplyForHelpAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "failReason",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "helpUserId",
                    "id": 3
                }
            ]
        },
        {
            "name": "FightApplyForHelpNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "reqHelpUserId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "reqHelpName",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "source",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "reqHelpUser",
                    "id": 5
                }
            ]
        },
        {
            "name": "FightAskForHelpResultReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isAgree",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "reqHelpUserId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "helpStageId",
                    "id": 3
                }
            ]
        },
        {
            "name": "FightAskForHelpResultAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isAgree",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "reqHelpUserId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "helpStageId",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "enterErr",
                    "id": 4
                }
            ]
        },
        {
            "name": "FightAskForHelpResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isAgree",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "helpUserId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 3
                }
            ]
        },
        {
            "name": "FightItemsAddNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "addSource",
                    "id": 3
                }
            ]
        },
        {
            "name": "FirstDropLoadReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "types",
                    "id": 1
                }
            ]
        },
        {
            "name": "FirstDropLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "types",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "GetDropItemInfo",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "AllDropItemGetCount",
                    "id": 3
                }
            ]
        },
        {
            "name": "GetFirstDropAwardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "GetFirstDropAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "types",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "GetDropItemInfo",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "DropItemGetCount",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "GetAllFirstDropAwardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "types",
                    "id": 1
                }
            ]
        },
        {
            "name": "GetAllFirstDropAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "types",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "GetDropItemInfo",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "DropItemGetCount",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "GetAllRedPacketReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "infos",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "GetAllRedPacketAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "usePacketNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "GetAllFirstDropAwardNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "DropItemGetCount",
                    "id": 1
                }
            ]
        },
        {
            "name": "FirstDropRedPointNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "items",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "FirstRechargeRewardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "day",
                    "id": 1
                }
            ]
        },
        {
            "name": "FirstRechargeRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "day",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "FirstRechargeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isRecharge",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "openDay",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitLvId",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitSkillUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitSkillUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillLv",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitSkillUpStarReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitSkillUpStarAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillStar",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitSkillChangeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillSlot",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitSkillChangeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillSlot",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitSkillResetReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitSkillResetAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillLv",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillStar",
                    "id": 3
                }
            ]
        },
        {
            "name": "FitFashionUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitFashionId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitFashionUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitFashionId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitFashionLv",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitFashionChangeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitFashionId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitFashionChangeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitFashionId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitSkillActiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitSkillActiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitSkillId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitEnterReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FitEnterAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cdStartTime",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cdEndTime",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitCancleReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FitCancleAck",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FitHolyEquipComposeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitHolyEquipComposeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "suitType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "FitHolyEquipUnit",
                    "name": "fitHolyEquip",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitHolyEquipDeComposeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagPos",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitHolyEquipDeComposeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitHolyEquipWearReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagPos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitHolyEquipWearAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "suitType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "FitHolyEquipUnit",
                    "name": "fitHolyEquip",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitHolyEquipRemoveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "suitType",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitHolyEquipRemoveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "suitType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "FitHolyEquipUnit",
                    "name": "fitHolyEquip",
                    "id": 2
                }
            ]
        },
        {
            "name": "FitHolyEquipSuitSkillChangeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "suitId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FitHolyEquipSuitSkillChangeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "suitId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendListReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FriendListAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "FriendInfo",
                    "name": "friendList",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendAddReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "userId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendAddAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "FriendInfo",
                    "name": "friendInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendDelReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "userId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendDelAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "userId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendBlockAddReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "userId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendBlockAddAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "FriendInfo",
                    "name": "friendList",
                    "id": 2
                }
            ]
        },
        {
            "name": "FriendSearchReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendSearchAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "FriendInfo",
                    "name": "friendList",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendBlockListReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FriendBlockListAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "FriendInfo",
                    "name": "friendList",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendBlockDelReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "userId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendBlockDelAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "userId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendMsgReadReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "friendId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendMsgReadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "friendId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendUserInfoReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "friendId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendUserInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "FriendUserInfo",
                    "name": "friendUserInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendMsgReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "friendId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendMsgAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "friendId",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "MsgLog",
                    "name": "msgLog",
                    "id": 2
                }
            ]
        },
        {
            "name": "FriendApplyAddReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "friendId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendApplyAddNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "friendId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendApplyAgreeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "friendId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendApplyAgreeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "friendId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendApplyRefuseReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "friendId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendApplyRefuseNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "friendId",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendApplyListReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FriendApplyListAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "FriendApplyInfo",
                    "name": "applyList",
                    "id": 1
                }
            ]
        },
        {
            "name": "FriendInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "userInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isOnline",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "outTime",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isRead",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "MsgLog",
                    "name": "lastMsg",
                    "id": 5
                }
            ]
        },
        {
            "name": "MsgLog",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "msg",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "time",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isMy",
                    "id": 3
                }
            ]
        },
        {
            "name": "FriendApplyInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "nickName",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "avatar",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sex",
                    "id": 6
                }
            ]
        },
        {
            "name": "ReportGtNoReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "gateNo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "gsNo",
                    "id": 2
                }
            ]
        },
        {
            "name": "ReportGtNoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "gateNo",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "gateAddr",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "gsNo",
                    "id": 4
                }
            ]
        },
        {
            "name": "GsBroadCastNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bytes",
                    "name": "msg",
                    "id": 1
                }
            ]
        },
        {
            "name": "GsMsgNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "indexes",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "bytes",
                    "name": "msg",
                    "id": 2
                }
            ]
        },
        {
            "name": "ReConnectReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "token",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "openId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "magic",
                    "id": 3
                }
            ]
        },
        {
            "name": "ReConnectAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fail",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "UserLoginInfo",
                    "name": "user",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "reConnectToken",
                    "id": 99
                }
            ]
        },
        {
            "name": "LogoutReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "reason",
                    "id": 1
                }
            ]
        },
        {
            "name": "LogoutAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 1
                }
            ]
        },
        {
            "name": "OfflineReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "src",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "CNo",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "reason",
                    "id": 4
                }
            ]
        },
        {
            "name": "ChatReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "typ",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "desId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pNo",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "msg",
                    "id": 4
                }
            ]
        },
        {
            "name": "ChatRsp",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "typ",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "srcId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "msg",
                    "id": 3
                }
            ]
        },
        {
            "name": "OnlineNumReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "gateNo",
                    "id": 1
                }
            ]
        },
        {
            "name": "OnlineNumAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "gsNo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "onlines",
                    "id": 2
                }
            ]
        },
        {
            "name": "MsgNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "typ",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "srcId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "msg",
                    "id": 3
                }
            ]
        },
        {
            "name": "PreJumpGsReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cNo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pNo",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 3
                }
            ]
        },
        {
            "name": "DoJumpGsReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cNo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pNo",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "desGsNo",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "reason",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "bytes",
                    "name": "args",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 6
                }
            ]
        },
        {
            "name": "JumpGsReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cNo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pNo",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "reason",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "arg1",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 6
                }
            ]
        },
        {
            "name": "BenchMarkReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "magic",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "magic2",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "magic3",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "magic4",
                    "id": 4
                }
            ]
        },
        {
            "name": "BenchMarkAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "magic",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "magic2",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "magic3",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "magic4",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "magic100",
                    "id": 100
                }
            ]
        },
        {
            "name": "KeepAliveRpt",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "mark",
                    "id": 1
                }
            ]
        },
        {
            "name": "OpenGiftReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "chooseItemId",
                    "id": 3,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "num",
                    "id": 4
                }
            ]
        },
        {
            "name": "OpenGiftAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                }
            ]
        },
        {
            "name": "GiftCodeRewardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "code",
                    "id": 1
                }
            ]
        },
        {
            "name": "GiftCodeRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "code",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "LimitedGiftNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "LimitedGiftInfo",
                    "name": "list",
                    "id": 1
                }
            ]
        },
        {
            "name": "LimitedGiftBuyReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                }
            ]
        },
        {
            "name": "LimitedGiftBuyAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 2
                }
            ]
        },
        {
            "name": "LimitedGiftReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "LimitedGiftInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "grade",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "startTime",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 5
                }
            ]
        },
        {
            "name": "OpenGiftBuyNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "openGiftId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 3
                }
            ]
        },
        {
            "name": "OpenGiftEndTimeReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "OpenGiftEndTimeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 1
                }
            ]
        },
        {
            "name": "GodEquipActiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 2
                }
            ]
        },
        {
            "name": "GodEquipActiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GodEquip",
                    "name": "godEquip",
                    "id": 2
                }
            ]
        },
        {
            "name": "GodEquipUpLevelReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 2
                }
            ]
        },
        {
            "name": "GodEquipUpLevelAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GodEquip",
                    "name": "godEquip",
                    "id": 2
                }
            ]
        },
        {
            "name": "GodEquipBloodReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "godEquipId",
                    "id": 2
                }
            ]
        },
        {
            "name": "GodEquipBloodAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GodEquip",
                    "name": "godEquip",
                    "id": 2
                }
            ]
        },
        {
            "name": "GrowFundBuyReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "GrowFundBuyAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isBuy",
                    "id": 1
                }
            ]
        },
        {
            "name": "GrowFundRewardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "GrowFundRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "EnterGuardPillarReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "GuardPillarResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rounds",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rank",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "roundGoods",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "rankGoods",
                    "id": 5
                }
            ]
        },
        {
            "name": "GuildLoadInfoReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "GuildLoadInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GuildInfo",
                    "name": "guildInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "CreateGuildReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "GuildName",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "GuildIcon",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "Notice",
                    "id": 3
                }
            ]
        },
        {
            "name": "CreateGuildAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GuildInfo",
                    "name": "guildInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "success",
                    "id": 2
                }
            ]
        },
        {
            "name": "JoinGuildCombatLimitReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "combat",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isAgree",
                    "id": 2
                }
            ]
        },
        {
            "name": "JoinGuildCombatLimitAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "success",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "limitCombat",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isAgree",
                    "id": 3
                }
            ]
        },
        {
            "name": "ModifyBulletinReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "content",
                    "id": 1
                }
            ]
        },
        {
            "name": "ModifyBulletinAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "success",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "content",
                    "id": 2
                }
            ]
        },
        {
            "name": "QuitGuildReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "QuitGuildAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "success",
                    "id": 1
                }
            ]
        },
        {
            "name": "KickOutReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "kickUserId",
                    "id": 1
                }
            ]
        },
        {
            "name": "KickOutAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "joinCd",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "kickUserId",
                    "id": 2
                }
            ]
        },
        {
            "name": "ImpeachPresidentReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ImpeachPresidentAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "success",
                    "id": 1
                }
            ]
        },
        {
            "name": "GuildCheckMemberInfoReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                }
            ]
        },
        {
            "name": "GuildCheckMemberInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GuildInfo",
                    "name": "guildInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "ApplyJoinGuildReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildId",
                    "id": 1
                }
            ]
        },
        {
            "name": "ApplyJoinGuildAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "success",
                    "id": 1
                }
            ]
        },
        {
            "name": "GuildAssignReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "position",
                    "id": 2
                }
            ]
        },
        {
            "name": "GuildAssignAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "success",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "assignUserId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "nowPosition",
                    "id": 3
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "positionCount",
                    "id": 4
                }
            ]
        },
        {
            "name": "AllGuildInfosReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "AllGuildInfosAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "GuildInfo",
                    "name": "guildInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "DissolveGuildReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "DissolveGuildAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "success",
                    "id": 1
                }
            ]
        },
        {
            "name": "JoinGuildDisposeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isAgree",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "applyUserId",
                    "id": 2
                }
            ]
        },
        {
            "name": "JoinGuildDisposeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "success",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "BriefUserInfo",
                    "name": "applyUserInfo",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isHaveJoinGuild",
                    "id": 3
                }
            ]
        },
        {
            "name": "GetApplyUserListReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "GetApplyUserListAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "BriefUserInfo",
                    "name": "applyUserInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "GuildInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "GuildName",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildLv",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "joinCd",
                    "id": 4
                },
                {
                    "rule": "repeated",
                    "type": "GuildMenberInfo",
                    "name": "guildMenberInfo",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "notice",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "combat",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isAutoAgree",
                    "id": 9
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "positionCount",
                    "id": 10
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "onlineUser",
                    "id": 11
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "GuildContributionValue",
                    "id": 12
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverId",
                    "id": 13
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildPeopleNum",
                    "id": 14
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "huiZhangLv",
                    "id": 15
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "huiZhangName",
                    "id": 16
                }
            ]
        },
        {
            "name": "GuildMenberInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "position",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "offlineTime",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildCapital",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildContribution",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "nickName",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "avatar",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "combat",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 10
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sex",
                    "id": 11
                }
            ]
        },
        {
            "name": "JoinGuildSuccessNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "success",
                    "id": 3
                }
            ]
        },
        {
            "name": "AllJoinGuildDisposeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isAgree",
                    "id": 1
                }
            ]
        },
        {
            "name": "AllJoinGuildDisposeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "BriefUserInfo",
                    "name": "applyUserInfo",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isFullState",
                    "id": 3
                }
            ]
        },
        {
            "name": "ApplyJoinGuildReDotNtf",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ImpeachPresidentNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "newHuiZhangUserId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "nowPosition",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "oldHuiZhangUserId",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "nowPosition1",
                    "id": 4
                }
            ]
        },
        {
            "name": "BroadcastGuildChangeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "types",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "userInfos",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "positionCount",
                    "id": 4
                },
                {
                    "rule": "repeated",
                    "type": "GuildMenberInfo",
                    "name": "guildMenberInfo",
                    "id": 5
                }
            ]
        },
        {
            "name": "GuildActivityOpenNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildActivityId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 2
                }
            ]
        },
        {
            "name": "GuildActivityLoadReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildActivityId",
                    "id": 1
                }
            ]
        },
        {
            "name": "GuildActivityLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildActivityId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isClose",
                    "id": 3
                }
            ]
        },
        {
            "name": "GuildBonfireLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "GuildBonfireLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "float",
                    "name": "expAddPercent",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "WoodPeople",
                    "name": "peopleList",
                    "id": 2
                }
            ]
        },
        {
            "name": "GuildBonfireAddExpReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "consumptionType",
                    "id": 1
                }
            ]
        },
        {
            "name": "GuildBonfireAddExpAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "float",
                    "name": "expAddPercent",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "WoodPeople",
                    "name": "peopleList",
                    "id": 2
                }
            ]
        },
        {
            "name": "WoodPeople",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "nickName",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "avatar",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "times",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "types",
                    "id": 4
                }
            ]
        },
        {
            "name": "EnterGuildBonfireFightReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "GuildBonfireFightNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 1
                }
            ]
        },
        {
            "name": "GuildBonfireOpenStateNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isOpen",
                    "id": 1
                }
            ]
        },
        {
            "name": "HellBossLoadReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "floor",
                    "id": 1
                }
            ]
        },
        {
            "name": "HellBossLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "floor",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "HellBossNtf",
                    "name": "list",
                    "id": 2
                }
            ]
        },
        {
            "name": "HellBossBuyNumReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "use",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "HellBossBuyNumAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "HellBossDareNumNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterHellBossFightReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "HellBossFightResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "winner",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isHelper",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "helpNum",
                    "id": 7
                }
            ]
        },
        {
            "name": "HellBossNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "float",
                    "name": "blood",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "reliveTime",
                    "id": 3
                }
            ]
        },
        {
            "name": "HolyBeastLoadInfoReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "HolyBeastLoadInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "HolyBeastInfos",
                    "keytype": "int32",
                    "name": "holyBeastInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "HolyBeastActivateReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 2
                }
            ]
        },
        {
            "name": "HolyBeastActivateAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "HolyBeastInfos",
                    "keytype": "int32",
                    "name": "holyBeastInfos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "HolyPoint",
                    "id": 2
                }
            ]
        },
        {
            "name": "HolyBeastUpStarReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 2
                }
            ]
        },
        {
            "name": "HolyBeastUpStarAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "HolyBeastInfos",
                    "keytype": "int32",
                    "name": "holyBeastInfos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "HolyPoint",
                    "id": 2
                }
            ]
        },
        {
            "name": "HolyBeastPointAddReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "useItemId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "count",
                    "id": 3
                }
            ]
        },
        {
            "name": "HolyBeastPointAddAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "HolyPoint",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                }
            ]
        },
        {
            "name": "HolyBeastChoosePropReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "index",
                    "id": 3
                }
            ]
        },
        {
            "name": "HolyBeastChoosePropAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "HolyBeastInfos",
                    "keytype": "int32",
                    "name": "holyBeastInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "HolyBeastRestReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 2
                }
            ]
        },
        {
            "name": "HolyBeastRestAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "HolyBeastInfos",
                    "keytype": "int32",
                    "name": "holyBeastInfos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "HolyPoint",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                }
            ]
        },
        {
            "name": "HolyBeastInfos",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "HolyBeastInfo",
                    "name": "holyBeastInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "allPonts",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 3
                }
            ]
        },
        {
            "name": "HolyBeastInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "star",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "chooseProperty",
                    "id": 3
                }
            ]
        },
        {
            "name": "HolyActiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "HolyActiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Holy",
                    "name": "holy",
                    "id": 1
                }
            ]
        },
        {
            "name": "HolyUpLevelReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 2
                }
            ]
        },
        {
            "name": "HolyUpLevelAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Holy",
                    "name": "holy",
                    "id": 1
                }
            ]
        },
        {
            "name": "HolySkillActiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "hid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "hlv",
                    "id": 2
                }
            ]
        },
        {
            "name": "HolySkillActiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Holy",
                    "name": "holy",
                    "id": 1
                }
            ]
        },
        {
            "name": "HolySkillUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "hid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "hlv",
                    "id": 2
                }
            ]
        },
        {
            "name": "HolySkillUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Holy",
                    "name": "holy",
                    "id": 1
                }
            ]
        },
        {
            "name": "InsideUpStarReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "InsideUpStarAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "InsideInfo",
                    "name": "insideInfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "InsideUpGradeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "InsideUpGradeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "InsideInfo",
                    "name": "insideInfo",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "res",
                    "id": 3
                }
            ]
        },
        {
            "name": "InsideUpOrderReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "InsideUpOrderAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "InsideInfo",
                    "name": "insideInfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "InsideSkillUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 2
                }
            ]
        },
        {
            "name": "InsideSkillUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "InsideInfo",
                    "name": "insideInfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "InsideAutoUpReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "InsideAutoUpAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "InsideInfo",
                    "name": "insideInfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "JewelMakeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "jewelPos",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 4
                }
            ]
        },
        {
            "name": "JewelMakeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "jewelPos",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "JewelInfo",
                    "name": "jewel",
                    "id": 4
                }
            ]
        },
        {
            "name": "JewelUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "jewelPos",
                    "id": 3
                }
            ]
        },
        {
            "name": "JewelUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "jewelPos",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "JewelInfo",
                    "name": "jewel",
                    "id": 4
                }
            ]
        },
        {
            "name": "JewelChangeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "jewelPos",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 4
                }
            ]
        },
        {
            "name": "JewelChangeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "jewelPos",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "JewelInfo",
                    "name": "jewel",
                    "id": 4
                }
            ]
        },
        {
            "name": "JewelRemoveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "jewelPos",
                    "id": 3
                }
            ]
        },
        {
            "name": "JewelRemoveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "equipPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "jewelPos",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "JewelInfo",
                    "name": "jewel",
                    "id": 4
                }
            ]
        },
        {
            "name": "JewelMakeAllReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "JewelMakeAllAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "JewelInfo",
                    "keytype": "int32",
                    "name": "jewels",
                    "id": 2
                }
            ]
        },
        {
            "name": "JuexueUpLevelReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "JuexueUpLevelAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Juexue",
                    "name": "juexue",
                    "id": 1
                }
            ]
        },
        {
            "name": "KillMonsterUniLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "KillMonsterUniLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "KillMonsterUniInfo",
                    "name": "list",
                    "id": 1
                }
            ]
        },
        {
            "name": "KillMonsterUniFirstDrawReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "KillMonsterUniFirstDrawAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "KillMonsterUniDrawReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "KillMonsterUniDrawAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "KillMonsterUniKillNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "klillUserId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "killUserName",
                    "id": 3
                }
            ]
        },
        {
            "name": "KillMonsterPerLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "KillMonsterPerLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "KillMonsterPerInfo",
                    "name": "list",
                    "id": 1
                }
            ]
        },
        {
            "name": "KillMonsterPerDrawReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "KillMonsterPerDrawAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "KillMonsterPerKillNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "KillMonsterMilLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "KillMonsterMilLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "KillMonsterMilInfo",
                    "name": "list",
                    "id": 1
                }
            ]
        },
        {
            "name": "KillMonsterMilDrawReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                }
            ]
        },
        {
            "name": "KillMonsterMilDrawAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                }
            ]
        },
        {
            "name": "KillMonsterMilKillNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "killNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "KillMonsterUniInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "klillUserId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "killUserName",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "serverFirstKill",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "serverKill",
                    "id": 5
                }
            ]
        },
        {
            "name": "KillMonsterPerInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "kill",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "receive",
                    "id": 3
                }
            ]
        },
        {
            "name": "KillMonsterMilInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "level",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "killNum",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "receive",
                    "id": 4
                }
            ]
        },
        {
            "name": "LabelUpReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "LabelUpAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "LabelTransferReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 1
                }
            ]
        },
        {
            "name": "LabelTransferAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "transfer",
                    "id": 2
                }
            ]
        },
        {
            "name": "LabelDayRewardReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "LabelDayRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "dayReward",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "LabelTaskReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "LabelTaskNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "labelId",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "LabelTaskUnit",
                    "keytype": "int32",
                    "name": "taskInfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "LabelTaskUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "taskId",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "value",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "cfgVal",
                    "id": 3,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isOver",
                    "id": 4
                }
            ]
        },
        {
            "name": "LotteryInfoReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "LotteryInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "LotteryInfo",
                    "name": "myLotteryInfos",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "LotteryInfo",
                    "name": "allLotteryInfos",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "LotteryInfo",
                    "name": "winLotteryInfos",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "PopUpState",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "goodLuckState",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "winUserInfo",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isGetAward",
                    "id": 7
                }
            ]
        },
        {
            "name": "GetGoodLuckReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "GetGoodLuckAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "state",
                    "id": 1
                }
            ]
        },
        {
            "name": "SetLotteryPopUpStateReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "state",
                    "id": 1
                }
            ]
        },
        {
            "name": "SetLotteryPopUpStateAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "state",
                    "id": 1
                }
            ]
        },
        {
            "name": "LotteryBuyNumsReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "num",
                    "id": 1
                }
            ]
        },
        {
            "name": "LotteryBuyNumsAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "LotteryInfo",
                    "name": "LotteryInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "BrocastBuyNumsNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "LotteryInfo",
                    "name": "LotteryInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "LotteryInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "userName",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "awardNumber",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "shareNum",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "combat",
                    "id": 5
                }
            ]
        },
        {
            "name": "LotteryEnd",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "LotteryInfo",
                    "name": "winLotteryInfos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "winUserInfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "LotteryInfo1Req",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "LotteryInfo1Ack",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isWin",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "itemUnit",
                    "name": "items",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "LotteryInfo",
                    "name": "winLotteryInfos",
                    "id": 3
                }
            ]
        },
        {
            "name": "LotteryGetEndAwardReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "LotteryGetEndAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "getState",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "itemUnit",
                    "name": "items",
                    "id": 2
                }
            ]
        },
        {
            "name": "MagicCircleUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "magicCircleType",
                    "id": 2
                }
            ]
        },
        {
            "name": "MagicCircleUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "magicCircleType",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "excelId",
                    "id": 3
                }
            ]
        },
        {
            "name": "MagicCircleChangeWearReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "magicCircleLvId",
                    "id": 2
                }
            ]
        },
        {
            "name": "MagicCircleChangeWearAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "magicCircleLvId",
                    "id": 2
                }
            ]
        },
        {
            "name": "MagicTowerEndNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rank",
                    "id": 1
                }
            ]
        },
        {
            "name": "MagicTowerGetUserInfoReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MagicTowerGetUserInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "score",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isGetAward",
                    "id": 2
                }
            ]
        },
        {
            "name": "MagicTowerlayerAwardReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MagicTowerlayerAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                }
            ]
        },
        {
            "name": "MailReadReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "MailReadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "MailRedeemReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "MailRedeemAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goodsChanges",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "MailNtf",
                    "name": "mail",
                    "id": 3
                }
            ]
        },
        {
            "name": "MailRedeemAllReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MailRedeemAllAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "ids",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goodsChanges",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "MailNtf",
                    "name": "mail",
                    "id": 3
                }
            ]
        },
        {
            "name": "MailNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "sender",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "title",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "Content",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "status",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "expireAt",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "createdAt",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "redeemedAt",
                    "id": 9
                },
                {
                    "rule": "repeated",
                    "type": "string",
                    "name": "args",
                    "id": 11
                },
                {
                    "rule": "repeated",
                    "type": "itemUnit",
                    "name": "items",
                    "id": 12
                }
            ]
        },
        {
            "name": "MailLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MailLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "MailNtf",
                    "name": "mails",
                    "id": 1
                }
            ]
        },
        {
            "name": "MailDeleteReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "MailDeleteAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "MailNtf",
                    "name": "mail",
                    "id": 1
                }
            ]
        },
        {
            "name": "MailDeleteAllReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MailDeleteAllAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "MailNtf",
                    "name": "mails",
                    "id": 1
                }
            ]
        },
        {
            "name": "MaterialStageLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MaterialStageLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "MaterialStage",
                    "keytype": "int32",
                    "name": "materialStage",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterMaterialStageFightReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "MaterialStageFightResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "materialType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "nowLayer",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lastLayer",
                    "id": 7
                }
            ]
        },
        {
            "name": "MaterialStageSweepReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "MaterialStageSweepAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "materialType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sweepNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                }
            ]
        },
        {
            "name": "MaterialStageBuyNumNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "materialType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "MaterialStageBuyNumReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "materialType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "use",
                    "id": 2
                }
            ]
        },
        {
            "name": "MaterialStageBuyNumAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "materialType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "MiJiUpReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "MiJiUpAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 2
                }
            ]
        },
        {
            "name": "MiningLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MiningLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "MiningInfo",
                    "name": "mining",
                    "id": 1
                }
            ]
        },
        {
            "name": "MiningUpMinerReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isMax",
                    "id": 1
                }
            ]
        },
        {
            "name": "MiningUpMinerAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "miner",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "luck",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isUp",
                    "id": 3
                }
            ]
        },
        {
            "name": "MiningBuyNumReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MiningBuyNumAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "MiningStartReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MiningStartAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "workTime",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "workNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "MiningRobReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "MiningRobAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "robNum",
                    "id": 3
                }
            ]
        },
        {
            "name": "MiningRobFightAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "MiningRobBackReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "MiningRobBackAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                }
            ]
        },
        {
            "name": "MiningRobBackFightAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "MiningRobListReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MiningRobListAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "MiningRob",
                    "keytype": "int64",
                    "name": "mineRob",
                    "id": 1
                }
            ]
        },
        {
            "name": "MiningListReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MiningListAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "MiningListInfo",
                    "keytype": "int64",
                    "name": "miningList",
                    "id": 1
                }
            ]
        },
        {
            "name": "MiningDrawLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MiningDrawLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "status",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "robId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "robName",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "robTime",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "rId",
                    "id": 5
                }
            ]
        },
        {
            "name": "MiningDrawReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MiningDrawAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                }
            ]
        },
        {
            "name": "MiningInReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MiningInAck",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "MiningInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "workTime",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "workNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "robNum",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "miner",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "luck",
                    "id": 6
                }
            ]
        },
        {
            "name": "MiningRob",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "combat",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "miner",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "robTime",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "id",
                    "id": 5
                }
            ]
        },
        {
            "name": "MiningListInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "uid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "combat",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "time",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "miner",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "id",
                    "id": 6
                }
            ]
        },
        {
            "name": "MonthCardBuyReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "MonthCardBuyAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "monthCardType",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "MonthCardUnit",
                    "name": "monthCard",
                    "id": 3
                }
            ]
        },
        {
            "name": "MonthCardDailyRewardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "monthCardType",
                    "id": 1
                }
            ]
        },
        {
            "name": "MonthCardDailyRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "monthCardType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "OfficialUpLevelReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "OfficialUpLevelAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "newLv",
                    "id": 1
                }
            ]
        },
        {
            "name": "OfflineAwardLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "OfflineAwardLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "offlineTimes",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "getExpNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "OfflineAwardGetReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "OfflineAwardGetAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isGet",
                    "id": 2
                }
            ]
        },
        {
            "name": "GetOnlineAwardInfoReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "GetOnlineAwardInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "onlineTime",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "getAwardId",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "GetOnlineAwardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "awardId",
                    "id": 1
                }
            ]
        },
        {
            "name": "GetOnlineAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "getAwardId",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "PanaceaUseReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "PanaceaUseAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "PanaceaInfo",
                    "name": "panacea",
                    "id": 2
                }
            ]
        },
        {
            "name": "PersonBossLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "PersonBossLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "PersonBoss",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterPersonBossFightReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterPersonBossFightAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "PersonBossFightResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "PersonBossSweepReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "PersonBossSweepAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isBagFull",
                    "id": 4
                }
            ]
        },
        {
            "name": "PersonBossDareNumNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "PetActiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "PetActiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "PetInfo",
                    "name": "petInfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "PetUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemNum",
                    "id": 3
                }
            ]
        },
        {
            "name": "PetUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 3
                }
            ]
        },
        {
            "name": "PetUpGradeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "PetUpGradeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "grade",
                    "id": 2
                }
            ]
        },
        {
            "name": "PetBreakReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "PetBreakAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "break",
                    "id": 2
                }
            ]
        },
        {
            "name": "PetChangeWearReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "PetChangeWearAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "petId",
                    "id": 1
                }
            ]
        },
        {
            "name": "PetAppendageReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "petId",
                    "id": 1
                }
            ]
        },
        {
            "name": "PetAppendageAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "petId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 2
                }
            ]
        },
        {
            "name": "PreviewFunctionLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "PreviewFunctionLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "haveBuyIds",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "havePointIds",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "PreviewFunctionGetReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "PreviewFunctionGetAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "haveBuyIds",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "havePointIds",
                    "id": 3,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "PreviewFunctionPointReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "PreviewFunctionPointAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "havePointIds",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "PrivilegeBuyReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "privilegeId",
                    "id": 1
                }
            ]
        },
        {
            "name": "PrivilegeBuyAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "privilegeId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "RankLoadReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                }
            ]
        },
        {
            "name": "RankLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "RankInfo",
                    "name": "ranks",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "self",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 3
                }
            ]
        },
        {
            "name": "RankWorshipReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "RankWorshipAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                }
            ]
        },
        {
            "name": "RechargFulfilledNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ingot",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "payMoney",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "vip",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "vipExp",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rechargedAll",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rechargeId",
                    "id": 6
                }
            ]
        },
        {
            "name": "RechargeApplyPayReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "payNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rechargeId",
                    "id": 2
                }
            ]
        },
        {
            "name": "RechargeApplyPayAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "result",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "payData",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rechargeId",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isPayToken",
                    "id": 4
                }
            ]
        },
        {
            "name": "MoneyPayReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "payType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "payNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "typeId",
                    "id": 3
                }
            ]
        },
        {
            "name": "MoneyPayAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "result",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "payData",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "payType",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "payNum",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "typeId",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isPayToken",
                    "id": 6
                }
            ]
        },
        {
            "name": "RechargeResetNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "recharge",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "ContRechargeCycleNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cycle",
                    "id": 1
                }
            ]
        },
        {
            "name": "ContRechargeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "recharge",
                    "id": 1
                }
            ]
        },
        {
            "name": "ContRechargeReceiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "contRechargeId",
                    "id": 1
                }
            ]
        },
        {
            "name": "ContRechargeReceiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "contRechargeId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "RechargeAllGetReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "RechargeAllGetAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "rechargeGetGetIds",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rechargeAll",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                }
            ]
        },
        {
            "name": "ReinActiveReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ReinActiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Rein",
                    "name": "rein",
                    "id": 1
                }
            ]
        },
        {
            "name": "ReincarnationReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ReincarnationAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Rein",
                    "name": "rein",
                    "id": 1
                }
            ]
        },
        {
            "name": "ReinCostBuyReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "num",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "use",
                    "id": 3
                }
            ]
        },
        {
            "name": "ReinCostBuyAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Rein",
                    "name": "rein",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "ReinCost",
                    "name": "reinCost",
                    "id": 2
                }
            ]
        },
        {
            "name": "ReinCostUseReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "ReinCostUseAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Rein",
                    "name": "rein",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "ReinCost",
                    "name": "reinCost",
                    "id": 2
                }
            ]
        },
        {
            "name": "ReinCostBuyNumRefNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "ReinCost",
                    "name": "reinCost",
                    "id": 1
                }
            ]
        },
        {
            "name": "RingWearReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ringPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagPos",
                    "id": 3
                }
            ]
        },
        {
            "name": "RingWearAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ringPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "Ring",
                    "name": "ring",
                    "id": 3
                }
            ]
        },
        {
            "name": "RingRemoveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ringPos",
                    "id": 2
                }
            ]
        },
        {
            "name": "RingRemoveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ringPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "Ring",
                    "name": "ring",
                    "id": 3
                }
            ]
        },
        {
            "name": "RingStrengthenReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ringPos",
                    "id": 2
                }
            ]
        },
        {
            "name": "RingStrengthenAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ringPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "Ring",
                    "name": "ring",
                    "id": 3
                }
            ]
        },
        {
            "name": "RingPhantomReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ringPos",
                    "id": 2
                }
            ]
        },
        {
            "name": "RingPhantomAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ringPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "Ring",
                    "name": "ring",
                    "id": 3
                }
            ]
        },
        {
            "name": "RingSkillUpReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ringPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "phantomPos",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 4
                }
            ]
        },
        {
            "name": "RingSkillUpAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ringPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "Ring",
                    "name": "ring",
                    "id": 3
                }
            ]
        },
        {
            "name": "RingFuseReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagPos1",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagPos2",
                    "id": 3
                }
            ]
        },
        {
            "name": "RingFuseAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "RingSkillResetReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ringPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "phantomPos",
                    "id": 3
                }
            ]
        },
        {
            "name": "RingSkillResetAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ringPos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "Ring",
                    "name": "ring",
                    "id": 3
                }
            ]
        },
        {
            "name": "Point",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "x",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "y",
                    "id": 2
                }
            ]
        },
        {
            "name": "SceneObj",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "Point",
                    "name": "point",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dir",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "teamId",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "hp",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "hpMax",
                    "id": 8
                },
                {
                    "rule": "repeated",
                    "type": "BuffInfo",
                    "name": "buffs",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "mp",
                    "id": 10
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "mpMax",
                    "id": 11
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverId",
                    "id": 12
                },
                {
                    "rule": "optional",
                    "type": "SceneUser",
                    "name": "user",
                    "id": 21
                },
                {
                    "rule": "optional",
                    "type": "SceneMonster",
                    "name": "monster",
                    "id": 23
                },
                {
                    "rule": "optional",
                    "type": "SceneItem",
                    "name": "item",
                    "id": 24
                },
                {
                    "rule": "optional",
                    "type": "ScenePet",
                    "name": "pet",
                    "id": 25
                },
                {
                    "rule": "optional",
                    "type": "SceneCollection",
                    "name": "collection",
                    "id": 26
                },
                {
                    "rule": "optional",
                    "type": "SceneFit",
                    "name": "fit",
                    "id": 27
                },
                {
                    "rule": "optional",
                    "type": "SceneSummon",
                    "name": "summon",
                    "id": 28
                },
                {
                    "rule": "optional",
                    "type": "SceneBuff",
                    "name": "buff",
                    "id": 29
                }
            ]
        },
        {
            "name": "SceneUser",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "Display",
                    "name": "display",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "vip",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lvl",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sex",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "combat",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "avatar",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 10
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildId",
                    "id": 12
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "guildName",
                    "id": 13
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "elfLv",
                    "id": 14
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "username",
                    "id": 15
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "usersex",
                    "id": 16
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userjob",
                    "id": 17
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "userHpTotal",
                    "id": 18
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "toHelpUserId",
                    "id": 19
                }
            ]
        },
        {
            "name": "SceneMonster",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "idx",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ownerUseId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "ownerUserName",
                    "id": 3
                }
            ]
        },
        {
            "name": "ScenePet",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "idx",
                    "id": 2
                }
            ]
        },
        {
            "name": "SceneCollection",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "collectionObjId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverTime",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "endTime",
                    "id": 4
                }
            ]
        },
        {
            "name": "SceneFit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fashionId",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fashionLv",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fitLv",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildId",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "guildName",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "leaderJob",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "leaderSex",
                    "id": 10
                }
            ]
        },
        {
            "name": "SceneSummon",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "summonId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 3
                }
            ]
        },
        {
            "name": "SceneItem",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemNum",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "owner",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ownerProtectedTime",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "disappearTime",
                    "id": 6
                }
            ]
        },
        {
            "name": "SceneBuff",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buffId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 2
                }
            ]
        },
        {
            "name": "SceneEnterNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "SceneObj",
                    "name": "objs",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "enterType",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isTower",
                    "id": 4
                }
            ]
        },
        {
            "name": "SceneEnterOverNtf",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "SceneLeaveNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "objIds",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "leaveType",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isTower",
                    "id": 4
                }
            ]
        },
        {
            "name": "SceneDieNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "SceneObj",
                    "name": "dropItems",
                    "id": 2
                }
            ]
        },
        {
            "name": "SceneMoveRpt",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "Point",
                    "name": "point",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "moveType",
                    "id": 3
                }
            ]
        },
        {
            "name": "SceneMoveNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "Point",
                    "name": "point",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "force",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "moveType",
                    "id": 4
                }
            ]
        },
        {
            "name": "SceneUserReliveNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "SceneObj",
                    "name": "obj",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "reliveType",
                    "id": 2
                }
            ]
        },
        {
            "name": "SceneUserUpdateNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "SceneUser",
                    "name": "objUser",
                    "id": 2
                }
            ]
        },
        {
            "name": "SceneUserElfUpdateNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "elfLv",
                    "id": 2
                }
            ]
        },
        {
            "name": "AttackRpt",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "objIds",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dir",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "Point",
                    "name": "point",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillLevel",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isElf",
                    "id": 7
                }
            ]
        },
        {
            "name": "AttackEffectNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "attackerId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dir",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "Point",
                    "name": "point",
                    "id": 4
                },
                {
                    "rule": "repeated",
                    "type": "HurtEffect",
                    "name": "hurts",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "Point",
                    "name": "MoveToPoint",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillLv",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "err",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "skillStartT",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "skillStopT",
                    "id": 10
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "serverTime",
                    "id": 11
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "mpNow",
                    "id": 12
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isElf",
                    "id": 13
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "attackerName",
                    "id": 14
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "hpNow",
                    "id": 15
                }
            ]
        },
        {
            "name": "HurtEffect",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "hp",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "changHp",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isDeath",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isDodge",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "hurtType",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "hurt",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "deathblow",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "cutHurt",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "Point",
                    "name": "MoveToPoint",
                    "id": 10
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "unBlock",
                    "id": 11
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "reflex",
                    "id": 12
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "reliveSelf",
                    "id": 13
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isWuDi",
                    "id": 14
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "killHurt",
                    "id": 15
                }
            ]
        },
        {
            "name": "SceneObjHpNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "hp",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "changeHp",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "totalHp",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "killerId",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "killerName",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "userHpTotal",
                    "id": 8
                }
            ]
        },
        {
            "name": "SceneObjMpNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "Mp",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "changeMp",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "totalMp",
                    "id": 4
                }
            ]
        },
        {
            "name": "FightHurtRankReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FightHurtRankAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "FightRankUnit",
                    "name": "ranks",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "FightRankUnit",
                    "name": "myRank",
                    "id": 2
                }
            ]
        },
        {
            "name": "FightRankUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rank",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "score",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "serverName",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 5
                }
            ]
        },
        {
            "name": "GetBossOwnerChangReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "BossOwnerChangNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ownerObjId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "userName",
                    "id": 4
                }
            ]
        },
        {
            "name": "BuffChangeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "BuffInfo",
                    "name": "buff",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "DelBuffInfo",
                    "name": "delBuffInfos",
                    "id": 2
                }
            ]
        },
        {
            "name": "BuffDelNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "DelBuffInfo",
                    "name": "delBuffInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "BuffInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ownerObjId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sourceObjId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "idx",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buffId",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "totalTime",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ownerUserId",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sourceUserId",
                    "id": 7
                }
            ]
        },
        {
            "name": "DelBuffInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ownerObjId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "idx",
                    "id": 2
                }
            ]
        },
        {
            "name": "BuffPropChangeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "propId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "total",
                    "id": 3
                }
            ]
        },
        {
            "name": "BuffHpChangeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "BuffHpChangeInfo",
                    "name": "buffHpChangeInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "BuffHpChangeInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ownerObjId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "idx",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "death",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "changeHp",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "totalHp",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "killerId",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "killerName",
                    "id": 7
                }
            ]
        },
        {
            "name": "MainCityEnterRpt",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "SceneObj",
                    "name": "obj",
                    "id": 1
                }
            ]
        },
        {
            "name": "MainCityMoveRpt",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "x",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "y",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dir",
                    "id": 5
                }
            ]
        },
        {
            "name": "MainCityLeaveRpt",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                }
            ]
        },
        {
            "name": "MainCityUpdateRpt",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "SceneObj",
                    "name": "obj",
                    "id": 1
                }
            ]
        },
        {
            "name": "FightEnterOkReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FightStartCountDownNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverTime",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "countDownTime",
                    "id": 2
                }
            ]
        },
        {
            "name": "FightStartCountDownOkReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "FightStartNtf",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "CollectionStatusChangeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userObjId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "startTime",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 4
                }
            ]
        },
        {
            "name": "FightTeamChangeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "userTeamIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "FightUserChangeToHelperNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "teamIndex",
                    "id": 2
                }
            ]
        },
        {
            "name": "FightNpcEventReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "npcId",
                    "id": 1
                }
            ]
        },
        {
            "name": "ExpStageKillInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "killMonsterNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "getExp",
                    "id": 2
                }
            ]
        },
        {
            "name": "PaodianTopUserReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "PaodianTopUserNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "userIds",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "PaoDianUserNumNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "userNums",
                    "id": 1
                }
            ]
        },
        {
            "name": "PaodianFightEnd",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "GetShabakeScoresReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ShabakeScoreRankNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "ShabakeUserScore",
                    "name": "userScores",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "ShabakeGuildScore",
                    "name": "guildScores",
                    "id": 2
                }
            ]
        },
        {
            "name": "ShabakeUserScore",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "userName",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "score",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rank",
                    "id": 4
                }
            ]
        },
        {
            "name": "ShabakeGuildScore",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "guildName",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "score",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverId",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "serverName",
                    "id": 5
                }
            ]
        },
        {
            "name": "ShabakeOccupiedNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isOccupy",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "userName",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildId",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "guildName",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "startTime",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "endTime",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fightStatus",
                    "id": 8
                }
            ]
        },
        {
            "name": "GetShabakeCrossScoresReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ShabakeCrossScoreRankNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "ShabakeCrossServerScore",
                    "name": "serverScores",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "ShabakeGuildScore",
                    "name": "guildScores",
                    "id": 2
                }
            ]
        },
        {
            "name": "ShabakeCrossServerScore",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "serverName",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "score",
                    "id": 3
                }
            ]
        },
        {
            "name": "ShabakeCrossOccupiedNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isOccupy",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "userName",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildId",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "guildName",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "startTime",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "endTime",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverId",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "serverName",
                    "id": 9
                }
            ]
        },
        {
            "name": "GuardPillarFightNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "FightHurtRankAck",
                    "name": "rank",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "wave",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "nextTime",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "monsterTotal",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "monsterless",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fightEndTime",
                    "id": 6
                }
            ]
        },
        {
            "name": "MagicTowerBossInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "bossName",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "layer",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "status",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "refreshTime",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "monsterId",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ownerUseId",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "ownerName",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "monsterObjId",
                    "id": 8
                }
            ]
        },
        {
            "name": "MagicTowerFightNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "ShabakeUserScore",
                    "name": "userScores",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "MagicTowerBossInfo",
                    "name": "bossInfos",
                    "id": 2
                }
            ]
        },
        {
            "name": "FightUserScoreNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "score",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "changeScore",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rankScore",
                    "id": 3
                }
            ]
        },
        {
            "name": "GetFightBossInfosReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "GetFightBossInfosAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "FightBossInfoUnit",
                    "name": "bossInfos",
                    "id": 1
                }
            ]
        },
        {
            "name": "FightBossInfoUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "monsterId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "Point",
                    "name": "point",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "reliveCD",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "hp",
                    "id": 5
                }
            ]
        },
        {
            "name": "BossReliveNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "objId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "reliveCd",
                    "id": 2
                }
            ]
        },
        {
            "name": "SevenInvestmentLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "SevenInvestmentLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "haveGetIds",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "activateDay",
                    "id": 2
                }
            ]
        },
        {
            "name": "GetSevenInvestmentAwardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "GetSevenInvestmentAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "haveGetIds",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "ShaBaKeInfoCrossReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ShaBaKeInfoCrossAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "Info",
                    "name": "WinGuildUserInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "firstGuildName",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isEnd",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverId",
                    "id": 4
                }
            ]
        },
        {
            "name": "Info",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "nickName",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sex",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "position",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "Display",
                    "name": "display",
                    "id": 5
                }
            ]
        },
        {
            "name": "EnterCrossShaBaKeFightReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "EnterCrossShaBaKeFightAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "state",
                    "id": 1
                }
            ]
        },
        {
            "name": "CrossShaBaKeFightEndNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "ShabakeRankScore",
                    "name": "serverRank",
                    "id": 1
                }
            ]
        },
        {
            "name": "CrossShabakeOpenNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isOpen",
                    "id": 1
                }
            ]
        },
        {
            "name": "ShabakeRankScore",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "score",
                    "id": 2
                }
            ]
        },
        {
            "name": "ShaBaKeInfoReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ShaBaKeInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "Info",
                    "name": "WinGuildUserInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "firstGuildName",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isEnd",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "winGuildServerId",
                    "id": 4
                }
            ]
        },
        {
            "name": "EnterShaBaKeFightReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ShaBaKeFightResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rank",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userRank",
                    "id": 2
                }
            ]
        },
        {
            "name": "ShabakeIsOpenNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isOpen",
                    "id": 1
                }
            ]
        },
        {
            "name": "ShopListReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "shopType",
                    "id": 1
                }
            ]
        },
        {
            "name": "ShopListAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "shopType",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "ShopInfo",
                    "name": "shopList",
                    "id": 2
                }
            ]
        },
        {
            "name": "ShopBuyReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "ShopBuyAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                }
            ]
        },
        {
            "name": "ShopWeekResetNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "ShopInfo",
                    "name": "shopInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "SignReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "SignAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "SignInfo",
                    "name": "signInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "SignRepairReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "repairDay",
                    "id": 1
                }
            ]
        },
        {
            "name": "SignRepairAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "SignInfo",
                    "name": "signInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "CumulativeSignReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cumulativeDay",
                    "id": 1
                }
            ]
        },
        {
            "name": "CumulativeSignAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "SignInfo",
                    "name": "signInfo",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "SignResetNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "SignInfo",
                    "name": "signInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "SkillUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 3
                }
            ]
        },
        {
            "name": "SkillUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillType",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "SkillUnit",
                    "name": "skill",
                    "id": 3
                }
            ]
        },
        {
            "name": "SkillChangePosReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 3
                }
            ]
        },
        {
            "name": "SkillChangePosAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillType",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "skillBags",
                    "id": 3
                }
            ]
        },
        {
            "name": "SkillChangeWearReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 2
                }
            ]
        },
        {
            "name": "SkillChangeWearAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillType",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "skillBags",
                    "id": 3
                }
            ]
        },
        {
            "name": "SkillResetReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillType",
                    "id": 2
                }
            ]
        },
        {
            "name": "SkillResetAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillType",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "skillBags",
                    "id": 3
                },
                {
                    "rule": "repeated",
                    "type": "SkillUnit",
                    "name": "skills",
                    "id": 4
                }
            ]
        },
        {
            "name": "SkillUseReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 2
                }
            ]
        },
        {
            "name": "SkillUseNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "skillId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "startTime",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 4
                }
            ]
        },
        {
            "name": "ClearSkillCdNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "SpecialEquipChangeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bagPos",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 4
                }
            ]
        },
        {
            "name": "SpecialEquipChangeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "SpecialEquipUnit",
                    "name": "specialEquip",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 5
                }
            ]
        },
        {
            "name": "SpecialEquipRemoveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 3
                }
            ]
        },
        {
            "name": "SpecialEquipRemoveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "pos",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "SpendRebatesRewardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "SpendRebatesRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "SpendRebatesNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "countIngot",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ingot",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "cycle",
                    "id": 3
                }
            ]
        },
        {
            "name": "StageFightStartReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "wave",
                    "id": 2
                }
            ]
        },
        {
            "name": "StageFightStartAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "wave",
                    "id": 2
                }
            ]
        },
        {
            "name": "StageFightEndReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "wave",
                    "id": 2
                }
            ]
        },
        {
            "name": "StageFightEndNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "wave",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "onlyUpdate",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 5
                }
            ]
        },
        {
            "name": "LeaveFightReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "LeaveFightAck",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "KillMonsterReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "monsterId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "killNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "KillMonsterAck",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "StartStageBossFightReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "StageBagChangeNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "hookupTime",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "itemUnit",
                    "name": "items",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isOnline",
                    "id": 3
                }
            ]
        },
        {
            "name": "StageGetHookMapRewardReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "StageGetHookMapRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "hookupTime",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "itemUnit",
                    "name": "items",
                    "id": 3
                }
            ]
        },
        {
            "name": "PingReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isActive",
                    "id": 1
                }
            ]
        },
        {
            "name": "PingAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "ts",
                    "id": 1
                }
            ]
        },
        {
            "name": "PreferenceSetReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "Preference",
                    "name": "preference",
                    "id": 1
                }
            ]
        },
        {
            "name": "PreferenceSetAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "Preference",
                    "name": "preference",
                    "id": 1
                }
            ]
        },
        {
            "name": "PreferenceLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "PreferenceLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "Preference",
                    "name": "preference",
                    "id": 1
                }
            ]
        },
        {
            "name": "FuncStateCloseNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "closeFuncId",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "TalentUpLvReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isMax",
                    "id": 3
                }
            ]
        },
        {
            "name": "TalentUpLvAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "TalentInfo",
                    "name": "talentInfo",
                    "id": 3
                }
            ]
        },
        {
            "name": "TalentResetReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 2
                }
            ]
        },
        {
            "name": "TalentResetAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "TalentInfo",
                    "name": "talentInfo",
                    "id": 3
                }
            ]
        },
        {
            "name": "TaskDoneReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "taskId",
                    "id": 1
                }
            ]
        },
        {
            "name": "TaskDoneAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 1
                }
            ]
        },
        {
            "name": "TaskNpcStateReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "TaskNpcStateAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "taskId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "process",
                    "id": 2
                }
            ]
        },
        {
            "name": "SetTaskInfoReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "taskId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "process",
                    "id": 2
                }
            ]
        },
        {
            "name": "SetTaskInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "taskId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "process",
                    "id": 2
                }
            ]
        },
        {
            "name": "TitleActiveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "titleId",
                    "id": 1
                }
            ]
        },
        {
            "name": "TitleActiveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "Title",
                    "name": "title",
                    "id": 1
                }
            ]
        },
        {
            "name": "TitleWearReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "titleId",
                    "id": 2
                }
            ]
        },
        {
            "name": "TitleWearAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "titleId",
                    "id": 2
                }
            ]
        },
        {
            "name": "TitleRemoveReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "TitleRemoveAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "titleId",
                    "id": 2
                }
            ]
        },
        {
            "name": "TitleLookReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "titleId",
                    "id": 1
                }
            ]
        },
        {
            "name": "TitleLookAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "titleId",
                    "id": 1
                }
            ]
        },
        {
            "name": "TitleAutoActiveNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "Title",
                    "name": "titleList",
                    "id": 1
                }
            ]
        },
        {
            "name": "TitleExpireNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "Title",
                    "name": "titleList",
                    "id": 1
                }
            ]
        },
        {
            "name": "OpenTowerReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "OpenTowerAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "towerLv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "dayAward",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lotteryNum",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lotterId",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "rankReward",
                    "id": 5
                }
            ]
        },
        {
            "name": "ToweryDayAwardReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "ToweryDayAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "dayAward",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "TowerLotteryReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "TowerLotteryAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lotteryNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lotteryId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "EnterTowerFightReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "TowerFightResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "towerLv",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                }
            ]
        },
        {
            "name": "TowerFightContinueReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "TowerSweepReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "TowerSweepAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "towerLv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "TowerRankRewardReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "TowerRankRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "rankReward",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "TowerLvNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "towerLv",
                    "id": 1
                }
            ]
        },
        {
            "name": "SetTreasurePopUpStateReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "state",
                    "id": 1
                }
            ]
        },
        {
            "name": "SetTreasurePopUpStateAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "state",
                    "id": 1
                }
            ]
        },
        {
            "name": "ChooseTreasureAwardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "index",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isReplace",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "replaceIndex",
                    "id": 4
                }
            ]
        },
        {
            "name": "ChooseTreasureAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "index",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "itemId",
                    "id": 3,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "map",
                    "type": "ChooseInfo",
                    "keytype": "int32",
                    "name": "choosItems",
                    "id": 4
                },
                {
                    "rule": "map",
                    "type": "ChooseInfo",
                    "keytype": "int32",
                    "name": "haveGetItems",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isReplace",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "replaceIndex",
                    "id": 7
                }
            ]
        },
        {
            "name": "ChooseInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "items",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "BuyTreasureItemReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "BuyTreasureItemAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "haveBuyTimes",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "TreasureApplyGetReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "TreasureApplyGetAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "items",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureTimes",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                },
                {
                    "rule": "map",
                    "type": "ChooseInfo",
                    "keytype": "int32",
                    "name": "choosItems",
                    "id": 4
                },
                {
                    "rule": "map",
                    "type": "ChooseInfo",
                    "keytype": "int32",
                    "name": "haveGetItems",
                    "id": 5
                },
                {
                    "rule": "repeated",
                    "type": "TreasureInfoUnit",
                    "name": "myTreasureInfo",
                    "id": 6
                },
                {
                    "rule": "repeated",
                    "type": "TreasureInfoUnit",
                    "name": "serverTreasureInfo",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "randomType",
                    "id": 8
                }
            ]
        },
        {
            "name": "TreasureInfosReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "TreasureInfosAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureTimes",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "TreasureInfoUnit",
                    "name": "myTreasureInfo",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "TreasureInfoUnit",
                    "name": "serverTreasureInfo",
                    "id": 3
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "haveGetRoundId",
                    "id": 4,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "HaveBuyTimes",
                    "id": 5
                },
                {
                    "rule": "map",
                    "type": "ChooseInfo",
                    "keytype": "int32",
                    "name": "choosItems",
                    "id": 6
                },
                {
                    "rule": "map",
                    "type": "ChooseInfo",
                    "keytype": "int32",
                    "name": "haveGetItems",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "season",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "popUpState",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "mergeMark",
                    "id": 10
                }
            ]
        },
        {
            "name": "TreasureInfoUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "count",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "time",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "userName",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "type",
                    "id": 5
                }
            ]
        },
        {
            "name": "TreasureInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "itemId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "time",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "userName",
                    "id": 3
                }
            ]
        },
        {
            "name": "GetTreasureIntegralAwardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "GetTreasureIntegralAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "haveGetIndex",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "treasureTimes",
                    "id": 2
                }
            ]
        },
        {
            "name": "TreasureDrawInfoReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "TreasureDrawInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "TreasureInfoUnit",
                    "name": "myTreasureInfo",
                    "id": 1
                },
                {
                    "rule": "repeated",
                    "type": "TreasureInfoUnit",
                    "name": "serverTreasureInfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "TreasureCloseNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isClose",
                    "id": 1
                }
            ]
        },
        {
            "name": "TreasureShopLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "TreasureShopLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "refreshFree",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "refreshTime",
                    "id": 3
                },
                {
                    "rule": "map",
                    "type": "bool",
                    "keytype": "int32",
                    "name": "shop",
                    "id": 4
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "car",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 6
                }
            ]
        },
        {
            "name": "TreasureShopCarChangeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "shopId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isAdd",
                    "id": 2
                }
            ]
        },
        {
            "name": "TreasureShopCarChangeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "car",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "shopId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isAdd",
                    "id": 3
                },
                {
                    "rule": "map",
                    "type": "bool",
                    "keytype": "int32",
                    "name": "shop",
                    "id": 4
                }
            ]
        },
        {
            "name": "TreasureShopBuyReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "shop",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "TreasureShopBuyAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "buyNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "bool",
                    "keytype": "int32",
                    "name": "shop",
                    "id": 4
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "car",
                    "id": 5
                }
            ]
        },
        {
            "name": "TreasureShopRefreshReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "TreasureShopRefreshNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "bool",
                    "keytype": "int32",
                    "name": "shop",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "refreshTime",
                    "id": 2
                },
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "car",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "refreshFree",
                    "id": 4
                }
            ]
        },
        {
            "name": "TrialTaskInfoReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "TrialTaskInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "TrialTaskInfo",
                    "keytype": "int32",
                    "name": "trialTaskInfos",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "endTime",
                    "id": 2
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "haveGetStageId",
                    "id": 3,
                    "options": {
                        "packed": true
                    }
                }
            ]
        },
        {
            "name": "TrialTaskInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "nowNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isGet",
                    "id": 2
                }
            ]
        },
        {
            "name": "TrialTaskGetAwardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "TrialTaskGetAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isGet",
                    "id": 2
                }
            ]
        },
        {
            "name": "TrialTaskGetStageAwardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                }
            ]
        },
        {
            "name": "TrialTaskGetStageAwardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isGet",
                    "id": 2
                }
            ]
        },
        {
            "name": "TrialTaskInfoNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "num",
                    "id": 2
                }
            ]
        },
        {
            "name": "EnterGameReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "openId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "loginKey",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "channel",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverId",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "origin",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "ip",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "deviceId",
                    "id": 7
                }
            ]
        },
        {
            "name": "EnterGameAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "UserLoginInfo",
                    "name": "user",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "ts",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "version",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "openServerDay",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "openServerTime",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "realServerId",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "realServerName",
                    "id": 7
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "closeFuncIds",
                    "id": 8,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "mergeOpenServerDay",
                    "id": 9
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "mergeOpenServerTime",
                    "id": 10
                },
                {
                    "rule": "map",
                    "type": "BriefServerInfo",
                    "keytype": "int32",
                    "name": "crossBriefServerInfo",
                    "id": 11
                }
            ]
        },
        {
            "name": "CreateUserReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "nickName",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "avatar",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 4
                }
            ]
        },
        {
            "name": "CreateUserAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "UserLoginInfo",
                    "name": "user",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "failReason",
                    "id": 2
                }
            ]
        },
        {
            "name": "RandNameReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sex",
                    "id": 1
                }
            ]
        },
        {
            "name": "RandNameAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "string",
                    "name": "names",
                    "id": 1
                }
            ]
        },
        {
            "name": "CreateHeroReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "sex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "job",
                    "id": 2
                }
            ]
        },
        {
            "name": "CreateHeroAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "HeroInfo",
                    "name": "hero",
                    "id": 1
                }
            ]
        },
        {
            "name": "KickUserNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "reason",
                    "id": 1
                }
            ]
        },
        {
            "name": "UserPropertyNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "HeroProp",
                    "keytype": "int32",
                    "name": "heroProps",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "userCombat",
                    "id": 2
                }
            ]
        },
        {
            "name": "DebugAddGoodsReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "id",
                    "id": 1,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "repeated",
                    "type": "int32",
                    "name": "count",
                    "id": 2,
                    "options": {
                        "packed": true
                    }
                },
                {
                    "rule": "repeated",
                    "type": "string",
                    "name": "args",
                    "id": 3
                }
            ]
        },
        {
            "name": "DebugAddGoodsAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 1
                }
            ]
        },
        {
            "name": "ChangeFightModelReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fightModel",
                    "id": 1
                }
            ]
        },
        {
            "name": "ChangeFightModelAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "fightModel",
                    "id": 1
                }
            ]
        },
        {
            "name": "ChangeHeroNameReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 2
                }
            ]
        },
        {
            "name": "ChangeHeroNameAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "HeroInfo",
                    "name": "heroInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "UserRechargeNumNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rechargeNum",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "redPacketNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dailyRecharge",
                    "id": 3
                }
            ]
        },
        {
            "name": "UserRedPacketGetNumNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "redPacketGetNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "UserOnlineNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "onlineTime",
                    "id": 2
                }
            ]
        },
        {
            "name": "UserOffLineNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "userId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "offLintTime",
                    "id": 2
                }
            ]
        },
        {
            "name": "VipCustomerReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "VipCustomerAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "flag",
                    "id": 1
                }
            ]
        },
        {
            "name": "UserInGameOkReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "CrossFightOpenNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "UserSubscribeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "subscribeId",
                    "id": 1
                }
            ]
        },
        {
            "name": "UserSubscribeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "subscribeId",
                    "id": 1
                }
            ]
        },
        {
            "name": "VipGiftGetReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 1
                }
            ]
        },
        {
            "name": "VipGiftGetAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 1
                }
            ]
        },
        {
            "name": "VipBossLoadReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "VipBossLoadAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "map",
                    "type": "int32",
                    "keytype": "int32",
                    "name": "vipBoss",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterVipBossFightReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "EnterVipBossFightAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "VipBossFightResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "VipBossSweepReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "VipBossSweepAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "isBagFull",
                    "id": 4
                }
            ]
        },
        {
            "name": "VipBossDareNumNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "dareNum",
                    "id": 2
                }
            ]
        },
        {
            "name": "WarOrderTaskNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "WarOrderTask",
                    "name": "task",
                    "id": 1
                },
                {
                    "rule": "map",
                    "type": "WarOrderTask",
                    "keytype": "int32",
                    "name": "weekTask",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 4
                }
            ]
        },
        {
            "name": "WarOrderOpenReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "WarOrderOpenAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "WarOrderTaskNtf",
                    "name": "warOrderInfo",
                    "id": 1
                }
            ]
        },
        {
            "name": "WarOrderTaskFinishReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "taskId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "week",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isWeekTask",
                    "id": 3
                }
            ]
        },
        {
            "name": "WarOrderTaskFinishAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "taskId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "week",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isWeekTask",
                    "id": 3
                }
            ]
        },
        {
            "name": "WarOrderTaskRewardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "taskId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "week",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isWeekTask",
                    "id": 3
                }
            ]
        },
        {
            "name": "WarOrderTaskRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "taskId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "week",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isWeekTask",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 5
                }
            ]
        },
        {
            "name": "WarOrderBuyLuxuryReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "WarOrderBuyLuxuryAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "isLuxury",
                    "id": 1
                }
            ]
        },
        {
            "name": "WarOrderBuyExpReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "WarOrderBuyExpAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 2
                }
            ]
        },
        {
            "name": "WarOrderLvRewardReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 1
                }
            ]
        },
        {
            "name": "WarOrderLvRewardAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 2
                }
            ]
        },
        {
            "name": "WarOrderExchangeReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exchangeId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "num",
                    "id": 2
                }
            ]
        },
        {
            "name": "WarOrderExchangeAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exchangeId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "num",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 4
                }
            ]
        },
        {
            "name": "WarOrderLvNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "lv",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "exp",
                    "id": 2
                }
            ]
        },
        {
            "name": "WarOrderResetNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "WarOrder",
                    "name": "warOrder",
                    "id": 1
                }
            ]
        },
        {
            "name": "WingUpLevelReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "WingUpLevelAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "Wing",
                    "name": "wing",
                    "id": 2
                }
            ]
        },
        {
            "name": "WingUseMaterialReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                }
            ]
        },
        {
            "name": "WingUseMaterialAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "Wing",
                    "name": "wing",
                    "id": 2
                }
            ]
        },
        {
            "name": "WingSpecialUpReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "specialType",
                    "id": 2
                }
            ]
        },
        {
            "name": "WingSpecialUpAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "WingSpecialNtf",
                    "name": "wingSpecial",
                    "id": 2
                }
            ]
        },
        {
            "name": "WingWearReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "wingId",
                    "id": 2
                }
            ]
        },
        {
            "name": "WingWearAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "heroIndex",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "wingId",
                    "id": 2
                }
            ]
        },
        {
            "name": "EnterWorldBossFightReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "WorldBossFightResultNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rank",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "GoodsChangeNtf",
                    "name": "goods",
                    "id": 3
                }
            ]
        },
        {
            "name": "LoadWorldLeaderReq",
            "syntax": "proto3",
            "fields": []
        },
        {
            "name": "LoadWorldLeaderAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "nowStageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildJoinNum",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bossHp",
                    "id": 3
                },
                {
                    "rule": "map",
                    "type": "WorldLeaderInfo",
                    "keytype": "int32",
                    "name": "worldLeaderInfoByStage",
                    "id": 4
                }
            ]
        },
        {
            "name": "WorldLeaderInfo",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildJoinNum",
                    "id": 1
                }
            ]
        },
        {
            "name": "GetWorldLeaderRankInfoReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "GetWorldLeaderRankInfoAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "repeated",
                    "type": "WorldLeaderRankUnit",
                    "name": "ranks",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bossHp",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 3
                }
            ]
        },
        {
            "name": "WorldLeaderStartNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "WorldLeaderEnterReq",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                }
            ]
        },
        {
            "name": "WorldLeaderEnterAck",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "bool",
                    "name": "enterState",
                    "id": 1
                }
            ]
        },
        {
            "name": "WorldLeaderEndRewardNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "BriefUserInfo",
                    "name": "owner",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rank",
                    "id": 3
                }
            ]
        },
        {
            "name": "WorldLeaderRankUnit",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "rank",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "guildId",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "guildName",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "int64",
                    "name": "score",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "serverId",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "serverName",
                    "id": 6
                }
            ]
        },
        {
            "name": "WorldLeaderBossHpNtf",
            "syntax": "proto3",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "stageId",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "bossHp",
                    "id": 2
                }
            ]
        }
    ],
    "isNamespace": true
}).build().pb;

(function() {
    var idMap = {
        "1": "ErrorAck",
        "2": "PingReq",
        "3": "PingAck",
        "10": "TopDataChangeNtf",
        "11": "BagDataChangeNtf",
        "12": "BagEquipDataChangeNtf",
        "98": "DebugAddGoodsReq",
        "99": "DebugAddGoodsAck",
        "100": "EnterGameReq",
        "101": "EnterGameAck",
        "102": "CreateUserReq",
        "103": "CreateUserAck",
        "104": "KickUserNtf",
        "105": "UserPropertyNtf",
        "106": "CreateHeroReq",
        "107": "CreateHeroAck",
        "108": "PreferenceSetReq",
        "109": "PreferenceSetAck",
        "110": "PreferenceLoadReq",
        "111": "PreferenceLoadAck",
        "112": "RandNameReq",
        "113": "RandNameAck",
        "114": "DisplayNtf",
        "115": "ChangeFightModelReq",
        "116": "ChangeFightModelAck",
        "117": "ResetNtf",
        "118": "CrossFightOpenNtf",
        "119": "ChangeHeroNameReq",
        "120": "ChangeHeroNameAck",
        "121": "UserRechargeNumNtf",
        "122": "RechargFulfilledNtf",
        "123": "RechargeApplyPayReq",
        "124": "RechargeApplyPayAck",
        "125": "UserRedPacketGetNumNtf",
        "126": "UserOnlineNtf",
        "127": "UserOffLineNtf",
        "128": "MoneyPayReq",
        "129": "MoneyPayAck",
        "130": "FuncStateCloseNtf",
        "131": "UserInGameOkReq",
        "132": "UserSubscribeReq",
        "133": "UserSubscribeAck",
        "200": "BagInfoReq",
        "201": "BagInfoNtf",
        "202": "BagSpaceAddReq",
        "203": "BagSpaceAddAck",
        "204": "BagSortReq",
        "205": "BagSortAck",
        "206": "GiftUseReq",
        "207": "GiftUseAck",
        "208": "EquipRecoverReq",
        "209": "EquipRecoverAck",
        "210": "ItemUseReq",
        "211": "ItemUseAck",
        "212": "EquipDestroyReq",
        "213": "EquipDestroyAck",
        "220": "EquipChangeReq",
        "221": "EquipChangeAck",
        "222": "EquipLockReq",
        "223": "EquipLockAck",
        "224": "EquipRemoveReq",
        "225": "EquipRemoveAck",
        "230": "EquipStrengthenReq",
        "231": "EquipStrengthenAck",
        "232": "EquipBlessNtf",
        "233": "EquipStrengthenAutoReq",
        "234": "EquipStrengthenAutoAck",
        "260": "FabaoActiveReq",
        "261": "FabaoActiveAck",
        "262": "FabaoUpLevelReq",
        "263": "FabaoUpLevelAck",
        "264": "FabaoSkillActiveReq",
        "265": "FabaoSkillActiveAck",
        "280": "WingUpLevelReq",
        "281": "WingUpLevelAck",
        "282": "WingSpecialUpReq",
        "283": "WingSpecialUpAck",
        "284": "WingSpecialNtf",
        "285": "WingUseMaterialReq",
        "286": "WingUseMaterialAck",
        "287": "WingWearReq",
        "288": "WingWearAck",
        "300": "StageFightStartReq",
        "301": "StageFightStartAck",
        "302": "StageFightEndReq",
        "303": "StageFightEndNtf",
        "304": "LeaveFightReq",
        "305": "LeaveFightAck",
        "306": "KillMonsterReq",
        "307": "KillMonsterAck",
        "308": "StartStageBossFightReq",
        "309": "StageBagChangeNtf",
        "310": "StageGetHookMapRewardReq",
        "311": "StageGetHookMapRewardAck",
        "320": "PersonBossLoadReq",
        "321": "PersonBossLoadAck",
        "322": "EnterPersonBossFightReq",
        "323": "PersonBossFightResultNtf",
        "324": "EnterPersonBossFightAck",
        "325": "PersonBossSweepReq",
        "326": "PersonBossSweepAck",
        "327": "PersonBossDareNumNtf",
        "330": "HolyActiveReq",
        "331": "HolyActiveAck",
        "332": "HolyUpLevelReq",
        "333": "HolyUpLevelAck",
        "334": "HolySkillActiveReq",
        "335": "HolySkillActiveAck",
        "336": "HolySkillUpLvReq",
        "337": "HolySkillUpLvAck",
        "340": "MiningLoadReq",
        "341": "MiningLoadAck",
        "342": "MiningUpMinerReq",
        "343": "MiningUpMinerAck",
        "344": "MiningBuyNumReq",
        "345": "MiningBuyNumAck",
        "346": "MiningStartReq",
        "347": "MiningStartAck",
        "348": "MiningRobReq",
        "349": "MiningRobAck",
        "350": "MiningRobListReq",
        "351": "MiningRobListAck",
        "352": "MiningListReq",
        "353": "MiningListAck",
        "354": "MiningDrawLoadReq",
        "355": "MiningDrawLoadAck",
        "356": "MiningDrawReq",
        "357": "MiningDrawAck",
        "358": "MiningRobBackReq",
        "359": "MiningRobBackAck",
        "360": "MiningInReq",
        "361": "MiningRobFightAck",
        "362": "MiningRobBackFightAck",
        "363": "MiningInAck",
        "370": "AreaUpLvReq",
        "371": "AreaUpLvAck",
        "380": "ReinActiveReq",
        "381": "ReinActiveAck",
        "382": "ReincarnationReq",
        "383": "ReincarnationAck",
        "384": "ReinCostBuyReq",
        "385": "ReinCostBuyAck",
        "386": "ReinCostUseReq",
        "387": "ReinCostUseAck",
        "388": "ReinCostBuyNumRefNtf",
        "400": "ShopListReq",
        "401": "ShopListAck",
        "402": "ShopBuyReq",
        "403": "ShopBuyAck",
        "404": "ShopWeekResetNtf",
        "411": "TowerRankRewardReq",
        "412": "TowerRankRewardAck",
        "416": "TowerSweepReq",
        "417": "TowerSweepAck",
        "420": "OpenTowerReq",
        "421": "OpenTowerAck",
        "422": "EnterTowerFightReq",
        "423": "TowerFightResultNtf",
        "424": "TowerFightContinueReq",
        "425": "ToweryDayAwardReq",
        "426": "ToweryDayAwardAck",
        "427": "TowerLotteryReq",
        "428": "TowerLotteryAck",
        "429": "TowerLvNtf",
        "430": "AtlasActiveReq",
        "431": "AtlasActiveAck",
        "432": "AtlasUpStarReq",
        "433": "AtlasUpStarAck",
        "434": "AtlasGatherActiveReq",
        "435": "AtlasGatherActiveAck",
        "446": "AtlasGatherUpStarReq",
        "447": "AtlasGatherUpStarAck",
        "448": "AtlasWearChangeReq",
        "449": "AtlasWearChangeAck",
        "450": "FieldBossLoadReq",
        "451": "FieldBossLoadAck",
        "452": "EnterFieldBossFightReq",
        "453": "FieldBossFightResultNtf",
        "454": "FieldBossBuyNumReq",
        "455": "FieldBossBuyNumAck",
        "456": "FieldBossNtf",
        "457": "EnterFieldBossFightAck",
        "458": "BeatBackInfoNtf",
        "459": "FieldBossFirstReq",
        "460": "FieldBossFirstAck",
        "470": "EnterWorldBossFightReq",
        "471": "WorldBossFightResultNtf",
        "472": "WorldBossInfoNtf",
        "490": "MaterialStageLoadReq",
        "491": "MaterialStageLoadAck",
        "492": "EnterMaterialStageFightReq",
        "493": "MaterialStageFightResultNtf",
        "494": "MaterialStageSweepReq",
        "495": "MaterialStageSweepAck",
        "496": "MaterialStageBuyNumNtf",
        "497": "MaterialStageBuyNumReq",
        "498": "MaterialStageBuyNumAck",
        "510": "VipBossLoadReq",
        "511": "VipBossLoadAck",
        "512": "EnterVipBossFightReq",
        "513": "VipBossFightResultNtf",
        "514": "EnterVipBossFightAck",
        "515": "VipBossSweepReq",
        "516": "VipBossSweepAck",
        "517": "VipBossDareNumNtf",
        "530": "ExpStageFightReq",
        "531": "ExpStageDareNumNtf",
        "532": "ExpStageFightResultNtf",
        "533": "ExpStageDoubleReq",
        "534": "ExpStageDoubleAck",
        "535": "ExpStageRefNtf",
        "536": "ExpStageBuyNumNtf",
        "537": "ExpStageSweepReq",
        "538": "ExpStageSweepAck",
        "539": "ExpStageBuyNumReq",
        "540": "ExpStageBuyNumAck",
        "550": "ArenaOpenReq",
        "551": "ArenaOpenAck",
        "552": "EnterArenaFightReq",
        "553": "ArenaFightNtf",
        "556": "BuyArenaFightNumReq",
        "557": "BuyArenaFightNumAck",
        "558": "RefArenaRankReq",
        "559": "RefArenaRankAck",
        "570": "DragonEquipUpLvReq",
        "571": "DragonEquipUpLvAck",
        "580": "SpecialEquipChangeReq",
        "581": "SpecialEquipChangeAck",
        "582": "SpecialEquipRemoveReq",
        "583": "SpecialEquipRemoveAck",
        "590": "DictateUpReq",
        "591": "DictateUpAck",
        "600": "PanaceaUseReq",
        "601": "PanaceaUseAck",
        "620": "SkillUpLvReq",
        "621": "SkillUpLvAck",
        "622": "SkillChangePosReq",
        "623": "SkillChangePosAck",
        "624": "SkillChangeWearReq",
        "625": "SkillChangeWearAck",
        "626": "SkillResetReq",
        "627": "SkillResetAck",
        "628": "SkillUseReq",
        "629": "SkillUseNtf",
        "630": "ClearSkillCdNtf",
        "640": "ComposeReq",
        "641": "ComposeAck",
        "642": "ComposeEquipReq",
        "643": "ComposeEquipAck",
        "644": "ComposeChuanShiEquipReq",
        "645": "ComposeChuanShiEquipAck",
        "660": "MagicCircleUpLvReq",
        "661": "MagicCircleUpLvAck",
        "662": "MagicCircleChangeWearReq",
        "663": "MagicCircleChangeWearAck",
        "670": "ClearReq",
        "671": "ClearAck",
        "680": "JewelMakeReq",
        "681": "JewelMakeAck",
        "682": "JewelUpLvReq",
        "683": "JewelUpLvAck",
        "684": "JewelChangeReq",
        "685": "JewelChangeAck",
        "686": "JewelRemoveReq",
        "687": "JewelRemoveAck",
        "688": "JewelMakeAllReq",
        "689": "JewelMakeAllAck",
        "690": "SignReq",
        "691": "SignAck",
        "692": "SignRepairReq",
        "693": "SignRepairAck",
        "694": "CumulativeSignReq",
        "695": "CumulativeSignAck",
        "696": "SignResetNtf",
        "710": "MailReadReq",
        "711": "MailReadAck",
        "712": "MailRedeemReq",
        "713": "MailRedeemAck",
        "714": "MailNtf",
        "715": "MailLoadReq",
        "716": "MailLoadAck",
        "717": "MailRedeemAllReq",
        "718": "MailRedeemAllAck",
        "719": "MailDeleteReq",
        "720": "MailDeleteAck",
        "721": "MailDeleteAllReq",
        "722": "MailDeleteAllAck",
        "731": "RankLoadReq",
        "732": "RankLoadAck",
        "733": "RankWorshipReq",
        "734": "RankWorshipAck",
        "751": "ChatMessageNtf",
        "752": "ChatMessageListReq",
        "753": "ChatMessageListAck",
        "754": "ChatSendReq",
        "755": "ChatSendAck",
        "756": "ChatBanNtf",
        "757": "ChatBanRemoveNtf",
        "781": "InsideUpStarReq",
        "782": "InsideUpStarAck",
        "783": "InsideUpGradeReq",
        "784": "InsideUpGradeAck",
        "785": "InsideUpOrderReq",
        "786": "InsideUpOrderAck",
        "787": "InsideSkillUpLvReq",
        "788": "InsideSkillUpLvAck",
        "789": "InsideAutoUpReq",
        "790": "InsideAutoUpAck",
        "800": "TaskInfoNtf",
        "801": "TaskDoneReq",
        "802": "TaskDoneAck",
        "803": "TaskNpcStateReq",
        "804": "TaskNpcStateAck",
        "805": "SetTaskInfoReq",
        "806": "SetTaskInfoAck",
        "820": "GetOnlineAwardInfoReq",
        "821": "GetOnlineAwardInfoAck",
        "822": "GetOnlineAwardReq",
        "823": "GetOnlineAwardAck",
        "830": "FashionUpLevelReq",
        "831": "FashionUpLevelAck",
        "832": "FashionWearReq",
        "834": "FashionWearAck",
        "840": "OfficialUpLevelReq",
        "841": "OfficialUpLevelAck",
        "850": "RingWearReq",
        "851": "RingWearAck",
        "852": "RingRemoveReq",
        "853": "RingRemoveAck",
        "854": "RingStrengthenReq",
        "855": "RingStrengthenAck",
        "856": "RingPhantomReq",
        "857": "RingPhantomAck",
        "858": "RingSkillUpReq",
        "859": "RingSkillUpAck",
        "860": "RingFuseReq",
        "861": "RingFuseAck",
        "862": "RingSkillResetReq",
        "863": "RingSkillResetAck",
        "880": "GodEquipActiveReq",
        "881": "GodEquipActiveAck",
        "882": "GodEquipUpLevelReq",
        "883": "GodEquipUpLevelAck",
        "884": "GodEquipBloodReq",
        "885": "GodEquipBloodAck",
        "890": "JuexueUpLevelReq",
        "891": "JuexueUpLevelAck",
        "900": "PetActiveReq",
        "901": "PetActiveAck",
        "902": "PetUpLvReq",
        "903": "PetUpLvAck",
        "904": "PetUpGradeReq",
        "905": "PetUpGradeAck",
        "906": "PetBreakReq",
        "907": "PetBreakAck",
        "908": "PetChangeWearReq",
        "909": "PetChangeWearAck",
        "910": "PetAppendageReq",
        "911": "PetAppendageAck",
        "920": "CompetitveLoadReq",
        "921": "CompetitveLoadAck",
        "922": "BuyCompetitveChallengeTimesReq",
        "923": "BuyCompetitveChallengeTimesAck",
        "924": "RefCompetitveRankReq",
        "925": "RefCompetitveRankAck",
        "926": "GetCompetitveDailyRewardReq",
        "927": "GetCompetitveDailyRewardAck",
        "928": "EnterCompetitveFightReq",
        "929": "CompetitveFightNtf",
        "930": "CompetitveMultipleClaimReq",
        "931": "CompetitveMultipleClaimAck",
        "941": "FieldFightLoadReq",
        "942": "FieldFightLoadAck",
        "943": "EnterFieldFightReq",
        "944": "FieldFightNtf",
        "945": "BuyFieldFightChallengeTimesReq",
        "946": "BuyFieldFightChallengeTimesAck",
        "947": "RefFieldFightRivalUserReq",
        "948": "RefFieldFightRivalUserAck",
        "956": "DarkPalaceDareNumNtf",
        "961": "DarkPalaceLoadReq",
        "962": "DarkPalaceLoadAck",
        "963": "EnterDarkPalaceFightReq",
        "964": "DarkPalaceFightResultNtf",
        "965": "DarkPalaceBuyNumReq",
        "966": "DarkPalaceBuyNumAck",
        "967": "DarkPalaceBossNtf",
        "968": "EnterDarkPalaceHelpFightReq",
        "969": "DarkPalaceHelpFightResultNtf",
        "970": "ExpPoolLoadReq",
        "971": "ExpPoolLoadAck",
        "972": "ExpPoolUpGradeReq",
        "973": "ExpPoolUpGradeAck",
        "981": "WarehouseInfoReq",
        "982": "WarehouseInfoNtf",
        "983": "WareHouseSpaceAddReq",
        "984": "WareHouseSpaceAddAck",
        "985": "WarehouseAddReq",
        "986": "WarehouseAddAck",
        "987": "WarehouseShiftOutReq",
        "988": "WarehouseShiftOutAck",
        "989": "WarehouseSortReq",
        "990": "WarehouseSortAck",
        "1001": "TalentUpLvReq",
        "1002": "TalentUpLvAck",
        "1003": "TalentResetReq",
        "1004": "TalentResetAck",
        "1020": "GuildLoadInfoReq",
        "1021": "GuildLoadInfoAck",
        "1022": "CreateGuildReq",
        "1023": "CreateGuildAck",
        "1026": "QuitGuildReq",
        "1027": "QuitGuildAck",
        "1028": "KickOutReq",
        "1029": "KickOutAck",
        "1030": "ImpeachPresidentReq",
        "1031": "ImpeachPresidentAck",
        "1032": "GuildCheckMemberInfoReq",
        "1033": "GuildCheckMemberInfoAck",
        "1034": "ApplyJoinGuildReq",
        "1035": "ApplyJoinGuildAck",
        "1036": "GuildAssignReq",
        "1037": "GuildAssignAck",
        "1038": "AllGuildInfosReq",
        "1039": "AllGuildInfosAck",
        "1040": "DissolveGuildReq",
        "1041": "DissolveGuildAck",
        "1042": "JoinGuildCombatLimitReq",
        "1043": "JoinGuildCombatLimitAck",
        "1044": "JoinGuildDisposeReq",
        "1045": "JoinGuildDisposeAck",
        "1046": "GetApplyUserListReq",
        "1047": "GetApplyUserListAck",
        "1048": "ModifyBulletinReq",
        "1049": "ModifyBulletinAck",
        "1050": "GuildBonfireLoadReq",
        "1051": "GuildBonfireLoadAck",
        "1052": "GuildBonfireAddExpReq",
        "1053": "GuildBonfireAddExpAck",
        "1054": "EnterGuildBonfireFightReq",
        "1055": "GuildBonfireFightNtf",
        "1056": "GuildBonfireOpenStateNtf",
        "1057": "JoinGuildSuccessNtf",
        "1058": "AllJoinGuildDisposeReq",
        "1059": "AllJoinGuildDisposeAck",
        "1060": "ApplyJoinGuildReDotNtf",
        "1061": "ImpeachPresidentNtf",
        "1062": "BroadcastGuildChangeNtf",
        "1071": "EnterDailyActivityReq",
        "1072": "DailyActivityResultNtf",
        "1073": "DailyActivityListReq",
        "1074": "DailyActivityListAck",
        "1075": "PaodianFightEnd",
        "1100": "ShaBaKeInfoReq",
        "1101": "ShaBaKeInfoAck",
        "1102": "EnterShaBaKeFightReq",
        "1103": "ShaBaKeFightResultNtf",
        "1104": "ShabakeIsOpenNtf",
        "1111": "VipGiftGetReq",
        "1112": "VipGiftGetAck",
        "1130": "AuctionInfoReq",
        "1131": "AuctionInfoNtf",
        "1132": "BidInfoReq",
        "1133": "BidInfoNtf",
        "1134": "BidReq",
        "1135": "BidNtf",
        "1136": "MyBidReq",
        "1137": "MyBidNtf",
        "1138": "BidItemUpdateNtf",
        "1139": "AuctionPutawayItemReq",
        "1140": "AuctionPutawayItemNtf",
        "1141": "MyPutAwayItemInfoReq",
        "1142": "MyPutAwayItemInfoAck",
        "1143": "MyBidInfoItemReq",
        "1144": "MyBidInfoItemAck",
        "1145": "RedPointStateNtf",
        "1146": "ConversionGoldIngotReq",
        "1147": "ConversionGoldIngotAck",
        "1161": "FriendListReq",
        "1162": "FriendListAck",
        "1163": "FriendAddReq",
        "1164": "FriendAddAck",
        "1165": "FriendDelReq",
        "1166": "FriendDelAck",
        "1167": "FriendBlockAddReq",
        "1168": "FriendBlockAddAck",
        "1169": "FriendSearchReq",
        "1170": "FriendSearchAck",
        "1171": "FriendBlockListReq",
        "1172": "FriendBlockListAck",
        "1173": "FriendBlockDelReq",
        "1174": "FriendBlockDelAck",
        "1175": "FriendMsgReadReq",
        "1176": "FriendMsgReadAck",
        "1177": "FriendUserInfoReq",
        "1178": "FriendUserInfoAck",
        "1179": "FriendMsgReq",
        "1180": "FriendMsgAck",
        "1181": "FriendApplyAddReq",
        "1182": "FriendApplyAddNtf",
        "1183": "FriendApplyAgreeReq",
        "1184": "FriendApplyAgreeNtf",
        "1185": "FriendApplyRefuseReq",
        "1186": "FriendApplyRefuseNtf",
        "1187": "FriendApplyListReq",
        "1188": "FriendApplyListAck",
        "1190": "FitUpLvReq",
        "1191": "FitUpLvAck",
        "1192": "FitSkillUpLvReq",
        "1193": "FitSkillUpLvAck",
        "1194": "FitSkillUpStarReq",
        "1195": "FitSkillUpStarAck",
        "1196": "FitSkillChangeReq",
        "1197": "FitSkillChangeAck",
        "1198": "FitSkillResetReq",
        "1199": "FitSkillResetAck",
        "1200": "FitFashionUpLvReq",
        "1201": "FitFashionUpLvAck",
        "1202": "FitFashionChangeReq",
        "1203": "FitFashionChangeAck",
        "1204": "FitSkillActiveReq",
        "1205": "FitSkillActiveAck",
        "1206": "FitEnterReq",
        "1207": "FitEnterAck",
        "1208": "FitCancleReq",
        "1209": "FitCancleAck",
        "1210": "RechargeAllGetReq",
        "1211": "RechargeAllGetAck",
        "1212": "RechargeResetNtf",
        "1221": "DailyTaskLoadReq",
        "1222": "DailyTaskLoadAck",
        "1223": "BuyChallengeTimeReq",
        "1224": "BuyChallengeTimeAck",
        "1225": "GetAwardReq",
        "1226": "GetAwardAck",
        "1227": "ResourcesBackGetRewardReq",
        "1228": "ResourcesBackGetRewardAck",
        "1229": "GetExpReq",
        "1230": "GetExpAck",
        "1231": "ResourcesBackGetAllRewardReq",
        "1232": "ResourcesBackGetAllRewardAck",
        "1241": "MonthCardBuyReq",
        "1242": "MonthCardBuyAck",
        "1243": "MonthCardDailyRewardReq",
        "1244": "MonthCardDailyRewardAck",
        "1260": "FirstRechargeRewardReq",
        "1261": "FirstRechargeRewardAck",
        "1262": "FirstRechargeNtf",
        "1280": "DailyRankLoadReq",
        "1281": "DailyRankLoadAck",
        "1282": "DailyRankGetMarkRewardReq",
        "1283": "DailyRankGetMarkRewardAck",
        "1284": "DailyRankBuyGiftReq",
        "1285": "DailyRankBuyGiftAck",
        "1291": "OpenGiftReq",
        "1292": "OpenGiftAck",
        "1301": "SpendRebatesRewardReq",
        "1302": "SpendRebatesRewardAck",
        "1303": "SpendRebatesNtf",
        "1320": "OfflineAwardLoadReq",
        "1321": "OfflineAwardLoadAck",
        "1322": "OfflineAwardGetReq",
        "1323": "OfflineAwardGetAck",
        "1331": "AchievementLoadReq",
        "1332": "AchievementLoadAck",
        "1333": "AchievementGetAwardReq",
        "1334": "AchievementGetAwardAck",
        "1335": "ActiveMedalReq",
        "1336": "ActiveMedalAck",
        "1337": "AchievementTaskInfoNtf",
        "1341": "GiftCodeRewardReq",
        "1342": "GiftCodeRewardAck",
        "1360": "ChallengeInfoReq",
        "1361": "ChallengeInfoAck",
        "1362": "ApplyChallengeReq",
        "1363": "ApplyChallengeAck",
        "1364": "ChallengeEachRoundPeopleReq",
        "1365": "ChallengeEachRoundPeopleAck",
        "1366": "BottomPourReq",
        "1367": "BottomPourAck",
        "1368": "ChallengeOpenNtf",
        "1369": "ChallengeRoundEndNtf",
        "1370": "ChallengeApplyUserInfoNtf",
        "1381": "LimitedGiftNtf",
        "1382": "LimitedGiftBuyReq",
        "1383": "LimitedGiftBuyAck",
        "1394": "LimitedGiftReq",
        "1400": "DailyPackBuyReq",
        "1401": "DailyPackBuyAck",
        "1420": "GrowFundBuyReq",
        "1421": "GrowFundBuyAck",
        "1422": "GrowFundRewardReq",
        "1423": "GrowFundRewardAck",
        "1440": "WorldLeaderStartNtf",
        "1441": "LoadWorldLeaderReq",
        "1442": "LoadWorldLeaderAck",
        "1443": "WorldLeaderEnterReq",
        "1444": "WorldLeaderEnterAck",
        "1445": "GetWorldLeaderRankInfoReq",
        "1446": "GetWorldLeaderRankInfoAck",
        "1447": "WorldLeaderEndRewardNtf",
        "1448": "WorldLeaderBossHpNtf",
        "1461": "WarOrderTaskNtf",
        "1462": "WarOrderTaskFinishReq",
        "1463": "WarOrderTaskFinishAck",
        "1464": "WarOrderTaskRewardReq",
        "1465": "WarOrderTaskRewardAck",
        "1466": "WarOrderBuyLuxuryReq",
        "1467": "WarOrderBuyLuxuryAck",
        "1468": "WarOrderBuyExpReq",
        "1469": "WarOrderBuyExpAck",
        "1470": "WarOrderLvRewardReq",
        "1471": "WarOrderLvRewardAck",
        "1472": "WarOrderExchangeReq",
        "1473": "WarOrderExchangeAck",
        "1474": "WarOrderLvNtf",
        "1475": "WarOrderOpenReq",
        "1476": "WarOrderOpenAck",
        "1477": "WarOrderResetNtf",
        "1500": "ShaBaKeInfoCrossReq",
        "1501": "ShaBaKeInfoCrossAck",
        "1502": "EnterCrossShaBaKeFightReq",
        "1503": "EnterCrossShaBaKeFightAck",
        "1504": "CrossShabakeOpenNtf",
        "1521": "ElfFeedReq",
        "1522": "ElfFeedAck",
        "1523": "ElfSkillUpLvReq",
        "1524": "ElfSkillUpLvAck",
        "1525": "ElfSkillChangePosReq",
        "1526": "ElfSkillChangePosAck",
        "1540": "CardActivityApplyGetReq",
        "1541": "CardActivityApplyGetAck",
        "1542": "CardActivityInfosReq",
        "1543": "CardActivityInfosAck",
        "1544": "CardInfoNtf",
        "1545": "GetIntegralAwardReq",
        "1546": "GetIntegralAwardAck",
        "1547": "CardCloseNtf",
        "1561": "SetTreasurePopUpStateReq",
        "1562": "SetTreasurePopUpStateAck",
        "1563": "ChooseTreasureAwardReq",
        "1564": "ChooseTreasureAwardAck",
        "1565": "BuyTreasureItemReq",
        "1566": "BuyTreasureItemAck",
        "1567": "TreasureApplyGetReq",
        "1568": "TreasureApplyGetAck",
        "1569": "TreasureInfosReq",
        "1570": "TreasureInfosAck",
        "1571": "GetTreasureIntegralAwardReq",
        "1572": "GetTreasureIntegralAwardAck",
        "1573": "TreasureInfoNtf",
        "1574": "TreasureDrawInfoReq",
        "1575": "TreasureDrawInfoAck",
        "1576": "TreasureCloseNtf",
        "1581": "CutTreasureUpLvReq",
        "1582": "CutTreasureUpLvAck",
        "1583": "CutTreasureUseReq",
        "1584": "CutTreasureUseAck",
        "1600": "HolyBeastLoadInfoReq",
        "1601": "HolyBeastLoadInfoAck",
        "1602": "HolyBeastActivateReq",
        "1603": "HolyBeastActivateAck",
        "1604": "HolyBeastUpStarReq",
        "1605": "HolyBeastUpStarAck",
        "1606": "HolyBeastPointAddReq",
        "1607": "HolyBeastPointAddAck",
        "1608": "HolyBeastChoosePropReq",
        "1609": "HolyBeastChoosePropAck",
        "1610": "HolyBeastRestReq",
        "1611": "HolyBeastRestAck",
        "1621": "FitHolyEquipComposeReq",
        "1622": "FitHolyEquipComposeAck",
        "1623": "FitHolyEquipDeComposeReq",
        "1624": "FitHolyEquipDeComposeAck",
        "1625": "FitHolyEquipWearReq",
        "1626": "FitHolyEquipWearAck",
        "1627": "FitHolyEquipRemoveReq",
        "1628": "FitHolyEquipRemoveAck",
        "1629": "FitHolyEquipSuitSkillChangeReq",
        "1630": "FitHolyEquipSuitSkillChangeAck",
        "1650": "ChuanShiWearReq",
        "1651": "ChuanShiWearAck",
        "1652": "ChuanShiRemoveReq",
        "1653": "ChuanShiRemoveAck",
        "1654": "ChuanShiDeComposeReq",
        "1655": "ChuanShiDeComposeAck",
        "1656": "ChuanshiStrengthenReq",
        "1657": "ChuanshiStrengthenAck",
        "1700": "PreviewFunctionLoadReq",
        "1701": "PreviewFunctionLoadAck",
        "1702": "PreviewFunctionGetReq",
        "1703": "PreviewFunctionGetAck",
        "1704": "PreviewFunctionPointReq",
        "1705": "PreviewFunctionPointAck",
        "1721": "SevenInvestmentLoadReq",
        "1722": "SevenInvestmentLoadAck",
        "1723": "GetSevenInvestmentAwardReq",
        "1724": "GetSevenInvestmentAwardAck",
        "1731": "ContRechargeCycleNtf",
        "1732": "ContRechargeNtf",
        "1733": "ContRechargeReceiveReq",
        "1734": "ContRechargeReceiveAck",
        "1750": "OpenGiftBuyNtf",
        "1751": "OpenGiftEndTimeReq",
        "1752": "OpenGiftEndTimeAck",
        "1770": "VipCustomerReq",
        "1771": "VipCustomerAck",
        "1790": "PaoMaDengNtf",
        "1800": "AncientBossLoadReq",
        "1801": "AncientBossLoadAck",
        "1802": "AncientBossBuyNumReq",
        "1803": "AncientBossBuyNumAck",
        "1804": "EnterAncientBossFightReq",
        "1805": "EnterAncientBossFightAck",
        "1806": "AncientBossFightResultNtf",
        "1807": "AncientBossNtf",
        "1808": "AncientBossOwnerReq",
        "1809": "AncientBossOwnerAck",
        "1830": "AncientSkillUpLvReq",
        "1831": "AncientSkillUpLvAck",
        "1832": "AncientSkillUpGradeReq",
        "1833": "AncientSkillUpGradeAck",
        "1834": "AncientSkillActiveReq",
        "1835": "AncientSkillActiveAck",
        "1850": "EnterGuardPillarReq",
        "1851": "GuardPillarResultNtf",
        "1870": "TitleActiveReq",
        "1871": "TitleActiveAck",
        "1872": "TitleWearReq",
        "1873": "TitleWearAck",
        "1874": "TitleAutoActiveNtf",
        "1875": "TitleExpireNtf",
        "1876": "TitleRemoveReq",
        "1877": "TitleRemoveAck",
        "1878": "TitleLookReq",
        "1879": "TitleLookAck",
        "1890": "GuildActivityOpenNtf",
        "1891": "GuildActivityLoadReq",
        "1892": "GuildActivityLoadAck",
        "1920": "MiJiUpReq",
        "1921": "MiJiUpAck",
        "1941": "AncientTreasuresLoadReq",
        "1942": "AncientTreasuresLoadAck",
        "1943": "AncientTreasuresZhuLinReq",
        "1944": "AncientTreasuresZhuLinAck",
        "1945": "AncientTreasuresUpStarReq",
        "1946": "AncientTreasuresUpStarAck",
        "1947": "AncientTreasuresJueXingReq",
        "1948": "AncientTreasuresJueXingAck",
        "1949": "AncientTreasuresActivateReq",
        "1950": "AncientTreasuresActivateAck",
        "1951": "AncientTreasuresResertReq",
        "1952": "AncientTreasuresResertAck",
        "1953": "AncientTreasuresCondotionInfosReq",
        "1954": "AncientTreasuresCondotionInfosAck",
        "1961": "KillMonsterUniLoadReq",
        "1962": "KillMonsterUniLoadAck",
        "1963": "KillMonsterUniFirstDrawReq",
        "1964": "KillMonsterUniFirstDrawAck",
        "1965": "KillMonsterUniDrawReq",
        "1966": "KillMonsterUniDrawAck",
        "1967": "KillMonsterUniKillNtf",
        "1968": "KillMonsterPerLoadReq",
        "1969": "KillMonsterPerLoadAck",
        "1970": "KillMonsterPerDrawReq",
        "1971": "KillMonsterPerDrawAck",
        "1972": "KillMonsterPerKillNtf",
        "1973": "KillMonsterMilLoadReq",
        "1974": "KillMonsterMilLoadAck",
        "1975": "KillMonsterMilDrawReq",
        "1976": "KillMonsterMilDrawAck",
        "1977": "KillMonsterMilKillNtf",
        "2000": "TreasureShopLoadReq",
        "2001": "TreasureShopLoadAck",
        "2002": "TreasureShopCarChangeReq",
        "2003": "TreasureShopCarChangeAck",
        "2004": "TreasureShopBuyReq",
        "2005": "TreasureShopBuyAck",
        "2006": "TreasureShopRefreshReq",
        "2007": "TreasureShopRefreshNtf",
        "2020": "HellBossLoadReq",
        "2021": "HellBossLoadAck",
        "2022": "HellBossBuyNumReq",
        "2023": "HellBossBuyNumAck",
        "2024": "HellBossDareNumNtf",
        "2025": "EnterHellBossFightReq",
        "2026": "HellBossFightResultNtf",
        "2027": "HellBossNtf",
        "2050": "MagicTowerEndNtf",
        "2051": "MagicTowerGetUserInfoReq",
        "2052": "MagicTowerGetUserInfoAck",
        "2053": "MagicTowerlayerAwardReq",
        "2054": "MagicTowerlayerAwardAck",
        "2061": "LotteryInfoReq",
        "2062": "LotteryInfoAck",
        "2063": "LotteryBuyNumsReq",
        "2064": "LotteryBuyNumsAck",
        "2065": "GetGoodLuckReq",
        "2066": "GetGoodLuckAck",
        "2067": "LotteryEnd",
        "2068": "SetLotteryPopUpStateReq",
        "2069": "SetLotteryPopUpStateAck",
        "2070": "LotteryInfo1Req",
        "2071": "LotteryInfo1Ack",
        "2072": "LotteryGetEndAwardReq",
        "2073": "LotteryGetEndAwardAck",
        "2074": "BrocastBuyNumsNtf",
        "2081": "TrialTaskInfoReq",
        "2082": "TrialTaskInfoAck",
        "2083": "TrialTaskGetAwardReq",
        "2084": "TrialTaskGetAwardAck",
        "2085": "TrialTaskGetStageAwardReq",
        "2086": "TrialTaskGetStageAwardAck",
        "2087": "TrialTaskInfoNtf",
        "2101": "DaBaoEquipUpReq",
        "2102": "DaBaoEquipUpAck",
        "2103": "EnterDaBaoMysteryReq",
        "2104": "DaBaoMysteryResultNtf",
        "2105": "DaBaoMysteryEnergyItemBuyReq",
        "2106": "DaBaoMysteryEnergyAddReq",
        "2107": "DaBaoMysteryEnergyAddAck",
        "2108": "DaBaoMysteryEnergyNtf",
        "2140": "EnterAppletsReq",
        "2141": "AppletsEnergyNtf",
        "2142": "AppletsReceiveReq",
        "2143": "AppletsReceiveAck",
        "2144": "CronGetAwardReq",
        "2145": "CronGetAwardAck",
        "2146": "EndResultReq",
        "2147": "EndResultAck",
        "2180": "LabelUpReq",
        "2181": "LabelUpAck",
        "2182": "LabelTransferReq",
        "2183": "LabelTransferAck",
        "2184": "LabelDayRewardReq",
        "2185": "LabelDayRewardAck",
        "2186": "LabelTaskReq",
        "2187": "LabelTaskNtf",
        "2220": "PrivilegeBuyReq",
        "2221": "PrivilegeBuyAck",
        "2250": "GetBossFamilyInfoReq",
        "2251": "GetBossFamilyInfoAck",
        "2252": "EnterBossFamilyReq",
        "2260": "FirstDropLoadReq",
        "2261": "FirstDropLoadAck",
        "2262": "GetFirstDropAwardReq",
        "2263": "GetFirstDropAwardAck",
        "2264": "GetAllFirstDropAwardReq",
        "2265": "GetAllFirstDropAwardAck",
        "2266": "GetAllRedPacketReq",
        "2267": "GetAllRedPacketAck",
        "2268": "GetAllFirstDropAwardNtf",
        "2269": "FirstDropRedPointNtf",
        "8000": "EnterPublicCopyReq",
        "8001": "EnterPublicCopyAck",
        "8102": "FightPickUpReq",
        "8103": "FightPickUpAck",
        "8104": "FightUserReliveReq",
        "8105": "FightUserReliveAck",
        "8106": "FightCheerReq",
        "8107": "FightCheerAck",
        "8108": "FightPotionReq",
        "8109": "FightPotionAck",
        "8110": "FightCollectionReq",
        "8111": "FightCollectionAck",
        "8112": "FightGetCheerNumReq",
        "8113": "FightGetCheerNumNtf",
        "8114": "FightCollectionNtf",
        "8115": "FightPotionCdReq",
        "8116": "FightPotionCdAck",
        "8117": "FightCollectionCancelReq",
        "8118": "FightCollectionCancelAck",
        "8119": "FightApplyForHelpReq",
        "8120": "FightApplyForHelpAck",
        "8121": "FightApplyForHelpNtf",
        "8122": "FightAskForHelpResultReq",
        "8123": "FightAskForHelpResultNtf",
        "8124": "FightTeamChangeNtf",
        "8125": "FightUserChangeToHelperNtf",
        "8126": "FightAskForHelpResultAck",
        "8127": "FightItemsAddNtf",
        "8128": "FightNpcEventReq",
        "9001": "SceneEnterNtf",
        "9002": "SceneLeaveNtf",
        "9003": "SceneMoveRpt",
        "9004": "SceneMoveNtf",
        "9005": "AttackRpt",
        "9007": "AttackEffectNtf",
        "9008": "SceneDieNtf",
        "9009": "SceneEnterOverNtf",
        "9010": "FightHurtRankReq",
        "9011": "FightHurtRankAck",
        "9013": "SceneUserUpdateNtf",
        "9014": "SceneUserReliveNtf",
        "9015": "SceneObjHpNtf",
        "9016": "GetBossOwnerChangReq",
        "9017": "BossOwnerChangNtf",
        "9018": "SceneObjMpNtf",
        "9019": "CollectionStatusChangeNtf",
        "9020": "FightEnterOkReq",
        "9021": "FightStartCountDownNtf",
        "9022": "FightStartCountDownOkReq",
        "9023": "FightStartNtf",
        "9024": "SceneUserElfUpdateNtf",
        "9025": "FightCheerNumChangeNtf",
        "9026": "GuardPillarFightNtf",
        "9027": "MagicTowerFightNtf",
        "9030": "BuffChangeNtf",
        "9031": "BuffDelNtf",
        "9032": "BuffHpChangeNtf",
        "9033": "BuffPropChangeNtf",
        "9040": "FightUserScoreNtf",
        "9050": "ExpStageKillInfoNtf",
        "9051": "PaodianTopUserReq",
        "9052": "PaodianTopUserNtf",
        "9053": "GetShabakeScoresReq",
        "9054": "ShabakeScoreRankNtf",
        "9055": "PaoDianUserNumNtf",
        "9056": "ShabakeOccupiedNtf",
        "9060": "GetShabakeCrossScoresReq",
        "9061": "ShabakeCrossScoreRankNtf",
        "9062": "ShabakeCrossOccupiedNtf",
        "9063": "CrossShaBaKeFightEndNtf",
        "9070": "GetFightBossInfosReq",
        "9071": "GetFightBossInfosAck",
        "9072": "BossReliveNtf",
    };
    pb.idMap = idMap;
    pb.nameMap = {};
    for (var key in idMap) {
        pb.nameMap[idMap[key]] = +key;
    }
})();
var api = {
};
window.pb = pb;
window.api = api;

