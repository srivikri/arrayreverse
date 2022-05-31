// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, sri")
	//sriArr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} -- shorthand declaration with num of elements and initialisation
	//var sriArr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} -- proper declaration with initialisation (slice)
	sriArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //shorthand declaration with elipsis
	//var sriArr [10]int -- decaration without initialisation
	//sriArr[0] = 1 -- assigning value

	fmt.Println("Printing using simple for loop:")
	for i := 0; i <= 9; i++ {
		fmt.Println(sriArr[i])
	}
	fmt.Println("Printing using range;")
	for i, j := range sriArr { //here i stores index, j stores value
		fmt.Println(i, j)
	}
	fmt.Println("Loop using condition only;")
	i := 0
	for i < 10 {
		fmt.Println(sriArr[i])
		i++
	}
	fmt.Println("Reversing....")
	sriArr2 := make([]int, len(sriArr))
	j := 0
	for i := len(sriArr) - 1; i >= 0; i-- {
		sriArr2[j] = sriArr[i]
		j++
	}
	fmt.Println("Printing reversed array")
	for i, j := range sriArr2 {
		fmt.Println(i, j)
	}
}
