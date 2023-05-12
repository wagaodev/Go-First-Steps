package main

import "fmt"

type User struct {
	name string
	age uint8
}

func arrayAndSlice() {
	arr := []int8 { 1,2,3 }
	arr = append(arr, 5) // 1,2,3,5 -> tamanho 4, capacidade 8 ( dobrou o tamanho )
	arr = append(arr, 4) // 1,2,3,5,4 -> tamanho 5, capacidade 8 ( continua a quantidade para não ter a necessidade voltar na memória e abrir mais espaço)
	arr = append(arr, 12) // 1,2,3,5,4,12 -> tamanho do slice 6, capacidade do slice em 8 ( continua ainda a quantidade de 8) 
	fmt.Println(arr, len(arr), cap(arr))
}
func mapObj() {
	obj := map[string]string {"name": "luis"}
	fmt.Println(obj)
}

func userStruct() {
	user := User {name: "Wagner", age: 29}
	fmt.Println(user)
}

func main() {
	arrayAndSlice()
	mapObj()
	userStruct()
}