package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	IsCpu, Ismem := false, true

	if IsCpu {
		cpu()
	} else if Ismem {
		headMem()
	}
}

// func randGen() {

// 	for i := 0; i < 1000000; i++ {
// 		n := rand.Intn(100)
// 		s := square(n)

// 		fmt.Printf("%d^2: %d\n", n, s)

//		}
//	}
func randGen() {
	var builder strings.Builder

	for i := 0; i < 1000000; i++ {
		n := rand.Intn(100)
		s := square(n)

		builder.WriteString(fmt.Sprintf("%d^2: %d\n", n, s))
	}

	fmt.Print(builder.String())
}

func square(n int) int {
	return n * n
}

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	IsCpu, Ismem := false, true

// 	if IsCpu {
// 		cpu()
// 	} else if Ismem {
// 		headMem()
// 	}
// }

// func randGen() {
// 	var buffer bytes.Buffer

// 	for i := 0; i < 1000000; i++ {
// 		n := rand.Intn(100)
// 		s := square(n)

// 		buffer.WriteString(fmt.Sprintf("%d^2: %d\n", n, s))  // The use of buffer is for good memory optomization howerver is this case it takes lot of memory as the bytesSlices is growing per each loop.

// 	}
// 	fmt.Println(buffer.String())
// }

// func square(n int) int {
// 	return n * n
// }
