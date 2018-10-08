package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)



func main() {

	// 读物配置文件
	file, err := os.Open("./conf/app.cnf")

	if err != nil {
		fmt.Printf("error %+v\n", err)
		return
	}

	defer file.Close()

	m := splitLine(file)

	fmt.Printf("%+v\n", m)

	fmt.Println("rsa =",m[`rsa`])

}


// 按照行读取文件 读取配置文件。。 配置文件的格式使用 = 号作 分割
func splitLine(reader io.Reader) (m map[string]string) {

	m = make(map[string]string)
	buff := bufio.NewReader(reader)

	for {
		line, _, err := buff.ReadLine()

		if err == nil {
			split := strings.Split(string(line), "=")

			if len(split) == 2 {
				m[split[0]] = split[1]
			}

		} else {
			break
		}
	}
	return m
}
