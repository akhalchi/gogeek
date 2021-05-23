package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите последовательность целых чисел: ")
	scanner.Scan()
	input :=scanner.Text()

	arrIn:= strings.Split(input, " ")
	fmt.Println(intSort(parseStringArrayToInt(arrIn)))

}


func intSort(arr []int) []int{
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for ; j >= 1 && arr[j-1] > x; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = x
	}
	return arr
}

func parseStringArrayToInt(arrStr []string)  []int{

	var arrInt = []int{}
	for _, i := range arrStr {
		j,err:=strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		arrInt = append(arrInt, j)
	}
	return arrInt
}