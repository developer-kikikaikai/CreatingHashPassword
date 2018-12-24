mysql -u root -p$1 -D $2 -e "INSERT INTO account_tbl values('test', 'testpass')"
mysql -u root -p$1 -D $2 -e "INSERT INTO account_tbl values('test2', 'testpass2')"
mysql -u root -p$1 -D $2 -e "INSERT INTO password_tbl values('test', 'sha256', 'seed1', 'testtitle1')"
mysql -u root -p$1 -D $2 -e "INSERT INTO password_tbl values('test', 'sha512', 'seed2', 'testtitle2')"
