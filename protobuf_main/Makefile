

#
#使用帮助
# make db
# 需要安装python 的相关模块
#sudo apt-get install python-setuptools python-dev build-essential
#sudo easy_install pip
#pip install pyparsing
#sudo pip install pyparsing
#debian base 运行以下
#sudo apt-get install libmysqlclient-dev
#centos 如果运行不成功，多运行几遍
#或者运行 yum install mysql-devel
#sudo pip install MySQL-python
#macos
#如果报找不到下面文件，建一个软连
#sudo ln -s /usr/local/mysql/lib/libmysqlclient.21.dylib /usr/local/lib/libmysqlclient.18.dylib
#mysql 创建好数据库后，要导入init_data.sql
#如果使用mysql 5.7+
#在 /etc/mysql/mysql.conf.d/mysqld.cnf 中加入：
#sql_mode = ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
#mac os mysql8.0可能要去掉 NO_AUTO_CREATE_USER
#

-include localEnv

export DIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
export DIR_SOURCE=${DIR}source/
export DIR_SERVER_OUT=D:/chuanqiH5/server/src/cqserver/protobuf/
export DIR_CLIEN_OUT=${DIR_}out/client/


generate_proto proto:
	@echo 001
	@cd ${DIR_SOURCE}
	@echo 002
	@protoc -I=${DIR_SOURCE}pb/ --gofast_out=${DIR_SERVER_OUT}pb/ ${DIR_SOURCE}pb/*.proto
	@echo 003
	@protoc -I=${DIR_SOURCE}pbgt/ --gofast_out=${DIR_SERVER_OUT}pbgt/ ${DIR_SOURCE}pbgt/*.proto
	@echo 004
	@protoc -I=${DIR_SOURCE}pbserver/ --gofast_out=${DIR_SERVER_OUT}pbserver/ ${DIR_SOURCE}pbserver/*.proto
	@echo 005 %cd%
	@node ${DIR}genproto\index.js -f ${DIR_SOURCE}pb/main.idmap -g json -o ${DIR_CLIEN_OUT}
	@echo 006
	@node ${DIR}genproto\index.js -f ${DIR_SOURCE}pb/main.idmap -g ts -o ${DIR_CLIEN_OUT}
	@echo 007
	@node ${DIR}genproto\index.js -f ${DIR_SOURCE}pb/main.idmap -g go -o ${DIR_SERVER_OUT}pb/
	@echo 008
	@node ${DIR}genproto\index.js -f ${DIR_SOURCE}pbserver/main.idmap -g go -o ${DIR_SERVER_OUT}pbserver/
	@echo 009
	@node ${DIR}genproto\index.js -f ${DIR_SOURCE}pbgt/main.idmap -g go -o ${DIR_SERVER_OUT}pbgt/
	@echo 010 ok

