@echo off

::将excel中的前四列转化为struct
::第一列字段类型		如 int
::第二列字段备注		如 符文等级上限
::第三列字段名			如 id
::第四列s,c,all			s表示服务端使用, c表示客户端使用, all表示都使用

REM 数据类型	示例			说明				go
REM int			1000			整形	
REM string		无生无灭炉		字符串	
REM float64		1.5				浮点数	
REM bool		true,false		布尔型	
			
REM IntSlice 	1|2|3|4			int一维数组			[]int{1,2,3,4}
REM IntSlice2 	1,100|2,100|3,100	int二维数组		[][]int{[1,100],[2,100],[3,100]}
REM IntMap		1,100|2,100|3,100	k和v都是int的集合  	map[int]int{1:100,2:100,3:100}
			
REM PropInfo	100|10000		k int,v int  		PropInfo{K: 100,V: 1000}
REM PropInfos	3200071,1|3200072,1	PropInfo的一维数组	
REM ItemInfo	100|10000		物品信息，会进行物品道具检查	ItemInfo{ItemId: 100,Count: 1000}
REM ItemInfos	3200071,1|3200072,1	ItemInfo的一维数组 	
					
REM StringSlice	无生无灭炉|雕花青铜炉	字符串的一维数组  	[]string{"无生无灭炉","雕花青铜炉"}
			
REM HmsTime	06:00:00			时间类型 	

set PROJECT_ROOT_DIR=%cd%/../../../
SET GOPATH=D:\goPath;%PROJECT_ROOT_DIR%
SET SERVER_DIR=%PROJECT_ROOT_DIR%src\cqserver\

::生成objs.go文件的路径
set savePath=%PROJECT_ROOT_DIR%src\cqserver\gamelibs\gamedb
::目标excel文件路径
set confPath=%PROJECT_ROOT_DIR%..\config
::所有的字段类型
set allType=int,IntSlice,IntSlice2,IntSlice3,IntMap,string,StringSlice,StringSlice2,float64,ItemInfo,ItemInfos,PropInfo,PropInfos,HmsTime,HmsTimes,bool,FloatSlice



set cmd_gamedb_struct=1
set cmd_build_genconf=2
set cmd_crate_gamedb_data=3
set cmd_create_language=4
set cmd_gamedb_struct_and_build_genconf=5
set cmd_build_genconf_and_gamedbdata=6
set cmd_all=7
set cmd_build_server=8
set cmd_start_server=9
set cmd_start_all_server=a


set list_len=10

set list[1].cmd=%cmd_gamedb_struct%
set list[1].dec=只生gamedbBase.go
set list[2].cmd=%cmd_build_genconf%
set list[2].dec=编译genconf.exe
set list[3].cmd=%cmd_crate_gamedb_data%
set list[3].dec=生成gamedb.dat
set list[4].cmd=%cmd_create_language%
set list[4].dec=生成language文件
set list[5].cmd=%cmd_gamedb_struct_and_build_genconf%
set list[5].dec=生成gamedbbase.go并编译genconf
set list[6].cmd=%cmd_build_genconf_and_gamedbdata%
set list[6].dec=并编译genconf,打包gamedb.dat
set list[7].cmd=%cmd_all%
set list[7].dec=生成gamedbbase.go并编译genconf,打包gamedb.dat
set list[8].cmd=%cmd_build_server%
set list[8].dec=编译服务器
set list[9].cmd=%cmd_start_server%
set list[9].dec=启动一组服务器
set list[10].cmd=%cmd_start_all_server%
set list[10].dec=启动所有服务器

:_reChoose

echo.
echo	选择你的操作
(for /L %%a in (1,1,%list_len%-1) do (

	call echo %%list[%%a].cmd%% %%list[%%a].dec%%
))
echo.


rem 等待用户输入
set /p chooseResult=请输入操作标识：
set inpuCmd=0
(for /L %%a in (1,1,%list_len%-1) do (

	FOR /F "usebackq delims==. tokens=1-3" %%I IN (`SET list[%%a]`) DO (
		if %%J==cmd (
			if %%K==%chooseResult% (
				set inpuCmd=%chooseResult%
			)
		)
	)
))


if /i %inpuCmd%==0 (
	echo 输入命令错误，重新输入
	goto _reChoose
)


if %chooseResult% ==%cmd_gamedb_struct% (
	call:genGameDBStuct
)
if %chooseResult%==%cmd_build_genconf% (
	call:genconf
)
if %chooseResult% ==%cmd_crate_gamedb_data% (
	call:genGameDbdat
)
if %chooseResult% ==%cmd_gamedb_struct_and_build_genconf% (
	call:genGameDBStuct
	call:genconf
)
if %chooseResult% ==%cmd_all% (
	call:genGameDBStuct
	call:genconf
	call:genGameDbdat
)
if %chooseResult% ==%cmd_build_genconf_and_gamedbdata% (
	call:genconf
	call:genGameDbdat
)
if %chooseResult% ==%cmd_create_language% (
	call:genLanguage
)
if %chooseResult% ==%cmd_build_server% (
	call buildserver_to_windows.bat 1.1.1
)
if %chooseResult% ==%cmd_start_server% (
	cd .\build_servers
	start cmd /k call Win_run_server.bat
	cd ..\
)
if %chooseResult% ==%cmd_start_all_server% (
	cd .\build_servers
	call Win_run_all.bat
)
echo\ 
echo\ 
echo ==================
echo ---本次操作完成---
echo ==================

cd %SERVER_DIR%tools
goto _reChoose

TIMEOUT /T 99
:over
goto:eof

:genGameDBStuct
echo 开始生成配置文件类结构
.\generateStruct\generateStruct.exe -savePath=%savePath% -readPath=%confPath%\excel -allType=%allType%
goto:eof

:genLanguage
echo 开始生成配置文件类结构
.\generateStruct\generateStruct.exe -savePath=%savePath% -readPath=%confPath%\excel -l=true
goto:eof

:genconf
echo 编译genconf.exe:
cd %SERVER_DIR%tools\genconf
go build -o %confPath%\genconf.exe
echo 编译genconf.exe完成:
goto:eof

:genGameDbdat
echo 生成gamedb.dat:
%confPath%\genconf.exe -gamedb=%confPath%
goto:eof

