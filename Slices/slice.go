package main

import (
	"fmt"
)
func x(){
	//creating a slice of multiple dt
	s1:= []int{1,2,3,4,5}
	fmt.Println(s1)
	s2:= []bool{true,false,true,true,false}
	fmt.Println(s2)

	s:= []struct{
       i int
	   b bool
	}{
		{s1[0],s2[0]},
		{s1[1],s2[1]},
		{s1[2],s2[2]},
		{s1[3],s2[3]},
		{s1[4],s2[4]},
	}
	fmt.Println(s)
}
func main(){
	names:= [4] string{
		"Ajit",
		"sasmita",
		"Nanu",
		"jyoti",
	}
	a:=names[0:3]
	b:=names[1:2]
	fmt.Println(a,b)

	b[0]="goals"
	fmt.Println(a,b)
	fmt.Println(names)
	x();
}