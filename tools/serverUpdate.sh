#!/bin/bash
#

serverPWD="/opt/www/chuanqiH5"
logPWD="/opt/log"
sourcePWD="/opt/data"

client_update(){
    cd $sourcePWD  
    if [ -d $serverPWD/client/ ];then
        echo "$(date +%F' '%T) - INFO: client更新开始" |tee -a $logPWD/update.log
		#rsync -avz --exclude-from=/opt/tools/serverTools/server.exclude  hrl_client/*  /opt/www/hrl/hrl_client/
		rsync -avz  h5  $serverPWD/client/
        echo "$(date +%F' '%T) - INFO: client更新完成" |tee -a $logPWD/update.log
    else
        echo "$(date +%F' '%T) - INFO: client未发现，跳过更新" |tee -a $logPWD/update.log
    fi
}

cross_update(){
    cd $sourcePWD 
    if [ -d $serverPWD/crosscenter ];then
        echo "$(date +%F' '%T) - INFO: crosscenterserver更新开始" |tee -a $logPWD/update.log
        if [[ -n "`ls cross*`" ]]; then	
            chmod +x ./crosscenter/crosscenterserver
        fi
        rsync -avz  ./crosscenter/crosscenterserver  $$serverPWD/crosscenter/
		rsync -avz  gamedb.dat maps.json  $$serverPWD/crosscenter/config/
        #/usr/local/pythonbrew/pythons/Python-2.7.3/bin/supervisorctl restart hrl_cross
         echo "$(date +%F' '%T) - INFO: cross更新完成" |tee -a $logPWD/update.log
    else
        echo "$(date +%F' '%T) - INFO: cross未发现，跳过更新" |tee -a $logPWD/update.log
    fi
}

login_update(){
     cd $sourcePWD 
    if [ -d $serverPWD/login/ ];then
        echo "$(date +%F' '%T) - INFO: login更新开始" |tee -a $logPWD/update.log
        if [[ -n "`ls login*`" ]]; then	
            chmod +x ./login/loginserver
        fi
        rsync -avz ./login/loginserver  $serverPWD/login/
        #/usr/local/pythonbrew/pythons/Python-2.7.3/bin/supervisorctl restart hrl_login
        echo "$(date +%F' '%T) - INFO: login更新完成" |tee -a $logPWD/update.log
    else
        echo "$(date +%F' '%T) - INFO: login未发现，跳过更新" |tee -a $logPWD/update.log
    fi
}

fightcenter_update(){
     cd $sourcePWD 
    if [ -d $serverPWD/fightcenter/ ];then
        echo "$(date +%F' '%T) - INFO: fightcenter更新开始" |tee -a $logPWD/update.log
        if [[ -n "`ls fightcenter*`" ]]; then	
            chmod +x ./fightcenter/fightcenterserver
        fi
        rsync -avz ./fightcenter/fightcenterserver  $serverPWD/fightcenter/
        echo "$(date +%F' '%T) - INFO: hrl_fightcenter更新完成" |tee -a $logPWD/update.log
    else
        echo "$(date +%F' '%T) - INFO: hrl_fightcenter未发现，跳过更新" |tee -a $logPWD/update.log
    fi
}

gs_update(){
    cd /opt/data/
    chmod +x ./gs/*server
    echo "$(date +%F' '%T) - INFO: gs server更新开始" |tee -a $logPWD/update.log
    for i in `ls -d $serverPWD/gs*`
    do
        rsync -avz   ./gs/*server $i
		rsync -avz  gamedb.dat maps.json  $i/config/
        echo "更新服务器,$i" |tee -a $logPWD/update.log
    done
    echo "$(date +%F' '%T) - INFO: gs server更新完成" |tee -a $logPWD/update.log
}

fightcross_update(){
    cd /opt/data/
	chmod +x ./fightcross/fightserver
    echo "$(date +%F' '%T) - INFO: fightcross更新开始" |tee -a $logPWD/update.log
	for i in `ls -d $serverPWD/fightcross*`
    do
        rsync -avz   ./fightcross/fightserver $i
		rsync -avz  gamedb.dat maps.json  $i/config/
        echo "更新fightcross服务器,$i" |tee -a $logPWD/update.log
    done    
    echo "$(date +%F' '%T) - INFO: fightcross更新完成" |tee -a $logPWD/update.log
}

client_update
gs_update
cross_update
login_update
fightcross_update
fightcenter_update