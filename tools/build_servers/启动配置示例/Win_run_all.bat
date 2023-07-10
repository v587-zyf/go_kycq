@echo off
setlocal EnableDelayedExpansion


set PROJECT_DIR=D:\chuanqiH5\server\src\gitlab.hd.com\

set SERVER_DIR=D:\chuanqiH5\server\bin\

set CONF_DIR=conf\loginserver\

set CLIENT_DIR=D:\goProject\yulong\src\gitlab.hd.com\yulong\data\client\

echo 输入要进行的操作0:生成gamedb.dat启动,1:全部生成启动,其他直接启动
set t1=2
set /p bj=%t1%
if "%bj%"=="1" (
  %PROJECT_DIR%yulong\data\genconf.exe -gamedb=%PROJECT_DIR%yulong\data\configs -tplPath=%PROJECT_DIR%yulong\data -client=true -lang=zh-cn -cntjs=./cn-t.js
	move /y data.d.ts %CLIENT_DIR%
	move /y data.js %CLIENT_DIR%
	move /y data.json %CLIENT_DIR%
) else (
	if "%bj%"=="0" (
		%PROJECT_DIR%yulong\data\genconf.exe -gamedb=%PROJECT_DIR%yulong\data\configs -tplPath=%PROJECT_DIR%yulong\data -client=false -lang=zh-cn -cntjs=./cn-t.js
		echo create gamedb.dat
  )
)



start "xk_Login" %SERVER_DIR%loginserver.exe -conf=conf\loginserver\loginserver.conf -logconf=conf\loginserver\logger.json -profileport=6161

start "xk_CrossCenter" %SERVER_DIR%crosscenterserver.exe -conf=conf\crosscenterserver\crosscenterserver.conf -gamedb=%PROJECT_DIR%yulong\data\configs\gamedb.dat -logconf=conf\crosscenterserver\logger.json -profileport=6170

start "xk_FsCc" %SERVER_DIR%fightcenterserver.exe -conf=conf\fightserver\fightcenterserver.conf -logconf=conf\fightserver\loggerfcc.json

start "xk_FsCross1" %SERVER_DIR%fightserver.exe -conf=conf\fightserver\crossfightserver.conf -gamedb=%PROJECT_DIR%yulong\data\configs\gamedb.dat -logconf=conf\fightserver\loggercf2.json -profileport=6167

start "xk_DynamicCross1" %SERVER_DIR%fightserver.exe -conf=conf\fightserver\crossfightserverFix1.conf -gamedb=%PROJECT_DIR%yulong\data\configs\gamedb.dat -logconf=conf\fightserver\loggerfix1.json -profileport=6168

#start "xk_DynamicCross2" %SERVER_DIR%fightserver.exe -conf=conf\fightserver\crossfightserverFix2.conf -gamedb=%PROJECT_DIR%yulong\data\configs\gamedb.dat -logconf=conf\fightserver\loggerfix2.json -profileport=6169



start "xk_Gate1" %SERVER_DIR%gateserver.exe -conf=conf\gateserver\gateserver.conf -logconf=conf\gateserver\logger.json -profileport=6162

start "xk_Game1" %SERVER_DIR%gameserver.exe -conf=conf\gameserver\gameserver.conf -gamedb=%PROJECT_DIR%yulong\data\configs\gamedb.dat -logconf=conf\gameserver\logger.json -profileport=6163

#start "xk_Center1" %SERVER_DIR%centerserver.exe -conf=conf\centerserver\centerserver.conf -logconf=conf\centerserver\logger.json -profileport=6164

start "xk_Fight1" %SERVER_DIR%fightserver.exe -conf=conf\fightserver\fightserver.conf -gamedb=%PROJECT_DIR%yulong\data\configs\gamedb.dat -logconf=conf\fightserver\logger.json -profileport=6165

#start "xk_FsSolo1" %SERVER_DIR%fightserver.exe -conf=conf\fightserver\crossfightserversolo.conf -gamedb=%PROJECT_DIR%yulong\data\configs\gamedb.dat -logconf=conf\fightserver\loggercf.json -profileport=6166


TIMEOUT /T 60
