// goex
package main

import (
	"errors"
	"fmt"
)

const (
	c1 = iota
	c2 = iota
	c3 = iota
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil
}

func main() {
	var i int = 0
	fmt.Printf("Hello World, %d, %d, %d, %d\n", i, c1, c2, c3)
	/*str := "我是好爸爸！"
	for i, ch := range str {
		fmt.Println(i, ch)
	}*/
	/*var arr [8]int = [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s_arr []int = make([]int, 5, 10)
	s_arr[3] = 19
	var s_arr1 []int = arr[:]
	s_arr = append(s_arr, s_arr1...)
	for i, v := range s_arr {
		fmt.Println("Array element[", i, "]=", v)
	}
	for i, v := range arr {
		fmt.Println("Array element[", i, "]=", v)
	}*/
	/*var personDB map[string]PersonInfo = make(map[string]PersonInfo)

	personDB["Tom"] = PersonInfo{"1", "Tom", "Addr12345"}
	personDB["Mike"] = PersonInfo{"2", "Mike", "Addr67890"}

	person, ok := personDB["Tom"]
	if t := true; ok && t {
		fmt.Println("Found person with id:", person.ID, " and name:", person.Name)
	} else {
		fmt.Println("Person not found.")
	}*/
	/*var arr [8]int = [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	for i, v := range arr {
		fmt.Println("Array element[", i, "]=", v)
	}*/
	sum, err := Add(1, 2)
	if err != nil {
		fmt.Printf("err = %s", err.Error())
	} else {
		fmt.Printf("sum = %d", sum)
	}

}
