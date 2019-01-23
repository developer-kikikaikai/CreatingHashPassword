#!/bin/bash
source load_setting.sh
mysql -u${USER} -p${PASS} -e "DROP DATABASE ${DBNAME}"
