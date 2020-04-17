package main

import (
	"fmt"
)

func main() {

	//var identityMAtrix [3][3]int = {[3][3]int{1,0,0},[3]int{0,1,0},[3]int{0,01}}
	//THe easiest way to see it is:
	var identityMAtrix [3][3]int
	identityMAtrix[0] = [3]int{1, 0, 0}
	identityMAtrix[1] = [3]int{0, 1, 0}
	identityMAtrix[2] = [3]int{0, 0, 1}

	fmt.Println(identityMAtrix)

}
