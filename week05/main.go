package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input Score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore := reader.ReadString('\n') // 1 variable but reader.ReadString returns 2 values
	fmt.Println(inputScore)
}
