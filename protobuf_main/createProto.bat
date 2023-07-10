@echo off
set DIR=./
set DIR_SOURCE=%DIR%source/
set DIR_SERVER_OUT=%DIR%out/server/
set DIR_CLIEN_OUT=%DIR%out/client/



echo �������ͻ���Э������
protoc.exe -I=%DIR_SOURCE%pb/ --gofast_out=%DIR_SERVER_OUT%pb/ %DIR_SOURCE%pb/*.proto
node %DIR%genproto\index.js -f %DIR_SOURCE%pb/main.idmap -g go -o %DIR_SERVER_OUT%pb/

echo ����Э������
protoc.exe -I=%DIR_SOURCE%pbgt/ --gofast_out=%DIR_SERVER_OUT%pbgt/ %DIR_SOURCE%pbgt/*.proto
node %DIR%genproto\index.js -f %DIR_SOURCE%pbgt/main.idmap -g go -o %DIR_SERVER_OUT%pbgt/

echo ������ͨѶЭ������
protoc.exe -I=%DIR_SOURCE%pbserver/ --gofast_out=%DIR_SERVER_OUT%pbserver/ %DIR_SOURCE%pbserver/*.proto
node %DIR%genproto\index.js -f %DIR_SOURCE%pbserver/main.idmap -g go -o %DIR_SERVER_OUT%pbserver/

echo �ͻ���Э��json����
node %DIR%genproto\index.js -f %DIR_SOURCE%pb/main.idmap -g json -o %DIR_CLIEN_OUT%
echo �ͻ���Э��ts����
node %DIR%genproto\index.js -f %DIR_SOURCE%pb/main.idmap -g ts -o %DIR_CLIEN_OUT%


echo ��Ϣ�������
TIMEOUT /T 99