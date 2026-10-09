/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn

// Takes condition and returns [Condition] branch.
//
// # Example:
//
// Simple return string:
//	s := If(a == 1).
//		Then("one").
//		Else("not one")
//
// Lazy evaluate by condition:
//	s := If(a != nil).
//		ThenF(func() string { return a.String() }).
//		Else("nil")
//
// Nested conditions:
//	s :=
// 		If(m == nil).Then("nil").Else(
//			If(len(m) == 0).Then("empty").Else(
//				If(len(m) == 1).Then("one item").Else(
// 					fmt.Sprintf("%m items", len(m)))))
func If(cond bool) Condition {
	return Condition(cond)
}

// The [Condition] provide methods to construct the [Branch].
type Condition bool

// Passes the value for [Condition] if it is true.
// Returns the [Branch].
func (c Condition) Then[T any](v T) Branch[T] {
	return Branch[T]{bool(c), v}
}

// Passes the closure returns value for [Condition] if it is true.
// Returns the [Branch].
func (c Condition) ThenF[T any](f func() T) Branch[T] {
	if c {
		return Branch[T]{c: true, v: f()}
	}
	return Branch[T]{c: false}
}

// The [Branch] provide methods to evaluate the [Condition].
type Branch[T any] struct {
	c bool
	v T
}

// Passes the value for [Condition] if it is false, evaluates the [Condition]
// and returns the result.
//
// If [Condition] is true, then returns the value early passed to [Condition.Then] (or to [Condition.ThenF]),
// elsewhere return passed value.
func (t Branch[T]) Else(v T) T {
	if t.c {
		return t.v
	}
	return v
}

// Passes the closure returns value for [Condition] if it is false, evaluates the [Condition]
// and returns the result.
//
// If [Condition] is true, then returns the value early passed to [Condition.Then] (or to [Condition.ThenF]),
// elsewhere returns the passed closure result.
func (t Branch[T]) ElseF(f func() T) T {
	if t.c {
		return t.v
	}
	return f()
}
