#!/bin/bash

SERVER="gemdale-api"
BASE_DIR=$PWD
INTERVAL=1
DATE=$(date "+%Y%m%d")

# 命令行参数，需要手动指定
ARGS=""

function start()
{
	if [ "`pgrep $SERVER -f -u $UID`" != "" ];then
		echo "$SERVER already running"
		exit 1
	fi

	go install
	go-assets-builder view/templates -o assets.go
	nohup $GOPATH/bin/chuanghehui_business >> log/"$DATE""_.log" 2>&1 &

	echo "starting..." &&  sleep $INTERVAL

	# check status
	if [ "`pgrep $SERVER -f -u $UID`" == "" ];then
		echo "$SERVER start failed"
		exit 1
	fi
}

function status()
{
    echo pgrep $SERVER -u $UID
	if [ "`pgrep $SERVER -f -u $UID`" != "" ];then
		echo $SERVER is running
	else
		echo $SERVER is not running
	fi
}

function stop()
{
	if [ "`pgrep $SERVER -f -u $UID`" != "" ];then
		kill -2 `pgrep $SERVER -f -u $UID`
	fi

	echo "sleeping..." &&  sleep $INTERVAL

	if [ "`pgrep $SERVER -f -u $UID`" != "" ];then
		echo "$SERVER stop failed"
		exit 1
	fi
}

function restart()
{
  go-assets-builder view/templates -o assets.go
	go install

	if [ "`pgrep $SERVER -f -u $UID`" != "" ];then
		kill -2 `pgrep $SERVER -f -u $UID`
	fi

	echo "sleeping..." &&  sleep $INTERVAL
	nohup $GOPATH/bin/chuanghehui_business >> log/"$DATE""_.log" 2>&1 &
	echo "starting..."

	# check status
	if [ "`pgrep $SERVER -f -u $UID`" == "" ];then
		echo "$SERVER start failed"
		exit 1
	fi
}

case "$1" in
	'start')
	start
	;;
	'docker')
	dockerMake
	;;
	'stop')
	stop
	;;
	'status')
	status
	;;
	'restart')
	restart
	;;
	*)
	echo "usage: $0 {start|stop|restart|status}"
	exit 1
	;;
esac

