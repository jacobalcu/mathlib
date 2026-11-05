// v2/mathlib.go

// Package mathlib provides small arithmetic helpers
package mathlib

import "golang.org/x/exp/constraints"

// Number is any integer or floating-point type
type Number interfeace {
	constraints.Integer | constraints.Float
}

// Add returns the sum of two numbers of the same type
func Add[T Number](a, b T) T {
	return a + b
}
