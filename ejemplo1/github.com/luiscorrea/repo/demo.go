package repo

import (
	"strconv"
	"fmt"
)
import "github.com/luiscorrea/repo/matrix"

//import "goweek/goweek"

func add(a, b int) int{
	return a + b
}

func Boolean() bool{
	var b bool
	b = true
	return b
}

func loop(){
	var entero int = 10
	fmt.Println("For")
	for i:=1; i <= entero; i++{
		fmt.Println(strconv.Itoa(i))
	}
}

func loop3(){
	var entero int = 10
	fmt.Println("For")
	for i:=1; i <= entero; i++{
		for j:=1; j <= i; j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func loop4(){
	var entero int = 10
	fmt.Println("For")
	for i:=entero; i >= 1; i--{
		for j:=1; j <= i; j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func addeven(){
	var entero int = 10
	fmt.Println("For")
	for i:=entero; i >= 1; i--{
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}


func primero(){
	fmt.Printf("Hello world!")
	//var b string
	//b := 
	var s string 
	s = strconv.Itoa(add(1,2))
	fmt.Println("Suma 1 + 2 = " +s)
	//fmt.Printf(b)
	if Boolean()==true {
		fmt.Println("bool = true")
	}
}

func divisores(a int) (int){
	var e1 int = 0
	var e2 int = a
	for i:=a; i >= 1; i--{
		if e2%i == 0 {
			e1++
		}
	}
	return e1
}

func numerosPrimos(a int){
	for i:=a; i >= 1; i--{
		if divisores(i) <= 2 {
			var s string 
			s = strconv.Itoa(i)
			fmt.Println(s)
		}
	}
}

/* func days(){
	week, -:=goweek(2015, 46)
	fmt.Println(week.Days)
} */
func arra1(){
	var a [7]int
	a[0] = 1
	a[6] = 7
	fmt.Println(a[4])
	fmt.Println(a[6])
	fmt.Println(a)
}
func arra2(){
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(len(a))
}
func main(){
	//numerosPrimos(15)
	//whatday(20)	
	//whatday(7)
	//arra2()
	matrix()
}

func whatday(n int){
	if n != 0 && n <= 7 {
		switch n {
			case 1:
				fmt.Println("Monday")
			case 2:
				fmt.Println("Tuesday")
			case 6:
				fmt.Println("Saturday")
			case 7:
				fmt.Println("Hola - Sunday")
			default:
				fmt.Println("It is weekday")
		}
	}else{
		fmt.Println("Wrong")
	}
}