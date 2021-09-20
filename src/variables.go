 package main

 import (
	 "fmt"
	 "strconv"
 )

 func main() {
	 var  (
		 		fname = "Jason"
				lname = "Black"
				birth = "01/23/54"
				role = "master teacher"
			)
	 var i int = 42
	 var j int = 58
	 var k int = 23
	 var m = strconv.Itoa(k)

	 // i := 123
	 // fmt.Println("Hellow, playground")
	 // fmt.Println(i)
	 fmt.Printf("%v, %T\n",j ,j ) //Prints the value and Type respectively
	 fmt.Printf("%v, %T\n",i ,i ) //Prints the value and Type respectively
	 fmt.Println(fname,lname,birth,role)
	 fmt.Println(string(k))
	 fmt.Printf("%v, %T\n", k, m)
 }
 
