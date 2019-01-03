#!/bin/bash
VALUE=(`ruby parse.rb`)
DBNAME=${VALUE[0]}
USER=${VALUE[1]}
PASS=${VALUE[2]}
#import DB
mysql -u${USER} -p${PASS} -e "CREATE DATABASE ${DBNAME}"
mysql -u${USER} -p${PASS} ${DBNAME} < ../db/dumpfile.backup
./utdata/insert_db.sh ${USER} ${PASS} ${DBNAME}
go run unittest.go
./utdata/cleanup_db.sh ${USER} ${PASS} ${DBNAME}
mysql -u${USER} -p${PASS} -e "DROP DATABASE ${DBNAME}"
