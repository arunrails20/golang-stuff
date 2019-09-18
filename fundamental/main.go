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

}

func constEx() {
  fmt.Println(s)
  const b string = "java"
  fmt.Println(b)
  // will raise error "cannot assign to b"
  //b = "py"
  //fmt.Println(b)
}
