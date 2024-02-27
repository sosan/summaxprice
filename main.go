package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func combinacionPrecios(precios []int, total int) []int {
	sort.Ints(precios)
	n := len(precios)
	dp := make([][]int, total+1)
	for i := range dp {
		dp[i] = make([]int, 0)
	}

	for i := 0; i < n; i++ {
		for j := total; j >= precios[i]; j-- {
			if len(dp[j-precios[i]]) != 0 || j == precios[i] {
				if len(dp[j]) == 0 || len(dp[j-precios[i]])+1 > len(dp[j]) {
					dp[j] = append(dp[j-precios[i]][:0:0], dp[j-precios[i]]...)
					dp[j] = append(dp[j], precios[i])
				}
			}
		}
	}

	for i := total; i >= 0; i-- {
		if len(dp[i]) != 0 {
			return dp[i]
		}
	}

	return nil
}

func readPrices() []int {
	fmt.Print("Introduce precios separados por comas 12, 13, 14, 15: \n")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("An error occured while reading input. Please try again")
	}
	
	input = strings.TrimSuffix(input, "\n")
	splitted := strings.Split(input, ",")
	var listPrices []int
	for _, value := range(splitted) {
		trimmedNumber := strings.TrimSpace(value)
		currentPrice, err := strconv.Atoi(trimmedNumber)
		if err != nil {
			log.Fatalf("An error occured while reading input. Please try again")
		}
		listPrices = append(listPrices, currentPrice)
	}
	return listPrices
}

func readMaxPrice() int {
	fmt.Print("Introduce el precio maximo que tiene que alcanzar: \n")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("An error occured while reading input. Please try again")
	}
	
	input = strings.TrimSuffix(input, "\n")
	numberWithoutspace := strings.TrimSpace(input)
	maxPrice, err := strconv.Atoi(numberWithoutspace)
	if err != nil {
		log.Fatalf("An error occured while reading input. Please try again")
	}
	return maxPrice

}

func main() {
	listPrices := readPrices()
	if listPrices == nil {
		fmt.Println("An error occured while reading input. Please try again")
	}

	maxPrice := readMaxPrice()
	resultado := combinacionPrecios(listPrices, maxPrice)
	fmt.Printf("La combinación de precios que se acerca lo más posible a %d sin pasarse es: %v\n", maxPrice, resultado)
}
