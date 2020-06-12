package models

//GetQuery : Modelo de GetQuery (modelo para hacer consultas)
type GetQuery struct {
	Tables     []string
	Selects    [][]string
	Params     [][]string
	Values     [][]interface{}
	Conditions []string
}
