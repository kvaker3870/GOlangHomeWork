package main

import (
	"fmt"
)

func sumArray(arr *[10]int, sum *int) int {
	for i := 0; i < len(arr); i++ {
		*sum += arr[i]
	}

	return *sum
}

func multiplySlice(s []int, multiplier int) []int {

	sCopy := make([]int, len(s))
	copy(sCopy, s)
	sCopy = append(sCopy, 6)

	for i := 0; i < len(s); i++ {
		s[i] = s[i] * multiplier
	}

	for i := 0; i < len(sCopy); i++ {
		sCopy[i] = sCopy[i] * (multiplier + 1)
	}

	return sCopy

}

func newSliceNoDuplicate(slice []int) []int {

	//тут я изучил вопрос и map занимает 1 байт если черерз true типа map[i] = true, через пустую структуру 0 байт занимает
	newSlice := []int{}
	seen := map[int]struct{}{}
	for _, i := range slice {
		if _, ok := seen[i]; !ok {
			seen[i] = struct{}{}
			newSlice = append(newSlice, i)
		}
	}
	return newSlice
}

func main() {

	//TODO task 1
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	sumArray(&arr, &sum)
	fmt.Println(sum)

	//TODO task2
	var sliceArr []int
	sliceArr = append(sliceArr, 1, 2, 3, 4, 5)
	result := multiplySlice(sliceArr, 3)
	//по ссылке
	fmt.Println(sliceArr)
	//черек копию +1 к мультиплаеру
	fmt.Println(result)

	//TODO task3
	numbers := make([]int, 0, 20)

	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	part := numbers[3:8]
	partCopy := make([]int, len(part))
	copy(partCopy, part)
	part = append(part, 22)
	fmt.Println("numbers:", numbers)
	fmt.Println("part:", part)
	fmt.Println("part capacity:", cap(part))
	fmt.Println("part len:", len(part))
	fmt.Println("partCopy:", partCopy)
	fmt.Println("partCopy capacity:", cap(partCopy))
	fmt.Println("partCopy len:", len(partCopy))

	//TODO task4 matrix
	matrix := [][]int{
		{0, 1, 2},
		{2, 3, 4},
		{4, 5, 6},
	}
	fmt.Println("matrix:", matrix)
	mainDiag := make([]int, 0)
	secondDiag := make([]int, 0)
	for i := 0; i < len(matrix); i++ {

		mainDiag = append(mainDiag, matrix[i][i])
		secondDiag = append(secondDiag, matrix[i][len(matrix)-1-i])
	}
	fmt.Println("mainDiag:", mainDiag)
	fmt.Println("secondDiag:", secondDiag)

	//TODO task5
	sliceWithDuplicates := []int{1, 1, 3, 5, 5, 6, 6, 8, 9, 10}
	fmt.Println(newSliceNoDuplicate(sliceWithDuplicates))
}
