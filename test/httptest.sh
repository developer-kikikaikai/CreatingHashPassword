#!/bin/bash
VALUE=(`ruby parse.rb`)
DBNAME=${VALUE[0]}
USER=${VALUE[1]}
PASS=${VALUE[2]}
#import DB
mysql -u${USER} -p${PASS} -e "CREATE DATABASE ${DBNAME}"
mysql -u${USER} -p${PASS} ${DBNAME} < ../db/dumpfile.backup
./httpdata/insert_db.sh ${USER} ${PASS} ${DBNAME}

#start main process
cd ../
go run main.go &
while [ true ];
do
	#this test can run with no go process
	http_pid=`ps aux | grep go-build | grep -v grep | awk -F" " '{print $2}'`
	if [ "x$http_pid" != "x" ]; then
		break
	fi
	sleep 1
done
cd -

#do test by using ruby
cd ./httpdata
ruby httptest.rb
cd -
#exit
kill $http_pid
./httpdata/cleanup_db.sh ${USER} ${PASS} ${DBNAME}
mysql -u${USER} -p${PASS} -e "DROP DATABASE ${DBNAME}"
