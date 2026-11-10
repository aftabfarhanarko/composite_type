
package main

import "fmt"

func main() {
	fmt.Println("Composite Type - Array")

	// 1. Array Declaration
	var arr1 [5]int

	// 2. Array Initialization
	arr2 := [6]int{1, 45, 67, 68, 78}

	// 3. Array with Automatic Length
	arr3 := [...]int{23, 45, 56, 657, 77}

	// 4. Find Array Length
	fmt.Println("\n--- Array Length ---")

	a1 := len(arr2)
	a2 := len(arr3)

	fmt.Println("Array 1:", arr1)
	fmt.Println("Array 2:", arr2)
	fmt.Println("Array 3:", arr3)

	fmt.Println("Length of arr2:", a1)
	fmt.Println("Length of arr3:", a2)

	// 5. Traditional For Loop
	fmt.Println("\n--- Traditional For Loop ---")

	for i := 0; i < len(arr3); i++ {
		fmt.Println("Index:", i, "Value:", arr3[i])
	}

	// 6. Range Loop
	fmt.Println("\n--- Range Loop ---")

	for index, value := range arr2 {
		fmt.Println("Index:", index, "Value:", value)
	}

	// 7. Array Copy and Modification
	fmt.Println("\n--- Array Copy ---")

	w := [4]int{3, 4, 9, 10}

	e := w
	e[2] = 56

	fmt.Println("Original Array:", w)
	fmt.Println("Copied Array:", e)
}
