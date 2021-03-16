package main

import (
	"fmt"
	"go/format"
	"log"
	"regexp"
	"strings"
)

type Field struct {
	// 字段名字
	Name string
	// 字段类型
	Kind string
	// 字段注释
	Note string
	// 是否必须
	Must bool
	// Tag字段
	Tag string
}

func parse(paramName, s string) (string, []*Field) {
	fields := getFields(s)
	if len(fields) == 0 {
		log.Fatal(paramName, "field is nil")
		return "", nil
	}
	s = getParam(paramName, fields)
	if s == "" {
		log.Fatal(paramName, "param entity is nil")
		return "", nil
	}
	return s, fields
}

// 获取字段
func getFields(s string) []*Field {
	// 判断是否是数字开头
	numberReg, _ := regexp.Compile(`^[\d]+`)
	// 根据空格分割
	reg, _ := regexp.Compile("\\s+")
	// 根据换行符分割
	lines := strings.Split(s, "\n")
	fields := make([]*Field, 0)
	// 当前field指针
	cursor := 0
	for i := 0; i < len(lines); i++ {
		cols := reg.Split(lines[i], -1)
		if len(cols) > 0 {
			if numberReg.MatchString(cols[0]) {
				fields[cursor].Note += "\n// " + lines[i]
				continue
			}
		}
		switch len(cols) {
		case 0:
			continue
		case 1:
			if cols[0] == "" {
				continue
			}
			fields[cursor].Note += "\n// " + cols[0]
		case 2:
			fields[cursor].Note += "\n// " + lines[i]
		case 3:
			fields[cursor].Note += "\n// " + lines[i]
		case 4:
			cursor = len(fields)
			field := &Field{
				Name: cols[0],
				Kind: getKind(cols[1]),
				Must: cols[2] == "是",
				Note: "// " + cols[3],
			}
			fields = append(fields, field)
		default:
			cursor = len(fields)
			note := ""
			for _, col := range cols[3:] {
				note += "\n// " + col
			}
			field := &Field{
				Name: cols[0],
				Kind: getKind(cols[1]),
				Must: cols[2] == "是",
				Note: note,
			}
			fields = append(fields, field)
		}
	}
	return fields
}

// 获取API参数
func getParam(paramName string, fields []*Field) string {
	s := fmt.Sprintf("type %s struct {\n", paramName)
	for _, field := range fields {
		s += fmt.Sprintln(field.Note)
		s += fmt.Sprintf("// 是否必须: %v\n", field.Must)
		if field.Must {
			s += fmt.Sprintf("%s  %s  `json:\"%s\"`\n\n", strings.Title(field.Name), field.Kind, field.Name)
		} else {
			s += fmt.Sprintf("%s  %s  `json:\"%s,omitempty\"`\n\n", strings.Title(field.Name), field.Kind, field.Name)
		}
	}
	s += fmt.Sprintln(`}`)
	b, err := format.Source([]byte(s))
	if err != nil {
		log.Println(err)
	}
	return string(b)
}

// 获取类型
func getKind(kind string) string {
	switch kind {
	case "String", "string":
		return "string"
	case "Int", "int", "Integer":
		return "int"
	case "Boolean", "boolean":
		return "bool"
	case "Long", "long":
		return "int64"
	case "Multipart":
		return "string"
	default:
		return "unknown"
	}
}
