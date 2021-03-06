package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [4]string{"A", "B", "C", "D"}                 //Shorthand declaration
	arr := [...]float32{0, 1.01, 2.04, 3.06, 4.04, 5.05} //Shorthand declaration without length mentioned

	//Print array values
	fmt.Printf("Print values using for arrays")
	fmt.Println("\nThe values in arr1 are:", arr1)
	fmt.Println("The values in arr2 are:", arr2)
	fmt.Println("The values in arr3 are:", arr)

	//Creating slices of arr1
	var slice_arr1 []int = arr1[1:2]

	fmt.Println("\nPrint values using for slice")
	fmt.Println("The values in slice_arr1 are:", slice_arr1)
	fmt.Println("The length of slice_arr1 is:", len(slice_arr1))
	fmt.Println("The capacity of slice_arr1 before is:", cap(slice_arr1))

	//Creating slices of arr2
	slice_arr2 := arr2[1:3]

	//Modify slice will change the original array
	fmt.Println("\nThe values in slice_arr2 before are:", slice_arr2)
	fmt.Println("The length of slice_arr2 before is:", len(slice_arr2))
	fmt.Println("The capacity of slice_arr2 before is:", cap(slice_arr2))
	fmt.Println("The values in arr2 are:", arr2) //Print the original Array for comparision

	//Change the value at position 1
	slice_arr2[1] = "G"

	fmt.Println("\nThe values in slice_arr2 after are:", slice_arr2)
	fmt.Println("The length of slice_arr2 after is:", len(slice_arr2))
	fmt.Println("The capacity of slice_arr2 after is:", cap(slice_arr2))

	//Increase length of slice
	slice_arr2 = slice_arr2[0:3]

	fmt.Println("\nThe values in slice_arr2 after are:", slice_arr2)
	fmt.Println("The length of slice_arr2 after is:", len(slice_arr2))
	fmt.Println("The capacity of slice_arr2 after is:", cap(slice_arr2))

	//Create slice using make func
	slice_arr3 := make([]int, 2, 5)

	fmt.Println("\nThe values in slice_arr3 are:", slice_arr3)
	fmt.Println("The length of slice_arr3 is:", len(slice_arr3))
	fmt.Println("The capacity of slice_arr3 is:", cap(slice_arr3))
	
	//Append a value to the slice_arr3
	slice_arr3 = append(slice_arr3, 6)
	
	fmt.Println("\nThe values in slice_arr3 after are:", slice_arr3)
	fmt.Println("The length of slice_arr3 after is:", len(slice_arr3))
	fmt.Println("The capacity of slice_arr3 after is:", cap(slice_arr3))

}
