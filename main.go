package main

import (
	"fmt"
	"log"
	"os"
	"example.com/accountbalance"
)

func main() {
	baseURL := "https://backend.testnet.alephium.org"
	address := os.Args[1] //"1DrDyTr9RpRsQnDnXo2YRiPzPW4ooHX5LLoqXrqfMrpQH"
	ab, err := accountbalance.GetAccountBalance(baseURL, address)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(ab)
}
