package main

import (
	
	"fmt"
)

func matrix(){
	var m[2][3]int
	for  i:=0;i<2;i++{
		for j:=0;j<3;j++{
			m[i][j] = 1
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
	fmt.Println("length is ", len(m))
	fmt.Println(m)
}