package File

import (
	"io/ioutil"
	"strings"
)

func FileRead(name string) string {
	if contents, err := ioutil.ReadFile(name); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		return result
	}
	return ""
}
func FileWrite(name, content string) bool {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		return true
	}
	return false
}

const JSONPATH = "$storage"

func Read(name string) string {
	if contents, err := ioutil.ReadFile(name); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		return result
	}
	return ""
}
func Write(name, content string) bool {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		return true
	}
	return false
}
