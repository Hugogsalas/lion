package utilities

import (
	"reflect"
)

//ObjectValues : Metodo que returna un array de los valores de un objeto
func ObjectValues(object interface{}) []interface{} {
	e := reflect.ValueOf(object)
	var arraybucket []interface{} 
	for i := 0; i < e.NumField(); i++ {
		arraybucket=append(arraybucket,e.Field(i))
	}
	return arraybucket
}

//ObjectFields : Metodo que retorna un array de los valores 
func ObjectFields(object interface{}) []string {
	e := reflect.ValueOf(object)
	var arraybucket []string 
	for i := 0; i < e.NumField(); i++ {
		arraybucket=append(arraybucket,e.Type().Field(i).Name)
	}
	return arraybucket
}
