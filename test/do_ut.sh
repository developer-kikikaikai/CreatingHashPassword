#!/bin/bash
(cd db && ./install_db.sh)
source load_setting.sh
#update DB
./utdata/insert_db.sh ${USER} ${PASS} ${DBNAME}
go run unittest.go
(cd db;./cleanup_db.sh)
