package main

import ("fmt";"unsafe")

func main() {

	//Integers Type
	var i1, i2 int // Integer Variable - Declarations
	i1 = -1 // Integer Variable - Declarations
	i2 = 2 // Integer Variable - Declarations
	i3 := 3 //Variable Short Initilization - Data interpreted based on value assigned i3 will be integer type due to value 3
	var i4 int8 = 44
	var i5 int16 = 55 
	var i6 uint16 = 66 

	//Print values
	fmt.Println("\n--------------Integers--------------")
	fmt.Printf("Data Type of i1 is %T, value is %d, size is %d", i1, i1, unsafe.Sizeof(i1))
	fmt.Printf("\nData Type of i2 is %T, value is %d, size is %d", i2, i2, unsafe.Sizeof(i2))
	fmt.Printf("\nData Type of i3 is %T, value is %d, size is %d", i3, i3, unsafe.Sizeof(i3))
	fmt.Printf("\nData Type of i4 is %T, value is %d, size is %d", i4, i4, unsafe.Sizeof(i4))
	fmt.Printf("\nData Type of i5 is %T, value is %d, size is %d", i5, i5, unsafe.Sizeof(i5))
	fmt.Printf("\nData Type of i5 is %T, value is %d, size is %d", i6, i6, unsafe.Sizeof(i6))
	
	
	//Float Type
	var f1, f2 float32 = -1.1, 2.2 // Float Variable - Declarations with Initilization
	f3:= 3.3 //Variable Short Initilization - Data interpreted based on value asigned f3 will be float type due to value 3.3
	var f4 float64
	
	//Print values
	fmt.Println("\n\n--------------Float--------------")
	fmt.Printf("Data Type of f1 is %T, value is %f, size is %d", f1, f1, unsafe.Sizeof(f1))
	fmt.Printf("\nData Type of f2 is %T, value is %f, size is %d", f2, f2, unsafe.Sizeof(f2))
	fmt.Printf("\nData Type of f3 is %T, value is %f, size is %d", f3, f3, unsafe.Sizeof(f3))
	fmt.Printf("\nData Type of f4 is %T, value is %f, size is %d", f4, f4, unsafe.Sizeof(f4))
	
	
	//Complex Type
	var c1 complex64
	c1 = complex(1,1)
	c2 := 2 + 2i
	
	//Print values
	fmt.Println("\n\n--------------Complex--------------")
	fmt.Printf("Data Type of c1 is %T, size is %d", c1, unsafe.Sizeof(c1))
	fmt.Printf("\nData Type of c2 is %T, size is %d", c2, unsafe.Sizeof(c2))
	
	
	//Boolean Type
	var t, f bool = true, false
	
	//Print values
	fmt.Println("\n\n--------------Boolean--------------")
	fmt.Printf("Data Type of t is %T, value is %v, size is %d", t, t, unsafe.Sizeof(t))
	fmt.Printf("\nData Type of f is %T, value is %v, size is %d", f, f, unsafe.Sizeof(f))

	//String Type
	var s1, s2 string = "Ankit", "Trivedi"
	
	//Print values
	fmt.Println("\n\n--------------String--------------")
	fmt.Printf("Data Type of s1 is %T, value is %s, size is %d", s1, s1, unsafe.Sizeof(s1))
	fmt.Printf("\nData Type of s2 is %T, value is %s, size is %d", s2, s2, unsafe.Sizeof(s2))
}
