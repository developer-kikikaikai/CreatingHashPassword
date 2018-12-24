package db

import (
	"errors"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var account_tbl = "account_tbl"
var pass_tbl = "password_tbl"

type Account struct {
	Username string
	Password string
}

type PasswordInfo struct {
	Username string
	Title string
	Algorithm string
	Seed string
}

func sql_command(query string) ([]map[string]string,error) {
	var dbname = "creatHashPassword"
	var mysql_user = "root"
	var mysql__password = ""

	db, err := sql.Open("mysql", mysql_user + ":" + mysql__password + "@/" + dbname)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()//close after return

	var results []map[string]string = make([]map[string]string, 0)
	fmt.Printf("query:%s\n", query)
	rows, err := db.Query(query)
	if err != nil {
		return results, errors.New("failed to call query")
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	/*update result*/
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var result map[string]string = make(map[string]string)
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col != nil {
				result[columns[i]]=string(col)
			}
		}
		results = append(results, result)
	}
	return results, nil
}

func GetAccount(username string) (Account, error) {
	result := Account{"",""}
	res , err:= sql_command("SELECT * FROM " + account_tbl + " WHERE username='" + username + "'")
	if err != nil {
		return result,  err
	}

	if len(res) != 1 {
		return result, errors.New("No account")
	}

	var ok bool
	result.Username, ok = res[0]["username"]
	if ok {
		result.Password = res[0]["password"]
	}
	return result, nil
}

func SetAccount(info Account) bool {
	var err error
	_, err = GetAccount(info.Username)
	if err != nil {
		//insert data
		_, err = sql_command("INSERT INTO " + account_tbl + " value('" + info.Username + "', '" + info.Password + "')")
	} else {
		//update data
		_, err = sql_command("UPDATE " + account_tbl + " SET password='" + info.Password + "' WHERE username='" + info.Username+ "'")
	}
	if err != nil {
		return false
	}

	return true
}

func DeleteAccount(username string) bool {
	_, err := sql_command("DELETE FROM " + account_tbl + " WHERE username='" + username + "'")
	return err == nil
}

func GetAllPassword(username string) ([]PasswordInfo, error) {
	var results []PasswordInfo = make([]PasswordInfo, 0)
	res , err := sql_command("SELECT * FROM " + pass_tbl + " WHERE username='" + username + "'")
	if err != nil {
		return results, err
	}

	for i := 0; i < len(res); i++  {
		result := PasswordInfo{"","","",""}
		result.Username = res[i]["username"]
		result.Title = res[i]["title"]
		result.Algorithm = res[i]["algorithm"]
		result.Seed = res[i]["seed"]
		results = append(results, result)
	}
	return results, nil
}

func GetPasswordInfo(username string, title string) (PasswordInfo,error) {
	var result PasswordInfo =PasswordInfo{"","","",""}

	res, err:= sql_command("SELECT * FROM " + pass_tbl + " WHERE username='" + username + "' and title='" + title + "'")
	if err != nil {
		return result, err
	}

	if len(res) != 1 {
		return result, errors.New("No password")
	}

	result.Username = res[0]["username"]
	result.Title = res[0]["title"]
	result.Algorithm = res[0]["algorithm"]
	result.Seed = res[0]["seed"]
	return result, nil
}

func SetPasswordInfo(info PasswordInfo) bool {
	res, err := GetPasswordInfo(info.Username, info.Title)
	if res.Username != info.Username || err != nil {
		//insert data
		_,err = sql_command("INSERT INTO " + pass_tbl + " value('" + info.Username + "', '" + info.Algorithm + "', '" + info.Seed + "', '" + info.Title + "')")
	} else {
		//update data
		_, err = sql_command("UPDATE " + pass_tbl + " SET algorithm='" + info.Algorithm + "', seed='" + info.Seed + "' WHERE username='" + info.Username + "' and title='" + info.Title + "'")
	}
	return err == nil
}

func DeletePasswordInfo(username string, title string) bool {
	_,err := sql_command("DELETE FROM " + pass_tbl + " WHERE username='" + username + "' and title='" + title + "'")
	return err == nil
}
