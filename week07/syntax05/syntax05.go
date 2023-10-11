package main

import (
	"fmt"
)

// shadowing
func main() {
	// 자료타입을 변수명으로 사용함 > shadowing 발생
	var test1 float64 = 9.1
	var test2 float64 = 6.9
	fmt.Println(test1 + test2)

	var univ string = "inha"
	fmt.Println(univ)

	var f1 string = "functions"
	var f2 = append([]string{}, "함수")
	fmt.Println(f1)
	fmt.Println(f2)
	/*----------------------------------------------------*/
	// 자료타입을 변수명으로 사용함 > shadowing 발생
	// var float64 float64 = 9.1
	// var test float64 = 6.9
	// fmt.Println(test)

	// var fmt string = "inha"
	// fmt.Println()

	// var append string = "functions"
	// var f = append([]string{}, "함수")
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// // v0.9
// func main() {
// 	var number int
// 	fmt.Print("점수 입력: ")
// 	n, err := fmt.Scanln(&number) // 입력받은 정수의 개수를 n에 넣음. n자리에 _ 를 넣으면 사용 안하고 비워둘 수 있음.

// 	fmt.Println(n)

// 	if /*number == 0 || number == 1*/ number < 2 {
// 		fmt.Println(number, "는(은) 정수가 아닙니다.")
// 		os.Exit(0)
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true

// 	for i := 2; i < number; i++ { // 1과 number 일 땐 반복문 실행x
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}
// 	if isPrime == true {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"log"
// )

// // v0.8 fmt 패키지의 scan 계열 함수 사용할것.
// func main() {
// 	var number int
// 	fmt.Print("점수 입력: ")
// 	n, err := fmt.Scanln(&number) // 입력받은 정수의 개수를 n에 넣음. n자리에 _ 를 넣으면 사용 안하고 비워둘 수 있음.

// 	fmt.Println(n)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true

// 	for i := 2; i < number; i++ { // 1과 number 일 땐 반복문 실행x
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}
// 	if isPrime == true {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // v0.7
// func main() {
// 	fmt.Print("점수 입력: ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n') // \n전 까지 읽어옴

// 	if err != nil { // error 발생 시. nil은 아무것도 없다는 뜻. 다른 언어의 null과 같음.
// 		log.Fatal(err)
// 	}
// 	in = strings.TrimSpace(in)

// 	number, err := strconv.Atoi(in) // Atoi함수는 int형으로 바꿔줌.
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true

// 	for i := 2; i < number; i++ { // 1과 number 일 땐 반복문 실행x
// 		if number%i == 0 {
// 			isPrime = false
// 		}
// 	}
// 	if isPrime == true {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }
