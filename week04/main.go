package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	var a int
	a = 4
	fmt.Println(a)

	var b int = 9
	fmt.Println(b)

	var c = 5
	fmt.Println(c)

	d := 15 // short form
	fmt.Println(d, reflect.TypeOf(d))

	// e := 2.71 < 이렇게 하면 double로 알아서 해줌
	var e float32 = 2.71
	fmt.Println(e, reflect.TypeOf(e))
	fmt.Println(float32(d) * e) // 그냥 d*e를 하면 d는 정수, e는 실수라서 계산 안됨.
	fmt.Println(float32(d) > e) // 15.0 > 2.71, 이런 형변환을 '캐스팅' 이라고 함.

	f := 'Z'
	fmt.Println(f, reflect.TypeOf(f))

	g, h := 4.10, 3.99
	fmt.Println(g, reflect.TypeOf(g), h, reflect.TypeOf(h))

	i := "문자열"
	fmt.Println(i, reflect.TypeOf(i))

	var j int     // 0 출력
	var k float32 // 0 출력
	var l bool    // false 출력
	var m string  // 빈 문자열이라 빈 칸 한 칸 출력됨
	N := "변수명이 대문자로 시작하면 다른 패키지에서 이 변수를 사용할 수 있음"
	koreanZombie := "정찬성"
	fmt.Println(j, k, l, m, N, koreanZombie)

	fmt.Printf("%d %f %s", j, k, m)

	//  fmt.Println(reflect.TypeOf(99))
	//  fmt.Println(reflect.TypeOf('대')) // rune, int32
	//  fmt.Println(reflect.TypeOf(2.7))
	//  fmt.Println(reflect.TypeOf(false))
	//  fmt.Println(reflect.TypeOf("Go!"))
	//  fmt.Println('A', 'a', '0', '김', '인', '하', '\n') // 65, 97, 48
	//  fmt.Println(math.Ceil(3.17))
	fmt.Println(math.Floor(3.17))
	//fmt.Println.(strings.Title("오픈소스 프로그래밍~"))
	fmt.Println(strings.Title("open source \t programming~\n\"go\"!"))
	// fmt.Println(math.Floor("go!"))
	// fmt.Println(strings.Title(3.14))

}
