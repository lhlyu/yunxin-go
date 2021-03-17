package main

import (
	"fmt"
	"strings"
)

type Name struct {
	// 请求地址的常量名
	ActionName string
	// 方法参数的名字
	ParamName string
	// 方法的名字
	FuncName string
	// 文件的名字
	FileName string
}

// 截取段
const apipath = "https://api.netease.im/nimserver/"

// 通过请求地址获取名字
func extractName(action string) Name {
	action = action[len(apipath):(len(action) - 7)]
	action = strings.ReplaceAll(action, "/", "_")
	name := snakeString(action)

	actionName := fmt.Sprintf("_API_%s_URL", strings.ToTitle(name))
	paramName := fmt.Sprintf("%sParam", camelString(name))
	funcName := fmt.Sprintf("Api%s", camelString(name))
	fileName := fmt.Sprintf("api_%s.go", name)

	return Name{
		ActionName: actionName,
		ParamName:  paramName,
		FuncName:   funcName,
		FileName:   fileName,
	}
}

// 获取蛇形名
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// 获取驼峰名
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// 字符串小驼峰
func littleCamelCase(s string) string {
	needCamel := false
	isFirst := true
	s = strings.Map(func(r rune) rune {
		if isFirst {
			isFirst = false
			if r < 65 || r > 90 {
				return r
			}
			return r + 32
		}
		if needCamel {
			if r == 32 || r == 95 {
				return -1
			}
			needCamel = false
			if r < 97 || r > 122 {
				return r
			}
			return r - 32
		}
		if r == 32 || r == 95 {
			needCamel = true
			return -1
		}
		return r
	}, s)
	return s
}
