mysql -u root -p$1 -D $2 -e "DELETE FROM account_tbl"
mysql -u root -p$1 -D $2 -e "DELETE FROM password_tbl"
