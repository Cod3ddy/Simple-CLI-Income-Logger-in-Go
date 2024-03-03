package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
	"log"
	"flag"
)

// COLOR

type Color string

const (
	ColorRed = "\u001b[31m"
	ColorReset = "\u001b[0m"
)

func alert(color Color, message string){
	fmt.Println(string(color), message, string(ColorReset))
}


func main() {

	var sourceFile string
	var resultFile string


	flag.StringVar(&sourceFile, "input", "",  "Input file")
	flag.StringVar(&resultFile, "results", "",  "results file")

	flag.Parse()


	// IF SRC AND RESULTFILE ARE EMPTY
	if sourceFile == "" || resultFile == ""{
		alert(ColorRed, "Input file or Result file is missing!");
		return
	}

	// ELSE

	filePtr, err := os.Open(sourceFile)
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
	storeResults(totalCredit, totalDebit, balance, resultFile)

	// GET RESULTS FROM FILE
	readResults(resultFile)
}


func storeResults(totalCredit float64, totalDebit float64, balance float64, outputFile string) {
	// WRITE TO results.txt FILE
	resPtr, err := os.Create(outputFile)
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