mysql -u $1 -p$2 -D $3 -e "DELETE FROM account_tbl"
mysql -u $1 -p$2 -D $3 -e "DELETE FROM passphrase_tbl"
