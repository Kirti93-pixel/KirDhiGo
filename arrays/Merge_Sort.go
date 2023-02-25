package main

import "fmt"

func merge(a, b []int, res *[]int) {
	i, j, k := 0, 0, 0
	al, bl := len(a), len(b)
	for i < al && j < bl {
		if a[i] > b[j] {
			(*res)[k] = b[j]
			j += 1
		} else {
			(*res)[k] = a[i]
			i += 1
		}
		k += 1
	}

	if i < al {
		for x := i; x < len(a); x++ {
			(*res)[k] = a[x]
			k += 1
		}
	}
	if j < bl {
		for x := j; x < len(b); x++ {
			(*res)[k] = b[x]
			k += 1
		}
	}
	return
}

func mergeSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	mid := (n + 1) / 2
	left, right := make([]int, 0), make([]int, 0)
	for i := 0; i < mid; i++ {
		left = append(left, arr[i])
	}
	for i := mid; i < n; i++ {
		right = append(right, arr[i])
	}
	mergeSort(left)
	mergeSort(right)
	merge(left, right, &arr)
}

func main() {
	c := []int{3, 2, 5, 1, 54, 23, 67, 34, 65, 90, 21, 43, 48}
	fmt.Println("array:::", c)
	mergeSort(c)
	fmt.Println("sorted array:::", c)
}
