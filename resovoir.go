package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ReservoirSample(S []int, R *[]int) {
	// caching length in k to prevent double call
	k := len(*R)-1
	// filling the reservoir array
	for index:=0; index<=k; index++ {
		(*R)[index] = S[index]
	}
	// replacing elements with gradually increasing probability
	for index:=k+1; index<len(S); index++ {
		randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
		j := randGen.Intn(index+1)
		if j<=k {
			(*R)[j] = S[index]
		}
	}
}

func getList() []int{
	return []int{1,2,3,4,5,6,7,8}
}

func main() {
	// using this to try out the function

	// defining lists, an indexed list and an empty list to fill
	S:=getList()
	R := make([]int, 5)
	// call to the function
	ReservoirSample(S, &R)
	// just displaying the result
	fmt.Println(R)
}