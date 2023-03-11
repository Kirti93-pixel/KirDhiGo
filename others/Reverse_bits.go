package others

import "fmt"

func reverseBits(num uint32) uint32 {
    if num == 0 {return 0}

    var res uint32 = 0
    for i:=0; i<32; i++ {
        res<<=1
        if ((num&1)==1) {
            res++
        }
        num>>=1
    }
    return res
}

func Run_Reverse_bits() {
	fmt.Println("Reversed num 4294967293 is", reverseBits(uint32(4294967293))) // binary 11111111111111111111111111111101
	fmt.Println("Reversed num 43261596 is", reverseBits(uint32(43261596))) // binary 00000010100101000001111010011100
}