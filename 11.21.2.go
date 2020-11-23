package main
import (
	"fmt"
//	"strings"
)
func main(){
	s:="XXVII"
	fmt.Printf("%d\n",R(s))
}
func R(s string) int{
	x:=0
	m := map[rune]int{'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000}
	//a := strings.Fields(s)
	for i,va :=range s{
		if i<len(s)-1 {
		if va=='I'&&(s[i+1]=='V'||s[i+1]=='X')||va=='C'&&(s[i+1]=='D'||s[i+1]=='M')||va=='X'&&(s[i+1]=='L'||s[i+1]=='C'){
			x-=m[va]
		}else{
			x+=m[va]
		}
		}else{
			x+=m[va]
		}
	}
	return x
}

