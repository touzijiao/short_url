package utils

/*
	本页面功能是把已经生成的id转换成更短的URL（string字符串）与字符串还原成相应id
*/
import (
	"strconv"
	"strings"
)

var (
	Tokens string //加密字符串
	Length int    //字符串长度
)

func init() {
	//0-9
	for i := 0; i <= 9; i++ {
		Tokens += strconv.Itoa(i)
	}

	//a-z
	for i := 0; i < 26; i++ {
		Tokens += string(byte('a') + byte(i))
	}

	//A-Z
	for i := 0; i < 26; i++ {
		Tokens += string(byte('A') + byte(i))
	}

	Length = len(Tokens)
}

//id加密成字符串
func IdToString(id int) string {
	//不断取模
	var res string
	for id > 0 {
		d := id % Length
		res = string(Tokens[d]) + res
		id /= Length
	}

	return res
}

//字符串还原成id
func StringToId(str string) int {
	var res = 0
	for _, s := range str {
		value := strings.Index(Tokens, string(s))
		res = res*Length + value
	}
	return res
}
