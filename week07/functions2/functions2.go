package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {

	if n < 2 {
		return false, fmt.Errorf("%d는(은) 소수가 아닙니다.", n)
	}

	for i := 2; i < n; i++ { // 1과 n 일 땐 반복문 실행x
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil // 소수판정값, 에러
}

func prime(n int) {
	p, err := isPrime(n)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if p {
		fmt.Println(n, "는(은) 소수입니다.")
	} else {
		fmt.Println(n, "는(은) 소수가 아닙니다.")
	}
}

func primeRange(a int, b int) {
	if a > b {
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isPrime(i)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if p {
			fmt.Print(i, " ")
		}
	}

}

// 구간 소수 판정 프로그램 v1.7
func main() {
	var menu int

	for true {
		fmt.Print("MENU : 1)소수판정 2)구간소수판정 : ")
		_, err := fmt.Scanln(&menu)

		if err != nil {
			log.Fatal(err)
		}

		switch menu {
		case 1:
			var in int
			fmt.Print("정수 1개 입력: ")
			fmt.Scanln(&in)
			prime(in)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		case 2:
			var n1, n2 int
			fmt.Print("정수 2개 입력: ")
			fmt.Scanln(&n1, &n2)
			primeRange(n1, n2)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		default:
			fmt.Println("프로그램을 종료합니다.")
			os.Exit(0)
		}
	}

}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool, error) {

// 	if n < 2 {
// 		return false, fmt.Errorf("%d는(은) 소수가 아닙니다.", n)
// 	}

// 	for i := 2; i < n; i++ { // 1과 n 일 땐 반복문 실행x
// 		if n%i == 0 {
// 			return false, nil
// 		}
// 	}
// 	return true, nil // 소수판정값, 에러
// }

// // 구간 소수 판정 프로그램 v1.4 isPrime()함수 속 변수를 줄이고 return 구문 추가. break
// func main() {
// 	var a, b int

// 	fmt.Print("점수 입력: ")
// 	_, err := fmt.Scanln(&a, &b)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if a > b {
// 		temp := a
// 		a = b
// 		b = temp
// 	}

// 	for i := a; i <= b; i++ {
// 		p, err := isPrime(i)

// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(0)
// 		}
// 		if p {
// 			fmt.Print(i, " ")
// 		}
// 	}

// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool, error) {
// 	prime := true

// 	if n < 2 {
// 		return false, fmt.Errorf("%d는(은) 소수가 아닙니다.", n)
// 	}

// 	for i := 2; i < n; i++ { // 1과 n 일 땐 반복문 실행x
// 		if n%i == 0 {
// 			prime = false
// 			break
// 		}
// 	}
// 	return prime, nil // 소수판정값, 에러
// }

// // 구간 소수 판정 프로그램 v1.3
// func main() {
// 	var a, b int

// 	fmt.Print("점수 입력: ")
// 	_, err := fmt.Scanln(&a, &b)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if a > b {
// 		temp := a
// 		a = b
// 		b = temp
// 	}

// 	for i := a; i <= b; i++ {
// 		p, err := isPrime(i)

// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(0)
// 		}
// 		if p {
// 			fmt.Print(i, " ")
// 		}
// 	}

// }
