package main

import (
	"fmt"
)

type queue []int

func main() {
	q := newQueue(10)
	q.print()
	q.checkEvenOdd()
}

func newQueue(count int) queue {
	q := queue{}

	for i := range count {
		q = append(q, i+1)
	}
	return q
}

func (q queue) print() {
	res := ""

	for _, v := range q {
		res += fmt.Sprintf("%d,", v)
	}

	fmt.Println(res)

}

func (q queue) checkEvenOdd() {
	for _, v := range q {
		if v%2 == 0 {
			fmt.Printf("%d is even\n", v)
		} else {
			fmt.Printf("%d is odd\n", v)
		}
	}
}
