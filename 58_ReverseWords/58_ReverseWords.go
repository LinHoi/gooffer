package main

import "fmt"

func ReverseWords(str string) string {
	strArray := make([]uint8,len(str))
	for i := range str{
		strArray[i] = str[i]
	}
	ReverseWord(strArray)
	low , high := 0,0
	for high < len(str) {
		if strArray[high] != ' '  {
			high ++
		} else {
			ReverseWord(strArray[low:high])
			low= high + 1
			high ++
		}
	}
	ReverseWord(strArray[low:])
	return string(strArray)

}

func ReverseWord(str []uint8) {
	if len(str) <= 1 {
		return
	}
	low , high := 0, len(str)-1
	for low < high {
		str[low], str[high] = str[high], str[low]
		low ++
		high --
	}
}

func main(){
	strings := []string{
		"hello world.",
		"I have a dream",
		"who  are you?",
		"",
	}
	for _,str := range strings {
		res := ReverseWords(str)
		fmt.Printf("reverse of '%s' is '%s'\n",str, res)
	}
}
/* output
reverse of 'hello world.' is 'world. hello'
reverse of 'I have a dream' is 'dream a have I'
reverse of 'who  are you?' is 'you? are  who'
reverse of '' is ''
 */
