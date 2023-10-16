package main

import (
	"fmt"
	"strings"
)

func main() {
	texts := "G@ M@ney"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)
	// 	zero value = 초기화 안했을 때 가지는 값.
	// 	var a int
	// // 	var b float64
	// // 	var c rune   // 작은따옴표는 rune이라는 타입임. 대입은 못하고 출력만 가능.
	// // 	var d bool   // false
	// // 	var e string // 빈 문자열
	// // 	// 변수 명 시작이 대문자 = 외부 패키지에서 접근 가능.
	// // 	// 소문자는 동일 패키지 내에서만.
	// // 	// 영문자로 시작해야함.
	// // 	var studentId string
	// // 	var i int32

	// // 	fmt.Println(a)
	// // 	fmt.Println(b)
	// // 	fmt.Println(c)
	// // 	fmt.Println(d)
	// // 	fmt.Println(e)
	// // 	fmt.Println(studentId)
	// // 	fmt.Println(i)
	// // 	fmt.Print("----------\n")
	// // 	fmt.Println(reflect.TypeOf(d))
	// // 	fmt.Println(reflect.TypeOf(e))
}
