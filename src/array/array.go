package array

import (
	"fmt"
	"sort"
)

func Demo1() {
	var a [5]int
	a[0] = 4
	a[1] = 5
	a[2] = 6
	a[3] = 7
	a[4] = 8

	fmt.Println(a)
	fmt.Println("Size: ", len(a))
	for i := 0; i < len(a); i++ {
		fmt.Println("\t", a[i])
	}
}

func Demo2() {
	a := []int{6, 2, 4, 6, 8, 9, 4, 5, 3}
	fmt.Println(a)
	fmt.Println("Size: ", len(a))
	for i := 0; i < len(a); i++ {
		fmt.Println("\t", a[i])
	}
}

func Demo3(a []int) {
	fmt.Println(len(a))
	fmt.Println("Size: ", len(a))
	for i := 0; i < len(a); i++ {
		fmt.Println("\t", a[i])
	}
}

func Demo4() {
	name := []string{"truong", "asdsa", "asds", "retty", "684564"}
	for i, v := range name {
		fmt.Println(i, " : ", v)
	}
}

func Demo5() {
	name := []string{"truong", "asdsa", "asds", "retty", "684564"}
	fmt.Println(name)
	names := append(name, "chun")
	fmt.Println(names)
}

func Demo5bis(a []int) []int {
	var result []int
	for _, v := range a {
		if v > 0 {
			result = append(result, v)
		}
	}
	return result
}

func Demo6() []int {
	a := []int{6, 2, -4, 6, 8, -9, 4, 5, -3}
	var result []int
	for _, v := range a {
		if v > 0 {
			result = append(result, v)
		}
	}
	return result
}

func Demo7(a []int) (result1 []int, result2 []int, result3 []int, result4 []int) {

	for _, v := range a {
		if v > 0 {
			result1 = append(result1, v)
		}
		if v < 0 {
			result2 = append(result2, v)
		}
		if v%2 == 0 {
			result3 = append(result3, v)
		}
		if v%2 != 0 {
			result4 = append(result4, v)
		}
	}
	return
}

func Demo8() {
	a := []int{6, 2, -4, 6, 8, -9, 4, 5, -3, -12, 434, 54, 65, -43}
	fmt.Println("ASC")
	sort.Ints(a)
	fmt.Println(a)

	names := []string{"abc", "xyz", "asd"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println("DESC")
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
}

func Demo9() {
	a := [2][3]int{
		{5, -2, 9},
		{10, 59, 99},
	}
	fmt.Println(a)
	fmt.Println("Size: ", len(a))
	for r := 0; r < len(a); r++ {
		for c := 0; c < len(a[r]); c++ {
			fmt.Print("\t", a[r][c])
		}
	}
}
