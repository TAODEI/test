package main
import "fmt"
type cc struct {
	a int
	b int
	c int
	s int
	v int
}
func (x cc) ss() int{
	return (x.a*x.b+x.b*x.c+x.c*x.b)*2
}

func (x cc) vv() int{
    return x.a*x.b*x.c
}

func main(){
	var a1 cc
	fmt.Scanf("%d%d%d",&a1.a,&a1.b,&a1.c)
	fmt.Printf("The V is: %d\n",a1.vv())
	fmt.Printf("The S is: %d\n",a1.ss())
}



