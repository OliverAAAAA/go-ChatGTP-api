#!/bin/bash
ss="go-chatgpt-api"

echo "stop ${ss}..."

pid=$(ps -ef | grep $ss |grep -v grep  | awk '{print $2}')
for item in $pid
do
        echo "kill -9 $item"
        kill -9 $item
done
echo "stop ${ss} success."
echo "start ${ss} ing..."
nohup ./${ss} &
echo "start ${ss} success..."
tail ./nohup.out -f