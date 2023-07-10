#!/bin/bash
#

gameserverId=2
gamedbName="cq_server_game_2"
redisdb=1
sql="DELETE FROM cq_accountdb.user where serverId=$gameserverId;
DROP TABLES $gamedbName.auction_bid,
$gamedbName.auction_item,
$gamedbName.card,
$gamedbName.guild,
$gamedbName.guild_auction_item,
$gamedbName.hero,
$gamedbName.mail,
$gamedbName.mining,
$gamedbName.orders,
$gamedbName.treasure,
$gamedbName.user;"


echo "清理数据库数据"
/usr/bin/mysql  -h127.0.0.1 -P3306 -uroot -p'123456' -e "$sql"


echo "清理redis数据"
/opt/redis/bin/redis-cli -h 127.0.0.1 -p 6379 -a 123456 <<END
select $redisdb
flushdb
END