package main

import (
	"flag"
	"fmt"
	"os"
	"csv"
)

func main(){
	csvFileName := flag.String("csv", "quiz.csv", "a csv file in the format of question, answer")
	flag.Parse()
	file, err := os.Open(*csvFileName)

	if err != nil{
		fmt.Printf("Failed to open csv file: %s", *csvFileName)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failed to parse the csv file: %s", csvFileName)
		os.Exit(1)
	}
	fmt.Printf(lines)
}