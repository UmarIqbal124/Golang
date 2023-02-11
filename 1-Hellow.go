package main

import (
	"fmt"
)

var value int = 100

func main() {
	//variable
	/*var a int
	a = 55
	var b string = "Umar"
	d := "Ali"
	fmt.Println("Hellow World ")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)
	*/

	/*scope of variables
	1) local (alwase use in function e.g myvalue)
	2) globle (outside the function(lawase use in pascal case e.g Myvalue-> use onely in same package -> must use var keyword))
	3) Package (use in camal case e.g myValue -> also use in other package)

	Shadowing (same name of local and package variable )

	Printf("%v , %t", varName, varName) -> to check value and type of variable

	Conversion ( one var type to other e.g var a int32 = 5   -> vab b int64 = int64(a)  Note other wisw not work as java work)
	if we use d:=56 then for conversion we must use var e int32 = int32(d)*/

	//var s int = 67

	//var b string = strconv.Itoa(s)
	//fmt.Println(value)
	c := 34 + 56i
	fmt.Printf("Real value is = %v \n Imagenary value is = %v", real(c), imag(c))
}
