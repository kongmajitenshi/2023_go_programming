package main

import "fmt"

func main() {
	// var games map[int]string
	// games = make(map[int]string)
	// games := make(map[int]string)
	games := map[int]string{
		games[456] = "성기훈",
		games[218] = "박해수",
		games[067] = "강새벽",
		games[001] = "오일남",
		games[199] = "알리",
		games[101] = "아이오아이",
}
	// 추가(append / map에는 append메소드 없음)
	// games[456] = "성기훈"
	// games[218] = "박해수"
	// games[067] = "강새벽"
	// games[001] = "오일남"
	// games[199] = "알리"
	// games[101] = "아이오아이"

	for _, v := range games {
		fmt.Println(v)
	}

	games[101] = "장덕수" // update
	delete(games, 199) // delete
	for k, v := range games {
		fmt.Println(k, v)
	}
}
