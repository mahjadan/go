package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//arr:= []int{7,3,7,3,1,3,4,1}
	arr := []int{2, 1, 1, 3, 2, 1, 1, 3}
	uniqueVal := map[int]int{}
	for _, v := range arr {
		uniqueVal[v] = 0
	}
	path := len(arr) + 1

	for i := 0; i < len(arr); i++ {
		if i <= len(uniqueVal) {
			if p, ok := visitAllUniqueValues(arr, i, uniqueVal); ok {
				fmt.Println("path found: ", p)
				if path > p {
					path = p
				}
			}
		}
	}
	fmt.Println("shortest path is :", path)
}

func Solution(A []int) int {
	// write your code in Go 1.4
	arr := []int{7, 3, 7, 3, 1, 3, 4, 1}
	uniqueVal := map[int]int{}
	for _, v := range arr {
		uniqueVal[v] = 0
	}
	path := len(A) + 1

	for i := 0; i < len(arr); i++ {
		if i <= len(uniqueVal) {
			if p, ok := visitAllUniqueValues(arr, i, uniqueVal); ok {
				if path > p {
					path = p
				}
			}
		}
	}
	return 0
}

func visitAllUniqueValues(arr []int, startFrom int, uniqueVal map[int]int) (int, bool) {
	fmt.Println("starting from: ", startFrom)
	pathCounter := 0
	uniqueValCounter := 0
	//uniqueVal:= map[int]int{}
	for i := startFrom; i < len(arr); i++ {
		pathCounter++

		if val, ok := uniqueVal[arr[i]]; ok && val == 0 {
			uniqueValCounter++
			val++
			uniqueVal[arr[i]] = val
		}

		//if count := uniqueVal[arr[i]]; count != 0{
		//	uniqueValCounter++
		//	count ++
		//	uniqueVal[arr[i]]= count
		//}
		if uniqueValCounter == len(uniqueVal) {
			return pathCounter, true
		}
	}
	resetCounter(uniqueVal)
	return 0, false
}

func resetCounter(val map[int]int) {
	for key, _ := range val {
		val[key] = 0
	}
}

func Solution1(N int) int {
	// write your code in Go 1.4
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
	}
	result := find_min(arr)
	return result
}

func found(arr []int, v int) bool {
	for _, i := range arr {
		if i == v {
			return true
		}
	}
	return false
}
func find_min(A []int) int {
	var result int = 0
	for i := 1; i < len(A); i++ {
		if result > A[i] {
			result = A[i]
		}
	}
	return result
}
