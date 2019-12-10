package main

import (
		"fmt"
)

func main(){
	bundeslandPop := map[string]int{
		"NRW": 120312,
		"BAY": 123213,
		"BER": 393223,
	}
	//m := map[[3]int]string{}
	fmt.Println(bundeslandPop["BER"])
}