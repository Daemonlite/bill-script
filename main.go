package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string,r *bufio.Reader) (string,error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}


func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name , _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)
	
	b := newBill(name)
	
	fmt.Println("Created Bill", b.name)

	return b
}

func main() {

	myBil := createBill()

	fmt.Println(myBil)
}