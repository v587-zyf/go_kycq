#!/bin/bash

mysql="/usr/local/mysql/bin/mysql"
mysql="mysql"
branch="branch"
host="127.0.0.1"
user="root"
pass=""
port=3306
source="*.sql"
mainsource="*.sql"
db="database"
tmpfile=$(mktemp /tmp/sql.XXXXXX)
verbose=false

help() {
    echo "Usage: $(basename "$0") options ()"
}


if [ $# -eq 0 ]
then
    help
fi


while getopts ":b:h:u:p:t:n:s:m:d:v" opt; do
    case  $opt in
        # b) branch=$OPTARG;;
        h) host=$OPTARG;;
        u) user=$OPTARG;;
        p) pass=$OPTARG;;
        t) port=$OPTARG;;
        n) name=$OPTARG;;
        s) source=$OPTARG;;
        # m) mainsource=$OPTARG;;
        d) db=$OPTARG;;
        v) verbose=true;;
        h) help ;;
        :) echo "Error: option ${OPTARG} requires an argument";;
        ?) echo "Invalid option: ${OPTARG}";;
    esac
done

shift $(($OPTIND - 1))

if [ $verbose = true ]
then
    echo $user,$pass,$port,$name,$sql,$verbose
    echo "Remaining args are: $*"
fi

# echo $user,$pass,$port,$name,$host,$db,$tmpfile
python dbsync/dbsync.py --host $host --port $port --user $user --pwd $pass --db $db -o $tmpfile
python dbsync/dbsync.py -s $source -t $tmpfile --db $db
cat output.sql
$mysql -h$host -u$user -P$port -p$pass $db < output.sql

# exit 0
# if [ "develop" != $branch ]; then
#     exit 0
# fi

# python dbsync/dbsync.py --host $host --port $port --user $user --db game-ldj-main -o $tmpfile
# python dbsync/dbsync.py -s $mainsource -t $tmpfile
# cat output.sql
# $mysql -uroot game-ldj-main < output.sql
