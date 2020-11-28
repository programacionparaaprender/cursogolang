package main

import (
	"strconv"
	"fmt"
	"sort"
	"io/ioutil"
	"os"
	"io"
	"time"
)
//import "github.com/luiscorrea/repo/matrix"

//import "goweek/goweek"

func main(){
	//fmt.Println(divisores(13))
	numerosPrimos(15)
	//whatday(20)	
	//whatday(7)
	//arra2()
	//matrix()
	//imprimirx(4)
	//slice3()
	//pointer1()
	//pointer2()
	//variadic()
	//recursion(25)
	//recursionpattern(25)
	//sortdemo()
	//structdemo()
	//addslice()
	//read1()
	//writefiledemo()
	//mapdemo()
	//methoddemo()
	//interfacedemo()
	//goroutinedemo()
	//channeldemo()
}

func numerosPrimos(a int){
	for i:=a; i >= 2; i--{
		if divisores(i) <= 2 {
			var s string 
			s = strconv.Itoa(i)
			fmt.Println(s)
		}
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


func channeldemo(){
	ch := make(chan string, 1)
	go func(msg string){
		ch <- msg
	}("hello")
	recv := <-ch 
	fmt.Println(recv)
}

func goroutinedemo(){
	//goroutinedemo1()
	goroutinedemo2()
}

func goroutinedemo1(){
	count(5)
}

func goroutinedemo2(){
	go count(5)
	time.Sleep(time.Second * 1)
}

func count(num int){
	for i:=0 ;i <= num; i++{
		fmt.Println(i)
	}
}

type userprofile interface {
	getname() string
	getaddress() string
	getcontact() map[string]string
}

func getuserinfo(u userprofile){
	fmt.Println("name: ", u.getname())
	fmt.Println("Address: ", u.getaddress())
	fmt.Println("Contact Details: ",u.getcontact())
	fmt.Println("Email Extraction: ", u.getcontact()["email"])
}

func interfacedemo(){
	s:= profile{
		name: "Sherlock",
		address: "2218 Baker String",
		email:"ejemplo@ejemplo.com",
		contactnum:"122343494",
	}
	getuserinfo(s)
}

type profile struct{
	name string
	address string
	email string
	contactnum string
}

func (p profile) getname() string{
	return p.name
}

func (p profile) getaddress() string{
	return p.address
}

func (p profile) getcontact() map[string]string{
	m:= map[string]string{"email":p.email,"contactnum":p.contactnum}
	return m
}

func methoddemo(){
	s:= profile{
		name: "Sherlock",
		address: "2218 Baker String",
		email:"ejemplo@ejemplo.com",
		contactnum:"122343494",
	}
	fmt.Println("name: ", s.name)
	fmt.Println("address: ", s.address)
	fmt.Println("email: ", s.email)
	fmt.Println("contactnum: ", s.contactnum)
}

func mapdemo(){
	rm := map[string]int{"first":1,"second":2}
	fmt.Println("Map nm: ", rm)
	m := make(map[string]int)
	m["a"] = 12
	m["b"] = 13
	m["c"] = 14
	fmt.Println("Map m: ", m)
	A := m["a"]
	fmt.Println("Deleted m['a'] from Map m:", A)
	
	delete(m, "b")
	fmt.Println("Deleted m['b'] from Map m:", m)

	B, is_B_present := m["b"]
	C, is_C_present := m["c"]

	fmt.Println("B from Map m:", B)
	fmt.Println("C from Map m:", C)
	fmt.Println("is_B_present:", is_B_present)
	fmt.Println("is_C_present:", is_C_present)

}


func writefiledemo(){
	//write1()
	write2()
}

func write1(){
	b:=[]byte("How did you like the review?\nHope you liked it.\n")
	ioutil.WriteFile("demo.txt", b, 0644)
	fmt.Println("done")
}

func write2(){
	f, _ := os.OpenFile("demo.txt", os.O_APPEND | os.O_WRONLY, 0777)
	b:= []byte("Adding new line")
	defer f.Close()

	f.Write(b)
	fmt.Println(f.Stat())
}

func readfiledemo(){
	//read1()
	//read2()
	read3()
}

func read1(){
	dat, _ := ioutil.ReadFile("dunkirk.txt")
	fmt.Println(dat)
}
func read2(){
	f, _ := os.Open("dunkirk.txt")
	b:= make([]byte, 100)
	for{
		r, _ := f.Read(b)
		if r == 0 {
			break;
		}
		fmt.Println(r)
		fmt.Println(string(b))
	}
}
func read3(){
	f, _ := os.Open("dunkirk.txt")
	b:= make([]byte, 100)
	for{
		r, err := f.Read(b)
		if err != nil && err != io.EOF{
			panic(err)
		}
		if r == 0 {
			break
		}
		fmt.Println(string(b[:r]))
	}
}

func addslice(){
	s := make([]int, 3)
	for i:= 0; i < 3; i++{
		s[i] = i + 1
	}
	s = append(s, 4,5,6,7,8,9)
	fmt.Println("s = ", s)

	b := make([]int,len(s))
	copy(b, s)
	fmt.Println("b = ", b)

	s = append(s[:2],s[4:]...)
	fmt.Println("cut(3,4) s=", s)

	//del by index
	s = delbyindex(2, s)
	fmt.Println("deleted index 2 in s = ", s)
}

func delbyindex(i int, a []int)[]int{
	a = append(a[:i],a[i+1:]...)
	return a
}

type Book struct{
	name string
	author string
	pages int
}

type Shelf struct{
	position int
	book Book
}

func structdemo(){
	book := Book{name:"Harry Potter",author:"J.K. Rowling",pages:800}
	fmt.Println(book)
	fmt.Println("Name:",book.name)
	fmt.Println("Author:",book.author)
	fmt.Println("Pages:",book.pages)
	fmt.Println()

	fmt.Println("Book Shelf")
	s:= Shelf{position:1,book:book}
	fmt.Println(s)
	fmt.Println("book position: ", s.position)
	fmt.Println("book details: ", s.book)

}

func sortdemo(){
	s := []string{"z", "x", "b", "a", "y"}
	sort.Strings(s)
	fmt.Println("Sorted string: ", s)
	nums := []int{7,5,4,1,2}
	sort.Ints(nums)
	fmt.Println("Sorted numbers: ", nums)
}

func recursionpattern(n int) int{
	if(n == 1){
		fmt.Print("*")
		return 1
	}else{
		for i:= 1; i<=n;i++{
			fmt.Print("*")
		}
		fmt.Println()
		return recursionpattern(n-1)
	}
}

func recursion(n int){
	z := fact(n)
	fmt.Printf("Factorial: %d\n", z)
}

func fact(n int) int{
	if(n == 1){
		fmt.Printf("1\n")
		return 1
	}else{
		fmt.Printf("%d*", n)
		return n * fact(n-1)
	}
}

func variadic(){
	sum(1,2,3,4)
	n:= []int{5,6,7}
	sum(n...)
}

func sum(nums ...int){
	fmt.Println("Nums Received", nums)
	total := 0
	for num:= range nums {
		fmt.Println("Num ", num)
		total += num
	}
	fmt.Printf("Total: %d\n", total)
}

func pointer2(){
	c := 5 
	increment(c)
	increment(c)
	fmt.Println("increment ", c)

	incrementbyptr(&c)
	incrementbyptr(&c)
	fmt.Println("incrementbyptr ", c)
}

func increment(c int){
	c++;
	fmt.Println("Called within increment", c)
}

func incrementbyptr(c *int){
	*c++;
	fmt.Println("Called within incrementbyptr", *c)
}

func pointer1(){
	var a *int 
	b := 6
	a = &b
	fmt.Println("Address ", a)
	fmt.Println("value ", *a)
}

func slice3(){
	s := make([]string, 3)
	fmt.Println("String-Slice", s)
	s[0] = "a"
	s[2] = "b"
	s = append(s, "c", "d")
	fmt.Println(s)
}


func slice2(){
	x := make([]int, 3, 10)
	x[1] = 2
	x = append(x, 1,2,3,4,5,6,7,8,9,10,11,12,13,14)
	fmt.Println("x->", x)
	fmt.Println("x[2:5] ->", x[2:5])
	fmt.Println("x[:5] ->", x[:5])
	fmt.Println("x[:] ->", x[:])
	fmt.Println("x[3:] ->", x[3:])
}

func slice1(){
	x := make([]int, 0)
	x = append(x, 1,2,3,4,5,6,7,8,9,10,11,12,13,14)
	fmt.Println(x)
}

func imprimirx(a int){
	var m[4][4]int
	for  i:=0;i<a;i++{
		for j:=0;j<a;j++{
			if i == j{
				m[i][j] = 1
			}else{
				m[i][j] = 0
			}
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
	fmt.Println("length is ", len(m))
	fmt.Println(m)
}

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