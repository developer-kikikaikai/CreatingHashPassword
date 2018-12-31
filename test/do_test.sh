VALUE=(`ruby parse.rb`)
DBNAME=${VALUE[0]}
USER=${VALUE[1]}
PASS=${VALUE[2]}
#import DB
mysql -u${USER} -p${PASS} -e "CREATE DATABASE ${DBNAME}"
mysql -u${USER} -p${PASS} ${DBNAME} < ../db/dumpfile.backup
./insert_db.sh ${USER} ${PASS} ${DBNAME}
go run unittest.go
./cleanup_db.sh ${USER} ${PASS} ${DBNAME}
mysql -u${USER} -p${PASS} -e "DROP DATABASE ${DBNAME}"
