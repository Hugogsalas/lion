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

//ExecuteQuery : Metodo de execucion de un query que retorna objetos
func ExecuteQuery(command string) (*sql.Rows, error) {
	db, err := sql.Open("mysql", "root:3$trella@tcp(127.0.0.1:3306)/lioness")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result, err := db.Query(command)
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
func GetObject(table []string, selects []string, params []string, values []interface{}) (*sql.Rows, error) {
	var command string

	command += "SELECT "
	if len(selects) == 0 {
		command += "* "
	} else {
		for i := 0; i < len(selects); i++ {
			if values[i] != nil {
				if i == (len(selects) - 1) {
					command += selects[i] + ","
				} else {
					command += selects[i]
				}
			}
		}
	}

	command += " FROM "
	for i := 0; i < len(table); i++ {
		if i == (len(selects) - 1) {
			command += table[i] + ","
		} else {
			command += table[i]
		}
	}

	where := false

	multiple := false

	for i := 0; i < len(values); i++ {
		if values[i] != nil {
			if where == false {
				command += " WHERE "
				where = true
			}
			if multiple {
				command += " AND "
			}
			var varType string = fmt.Sprintf("%T", values[i])
			var varString string = fmt.Sprintf("%v", values[i])
			if varType != "string" {
				command += params[i] + "=" + varString
			} else {
				command += params[i] + "='" + varString + "'"
			}
			multiple = true
		}
	}
	result, err := ExecuteQuery(command)

	if err != nil {
		return nil, err
	}

	return result, nil
}

//UpdateObject : metodo que actualiza parametros de objetos por filtros
func UpdateObject(table string, selects []string, filters []interface{}, params []string, values []interface{}) (bool, error) {
	var command string

	command += "UPDATE " + table + " SET "

	comma := false

	for i := 0; i < len(values); i++ {
		if values[i] != nil {
			var varType string = fmt.Sprintf("%T", values[i])
			var varString string = fmt.Sprintf("%v", values[i])
			if comma == false {
				comma = true
			} else {
				command += ","
			}
			if varType != "string" {
				command += params[i] + " = " + varString
			} else {
				command += params[i] + " = '" + varString + "'"
			}
		}
	}

	where := false

	for i := 0; i < len(filters); i++ {
		if filters[i] != nil {
			if where == false {
				command += " WHERE "
				where = true
			} else {
				command += " AND "
			}
			var varType string = fmt.Sprintf("%T", filters[i])
			var varString string = fmt.Sprintf("%v", filters[i])
			if varType != "string" {
				command += selects[i] + "=" + varString
			} else {
				command += selects[i] + "='" + varString + "'"
			}
		}
	}
	fmt.Println(command)
	result, err := ExecuteQuery(command)

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

//CallStorageProcedure : funcion que lla a un procedimiento almacenado
func CallStorageProcedure(Name string, Values []interface{}) (*sql.Rows, error) {
	command := "call " + Name + "("

	for i := 0; i < len(Values); i++ {
		var varType string = fmt.Sprintf("%T", Values[i])
		var varString string = fmt.Sprintf("%v", Values[i])
		if varType == "string" {
			command += "'" + varString + "'"
		} else if Values[i] == nil {
			command += "null"
		} else {
			command += varString
		}
		if i != (len(Values) - 1) {
			command += ","
		}

	}

	command += ")"

	result, err := ExecuteQuery(command)

	if err != nil {
		return nil, err
	}

	return result, nil

}
