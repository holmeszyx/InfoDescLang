package idl

import "os"


// 是否是空白字符
func IsSpace(c byte) bool{
	return c == '\n' || c == '\t' || c == '\f' || c == '\r' || c == ' ';
}

// 获取第一个不是空白字符的index
func getFirstNoSpaceIndex(line string) int{
	var first int = -1
	for i, v := range(line){
		if !IsSpace((byte)(v)) {
			first = i
			break
		}
	}
	return first
}

// 文件是否存在
func IsFileExist(file string) bool{
	if _, err := os.Stat(file); err != nil{
		return false
	}
	return true
}
