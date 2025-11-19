package main

import "fmt"

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

func main() {

	slice := []int{1, 2, 3, 4, 6, 18, 24, 654, 443}

	fmt.Println(sliceSum(slice))
	fmt.Println(FilterEven(slice))
	fmt.Println(Reverse(slice))
	fmt.Println(Delete(slice, 4))

}
