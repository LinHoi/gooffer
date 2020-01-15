package main

import (
	"fmt"
	"reflect"
)

func Power(base float64, exp int64) float64 {
	if equal(base,0.0) {
		return 0.0
	}
	if exp == 0 {
		return 1
	}
	if exp < 0 {
		return 1/power2(base,-exp)
	}else{
		return power2(base, exp)
	}
}
var miniNum = 0.0000001
func equal(a,b float64) bool {
	return (a-b) < miniNum && (a-b) > (-1)*miniNum
}

func power(base float64, exp int64) float64 {
	if exp == 0 {
		return 0.0
	}
	if exp == 1 {
		return base
	}

	result := power(base, exp >> 1 )
	result *= result
	if exp & 1 == 1 {
		result *= base
	}
	return  result
}

func power2(base float64, exp int64) float64 {
	result := 1.0
	for i := int64(1) ; i <= exp ; i++ {
		result *= base
	}
	return result
}
func main(){
	nums := [][2]interface{}{
		{0.0, 0},
		{0.0,1},
		{0.1, 0},
		{0.01, -1},
		{0.01, 9},
		{99.2, 3},
	}
	for _, num := range nums {
		result := Power(reflect.ValueOf(num[0]).Float(),reflect.ValueOf(num[1]).Int())
		fmt.Printf("%f Power %d equals %f \n",num[0],num[1], result)
	}
}

/* output
0.000000 Power 0 equals 0.000000
0.000000 Power 1 equals 0.000000
0.100000 Power 0 equals 1.000000
0.010000 Power -1 equals 100.000000
0.010000 Power 9 equals 0.000000
99.200000 Power 3 equals 976191.488000
*/