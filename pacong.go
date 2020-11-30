package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"os"
	"regexp"
)

func pa(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	//方法一
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return string(result)
	//fmt.Println(string(result))
}

func jie(s string)[]string {
	re:=regexp.MustCompile(`="noopener noreferrer">(?s:(.*?))<svg xmlns="http`)
	if re == nil {
		fmt.Println("错误")
	}
	x:=re.FindAllStringSubmatch(s,-1)
	var z []string
	for _,data:=range x{
		z=append(z,data[1])
	}
	return z
}

func main() {
	s:= "https://muxi-studio.github.io/101/cs/clang/c_50.html"
	result := pa(s)
	//解析标签
	ss:=jie(result)
	fmt.Println(ss)
	/*创建文件
	file, err := os.Create("哈哈哈")
	if err != nil {
		panic(err)
	}
	file.WriteString(string(ss))
	file.Close()
	fmt.Println("OK")*/
}
