合服启动参数配置：
//启动配置
-conf=./src/gitlab.hd.com/ylserver/tools/serverMerge/config.conf  	
	
//合服的serverid（不填写默认读config的mergeServerIds）
-server_ids = 1,2	

//合并到的新数据库dbname（不填写默认读config的dbconfigs.serverdbNew.dbname）
-db_name = yulong_server_game_1	
//合并到的新数据库host（不填写默认读config的dbconfigs.serverdbNew.host）
-db_host = 127.0.0.1
//合并到的新数据库db_port（不填写默认读config的dbconfigs.serverdbNew.port）
-db_port = 3306		
//合并到的新数据库db_user（不填写默认读config的dbconfigs.serverdbNew.uid）
-db_user = root
//合并到的新数据库db_pass（不填写默认读config的dbconfigs.serverdbNew.pwd）
-db_pass = 123456	


//合并到的新redis redis_address（不填写默认读config的dbconfigs.redis.address）
-redis_address = 127.0.0.1:6379	
//合并到的新redis redisDb（不填写默认读config的dbconfigs.redis.db）
-redis_db = 0	
//合并到的新redis redisAddr（不填写默认读config的dbconfigs.redis.password）
-redis_pass = 123456	
	
		
//log配置
-logconf=./src/gitlab.hd.com/ylserver/tools/serverMerge/logger.json 	

//表路径.默认为./config/excels
-csvConf = ./src/gitlab.hd.com/ylserver/tools/serverMerge/config/excels
