// v0.3
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().Unix() // seed 설정
	rand.Seed(seed)           // seed 설정

	isPrime := true
	number := rand.Intn(150) + 2 // '+2' 는 0과 1 제외. 2~151
	fmt.Println("임의로 추출된 수 :", number)

	for i := 2; i < number; i++ { // 1과 number 일 땐 반복문 실행x
		if number%i == 0 {
			isPrime = false
		}
	}
	if isPrime == true {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}

// v0.2
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	seed := time.Now().Unix() // seed 설정
// 	rand.Seed(seed)           // seed 설정

// 	count := 0
// 	number := rand.Intn(150) + 2 // '+2' 는 0과 1 제외. 2~151
// 	fmt.Println("임의로 추출된 수 :", number)

// 	for i := 2; i < number; i++ { // 1과 number 일 땐 반복문 실행x
// 		if number%i == 0 {
// 			count++ // count = count + 1
// 		}
// 	}
// 	if count == 0 {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }

// v 0.1
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	seed := time.Now().Unix() // seed 설정
// 	rand.Seed(seed)           // seed 설정

// 	count := 0
// 	number := rand.Intn(150) + 2 // '+2' 는 0과 1 제외. 2~151
// 	fmt.Println("임의로 추출된 수 :", number)

// 	for i := 1; i <= number; i++ {
// 		if number%i == 0 {
// 			count++
// 		}
// 	}
// 	if count == 2 {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	seed := time.Now().Unix() // seed 설정
// 	rand.Seed(seed)           // seed 설정

// 	for i := 1; i < 6; i++ {
// 		dice := rand.Intn(6) + 1
// 		fmt.Println(dice)
// 	}
// 	// number := rand.Intn(6) + 1 // 0에서부터 3전 까지(0,1,2) 랜덤 숫자
// 	// fmt.Println(number)
// }
