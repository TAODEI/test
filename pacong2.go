package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/anaskhan96/soup"
)

func pa(url1 string) string {
	url := "http://work.muxi-tech.xyz/" + url1
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return string(result)
}

func main() {
	fmt.Println("请输入选择:")
	fmt.Println(`
	1.项目
	2.进度
	3.动态
	4.成员
	`)
	var choice int
	fmt.Scanln(&choice)
	var s string
	switch choice {
	case 1:
		s = "project"
	case 2:
		s = "status"
	case 3:
		s = "feed"
	case 4:
		s = "teamMember"
	default:
		fmt.Println("错误")
		return
	}
	result := pa(s)
	fmt.Println(result)
	//解析标签
	var s string
	doc:=soup.HTMLParse(result)
    subDocs:=doc.FindAll("div","class","uk-margin")
    for _,subDoc:=range subDocs{
        link:=subDoc.Find("a")
        s+=string(link.Text())+"\n"//fmt.Println(link.Text())
    }
	//创建文件
	file, err := os.Create("略略略")
	if err != nil {
		panic(err)
	}
	file.WriteString(s)
	file.Close()
	fmt.Println("OK")
}
