package others

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomNum() {
	var min, max int

	fmt.Print("Enter minimum number: ")
	fmt.Scan(&min)

	fmt.Print("Enter maximum number: ")
	fmt.Scan(&max)

	// Seed random generator
	rand.Seed(time.Now().UnixNano())

	// Generate random number
	randomNum := rand.Intn(max-min+1) + min

	fmt.Println("Random Number is:", randomNum)
}

func Run_Random_Num_Generator() {
	fmt.Println("Generate Random Number....")
	GenerateRandomNum()
}
