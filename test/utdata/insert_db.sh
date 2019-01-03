mysql -u $1 -p$2 -D $3 -e "INSERT INTO account_tbl values('test', 'testpass')"
mysql -u $1 -p$2 -D $3 -e "INSERT INTO account_tbl values('test2', 'testpass2')"
mysql -u $1 -p$2 -D $3 -e "INSERT INTO passphrase_tbl values('test', 'sha256', 'seed1', 'testtitle1')"
mysql -u $1 -p$2 -D $3 -e "INSERT INTO passphrase_tbl values('test', 'sha512', 'seed2', 'testtitle2')"
