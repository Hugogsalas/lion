package utilities

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	models "../Models"
	_ "github.com/go-sql-driver/mysql"
)

//ExecuteCommand : Metodo de execucion de un query que no retorna nada
func ExecuteCommand(command string) (sql.Result, error) {
	fmt.Println(command)
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
	fmt.Println(command)
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
func InsertObject(table string, values []interface{}, fields []string) (int64, error) {
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
			if valueString == "<nil>" {
				command += "null"
			} else {
				command += valueString
			}
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
		return 0, err
	}
	IDinsert, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}
	return IDinsert, nil
}

//GetObject : metodo que retorna un objeto segun parametros
func GetObject(Query models.GetQuery) (*sql.Rows, error) {
	var command string

	command += "SELECT "
	if len(Query.Selects) == 0 {
		command += "* "
	} else {
		for i := 0; i < len(Query.Selects); i++ {
			for j := 0; j < len(Query.Selects[i]); j++ {
				if Query.Selects[i][j] != "" {
					if i == (len(Query.Selects)-1) && j == (len(Query.Selects[i])-1) {
						command += Query.Tables[i] + "." + Query.Selects[i][j]
					} else {
						command += Query.Tables[i] + "." + Query.Selects[i][j] + ","
					}
				}
			}
		}
	}

	command += " FROM "
	for i := 0; i < len(Query.Tables); i++ {
		if i == (len(Query.Tables) - 1) {
			command += Query.Tables[i]
		} else {
			command += Query.Tables[i] + ","
		}
	}

	where := false

	for i := 0; i < len(Query.Values); i++ {
		for j := 0; j < len(Query.Values[i]); j++ {
			if Query.Values[i][j] != nil {
				if !where {
					command += " WHERE "
					where = true
				} else {
					command += " AND "
				}
				var varType string = fmt.Sprintf("%T", Query.Values[i][j])
				var varString string = fmt.Sprintf("%v", Query.Values[i][j])
				if varType != "string" {
					command += Query.Tables[i] + "." + Query.Params[i][j] + "=" + varString
				} else {
					command += Query.Tables[i] + "." + Query.Params[i][j] + "='" + varString + "'"
				}
			}
		}
	}

	for i := 0; i < len(Query.Conditions); i++ {
		if i == 1 && !where {
			command += " WHERE "
		} else {
			command += " AND "
		}
		command += Query.Conditions[i]
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

	if !strings.Contains(command, "WHERE") {
		return false, errors.New("No se puede actualizar sin filtrado")
	}

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

//DeleteObject : metodo que elimina objetos por filtros
func DeleteObject(table string, selects []string, filters []interface{}) (bool, error) {
	var command string

	command += "DELETE FROM " + table

	comma := false

	for i := 0; i < len(selects); i++ {
		if filters[i] != nil {
			var varType string = fmt.Sprintf("%T", filters[i])
			var varString string = fmt.Sprintf("%v", filters[i])
			if comma == false {
				command += " WHERE "
				comma = true
			} else {
				command += " AND "
			}
			if varType != "string" {
				command += selects[i] + " = " + varString
			} else {
				command += selects[i] + " = '" + varString + "'"
			}
		}
	}

	if !strings.Contains(command, "WHERE") {
		return false, errors.New("No se puede borrar sin filtrado")
	}

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
