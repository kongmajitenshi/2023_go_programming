// after (multi return)
package main

import (
	"fmt"
	"log"
	"os"
)

func processScore(kor int, eng int, math int) (int, int) {
	totalScore := kor + eng + math
	average := totalScore / 3

	return totalScore, average
	// fmt.Printf("%s의 총점은 %d점 입니다. 평균은 %d점 입니다. \n", name, totalScore, average)
}

func isPrime(n int) (bool, error) {
	prime := true

	if n < 2 {
		return false, fmt.Errorf("%d는(은) 소수가 아닙니다.", n)
	}

	for i := 2; i < n; i++ { // 1과 number 일 땐 반복문 실행x
		if n%i == 0 {
			prime = false
			break
		}
	}
	return prime, nil // 소수판정값, 에러
}

func main() {
	var number int

	fmt.Print("점수 입력: ")
	_, err := fmt.Scanln(&number)

	if err != nil {
		log.Fatal(err)
	}

	p, err := isPrime(number)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if p == true {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}

	// // after (multi return)
	// package main

	// import (
	// 	"fmt"
	// 	"log"
	// 	"os"
	// )

	// func processScore(kor int, eng int, math int) (int, int) {
	// 	totalScore := kor + eng + math
	// 	average := totalScore / 3

	// 	return totalScore, average
	// 	// fmt.Printf("%s의 총점은 %d점 입니다. 평균은 %d점 입니다. \n", name, totalScore, average)
	// }

	// func isPrime(n int) bool {
	// 	prime := true

	// 	for i := 2; i < n; i++ { // 1과 number 일 땐 반복문 실행x
	// 		if n%i == 0 {
	// 			prime = false
	// 			break
	// 		}
	// 	}
	// 	return prime // true이면 소수, false이면 소수아님.
	// }

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

	// 	if isPrime(number) {
	// 		fmt.Println(number, "는(은) 소수입니다.")
	// 	} else {
	// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
	// 	}

	// var t int
	// var a int

	// t, a = processScore(100, 90, 93)
	// fmt.Printf("%s님의 총점은 %d점, 평균은%d점 입니다.\n", "홍길동", t, a)

	// t, a = processScore(100, 100, 97)
	// fmt.Printf("%s님의 총점은 %d점, 평균은%d점 입니다.\n", "최지은", t, a)
	// kor := 100
	// eng := 92
	// math := 93
	// name := "홍길동"

	// fmt.Println(name, "의 총점은", kor+eng+math, "입니다. 평균은", (kor+eng+math)/3.0, "입니다.")
}
