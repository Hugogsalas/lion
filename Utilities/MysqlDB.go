package utilities

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//ExecuteCommand : Metodo de execucion de un query que no retorna nada
func ExecuteCommand(command string) (interface{}, error) {
	db, err := sql.Open("mysql", "root:ainerbds117@@tcp(127.0.0.1:3306)/lioness")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result, err := db.Exec(command)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//InsertObject : inserta un objeto en la tabla especificada
func InsertObject(table string, values []interface{}, fields []string) (bool, error) {
	var command string

	command = "INSERT INTO " + table + " ("
	for i := 0; i < len(fields); i++ {
		command += fields[i]
		if i == (len(fields) - 1) {
			command += ")"
		} else {
			command += ","
		}
	}
	command += " VALUES ("
	for i := 0; i < len(values); i++ {
		var valueString string = fmt.Sprintf("%v", values[i])
		var typeVar string = fmt.Sprintf("%T", values[i])
		if typeVar != "string" {
			command += valueString
		} else {
			if valueString == "" {
				command += "null"
			} else {
				command += "'" + valueString + "'"
			}
		}
		if i == (len(fields) - 1) {
			command += ")"
		} else {
			command += ","
		}
	}
	result, err := ExecuteCommand(command)
	if err != nil {
		return false, err
	} else {
		if result != nil {
			return true, nil
		} else {
			return false, nil
		}
	}
}

//GetObject : metodo que retorna un objeto segun parametros
func GetObject(table string, selects []string, params []string, values []interface{}) (bool, interface{}) {
	var command string

	command += "SELECT "
	if len(selects) == 0 {
		command += "* "
	} else {
		for i := 0; i < len(selects); i++ {
			if i == (len(selects) - 1) {
				command += selects[i] + ","
			} else {
				command += selects[i]
			}
		}
	}

	command += " FROM " + table + " WHERE "

	for i := 0; i < len(values); i++ {
		var varType string = fmt.Sprintf("%T", values[i])
		var varString string = fmt.Sprintf("%v", values[i])
		if varType != "string" {
			command += params[i] + "=" + varString
		} else {
			command += params[i] + "='" + varString + "'"
		}
		if i != (len(values) - 1) {
			command += " AND "
		}
	}

	fmt.Println(command)

	return true,nil
}
