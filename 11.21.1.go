package main
import "fmt"
func main(){
	s :="asdaddsada"
	a :=byte('a')
	b :=byte('b')
	s=replace(s,a,b)
	fmt.Println(s)
}
func replace(s string,ta,re byte) string{
	var x []byte
	for i:=0;i<len(s);i++ {
		x=append(x,s[i])
	}
	for i,va :=range x{
		if va==ta {x[i]=re}
	}
	s=string(x)
	return s
}


