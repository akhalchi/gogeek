package lesson10

import (
	"errors"
	"strconv"
	"strings"
)

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

func calc(arg1,arg2 float64, operation string)  (float64,error) {
	switch operation {
	case "*":
		return arg1 * arg2, nil
	case "/":
		if arg2 == 0.0 {
			return 0.0, errors.New("На ноль делить нельзя")
		}
		return arg1 / arg2, nil
	case "+":
		return arg1 + arg2, nil
	case "-":
		return arg1 - arg2, nil
	default:
		return 0.0, errors.New("Функции calc не передана арифметическая операция")
	}

}