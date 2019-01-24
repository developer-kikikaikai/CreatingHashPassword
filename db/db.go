package db

import (
	"errors"
	_ "fmt"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	_ "github.com/go-sql-driver/mysql"
)

var account_tbl = "account_tbl"
var pass_tbl = "passphrase_tbl"

type Account struct {
	Username string
	Passphrase string
}

type PassphraseInfo struct {
	Username string
	Title string
	Algorithm string
	Seed string
	Length string
	DisableSymbol string
}

//define to read json setting, parameter have to define public
type DBsetting struct {
	DBname string
	User string
	Passphrase string
}

func sql_db_setting() (DBsetting,error) {
	var setting DBsetting
	bytes, err := ioutil.ReadFile("db/dbsetting.json")
	if err != nil { return setting, err }

	err = json.Unmarshal(bytes, &setting);
	return setting,err;
}

func sql_command(query string) ([]map[string]string,error) {
	setting,err := sql_db_setting()
	if err != nil {
		panic(err.Error())
	}

	db, err := sql.Open("mysql", setting.User + ":" + setting.Passphrase + "@/" + setting.DBname)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()//close after return

	results := make([]map[string]string, 0)
	//fmt.Printf("query:%s\n", query)
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
		result.Passphrase = res[0]["passphrase"]
	}
	return result, nil
}

func SetAccount(info Account) bool {
	var err error
	_, err = GetAccount(info.Username)
	if err != nil {
		//insert data
		_, err = sql_command("INSERT INTO " + account_tbl + " value('" + info.Username + "', '" + info.Passphrase + "')")
	} else {
		//update data
		_, err = sql_command("UPDATE " + account_tbl + " SET passphrase='" + info.Passphrase + "' WHERE username='" + info.Username+ "'")
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

func passphraseResult(res *map[string]string) PassphraseInfo{
	result := PassphraseInfo{"","","","","0","false"}
	result.Username = (*res)["username"]
	result.Title = (*res)["title"]
	result.Algorithm = (*res)["algorithm"]
	result.Seed = (*res)["seed"]
	result.Length = (*res)["length"]
	result.DisableSymbol = (*res)["disable_symbol"]
	return result
}

func GetAllPassphrase(username string) ([]PassphraseInfo, error) {
	//var results []PassphraseInfo = make([]PassphraseInfo, 0)
	results := make([]PassphraseInfo, 0)
	res , err := sql_command("SELECT * FROM " + pass_tbl + " WHERE username='" + username + "'")
	if err != nil {
		return results, err
	}

	for i := 0; i < len(res); i++  {
		result := passphraseResult(&res[i])
		results = append(results, result)
	}
	return results, nil
}

func GetPassphraseInfo(username string, title string) (PassphraseInfo,error) {
	result := PassphraseInfo{"","","","","0","false"}
	res, err:= sql_command("SELECT * FROM " + pass_tbl + " WHERE username='" + username + "' and title='" + title + "'")
	if err != nil {
		return result, err
	}

	if len(res) != 1 {
		return result, errors.New("No passphrase")
	}

	result = passphraseResult(&res[0])
	return result, nil
}

func SetPassphraseInfo(info PassphraseInfo) bool {
	res, err := GetPassphraseInfo(info.Username, info.Title)
	if res.Username != info.Username || err != nil {
		//insert data
		_,err = sql_command("INSERT INTO " + pass_tbl + " value('" + info.Username + "', '" + info.Algorithm + "', '" + info.Seed + "', '"  + info.Length + "', '" + info.DisableSymbol + "', '" + info.Title + "')")
	} else {
		//update data
		_, err = sql_command("UPDATE " + pass_tbl + " SET algorithm='" + info.Algorithm + "', seed='" + info.Seed + "', length='"  + info.Length + "', disable_symbol='" + info.DisableSymbol + "' WHERE username='" + info.Username + "' and title='" + info.Title + "'")
	}
	return err == nil
}

func DeletePassphraseInfo(username string, title string) bool {
	_,err := sql_command("DELETE FROM " + pass_tbl + " WHERE username='" + username + "' and title='" + title + "'")
	return err == nil
}
