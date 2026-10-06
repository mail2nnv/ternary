/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn

// Takes condition and returns [ThenIntf] branch.
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
//	s := If(s == nil).
//		Then("nil").
//		ElseIf(len(s) == 0).
//		Then("empty slice").
//		ElseIf(len(s) == 1).
//		Then("one element slice").
//		Else(fmt.Sprintf("%d-element slice", len(s)))
func If(cond bool) Then {
	return Then(cond)
}

// The [ThenIntf] branch provide methods [ThenIntf.Then] and [ThenIntf.ThenF]
// to pass result if condition is true and returns [ElseIntf] branch.
type ThenIntf[T any] interface {
	// Takes value for true condition and returns [Else].
	Then(T) ElseIntf[T]
	// Takes closure for true condition and returns [Else].
	ThenF(func() T) ElseIntf[T]
}

// The [ElseIntf] provide methods:
// 	- [ElseIntf.ElseIntf], [ElseIntf.ElseF] to pass result if condition is false, or
//	- [ElseIntf.ElseIf], [ElseIntf.ElseIfF] to continue with nested [If], or
// 	- [ElseIntf.Panic] to stop evaluation with panic.
type ElseIntf[T any] interface {
	// Takes value for false condition, finish evaluation and returns result.
	Else(T) T

	// Takes closure for false condition, finish evaluation and returns result.
	ElseF(func() T) T

	// Takes the condition for a nested [If], constructs it and returns its [ThenIntf].
	ElseIf(bool) ThenIntf[T]

	// Takes the closure what return a condition, constructs a nested [If] and returns its [ThenIntf].
	ElseIfF(func() bool) ThenIntf[T]

	// If condition is true then returns value passed to [Then], otherwise panics.
	ElsePanic(any) T
}

// Implements [ThenIntf] branch
type Then bool

// [Then.Then]
func (t Then) Then[T any](v T) ElseIntf[T] {
	if t {
		return ret1[T]{v}
	}
	return else1[T]{}
}

// [Then.ThenF]
func (t Then) ThenF[T any](f func() T) ElseIntf[T] {
	if t {
		return ret1[T]{f()}
	}
	return else1[T]{}
}

// Implements [ElseIntf] branch
type else1[T any] struct{}

// [ElseIntf.Else]
func (else1[T]) Else(v T) T {
	return v
}

// [ElseIntf.ElseF]
func (else1[T]) ElseF(f func() T) T {
	return f()
}

// [ElseIntf.ElseIf]
func (else1[T]) ElseIf(cond bool) ThenIntf[T] {
	return elseIfThen[T](cond)
}

// [ElseIntf.ElseIfF]
func (else1[T]) ElseIfF(f func() bool) ThenIntf[T] {
	return elseIfThen[T](f())
}

// [ElseIntf.ElsePanic]
func (else1[T]) ElsePanic(v any) T {
	panic(v)
}

// Implements [ThenIntf] branch for [ElseIntf.ElseIf] calls
type elseIfThen[T any] bool

// [Then.Then]
func (t elseIfThen[T]) Then(v T) ElseIntf[T] {
	if t {
		return ret1[T]{v}
	}
	return else1[T]{}
}

// [Then.ThenF]
func (t elseIfThen[T]) ThenF(f func() T) ElseIntf[T] {
	if t {
		return ret1[T]{f()}
	}
	return else1[T]{}
}

// Implements both branches ([ThenIntf] and [ElseIntf]) return for succussfully completed evaluation.
type ret1[T any] struct {
	v T
}

// [ElseIntf.Else]
func (r ret1[T]) Else(T) T {
	return r.v
}

// [ElseIntf.ElseF]
func (r ret1[T]) ElseF(func() T) T {
	return r.v
}

// [ElseIntf.ElseIf]
func (r ret1[T]) ElseIf(bool) ThenIntf[T] {
	return r
}

// [ElseIntf.ElseIfF]
func (r ret1[T]) ElseIfF(func() bool) ThenIntf[T] {
	return r
}

// [ElseIntf.ElsePanic]
func (r ret1[T]) ElsePanic(any) T {
	return r.v
}

// [Then.Then]
func (r ret1[T]) Then(T) ElseIntf[T] {
	return r
}

// [Then.ThenF]
func (r ret1[T]) ThenF(func() T) ElseIntf[T] {
	return r
}
