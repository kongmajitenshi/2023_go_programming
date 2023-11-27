package main

import "fmt"

func status(name string) {
	balls := map[string]int{"성기훈": 20, "오일남": 0}
	ball, exists := balls[name]
	// fmt.Println(ball, exists)
	if !exists {
		fmt.Println(name, "님은 참여자가 아닙니다.")
	} else if ball < 1 {
		fmt.Println(name, "님은", balls[name], "개로 탈락입니다.")
	} else {
		fmt.Println(name, "님은", balls[name], "개로 통과입니다.")
	}
}

func main() {
	status("오일남")
	status("강철")
	status("성기훈")

}
