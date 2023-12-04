package others

import "fmt"

func GasStation(gas, cost []int) int {
	n := len(gas)
	fuelLeft, globalFuelLeft, start := 0, 0, 0
	for i:=0; i<n; i++ {
		globalFuelLeft += gas[i]-cost[i]
		fuelLeft += gas[i]-cost[i]
		if fuelLeft < 0 {
			start=1+i
			fuelLeft = 0
		}
	}
	if globalFuelLeft < 0 {
		return -1
	}
	return start
}

func Run_Gas_Station() {
	fmt.Println("Gas station index:::", GasStation([]int{1,2,3,4,5}, []int{3,4,5,1,2}))
}
