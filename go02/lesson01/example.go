package main

import (
	"log"
)

func main() {
	names := []string{
		"lobster",
		"sea urchin",
		"sea cucumber",
	}

	i:= 7

	OutOfRange(names,i)

//	str:=returnArrEl(names,i)
//	fmt.Println(str)
}

func returnArrEl(arr []string, el int) (string){

	return arr[el]

}

func OutOfRange(arr []string, el int)  {

	defer func() {
		if err := recover(); err != nil {
		log.Println("panic occurred:", err)

		}
	}()
	returnArrEl(arr, el)

}