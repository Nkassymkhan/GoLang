package main

import "fmt"
	

func main(){
	a := [5]int {1, 3 ,4 ,5 ,6 }
	for i := 0; i < len(a); i++{
		fmt.Println(a[i])
	}

}