DBNAME=creatHashPassphrase
#import DB
mysql -uroot -p$1 -e "CREATE DATABASE ${DBNAME}"
mysql -uroot -p$1 ${DBNAME} < ../db/dumpfile.backup
./insert_db.sh $1 ${DBNAME}
go run test.go
./cleanup_db.sh $1 ${DBNAME}
mysql -uroot -p$1 -e "DROP DATABASE ${DBNAME}"
