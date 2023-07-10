@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

set PROJECT_ROOT_DIR=%cd%/../../../

SET GOPATH=D:\goPath;%PROJECT_ROOT_DIR%

SET SERVER_DIR=%PROJECT_ROOT_DIR%src\cqserver\
SET OUTPUT_DIR=%PROJECT_ROOT_DIR%src\cqserver\tools\build_servers\

echo Linux:
::genver %SERVER_DIR%loginserver %SERVER_DIR%gateserver %SERVER_DIR%gameserver %SERVER_DIR%fightserver %SERVER_DIR%fightCenterServer

::move /y server_game_update.sql server_linux/

echo 1.generate loginserver
cd %SERVER_DIR%loginserver
go build -o %OUTPUT_DIR%login\loginserver

echo 2.generate gateserver
cd %SERVER_DIR%gateserver
go build -o %OUTPUT_DIR%gs\gateserver

echo 3.generate gameserver
cd %SERVER_DIR%gameserver
go build -o %OUTPUT_DIR%gs\gameserver

echo 4.generate fightserver
cd %SERVER_DIR%fightserver
go build -o %OUTPUT_DIR%gs\fightserver
go build -o %OUTPUT_DIR%fightcross\fightserver

echo 5.generate fightCenterServer
cd %SERVER_DIR%fightCenterServer
go build -o %OUTPUT_DIR%fightcenter\fightcenterserver

echo 6.generate crosscenterserver
cd %SERVER_DIR%crosscenterserver
go build -o %OUTPUT_DIR%crosscenter\crosscenterserver

echo 7.generate servermerge:
cd %SERVER_DIR%tools\serverMerge
go build -o %OUTPUT_DIR%servermerge\servermerge

echo Done.
pause