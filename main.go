package main

import (
	"fmt"
	bank "training/bank"
	investment "training/investment_calculator"
	profit "training/profit_calculator"
)

func main() {
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Run Bank App")
		fmt.Println("2. Calculate Investment")
		fmt.Println("3. Calculate Profit")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			bank.App()
		case 2:
			profit.App()
		case 3:
			investment.App()
		case 4:
			return
		}
	}
}