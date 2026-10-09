/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn

// Takes condition and returns [IfBranch] branch.
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
func If(cond bool) IfBranch {
	return IfBranch(cond)
}

// The [IfBranch] provide methods to construct the [Then].
type IfBranch bool

// Passes the value for [IfBranch] if it is true.
// Returns the [Then].
func (i IfBranch) Then[V any](v V) Then[V] {
	return Then[V]{bool(i), v}
}

// Passes the closure returns value for [IfBranch] if it is true.
// Returns the [Then].
func (i IfBranch) ThenF[T any](f func() T) Then[T] {
	if i {
		return Then[T]{bool: true, v: f()}
	}
	return Then[T]{bool: false}
}

// The [Then] provide methods to evaluate the [IfBranch].
type Then[V any] struct {
	bool
	v V
}

// Passes the value for [IfBranch] if it is false, evaluates the [IfBranch]
// and returns the result.
//
// If [IfBranch] is true, then returns the value early passed to [IfBranch.Then] (or to [IfBranch.ThenF]),
// elsewhere return passed value.
func (t Then[V]) Else(v V) V {
	if t.bool {
		return t.v
	}
	return v
}

// Passes the closure returns value for [IfBranch] if it is false, evaluates the [IfBranch]
// and returns the result.
//
// If [IfBranch] is true, then returns the value early passed to [IfBranch.Then] (or to [IfBranch.ThenF]),
// elsewhere returns the passed closure result.
func (t Then[V]) ElseF(f func() V) V {
	if t.bool {
		return t.v
	}
	return f()
}
