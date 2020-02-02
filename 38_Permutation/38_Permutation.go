package main

import (
	"fmt"
	"strings"
)

func Permutation(str string) {
	if len(str) == 0 {
		return
	}
	permutation2(str, 0)
}

func permutation(str string,strIndex int) {
	if strIndex == len(str) - 1 {
		fmt.Println(str)
		return
	}
	strLength := len(str)
	for source := strIndex; source < strLength ; source ++ {
		for destination := source ; destination < strLength  ; destination ++ {
			str = switchString(str,source,destination)
			permutation(str, source + 1 )
			str = switchString(str,source,destination)
		}
	}
}

func permutation2(str string,strIndex int) {
	strLength := len(str)
	for source := strIndex; source < strLength ; source ++ {
		for destination := source+1; destination < len(str); destination ++ {
			str = switchString(str,source,destination)
			permutation(str, source + 1 )
			str = switchString(str,destination,source)
		}
	}
	fmt.Println(str)

}

func switchString(str string, source,destination int) string {
	if source == destination {
		return str
	}
	var b strings.Builder
	b.Grow(len(str))
	for i := range str {
		if i == source {
			b.WriteByte(str[destination])
		} else if i == destination {
			b.WriteByte(str[source])
		} else {
			b.WriteByte(str[i])
		}
	}
	return b.String()
}

func main() {
	str := "123"
	Permutation(str)
}


/*
213
231
321
312
132
123
 */