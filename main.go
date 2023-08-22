package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func txtToArray(txtPath string) {
	file, err := os.Open(txtPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建字符串数组来存储文件内容
	var textArray []string

	// 逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		textArray = append(textArray, line)
	}

	// 检查扫描过程中是否出错
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// 将字符串数组转换为多行字符串字面量
	arrayCode := strings.Join(textArray, `",`+"\n\t"+`"`)
	finalCode := fmt.Sprint("package sdict\n\nvar MyArray = []string{\n\t\"", arrayCode, "\",\n}")

	// 写入到新的Go文件
	filePath := "./sdict/myarray.go"
	file, err = os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(finalCode)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created:", filePath)
}

func main() {
	txtToArray("./dict/dict.txt")
}
