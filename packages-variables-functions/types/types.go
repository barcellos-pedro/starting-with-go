package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	age    int        = 25
)

func main() {
	// type conversion
	// var ageFloat float32 = float32(age)
	ageFloat := float32(age) // or shorthand

	// constants, cant use ':='
	const a = "a"

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println(age, ageFloat)
	fmt.Println(a)
	fmt.Println("E:", math.E)
	fmt.Println("PI:", math.Pi)
}

/** Go Types */

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32 | represents a Unicode code point

// float32 float64

// complex64 complex128
