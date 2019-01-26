#!/bin/bash
source load_setting.sh
#import DB
mysql -u${USER} -p${PASS} -e "CREATE DATABASE ${DBNAME}"
mysql -u${USER} -p${PASS} ${DBNAME} < ../db/dumpfile.backup
