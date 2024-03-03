package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
	"log"
)

func main() {
	filePtr, err := os.Open("./income.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer filePtr.Close()

	totalCredit := 0.0
	totalDebit := 0.0

	scanner := bufio.NewScanner(filePtr)
	for scanner.Scan() {
		line := scanner.Text()

		//get the num
		splittedString := strings.Split(line, "k")
		numberString := splittedString[0];

		// Parse the line
		amount, err := strconv.Atoi(numberString)
		if err != nil {
			fmt.Println("Error parsing line:", line)
			continue
		}

		// check the sign
		if strings.HasPrefix(line, "-") {
			totalCredit += math.Abs(float64(amount))
		}else if strings.HasPrefix(line, "+") {
			totalDebit += math.Abs(float64(amount))
		}
		
	}

	// GET THE BALANCE
	balance := totalCredit - totalDebit

	//STORE RESULTS IN RESULTS.TXT FILE
	storeResults(totalCredit, totalDebit, balance)

	// GET RESULTS FROM FILE
	readResults("./results.txt")
}


func storeResults(totalCredit float64, totalDebit float64, balance float64) {
	// WRITE TO results.txt FILE
	resPtr, err := os.Create("./results.txt")
	if err != nil{
		fmt.Println("Error creating the file:", err)
		return
	}
	defer resPtr.Close()

	resPtr.WriteString(fmt.Sprintf("Credit: K%.2f\n", totalCredit))
	resPtr.WriteString(fmt.Sprintf("Debit: K%.2f\n", totalDebit))
	resPtr.WriteString(fmt.Sprintf("Balance: K%.2f\n", balance))
}


func readResults(file string){
	// OPEN results.txt FILE
	resPtr, err := os.Open(file)

	if err != nil{
		log.Fatalf("failed to open the file: %s", err)
	}else{
		scanner := bufio.NewScanner(resPtr)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan(){
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil{
			log.Fatalf("failed to scan file: %s", err)
		}
	}
}