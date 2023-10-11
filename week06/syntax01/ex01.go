package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var c rune = '가' // 작은따옴표는 rune이라는 타입임. 대입은 못하고 출력만 가능.
	var a int16 = 7
	var b = 7
	d := 7.9 // 이 방법을 많이 사용함. 값을 보고 타입을 자동으로 결정함. 이후에 값이 변경은 가능하나, 타입이 변경되진 않음.
	e := int(d)
	fmt.Println(d)
	fmt.Printf("d의 타입: %T \n", d)
	f := false
	fmt.Println(f)
	fmt.Printf("%T \n", f)

	fmt.Println(e)

	fmt.Println(a)
	fmt.Printf("a의 타입: %T \n", a)
	fmt.Println(b)
	fmt.Printf("b의 타입: %T \n", b)

	fmt.Println(c)         // 유니코드(UTF-8) 값 출력
	fmt.Printf("%c \n", c) // c에 넣은 '가' 출력
	fmt.Printf("%T \n", c) // c의 타입 출력, rune은 int32의 별명이다.

	fmt.Println(math.Ceil(2.71))
	fmt.Println(math.Floor(2.71))
	fmt.Println(math.Round(2.71))
	fmt.Println("Hello Go~")
	fmt.Println(strings.Title("go git github java"))
}
