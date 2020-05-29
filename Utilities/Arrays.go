package utilities

//ObjectinArray : Metodo que te responde si un objeto se encuntra en un array
func ObjectinArray(array []interface{}, object interface{}) bool{
	for i := 0; i < len(array); i++ {
		if array[i]==object{
			return true
		}
	}
	return false
}

//Ιndexof : Metodo que te responde el indice de un objeto si se encuntra en un array
func Ιndexof(array []interface{}, object interface{}) int{
	for i := 0; i < len(array); i++ {
		if array[i]==object{
			return i
		}
	}
	return -1
}

