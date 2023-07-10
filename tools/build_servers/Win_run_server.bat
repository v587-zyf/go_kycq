@echo off
setlocal EnableDelayedExpansion

set PROJECT_ROOT_DIR=%cd%/../../../../
set SERVER_DIR=%PROJECT_ROOT_DIR%bin\

set GAMEDB_DIR=%PROJECT_ROOT_DIR%..\config\

if not exist %GAMEDB_DIR% (
	echo 请先配置游戏配置路径
	TIMEOUT /T 30
)

set CONF_DIR=%PROJECT_ROOT_DIR%src\cqserver\


echo ======================
echo 1 启动game
echo 2 启动fight
echo 3 启动全部
echo 4 关掉所有
echo ======================
set /p startCmd=输入启动命令

echo %startCmd%




if %startCmd% ==1 (
	start "cq2_Game2" %SERVER_DIR%gameserver.exe -conf=%CONF_DIR%gameserver\gameserver.conf -gamedb=%GAMEDB_DIR%gamedb.dat -logconf=%CONF_DIR%gameserver\logger.json -profileport=6163
)

if %startCmd% ==2 (
	start "cq2_Fight2" %SERVER_DIR%fightserver.exe -conf=%CONF_DIR%fightserver\fightserver.conf -gamedb=%GAMEDB_DIR%gamedb.dat -logconf=%CONF_DIR%fightserver\logger.json -profileport=6164
)

if %startCmd% ==3 (
	rem start "cq_FsCc" %SERVER_DIR%fightcenterserver.exe -conf=%CONF_DIR%\fightcenterserver\fightcenterserver.conf -logconf=conf\fightserver\logger.json
	rem start "cq_DynamicCross1" %SERVER_DIR%fightserver.exe -conf=%CONF_DIR%\fightserver\fightserverCross.conf -gamedb=%GAMEDB_DIR%gamedb.dat -logconf=%CONF_DIR%\fightserver\logger.json -profileport=6168
	start "cq2_Gate2" %SERVER_DIR%gateserver.exe -conf=%CONF_DIR%gateserver\gateserver.conf -logconf=%CONF_DIR%gateserver\logger.json -profileport=6162

	start "cq2_Game2" %SERVER_DIR%gameserver.exe -conf=%CONF_DIR%gameserver\gameserver.conf -gamedb=%GAMEDB_DIR%gamedb.dat -logconf=%CONF_DIR%gameserver\logger.json -profileport=6163

	start "cq2_Fight2" %SERVER_DIR%fightserver.exe -conf=%CONF_DIR%fightserver\fightserver.conf -gamedb=%GAMEDB_DIR%gamedb.dat -logconf=%CONF_DIR%fightserver\logger.json -profileport=6164
)

echo "cq2_Game2" %SERVER_DIR%gameserver.exe -conf=%CONF_DIR%gameserver\gameserver.conf -gamedb=%GAMEDB_DIR%gamedb.dat -logconf=%CONF_DIR%gameserver\logger.json -profileport=6163

if %startCmd% ==4 (
	taskkill /F /FI "WINDOWTITLE eq cq2_*"
)


TIMEOUT /T 60
