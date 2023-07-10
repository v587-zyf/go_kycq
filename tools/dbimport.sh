#!/bin/bash


# export from this databases
from_host="127.0.0.1"
from_user="root"
from_pass=""
from_port=3306
from_name=""

# import into this databases
to_host="localhost"
to_user="root"
to_pass=""
to_port=3306
to_name=""

verbose=false
# temp sql file
tmpfile=$(mktemp /tmp/sql.XXXXXX)


help() {
    echo "Usage: $(basename "$0") options ()"
}


if [ $# -eq 0 ]
then
    help
fi

while getopts ":h:u:p:t:n:H:U:P:T:N:v" opt; do
    case  $opt in
        h) from_host=$OPTARG;;
        u) from_user=$OPTARG;;
        p) from_pass=$OPTARG;;
        t) from_port=$OPTARG;;
        n) from_name=$OPTARG;;
        H) to_host=$OPTARG;;
        U) to_user=$OPTARG;;
        P) to_pass=$OPTARG;;
        T) to_port=$OPTARG;;
        N) to_name=$OPTARG;;
        v) verbose=true;;
        h) help ;;
        :) echo "Error: option ${OPTARG} requires an argument";;
        ?) echo "Invalid option: ${OPTARG}";;
    esac
done

shift $(($OPTIND - 1))


to_backup=$to_name$(date +%D_%T)
if [ $verbose = true ]
then
    echo "backup database: " $to_backup
    echo "temp file: "$tmpfile
    echo "from this: "$from_host,$from_user,$from_pass,$from_port,$from_name
    echo "to this: "$to_host,$to_user,$to_pass,$to_port,$to_name
    echo "Remaining args are: $*"
fi

echo "backuping the old database"
mysqldump -h$to_host -u$to_user -p$to_pass -P$to_port $to_name > $tmpfile
mysql -u$to_user -p$to_pass -P$to_port -e "DROP DATABASE IF EXISTS $to_name; create database $to_name"
mysql -u$to_user -p$to_pass -P$to_port -e "DROP DATABASE IF EXISTS \`$to_backup\`; create database \`$to_backup\`"
mysql -u$to_user -p$to_pass -P$to_port $to_backup < $tmpfile

echo "exporting from the target database"
mysqldump -h$from_host -u$from_user -p$from_pass -P$from_port $from_name > $tmpfile
echo "creating the new database"
mysql -u$to_user -p$to_pass -P$to_port $to_name < $tmpfile
