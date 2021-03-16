package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"
)

func write(model Model) string {
	var tl string
	if model.Args == "" {
		tl = render(t, model)
	} else {
		tl = render(t2, model)
	}
	b, err := format.Source([]byte(tl))
	if err != nil {
		log.Println(tl)
		log.Fatal(model.FuncName, err)
	}
	writeFile(model.FilePath, string(b))
	return string(b)
}

func render(tmpl string, v interface{}) string {
	t := template.Must(template.New("template.html").Delims("[[", "]]").Parse(tmpl))
	buf := bytes.NewBufferString("")
	err := t.Execute(buf, v)
	if err != nil {
		return ""
	}

	return buf.String()
}

func writeFile(filePath, content string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()
}
