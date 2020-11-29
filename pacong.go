package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/anaskhan96/soup"
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

	/*方法二
	a := make([]byte, 1024*4)
	var result string
	for n := 1; n != 0; {
		n, _ = response.Body.Read(a)
		result += string(a[:n])
	}*/

	return string(result)
	//fmt.Println(string(result))
}

func main() {
	fmt.Println("请输入爬取的网址:")
	var s string
	fmt.Scanf("%s", &s)
	result := pa(s)
	//创建文件
	file, err := os.Create("哈哈哈")
	if err != nil {
		panic(err)
	}
	//解析标签
	var s string
	doc:=soup.HTMLParse(result)
    subDocs:=doc.FindAll("div","class","uk-margin")
    for _,subDoc:=range subDocs{
        link:=subDoc.Find("a")
        s+=string(link.Text())+"\n"//fmt.Println(link.Text())
    }
    
	file.WriteString(s)
	file.Close()
	fmt.Println("OK")
}
