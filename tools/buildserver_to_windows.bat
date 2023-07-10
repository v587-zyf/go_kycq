@echo off
::SET CGO_ENABLED=0
::SET GOOS=linux
::SET GOARCH=amd64

set PROJECT_ROOT_DIR=%cd%/../../../

SET GOPATH=D:\goPath;%PROJECT_ROOT_DIR%

SET SERVER_DIR=%PROJECT_ROOT_DIR%src\cqserver\
SET OUTPUT_DIR=%PROJECT_ROOT_DIR%src\cqserver\tools\build_servers\

set version=%1%

if "%1" =="" (
    set version=1.0.0
)

echo Windows:
echo 1.generate loginserver
cd %SERVER_DIR%loginserver
go build -ldflags "-X main.APP_VERSION=%version%" -o %OUTPUT_DIR%login\loginserver.exe

echo 2.generate gateserver
cd %SERVER_DIR%gateserver
go build -ldflags "-X main.APP_VERSION=%version%" -o %OUTPUT_DIR%gs\gateserver.exe

echo 3.generate gameserver
cd %SERVER_DIR%gameserver
go build -ldflags "-X main.APP_VERSION=%version%" -o %OUTPUT_DIR%gs\gameserver.exe

echo 4.generate fightserver
cd %SERVER_DIR%fightserver
go build -ldflags "-X main.APP_VERSION=%version%" -o %OUTPUT_DIR%gs\fightserver.exe
go build -ldflags "-X main.APP_VERSION=%version%" -o %OUTPUT_DIR%fightcross\fightserver.exe

echo 5.generate fightCenterServer
cd %SERVER_DIR%fightcenterserver
go build -ldflags "-X main.APP_VERSION=%version%" -o %OUTPUT_DIR%fightcenter\fightcenterserver.exe

echo 6.generate crosscenterserver
cd %SERVER_DIR%crosscenterserver
go build -ldflags "-X main.APP_VERSION=%version%" -o %OUTPUT_DIR%crosscenter\crosscenterserver.exe

echo 7.generate servermerge
cd %SERVER_DIR%tools\serverMerge
go build -ldflags "-X main.APP_VERSION=%version%" -o %OUTPUT_DIR%servermerge\serverMerge.exe

echo Done.
pause