package main

import "fmt"

func main() {
	fmt.Println("start counting...")
	for i:=0;i<10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end of counting....")
}

// defer is like it exicute that part at the end of the program mean after all the programs run then this part will run , under the hood that part will push in a stack , after all other part completely run then the defer code fun in lifo order.
