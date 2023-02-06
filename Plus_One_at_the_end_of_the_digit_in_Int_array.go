package main

import "fmt"

func plusOne(digits []int) []int {
    lenDigits := len(digits)
    for i:=lenDigits-1; i>=0 ; i-- {
        if i == lenDigits-1 {
            digits[i]++
        }
        if digits[i] == 10 {
            digits[i] = 0
            if i != 0 {
                digits[i-1]++
            } else {
                digits = prepend(digits, 1)
            }
        }
    }
    return digits
}

func prepend(x []int, y int) []int {
    x = append(x, 0)
    copy(x[1:], x)
    x[0] = y
    return x
}

func main() {
	fmt.Println("After adding plus One to [1,4,5], the array is:::", plusOne([]int{1,4,5}))
	fmt.Println("After adding plus One to [1,4,9], the array is:::", plusOne([]int{1,4,9}))
	fmt.Println("After adding plus One to [9,0,9], the array is:::", plusOne([]int{9,0,9}))
	fmt.Println("After adding plus One to [9,9,9], the array is:::", plusOne([]int{9,9,9}))

}