package main

import "fmt"

func lengthOfLastWord(s string) int {
    count :=0
    lenS := len(s)
    for i:=lenS-1; i>=0; i-- {
        if string(s[i]) == " " && count == 0 {
            continue
        } else {
            if string(s[i]) == " " && count > 0 {
                break
            } else {
                count++
            }
        }
    }
    return count
}

func main() {
	fmt.Println("Length of the last word in \"Hello World\" is:::", lengthOfLastWord("Hello World"))
	fmt.Println("Length of the last word in \"Hello World KirtiDhi   \" is:::", lengthOfLastWord("Hello World KirtiDhi   "))
}