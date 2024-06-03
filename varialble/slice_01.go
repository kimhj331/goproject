// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.2} //선언 및 초기화
	fmt.Println(t)

	var slice []int
	fmt.Println(slice, nil == slice)
	//초기화 하지 않은 slice는 nil이다
	// [] true

	slice = []int{}
	fmt.Println(slice, nil == slice)

	var targetSlice = []int{10, 11, 12}
	fmt.Println(targetSlice, len(targetSlice), cap(targetSlice), &targetSlice)

}
