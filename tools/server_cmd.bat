@echo off

::��excel�е�ǰ����ת��Ϊstruct
::��һ���ֶ�����		�� int
::�ڶ����ֶα�ע		�� ���ĵȼ�����
::�������ֶ���			�� id
::������s,c,all			s��ʾ�����ʹ��, c��ʾ�ͻ���ʹ��, all��ʾ��ʹ��

REM ��������	ʾ��			˵��				go
REM int			1000			����	
REM string		��������¯		�ַ���	
REM float64		1.5				������	
REM bool		true,false		������	
			
REM IntSlice 	1|2|3|4			intһά����			[]int{1,2,3,4}
REM IntSlice2 	1,100|2,100|3,100	int��ά����		[][]int{[1,100],[2,100],[3,100]}
REM IntMap		1,100|2,100|3,100	k��v����int�ļ���  	map[int]int{1:100,2:100,3:100}
			
REM PropInfo	100|10000		k int,v int  		PropInfo{K: 100,V: 1000}
REM PropInfos	3200071,1|3200072,1	PropInfo��һά����	
REM ItemInfo	100|10000		��Ʒ��Ϣ���������Ʒ���߼��	ItemInfo{ItemId: 100,Count: 1000}
REM ItemInfos	3200071,1|3200072,1	ItemInfo��һά���� 	
					
REM StringSlice	��������¯|����ͭ¯	�ַ�����һά����  	[]string{"��������¯","����ͭ¯"}
			
REM HmsTime	06:00:00			ʱ������ 	

set PROJECT_ROOT_DIR=%cd%/../../../
SET GOPATH=D:\goPath;%PROJECT_ROOT_DIR%
SET SERVER_DIR=%PROJECT_ROOT_DIR%src\cqserver\

::����objs.go�ļ���·��
set savePath=%PROJECT_ROOT_DIR%src\cqserver\gamelibs\gamedb
::Ŀ��excel�ļ�·��
set confPath=%PROJECT_ROOT_DIR%..\config
::���е��ֶ�����
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
set list[1].dec=ֻ��gamedbBase.go
set list[2].cmd=%cmd_build_genconf%
set list[2].dec=����genconf.exe
set list[3].cmd=%cmd_crate_gamedb_data%
set list[3].dec=����gamedb.dat
set list[4].cmd=%cmd_create_language%
set list[4].dec=����language�ļ�
set list[5].cmd=%cmd_gamedb_struct_and_build_genconf%
set list[5].dec=����gamedbbase.go������genconf
set list[6].cmd=%cmd_build_genconf_and_gamedbdata%
set list[6].dec=������genconf,���gamedb.dat
set list[7].cmd=%cmd_all%
set list[7].dec=����gamedbbase.go������genconf,���gamedb.dat
set list[8].cmd=%cmd_build_server%
set list[8].dec=���������
set list[9].cmd=%cmd_start_server%
set list[9].dec=����һ�������
set list[10].cmd=%cmd_start_all_server%
set list[10].dec=�������з�����

:_reChoose

echo.
echo	ѡ����Ĳ���
(for /L %%a in (1,1,%list_len%-1) do (

	call echo %%list[%%a].cmd%% %%list[%%a].dec%%
))
echo.


rem �ȴ��û�����
set /p chooseResult=�����������ʶ��
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
	echo �������������������
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
echo ---���β������---
echo ==================

cd %SERVER_DIR%tools
goto _reChoose

TIMEOUT /T 99
:over
goto:eof

:genGameDBStuct
echo ��ʼ���������ļ���ṹ
.\generateStruct\generateStruct.exe -savePath=%savePath% -readPath=%confPath%\excel -allType=%allType%
goto:eof

:genLanguage
echo ��ʼ���������ļ���ṹ
.\generateStruct\generateStruct.exe -savePath=%savePath% -readPath=%confPath%\excel -l=true
goto:eof

:genconf
echo ����genconf.exe:
cd %SERVER_DIR%tools\genconf
go build -o %confPath%\genconf.exe
echo ����genconf.exe���:
goto:eof

:genGameDbdat
echo ����gamedb.dat:
%confPath%\genconf.exe -gamedb=%confPath%
goto:eof

