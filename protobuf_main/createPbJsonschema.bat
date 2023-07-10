@echo off
set DIR=./
set DIR_SOURCE=%DIR%source/
set DIR_SERVER_OUT=%DIR%out/server/
set DIR_CLIEN_OUT=%DIR%out/client/


@node %DIR%genproto\index.js -f %DIR_SOURCE%pb/main.idmap -g jsonschema -o %DIR_CLIEN_OUT%


echo 消息生成完成
TIMEOUT /T 99