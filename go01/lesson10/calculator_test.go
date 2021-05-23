package lesson10

import (
	"testing"
)

var tests = []struct {
	a    string
	exp1 string
	exp2 int
	exp3 error
}{
	{"4*4", "*", 1, nil},
	{"44+4", "+", 2, nil},
	{"43-22", "-", 2, nil},
	{"444/2", "/", 3, nil},
	//{"444*444", "*", 9, nil}, //ошибочный кейс
	{"444/2", "/", 3, nil},
	//{"3333", "-", 56, nil},// ошибочный кейс
}

func TestCheckOperation(t *testing.T) {
	for _, e := range tests {
		res1,res2,res3 := checkOperation(e.a)
		if res1 != e.exp1 || res2 != e.exp2 || res3 != e.exp3 {
			t.Errorf("checkOperation(%s) = %s,%v,%#v, expected %s,%v,%#v ",
				e.a, res1, res2, res3, e.exp1,e.exp2,e.exp3)
		}
	}

}
func BenchmarkCheckOperation(b *testing.B) {
	for _, e := range tests {
		res1, res2, res3 := checkOperation(e.a)
		if res1 != e.exp1 || res2 != e.exp2 || res3 != e.exp3 {
			b.Errorf("checkOperation(%s) = %s,%v,%#v, expected %s,%v,%#v ",
				e.a, res1, res2, res3, e.exp1, e.exp2, e.exp3)
		}
	}
}