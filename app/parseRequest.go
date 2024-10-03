package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseRequest(data []byte) ([]string, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("Empty data")
	}

	if len(data) < 5 || data[0] != '*' {
		return nil, fmt.Errorf("Invalid data")
	}

	// 解析数组长度
	arrayLengthEnd := strings.Index(string(data), "\r\n")
	if arrayLengthEnd == -1 {
		return nil, fmt.Errorf("Invalid data format")
	}
	arrayLength, err := strconv.Atoi(string(data[1:arrayLengthEnd]))
	if err != nil {
		return nil, fmt.Errorf("Invalid data while parsing array length")
	}

	curIndex := arrayLengthEnd + 2
	var result []string

	// 循环解析命令和参数
	for i := 0; i < arrayLength; i++ {
		// 查找下一个 $
		if data[curIndex] != '$' {
			return nil, fmt.Errorf("Invalid format, expected '$'")
		}
		curIndex++

		// 解析当前命令或参数的长度
		lengthEnd := strings.Index(string(data[curIndex:]), "\r\n")
		if lengthEnd == -1 {
			return nil, fmt.Errorf("Invalid data format while parsing length")
		}
		length, err := strconv.Atoi(string(data[curIndex : curIndex+lengthEnd]))
		if err != nil {
			return nil, fmt.Errorf("Invalid data while parsing length")
		}
		curIndex += lengthEnd + 2

		// 解析命令或参数内容
		value := string(data[curIndex : curIndex+length])
		result = append(result, strings.ToLower(value))
		curIndex += length + 2 // 跳过内容和结尾的 \r\n
	}

	return result, nil
}
