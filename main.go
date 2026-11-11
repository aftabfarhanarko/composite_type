package main

import "fmt"

func main() {
	fmt.Println("Hello yujtyttj5j")

	// Array
	arr1 := [5]int{34, 45, 45, 6, 65}

	fmt.Println("Array  Prient", arr1)

	// Slice
	sl1 := []int{3, 4, 4, 5, 6, 6, 7, 7, 7, 23, 5, 45, 434}
	fmt.Println("Slice Prient","Length",len(sl1), sl1)
	// Slice data add
	sl2 := append(sl1, 56, 67, 68678, 67, 89, 768, 67, 8, 678, 67, 8, 678, 6, 45)
	fmt.Println("Slice Prient","Length",len(sl2), sl2)

	// way to decler slices
	s3 := make([]int , 4)
	fmt.Println(s3)

	s4 := make([]float64, 6, 60)
	fmt.Println("Lingth",len(s4))
	fmt.Println("Capecitey", cap(s4))

	fmt.Println(sl1[1])
	fmt.Println(append(sl1 , 555,666,777,888,999))

	fmt.Println(append(sl1,s3...))

	for i := 0 ; i < len(sl1); i++ {
		fmt.Println(sl1[i])
	}

	for i , v := range sl2 {
		fmt.Println(i, v)
	}
}
