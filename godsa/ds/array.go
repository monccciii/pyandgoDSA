package main

import "fmt"

//Given an array of integers, write a function that
// finds the second largest element in the array.
//The function should return -1 if the array is too small
// (i.e. less than 2 elements).
// How would you implement this in Go?

//Pseudo code:
//initialize the array, create a function which takes an array as the params
//the function checks the length of the array and verifies it's not less than 2
//if the function is not less than two, continue with finding the largest element
// to find the largest element, max will be declared as the first element
// each element from the first element on will be checked to see if it is larger than max
// if larger than max, max gets set to be that element
// element max gets popped from array, search begins again.

func calcLargestElement(arr []int) int {
	if len(arr) < 2 {
		return -1
	} else {
		max := arr[0]
		secondmax := arr[0]
		for i:= 0;i<len(arr);i++{
			if max<arr[i]{
				secondmax=max
				max=arr[i]
			}
		}
		return secondmax
	}
}
func main() {
	var myArray = []int{1, 4, 5, 3, 2}
	fmt.Println(calcLargestElement(myArray))
}
