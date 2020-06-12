package utilities

import (
	"reflect"
)

//ObjectFields : Metodo que retorna un array de los valores
func ObjectFields(object interface{}) ([]string, []interface{}) {
	e := reflect.ValueOf(object)
	var arrayNames []string
	var arrayValues []interface{}
	for i := 0; i < e.NumField(); i++ {
		Value := CleanValue(e.Field(i).Interface())
		arrayValues = append(arrayValues, Value)
		arrayNames = append(arrayNames, e.Type().Field(i).Name)

	}
	return arrayNames, arrayValues
}

//CleanValue : Metodo que recibe un objeto y lo pasa a nil si este biene con un valor no valido
func CleanValue(value interface{}) interface{} {
	var realValue interface{}
	realValue = nil
	switch value.(type) {
	case int, int8, int16, int32, int64:
		if value != 0 {
			realValue = value
		}
		break
	case string:
		if value != "" {
			realValue = value
		}
		break
	case float32, float64:
		if value != 0.0 {
			realValue = value
		}
		break

	}
	return realValue
}
