// 포인터

package main

import (
	"fmt"
)

func main() {
	a := 10
	var b int = 20
	var pa *int = &a

	fmt.Printf("%T %T %T\n", &a, pa, &pa)
	fmt.Printf("%x %x %x\n", &a, pa, &pa)
	fmt.Println(&a, pa, &pa)
	fmt.Println(*pa)
	pa = &b
	fmt.Println(*pa)
	
}
// package main

// import (
// 	"fmt"
// )
// func swap(n1 *int, n2 *int) {
// 	fmt.Println(n1,n2)
// 	temp := *n1
// 	*n1 = *n2
// 	*n2 = temp
// }

// func main() {
// 	a := 10
// 	b := 20

// 	c := &a // var c *int = &a
// 	fmt.Printf("%T\n", c)

// 	fmt.Println(a, b)
// 	swap(&a,&b) // pass by pionter
// 	fmt.Println(a,b)
// }


// package main

// import (
// 	"fmt"
// )

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

// func double(n int) {
// 	n = n * 2
// }
// func main() {
// 	var a int = 6
// 	double(a)
// 	fmt.Println(a)
// }
