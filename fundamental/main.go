package main

import "fmt"

const s string = "pakage level const"

func main() {

	// string
	var a = "this is string"
	fmt.Println(a)
	// int
	var i = 20
	fmt.Println(i)
	// boolean
	var s = true
	fmt.Println(s)
	// default value as 0
	var l int
	fmt.Println(l)
	//calling function
	fmt.Println("=========Const Ex==================")
	constEx()
	// if you call only constEx,
	// will error msg "constEx evaluated but not used"
	//constEx
	fmt.Println("=========Slices and Array==================")
	sliceAndArrEx()
	fmt.Printf("==================Range Example===============\n")
	rangeEx()
}

func constEx() {
	fmt.Println(s)
	const b string = "java"
	fmt.Println(b)
	// will raise error "cannot assign to b"
	//b = "py"
	//fmt.Println(b)
}

func sliceAndArrEx() {
	slice := make([]string, 3)
	var arr = [3]int{3, 4, 5}
	// array index 3 out of bounds [0:3]
	// var arr = [3]int {3,4,5,6}
	slice[0] = "Hi"
	slice[1] = "Hello"
	slice[2] = "bye"
	// panic: runtime error: index out of range [3] with length 3
	// s[3] = "Tata"
	// append(slice, "Tata") evaluated but not used
	// append(slice, "Tata")
	slice = append(slice, "Tata")
	slice = append(slice, "see")
	slice = append(slice, "you")
	fmt.Println(slice, arr)
}

func rangeEx() {
	m := make(map[string]int)
	m["k2"] = 2
	m["k1"] = 1
	m["k3"] = 3
	fmt.Println(m)
	delete(m, "k3")
	fmt.Println(m)
	// the second optinal argument(p) bool type
	// will indicates the key is present or not
	v, p := m["k9"]
	fmt.Println(v, p)

	for k, v := range m {
		fmt.Printf("%s => %d\n", k, v)
	}

	nums := []int{2, 4, 5, 6}
	for _, val := range nums {
		fmt.Println(val)
	}
}
