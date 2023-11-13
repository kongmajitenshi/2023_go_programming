package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	// s := make([]int, 5) // 단축 연산자. make함수를 이용해 슬라이스 생성, 메모리 할당, zero value값 적용

	s := []int{0, 0, 0, 0, 0} // 슬라이스 리터럴을 이용해 슬라이스 생성, 메모리 할당, 초기화 진행

	s[4] = 99
	s[2] = 91
	for _, value := range s {
		fmt.Println(value)
	}
}
