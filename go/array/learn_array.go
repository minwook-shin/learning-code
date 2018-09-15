package main

import (
	"fmt"
)

func printArr(temp[3]int){
	fmt.Println(temp)
}

func edit(temp[]int){
	for i,j := range temp {
		fmt.Println(i,j)
		temp[i] = 0
	}
}

func main() {
	var arr[10] int
	fmt.Println(arr)

	var changeArr[10] int
	changeArr[0] = 1
	changeArr[9] = 1
	fmt.Println(changeArr)

	var insertArr = [10] int{1,2,3}
	fmt.Println(insertArr)

	var resizeArr = [...]int{1,2,3}
	fmt.Println(resizeArr)
	
	var tmp = [3] int {1,2,3}
	var copyArr = tmp
	fmt.Println(copyArr)

	var temp = [3] int {1,2,3}
	printArr(temp)
	fmt.Println(temp)

	var lengthTest [10]int
	fmt.Println(len(lengthTest))

	var lengthTest2 = [...]int{1,2,3}
	fmt.Println(len(lengthTest2))

	var tmp2 = [3] int {1,2,3}
	for index := 0; index < len(tmp2); index++ {
		fmt.Println(tmp2[index])
	}

	var multiArr [2][3]int
	fmt.Println(multiArr)

	sliceArr := [10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(sliceArr[:],sliceArr[0:5],len(sliceArr[0:5]),cap(sliceArr[0:5]))

	makeArr := make([]int, 3, 10)
	fmt.Println(makeArr)
	
	realArr := []int{1,2,3}
	addArr := []int{1,2,3}
	fmt.Println(append(realArr,addArr...))


	editArr := []int{3,2,1}
	edit(editArr)
	fmt.Println(editArr)

	copyArr2 := []int{1,2,3,4,5,6,7,8}
	tmpArr := make([]int,2)
	fmt.Println(copyArr2)
	fmt.Println(tmpArr)
	copy(tmpArr, copyArr2)
	fmt.Println(tmpArr)	
}