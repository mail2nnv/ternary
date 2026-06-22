/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn

// Takes condition and returns [Then] branch.
//
// # Example:
//
// Simple return string:
//	s := If[string](a == 1).
//		Then("one").
//		Else("not one")
//
// Lazy evaluate by condition:
//	s := If[string](a != nil).
//		ThenF(func() string { return a.String() }).
//		Else("nil")
//
// Nested conditions:
//	s := If[string](s == nil).
//		Then("nil").
//		ElseIf(len(s) == 0).
//		Then("empty slice").
//		ElseIf(len(s) == 1).
//		Then("one element slice").
//		Else(fmt.Sprintf("%d-element slice", len(s)))
func If[T any](cond bool) Then[T] {
	return then[T](cond)
}

// The [Then] branch provide methods [Then.Then] and [Then.ThenF]
// to pass result if condition is true and returns [Else] branch.
type Then[T any] interface {
	// Takes value for true condition and returns [Else].
	Then(T) Else[T]
	// Takes closure for true condition and returns [Else].
	ThenF(func() T) Else[T]
}

// The [Else] provide methods:
// 	- [Else.Else], [Else.ElseF] to pass result if condition is false, or
//	- [Else.ElseIf], [Else.ElseIfF] to continue with nested [If], or
// 	- [Else.Panic] to stop evaluation with panic.
type Else[T any] interface {
	// Takes value for false condition, finish evaluation and returns result.
	Else(T) T

	// Takes closure for false condition, finish evaluation and returns result.
	ElseF(func() T) T

	// Takes the condition for a nested [If], constructs it and returns its [Then].
	ElseIf(bool) Then[T]

	// Takes the closure what return a condition, constructs a nested [If] and returns its [Then].
	ElseIfF(func() bool) Then[T]

	// If condition is true then returns value passed to [Then], otherwise panics.
	ElsePanic(any) T
}

// Implements [Then] branch
type then[T any] bool

// [Then.Then]
func (t then[T]) Then(v T) Else[T] {
	if t {
		return ret[T]{v}
	}
	return else1[T]{}
}

// [Then.ThenF]
func (t then[T]) ThenF(f func() T) Else[T] {
	if t {
		return ret[T]{f()}
	}
	return else1[T]{}
}

// Implements [Else] branch
type else1[T any] struct{}

// [Else.Else]
func (else1[T]) Else(v T) T {
	return v
}

// [Else.ElseF]
func (else1[T]) ElseF(f func() T) T {
	return f()
}

// [Else.ElseIf]
func (else1[T]) ElseIf(cond bool) Then[T] {
	return then[T](cond)
}

// [Else.ElseIfF]
func (else1[T]) ElseIfF(f func() bool) Then[T] {
	return then[T](f())
}

// [Else.ElsePanic]
func (else1[T]) ElsePanic(v any) T {
	panic(v)
}

// Implements both branches ([Then] and [Else]) return for succussfully completed evaluation.
type ret[T any] struct {
	v T
}

// [Else.Else]
func (r ret[T]) Else(T) T {
	return r.v
}

// [Else.ElseF]
func (r ret[T]) ElseF(func() T) T {
	return r.v
}

// [Else.ElseIf]
func (r ret[T]) ElseIf(bool) Then[T] {
	return r
}

// [Else.ElseIfF]
func (r ret[T]) ElseIfF(func() bool) Then[T] {
	return r
}

// [Else.ElsePanic]
func (r ret[T]) ElsePanic(any) T {
	return r.v
}

// [Then.Then]
func (r ret[T]) Then(T) Else[T] {
	return r
}

// [Then.ThenF]
func (r ret[T]) ThenF(func() T) Else[T] {
	return r
}
