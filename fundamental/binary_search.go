package main

import (
	"fmt"
	"sort"
)

func main() {
  arr := []int{22,33,12,11,4,5,6,80}
	sort.Ints(arr)
	fmt.Println(binary_search(arr))
}

func binary_search(arr []int) bool{
  size := len(arr)
  k := 33
  low := 0
  high := size - 1
  mid := 0
  for low <= high {
    mid = low + (high - low) / 2
    if arr[mid] == k {
      return true
    } else if k < arr[mid] {
      high = mid - 1
    } else {
      low = mid + 1
    }
  }
  return false
}
