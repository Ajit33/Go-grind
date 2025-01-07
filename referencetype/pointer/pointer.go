package main
import "fmt"

func main(){
	i,j:= 45,55;
	p:=&i;  //point to address of i
	fmt.Println(*p); //read i through i
	*p=21;
	fmt.Println(i)
	p=&j
	*p=*p/37
	fmt.Println(j)
}