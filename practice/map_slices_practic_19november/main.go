package main

import (
	"fmt"
)

func sliceSum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func FilterEven(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, v := range nums {
		if v%2 == 0 {
			res = append(res, v)
		}
	}

	return res

}

func Reverse(nums []int) (res []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	res = nums
	return
}

func Delete(slice []int, index int) []int {
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

func insertByIndex(slice []int, index int, val int) []int {
	return append(slice[:index], append([]int{val}, slice[index+1:]...)...)

}
func insertAfterIndex(slice []int, index int, val int) []int {
	return append(slice[:index], append([]int{val}, slice[index:]...)...)
}

func makeChunks(slice []int, chunkSize int) [][]int {
	var chunks [][]int

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

type User struct {
	Name string
	City string
}

func GroupByCity(users []User) map[string][]string {
	groupByCity := make(map[string][]string)
	for _, user := range users {
		groupByCity[user.City] = append(groupByCity[user.City], user.Name)
	}
	return groupByCity
}

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(sliceSum(slice))
	fmt.Println(FilterEven(slice))
	fmt.Println(Reverse(slice))
	fmt.Println(Delete(slice, 4))
	fmt.Println(insertByIndex(slice, 2, 222))
	fmt.Println(insertAfterIndex(slice, 2, 444))
	//чанки
	fmt.Println(makeChunks(slice, 3))
}
