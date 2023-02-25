package others

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	aLen, bLen := len(a), len(b)
	carry, sum, i, j := 0, 0, aLen-1, bLen-1
	output := ""

	for i >= 0 && j >= 0 {
		first := int(a[i] - '0')
		sec := int(b[j] - '0')
		// sum, carry = binarySum(first, sec, carry)
		sum, carry = binarySumUsingMaths(first, sec, carry)
		output = strconv.Itoa(sum) + output
		i--
		j--
	}

	for i >= 0 {
		first := int(a[i] - '0')
		// sum, carry = binarySum(first, 0, carry)
		sum, carry = binarySumUsingMaths(first, 0, carry)
		output = strconv.Itoa(sum) + output
		i--
	}

	for j >= 0 {
		sec := int(b[j] - '0')
		// sum, carry = binarySum(0, sec, carry)
		sum, carry = binarySumUsingMaths(0, sec, carry)
		output = strconv.Itoa(sum) + output
		j--
	}

	if carry > 0 {
		output = strconv.Itoa(1) + output
	}

	return output
}

func binarySum(a, b, carry int) (int, int) { //sum,carry
	output := a + b + carry

	switch output {
	case 0:
		return 0, 0
	case 1:
		return 1, 0
	case 2:
		return 0, 1
	case 3:
		return 1, 1
	}
	return 0, 0
}

func binarySumUsingMaths(a,b,carry int) (int,int) { //sum(using modulus),carry(using division)
    return (a + b + carry)%2, (a + b + carry)/2
}

func Run_Binary_Sum() {
	fmt.Println("Binary sum of 1100 & 10 is:::", addBinary("1100", "10"))
	fmt.Println("Binary sum of 1111 & 1111 is:::", addBinary("1111", "1111"))
	fmt.Println("Binary sum of 110010 & 10001 is:::", addBinary("110010", "10001"))
	fmt.Println("Binary sum of 11 & 10 is:::", addBinary("11", "10"))
}
