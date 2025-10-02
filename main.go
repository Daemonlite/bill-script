package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string,r *bufio.Reader) (string,error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}


func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)
	
	b := newBill(name)
	
	fmt.Println("Created Bill", b.name)

	return b
}

func proptOptions(b bill) {

	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)	
		price, _ := getInput("Item price: ", reader)

		p,err := strconv.ParseFloat(price,64)

		if err != nil {
			fmt.Println("Cannot convert price to float")
			proptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added...", name, price)

		proptOptions(b)

	case "s":
		fmt.Print("You chose to save the bill s: ",b)
		// save bill to txt
		b.save()
		fmt.Println("Bill saved to", b.name+".txt")

	case "t":
		tip, _ := getInput("Tip amount: ", reader)
		
		p,err := strconv.ParseFloat(tip,64)

		if err != nil {
			fmt.Println("Cannot convert tip to float")
			proptOptions(b)
		}

		b.updateTip(p)

		fmt.Println("Tip updated...", tip)

		proptOptions(b)
	
	default:
		fmt.Println("Not a valid option...")
		proptOptions(b)
	}

	
	
}

func main() {

	myBil := createBill()

	proptOptions(myBil)

}