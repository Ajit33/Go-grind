package main

import (
	"fmt";
)

func main(){
	var array1 [2] string;
	array1[0]= "helllo";
	array1[1]="Array";
 primes:=[6]int{1,2,3,3,4,4}


	fmt.Println(array1)
	fmt.Println(primes)
	fmt.Println(array1[0])
	fmt.Printf("types of primes: %T\n", primes)
}