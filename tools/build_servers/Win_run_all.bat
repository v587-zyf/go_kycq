@echo off
setlocal EnableDelayedExpansion

set PROJECT_ROOT_DIR=%cd%/../../../../
set SERVER_DIR=%PROJECT_ROOT_DIR%src\cqserver\tools\build_servers\

set GAMEDB_DIR=%PROJECT_ROOT_DIR%..\config\

if not exist %GAMEDB_DIR% (
	echo «Îœ»≈‰÷√”Œœ∑≈‰÷√¬∑æ∂
	TIMEOUT /T 30
)



start "cq_Login" %SERVER_DIR%login\loginserver.exe -conf=%SERVER_DIR%login\loginserver.conf -logconf=%SERVER_DIR%login\logger.json -profileport=6161

start "cq_CrossCenter" %SERVER_DIR%crosscenter\crosscenterserver.exe -conf=%SERVER_DIR%crosscenter\crosscenterserver.conf -gamedb=%GAMEDB_DIR%gamedb.dat -logconf=%SERVER_DIR%crosscenter\logger.json -profileport=6170

start "cq_fightcenter" %SERVER_DIR%fightcenter\fightcenterserver.exe -conf=%SERVER_DIR%fightcenter\fightcenterserver.conf -logconf=%SERVER_DIR%fightcenter\logger.json

start "cq_FsCross1" %SERVER_DIR%fightcross\fightserver.exe -conf=%SERVER_DIR%fightcross\fightserver.conf -gamedb=%GAMEDB_DIR%gamedb.dat -logconf=%SERVER_DIR%fightcross\logger.json -profileport=6167


start "cq_Gate1" %SERVER_DIR%gs\gateserver.exe -conf=%SERVER_DIR%gs\gateserver.conf -logconf=%SERVER_DIR%gs\gateserver_logger.json -profileport=6162

start "cq_Fight1" %SERVER_DIR%gs\fightserver.exe -conf=%SERVER_DIR%gs\fightserver.conf -gamedb=%GAMEDB_DIR%gamedb.dat -logconf=%CONF_DIR%%SERVER_DIR%gs\fightserver_logger.json -profileport=6164

start "cq_Game1" %SERVER_DIR%gs\gameserver.exe -conf=%SERVER_DIR%gs\gameserver.conf -gamedb=%GAMEDB_DIR%gamedb.dat -logconf=%SERVER_DIR%gs\gameserver_logger.json -profileport=6163
rem D:\chuanqiH5\server\src\cqserver\tools\server_windows\gameserver.exe -conf=conf\gameserver\gameserver.conf -gamedb=D:\chuanqiH5\config\gamedb.dat -logconf=conf\gameserver\logger.json -profileport=6163

::start "xk_Center1" %SERVER_DIR%centerserver.exe -conf=conf\centerserver\centerserver.conf -logconf=conf\centerserver\logger.json -profileport=6164

::start "xk_Fight1" %SERVER_DIR%fightserver.exe -conf=conf\fightserver\fightserver.conf -gamedb=%PROJECT_DIR%yulong\data\configs\gamedb.dat -logconf=conf\fightserver\logger.json -profileport=6165

::start "xk_FsSolo1" %SERVER_DIR%fightserver.exe -conf=conf\fightserver\crossfightserversolo.conf -gamedb=%PROJECT_DIR%yulong\data\configs\gamedb.dat -logconf=conf\fightserver\loggercf.json -profileport=6166


TIMEOUT /T 1
