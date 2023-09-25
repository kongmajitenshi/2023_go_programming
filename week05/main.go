package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input Score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	inputScoreString = strings.TrimSpace(inputScoreString)      // remove spacebar
	inputScore, err := strconv.ParseFloat(inputScoreString, 32) // casting
	if err != nil {
		log.Fatal(err)
	}

	if inputScore >= 90 { 
		grade := "A grade!"
	} else {
		grade := "BCDE grade~"
	}
	fmt.Println(inputScore)
	fmt.Println("You got" + grade)

