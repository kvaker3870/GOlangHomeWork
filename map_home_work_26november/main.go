package main

import (
	"fmt"
	"strings"
)

func CountInts(nums []int) map[int]int {
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	return counter
}

func WordCount(s string) map[string]int {
	sArray := strings.Split(s, ",")
	counter := map[string]int{}
	for _, s := range sArray {
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)
		counter[s]++
	}

	return counter

}

func IsAnagram(a, b string) bool {
	aRunes := []rune(a)
	bRunes := []rune(b)
	counter := 0
	if len(aRunes) != len(bRunes) {
		return false
	}

	for i := 0; i < len(aRunes); i++ {
		for j := 0; j < len(bRunes); j++ {
			if aRunes[i] == bRunes[j] {
				counter++
				break
			}
		}
	}

	return counter == len(aRunes)

}

func Invert(m map[string]int) map[int]string {
	inverted := make(map[int]string)
	for k, v := range m {
		inverted[v] = k
	}

	return inverted
}

func mergeMaps(dst, src map[string]int) map[string]int {
	for k, v := range src {
		dst[k] = v
	}

	return dst
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

	//дом. задание

	//task1
	nums := []int{2, 3, 4, 5, 3, 2, 2, 4, 5, 6, 7, 5, 3, 2, 2, 4, 5, 10, 10, 10, 10, 45, 45, 45}
	fmt.Println("Task #1", "\n", "In:", nums, "\n", "Out:", CountInts(nums))
	//task2
	s := "Ехал,Грека,через,реку,видит,Грека,в,реке,рак,Сунул,Грека,руку,в,реку,рак ,за,руку, Греку, цап"
	fmt.Println("Task #2", "\n", "In:", s, "\n", "Out:", WordCount(s))
	//task3
	a := "апельсин"
	b := "спаниель"
	fmt.Println("Task #3", "\n", "In:", a, ",", b, "\n", "Out:", IsAnagram(a, b))
	//task4
	c := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fmt.Println("Task #4", "\n", "In:", c, "\n", "Out", Invert(c))
	//task5
	d := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	e := map[string]int{"h": 1, "j": 2, "v": 3, "z": 4, "y": 5}
	fmt.Println("Task #5", "\n", "In:", d, e, "\n", "Out:", mergeMaps(d, e))
	//task6
	users := []User{
		{
			Name: "Лена",
			City: "Москва",
		},
		{
			Name: "Дима",
			City: "Алматы",
		},
		{
			Name: "Кирилл",
			City: "Алматы",
		},
		{
			Name: "Саша",
			City: "Алматы",
		},
		{
			Name: "Егор",
			City: "Москва",
		},
		{
			Name: "Вася",
			City: "Питербург",
		},
	}

	fmt.Println("Task #6", "\n", "In:", users, "\n", "Out:", GroupByCity(users))

}
