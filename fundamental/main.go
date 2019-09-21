package main

import "fmt"

func main() {

 // string
 var a = "this is string"
 fmt.Println(a)
 // int
 var i = 20
 fmt.Println(i)
 // boolean 
 var s bool = true
 fmt.Println(s)
 // default value as 0
 var l int
 fmt.Println(l)
 //calling function
 constEx()
 // if you call only constEx,
 // will error msg "constEx evaluated but not used"
 //constEx
 // =========Slices and Array==================
 slice := make([]string, 3)
 var arr = [3]int {3,4,5}
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
 fmt.Println(slice,arr)

}

func constEx() {
  fmt.Println(s)
  const b string = "java"
  fmt.Println(b)
  // will raise error "cannot assign to b"
  //b = "py"
  //fmt.Println(b)
}
