#!/bin/bash
(cd db && ./install_db.sh)
source load_setting.sh

./httpdata/insert_db.sh ${USER} ${PASS} ${DBNAME}

#start main process
cd ../
go run main.go &
while [ true ];
do
	#this test can run with no go process
	http_pid=`ps aux | grep go-build | grep -v grep | awk -F" " '{print $2}'`
	http_pname=`ps aux | grep go-build | grep -v grep | awk -F" " '{print $11}' | grep go-build`
	if [ "x$http_pname" != "x" -a "x$http_pid" != "x" ]; then
		break
	fi
	sleep 1
done
cd -

ps aux | grep go
sleep 10

#do test by using ruby
cd ./httpdata
ruby httptest.rb
cd -
#exit
kill $http_pid
(cd db;./cleanup_db.sh)
