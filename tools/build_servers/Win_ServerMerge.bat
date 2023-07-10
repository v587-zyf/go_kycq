@echo off
setlocal EnableDelayedExpansion


set PROJECT_ROOT_DIR=%cd%/../../../../
set SERVER_DIR=%PROJECT_ROOT_DIR%src\cqserver\tools\build_servers\

set GAMEDB_DIR=%PROJECT_ROOT_DIR%..\config\

if not exist %GAMEDB_DIR% (
	echo 请先配置游戏配置路径
	TIMEOUT /T 30
)



start  "cq_serverMerge"  %SERVER_DIR%servermerge\serverMerge.exe  -conf=%SERVER_DIR%servermerge\serverMerge.conf  -gamedb=%GAMEDB_DIR%gamedb.dat -logconf=%SERVER_DIR%servermerge\serverMerge_logger.json

echo %SERVER_DIR%servermerge\serverMerge.exe  -conf=%SERVER_DIR%servermerge\serverMerge.conf  -gamedb=%GAMEDB_DIR%gamedb.dat -logconf=%SERVER_DIR%servermerge\serverMerge_logger.json

pause