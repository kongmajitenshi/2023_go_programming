package main

import (
	"fmt"
)
func swap(n1 *int, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

func main() {
	a := 10
	b := 20

	c := &a // var c *int = &a
	fmt.Printf("%T\n", c)

	fmt.Println(a, b)
	swap(&a,&b) // pass by pionter
	fmt.Println(a,b)
}


// package main

// import (
// 	"fmt"
// )

// // 포인터
// func double(n *int) {
// 	*n = *n * 2
// }
// func main() {
// 	var a int = 6
// 	double(&a)
// 	fmt.Println(a)
// }


// package main

// import (
// 	"fmt"
// )

// // 포인터
// func double(n int) {
// 	n = n * 2
// }
// func main() {
// 	var a int = 6
// 	double(a)
// 	fmt.Println(a)
// }
