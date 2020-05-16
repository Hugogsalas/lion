package utilities

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//ExecuteCommand : Metodo de execucion de un query que no retorna nada
func ExecuteCommand(command string) (interface{}, error) {
	db, err := sql.Open("mysql", "root:3$trella@tcp(127.0.0.1:3306)/lioness")
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
		fmt.Println(command)
	}
	command += " VALUES ("
	for i := 0; i < len(values); i++ {
		var valueString string = fmt.Sprintf("%v", values[i])
		command += valueString
		if i == (len(fields) - 1) {
			command += ")"
		} else {
			command += ","
		}
		fmt.Println(command)
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
