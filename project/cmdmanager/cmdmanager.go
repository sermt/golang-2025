package cmdmanager

import "fmt"

type CMDManager struct {
	// Command string
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	var prices []string
	fmt.Println("Please enter your prices. Confirm every price with Enter.")
	for {
		var price string
		fmt.Print("Enter price: ")
		fmt.Scan(&price)
		if price == "0" {
			fmt.Println("Exiting...")
			break
		}

		prices = append(prices, price)

	}
	return prices, nil
}

func (cmd CMDManager) WriteJSON(data interface{}) error {
	fmt.Println(data)
	return nil
}

func NewCMDManager() CMDManager {
	return CMDManager{}
}
