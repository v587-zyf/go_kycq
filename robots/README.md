# 机器人
```
该机器人实现了国战与主城跑地图.
支持家族boss战、国家boss战、护国寺、家族守护佳人功能。
```

## 主城
```
./robots -prefix=wei -nums=100 -line=1 -addr=192.168.5.25:7100
```

## 国战启动参数 手工开启后运行机器人
```
./robots -city=11 -prefix=wei -kingdom=1 -nums=100 -battletype=1 -gate=1  -addr=192.168.5.25:7100 -serverIdx=1 -startId=1000
./robots -city=11 -prefix=shu -kingdom=2 -nums=200 -battletype=1 -gate=1  -addr=192.168.5.25:7100 -serverIdx=1 -startId=1000
```

## 家族boss 手工开启后运行机器人
```
 ./robots -prefix=wei -kingdom=1 -battletype=2 -nums=30 -addr=192.168.5.25:7100 -guildname="XXX"
 ```

## 国家boss 手工开启后运行机器人
```
 ./robots -prefix=wei -kingdom=1 -battletype=3 -nbosstype=2  -nums=30 -addr=192.168.5.25:7100
```
## 护国寺
```
 ./robots -prefix=wei -kingdom=1 -battletype=4 -nums=30 -addr=192.168.5.25:7100
 ```

 ## 家族守护佳人 手工开启后运行机器人
```
 ./robots -prefix=wei -kingdom=1 -battletype=5 -nums=30 -addr=192.168.5.25:7100 -guildname="XXX"
 ```

## 支持参数

```
Usage of ./robots:
  -addr string
        http service address (default "127.0.0.1:7100")
  -batch int
        batch num (default 100)
  -batchcd int
        batch log cd (default 850)
  -benchmark
        print elapse of req/ack
  -city int
        national fight city
  -entercd int
        re-entergame cd (default 10)
  -exedura int
        exe dura time (default 1500)
  -kingdom int
        national fight kingdom
  -line int
        maincity line (default 1)
  -nums int
        robots nums (default 1)
  -prefix string
        robot openid prefix (default "rot")
  -serverIdx int
        server idx (default 1)
  -startId int
        robots start openid num (default 1000)
  -step int
        move step (default 30)
  -usews
        robot use ws protocol (default true)
　-battletype int
        battle type, 0-主城　１-国战　2－家族boo战 3-国家boss战　4－护国寺　(default 0)
　-gate int
        国战城门1-4
  -guildName  string
        家族boo战家族名 (default "",非家族boss战不需要)
  - nbosstype int
        国家boss战类型, 1 黄巾起义,2 南蛮入侵,3	匈奴南下 (default 2)
```


## 注
* _prefix参数优先级高于kingdom_
* _相同国家、同段的帐号，不要同时登录，否则会相互踢人。这个问题通过设置startId来避免_






