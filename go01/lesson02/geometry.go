package main

import "fmt"
import "math"

func main()  {

	var in int

	fmt.Print("Второй урок \n1. Площадь прямоугольника \n2. Диаметр и длина окружности\n3. Трёхзначное число\nВведите 1, 2 или 3: ")
	fmt.Scanln(&in)
	switch in {
	case 1: rectangle()
	case 2: сircle()
	case 3: threeDigitNumber()
	default: fmt.Println("Не выбран ни один вариант, запустите повторно")
	}
}

func rectangle() {
	var a, b float64

	fmt.Println("Введите длину сторон прямоугольника:")
	fmt.Scanln(&a, &b)
	fmt.Println("Площадь прямоугольника:",a*b)
}

func сircle() {
	var circleArea float64
	fmt.Println("Введите площадь окружности:")
	fmt.Scanln(&circleArea)
	fmt.Println("Диаметр круга:",math.Sqrt((circleArea*4)/math.Pi))
	fmt.Println("Длина окружности:",math.Sqrt((circleArea*4*math.Pi)))
}

func threeDigitNumber() {
	var number int
	fmt.Println("Введите трёхзначное число:")
	fmt.Scanln(&number)

	fmt.Println("Количество сотен:",number/100)
	fmt.Println("Количество десятков:",number/10)
	fmt.Println("Количество единиц:",number/1)
}
