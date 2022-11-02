// package main

// import "fmt"

// const (
// 	Low = iota
// 	Medium
// 	High
// )

// type PriorityQueue[P comparable, V any] struct {
// 	items      map[P][]V
// 	priorities []P
// }

// func (pq *PriorityQueue[P, V]) Add(priority P, value V) {
// 	pq.items[priority] = append(pq.items[priority], value)
// }

// func (pq *PriorityQueue[P, V]) Next() (V, bool) {
// 	for i := 0; i < len(pq.priorities); i++ {
// 		priority := pq.priorities[i]
// 		items := pq.items[priority]
// 		if len(items) > 0 {
// 			next := items[0]
// 			pq.items[priority] = items[1:]
// 			return next, true
// 		}
// 	}
// 	var empty V
// 	return empty, false
// }

// func NewPriorityQueue[P comparable, V any](priorities []P) PriorityQueue[P, V] {
// 	return PriorityQueue[P, V]{items: make(map[P][]V), priorities: priorities}
// }

// func main() {
// 	queue := NewPriorityQueue[int, string]([]int{High, Medium, Low})
// 	queue.Add(Low, "L-1")
// 	queue.Add(High, "H-1")

// 	fmt.Println(queue.Next())

// 	queue.Add(Medium, "M-1")
// 	queue.Add(High, "H-2")
// 	queue.Add(High, "H-3")

// 	fmt.Println(queue.Next())
// 	fmt.Println(queue.Next())
// 	fmt.Println(queue.Next())
// 	fmt.Println(queue.Next())
// 	fmt.Println(queue.Next())
// }

//--Summary:
//  Write a generic function to complete the existing code.
//
//--Requirements:
//* Write a single function named `clamp` that can restrict a value
//  to a specific range
//  - The function should work with floating point numbers, integers
//    and arbitrary type aliases
//  - Use the existing `clamp` function signature and comments as a
//    starting point
//
//--Notes:
//* The existing code relies on a proper implementation of the clamp function
//* The program will print "Exercise complete!" when properly implemented

package main

import (
	"fmt"
	"constraints"
)

type Distance int32
type Velocity float64

type Number interface {
	constraints.Float;
	constraints.Interger
}
// The `clamp` function returns a value that has been "clamped"
// within a specific range of numbers. The `min` value is the
// minimum allowable and the `max` value is the maximum allowable.
// If the initial value lies within the range of `min` and `max`,
// then it should be returned as-is.
// Mathematically:
//   min <= value <= max


func clamp(T Number)(value, min, max T) T {
if value>max {
	return max
}else if value <min{
	return min
}else{
	return value
}
}


func testClampInt8() {
	var (
		min int8 = -10
		max int8 = 10
	)
	if clamp(-30, min, max) != min {
		panic("clamp failed for int8")
	}
}

func testClampUint32() {
	var (
		min uint32 = 0
		max uint32 = 10
	)
	if clamp(20, min, max) != max {
		panic("clamp failed for uint32")
	}
}

func testClampFloat32() {
	var (
		min float32 = 5.5
		max float32 = 9.9
	)
	if clamp(0, min, max) != min {
		panic("clamp failed for float32")
	}
}

func testClampDistance() {
	var (
		min Distance = 0
		max Distance = 100
	)
	if clamp(Distance(99), min, max) != 99 {
		panic("clamp failed for Distance")
	}
}

func testClampVelocity() {
	var (
		min Velocity = 0.0
		max Velocity = 99.9
	)
	if clamp(Velocity(100), min, max) != max {
		panic("clamp failed for Velocity")
	}
}

func main() {
	testClampInt8()
	testClampUint32()
	testClampFloat32()
	testClampDistance()
	testClampVelocity()
	fmt.Println("Exercise complete!")
}
