package others

import "fmt"

func hammingWeight(num uint32) int {
    cnt := 0
    for num != 0 {
        cnt = cnt + int(num%2)
        num >>= 1
    }
    return cnt
}

func Run_Hamming_weight() {
	fmt.Println("Hamming weight of 00000000000000000000000000001011 is", hammingWeight(00000000000000000000000000001011))
	fmt.Println("Hamming weight of 00000000000000000000001000100101 is", hammingWeight(00000000000000000000001000100101))
}