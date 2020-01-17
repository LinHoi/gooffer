package main

import "fmt"

func IsNumber2(str string) bool {
	if str == "" {
		return false
	}
	index := 0
	if str[0] == '+' || str[0] == '-' {
		index ++
	}
	countDot := 0
	for index < len(str) && str[index] != 'e' && str[index] != 'E' {
		if str[index] <= '9' && str[index] >= '0' {
			index ++
		}else if str[index] == '.'{
				countDot ++
				if countDot > 1 {
					return false
				}
				index ++
		} else {
					return false
		}
	}
	if index == len(str) {
		return true
	}else if index == (len(str)-1) && (str[index] == 'e' || str[index] == 'E') {
		return false
	}else {
		index ++
	}


	return true
}

func IsNumberic( str string) bool {
	length := len(str)
	index := locateE(str)
	return IsDecimal(str, 0,index -1,) && IsNumber(str,index + 1,length-1)

}
func locateE(str string) int {
	for index := range str {
		if str[index] == 'e' || str[index] == 'E' {
			return index
		}
	}
	return len(str)
}
func IsDecimal(str string, low, high int) bool {
	countdot := 0
	for index := low; index <= high ; index ++ {
		if index == low && (str[index] == '-' || str[index] == '+'){
			continue
		} else if str[index] == '.' {
			countdot ++
			if countdot > 1 { return false }
		}else if str[index] < '0' || str[index] >'9' {
			return false
		} else {
			continue
		}
	}
	return true
}

func IsNumber(str string, low, high int) bool {
	if low > high { return true }
	for index := low; index <= high; index++ {
		if index == low && (str[index] == '-' || str[index] == '+'){
			continue
		}else if str[index] < '0' || str[index] > '9'{
			return false
		}else {
			continue
		}
	}
	return true
}

func main() {
	numbers := []string{
		"+12345",
		"+-123.45",
		"+12.3.4",
		"-123.4e",
		"+123.45e12",
		"12.45e-12",
		"12.3.4e1.2.3",
		"12.34e12.2",
		"12a44h33f.m",
		"123.E.2",
	}
	for _,number := range numbers {
		if IsNumberic(number) {
			fmt.Printf("%s   is a number\n",number)
		}else {
			fmt.Printf("%s   is not a number\n",number)
		}
	}
}

/*OutPut
+12345   is a number
+-123.45   is not a number
+12.3.4   is not a number
-123.4e   is a number
+123.45e12   is a number
12.45e-12   is a number
12.3.4e1.2.3   is not a number
12.34e12.2   is not a number
12a44h33f.m   is not a number
123.E.2   is not a number
*/