package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

)

func main()  {

	var in string

	fmt.Println("Консольный калькулятор. Введите операцию:")
	fmt.Scanln(&in)
	operation, indexOperation, err:=checkOperation(in)

	if err == nil {

		num1,err1:=parseArg(in[:indexOperation])
		num2,err2:=parseArg(in[indexOperation+1:])
		if err1 == nil && err2 == nil {
			result,err3:=calc(num1,num2,operation)
				if err3==nil {
					fmt.Println("Результат:",result)
				} else {
					fmt.Println(err3)
				}
			} else {
				fmt.Println(err1)
				fmt.Println(err2)
			}

		}  else {
		fmt.Println(err)
	}

}

func checkOperation(checkString string) (string,int,error){

	var operations=[4]string{"+","-","*","/"}

	for _, i := range operations {
		if strings.Contains(checkString,i)==true {
			return i,strings.Index(checkString, i),nil
		}
	}
	return "",0,errors.New("Отсутствует знак арифметической операции: +,-,*,/")

}

func parseArg(arg string) (float64,error) {

	num, err := strconv.ParseFloat(arg,64)
	if err != nil {
		return 0.0, err
	}
	return num, nil

}

func calc(arg1,arg2 float64, operation string)  (float64,error){
	switch operation {
	case "*":
		return arg1 * arg2,nil
	case "/":
		if arg2 == 0.0 {
			return 0.0, errors.New("На ноль делить нельзя")
		}
		return arg1 / arg2,nil
	case "+":
		return arg1 + arg2,nil
	case "-":
		return arg1 - arg2,nil
	default:
		return 0.0, errors.New("Функции calc не передана арифметическая операция")
	}

}
