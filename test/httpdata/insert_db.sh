PASSHASH=`echo -n "test:CreatingHashPassword:testpass" | md5sum | awk -F" " '{print $1}'`
mysql -u $1 -p$2 -D $3 -e "INSERT INTO account_tbl values('test', '${PASSHASH}')"
PASSHASH=`echo -n "test2:CreatingHashPassword:testpass2" | md5sum | awk -F" " '{print $1}'`
mysql -u $1 -p$2 -D $3 -e "INSERT INTO account_tbl values('test2', '${PASSHASH}')"
mysql -u $1 -p$2 -D $3 -e "INSERT INTO passphrase_tbl values('test', 'sha256', 'seed1', 'testtitle1')"
mysql -u $1 -p$2 -D $3 -e "INSERT INTO passphrase_tbl values('test', 'sha512', 'seed2', 'testtitle2')"
