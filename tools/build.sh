#!/bin/bash

help() {
    echo "Usage: $(basename "$0") "
    # exit 1
}


BASEDIR=$(dirname "$0")
set -x
go build -o $BASEDIR/../apps/gameserver/gameserver gitlab.hd.com/yulong/server/apps/gameserver
set +x
pwd
