package main

import (
	"bufio"
	"fmt"
	"log" // Fatal Function
	"os"
	"strconv" // TrimSpace()를 사용하기 위함, 공백 엔터 제거
	"strings" // ParseInt()를 사용하기 위함.
)

func main() {
	fmt.Print("단 입력: ")
	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n')

	if err != nil { // error 발생 시. nil은 아무것도 없다는 뜻. 다른 언어의 null과 같음.
		log.Fatal(err)
	}
	in = strings.TrimSpace(in)
	// dan, err := strconv.ParseInt(in, 10, 32) // base int는 몇 진수냐 이런거임.\
	dan, err := strconv.Atoi(in) // Atoi함수는 int형으로 바꿔줌.
	if err != nil {
		log.Fatal(err)
	}
	/* for문으로 while문 구현. Go에는 while문이 없다.
	i := 1
	for i < 10 {
		fmt.Printf("%d * %d = %d \n", dan, i, (dan * i))
		i++
	}
	*/
	for i := 1; i < 10; i++ {
		// fmt.Println(dan, "*", i, "=", (int(dan) * i))
		// fmt.Println(dan, "*", i, "=", (dan * i))
		fmt.Printf("%d * %d = %d \n", dan, i, (dan * i))
	}

}
