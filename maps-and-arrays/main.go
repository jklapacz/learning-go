package main

import (
	"fmt"
	"sort"
)

type NameAge struct {
	Name string
	Age int
}

func main() {
	intSlice := []int{1,5,5,5,5,7,8,6,6,6}
	uniqueIntSlice := unique(intSlice)
	fmt.Println(uniqueIntSlice)
	str := []string{"bob", "john", "kate", "andy"}
	for i, v := range str {
		if v == "john" {
			fmt.Println(i)
		}
	}
	sortedList := sort.StringSlice(str)
	sortedList.Sort()
	fmt.Println(sortedList)
	index := sortedList.Search("andy")
	fmt.Println(index)

	for i, v := range intSlice {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// map to array of keys and values

	var nameAgeSlice []NameAge
	nameAges := map[string]int {
		"Michael": 30,
		"John": 25,
		"Jessica": 26,
		"Ali": 18,
	}
	for key, value := range nameAges {
		nameAgeSlice = append(nameAgeSlice, NameAge{ key, value})
	}
	fmt.Println(nameAgeSlice)
	// merging arrays
	items1 := []int{3, 4}
	items2 := []int{1, 2}
	result := append(items1, items2...)
	fmt.Println(result)

	// merging maps
	map1 := map[string]int {
		"Michael": 10,
		"Jessica": 15,
	}
	map2 := map[string]int {
		"Tom": 20,
		"Brad": 25,
	}
	fmt.Println(map1)
	for key, value := range map2 {
		map1[key] = value
	}
	fmt.Println(map1)

	// test for key in map
	fmt.Println(nameAges["bob"])
	if _, exists := nameAges["Michael"]; exists {
		fmt.Println("bob exists")
	}


}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	uniqueElements := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value{
			keys[entry] = true
			uniqueElements = append(uniqueElements, entry)
		}
	}
	return uniqueElements
}
