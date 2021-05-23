package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

)

type expression struct {
	arg1    float64
	arg2 	float64
	op string
}

func main()  {

	var in string


	fmt.Println("Консольный калькулятор. Введите операцию:")
	fmt.Scanln(&in)
	operation, indexOperation, err:=checkOperation(in)

	if err == nil {

		num1,err1:=parseArg(in[:indexOperation])
		num2,err2:=parseArg(in[indexOperation+1:])
		g:=expression{arg1: num1,arg2: num2,op: operation}

		if err1 == nil && err2 == nil {
			result,err3:=calc(g)
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

func calc(exp expression)  (float64,error){
	switch exp.op {
	case "*":
		return exp.arg1 * exp.arg2,nil
	case "/":
		if exp.arg2 == 0.0 {
			return 0.0, errors.New("На ноль делить нельзя")
		}
		return exp.arg1 / exp.arg2,nil
	case "+":
		return exp.arg1 + exp.arg2,nil
	case "-":
		return exp.arg1 - exp.arg2,nil
	default:
		return 0.0, errors.New("Функции calc не передана арифметическая операция")
	}

}
