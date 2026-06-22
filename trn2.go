/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn

// Takes condition and returns [Then2] branch.
//
// # Example:
//
// Simple return string and error:
//	s, err := If2[string, error](a == 1).
//		Then("one", nil).
//		Else("not one", errors.New("too more"))
//
// Lazy evaluate by condition:
//	s, err := If2[string, error](a != nil).
//		ThenF(func() (string, error) { return a.String(), nil }).
//		Else("nil", errors.New("missed"))
//
// Nested conditions:
//	s, err := If2[string, error](s == nil).
//		Then("nil", errors.New("nil")).
//		ElseIf(len(s) == 0).
//		Then("empty slice", errors.New("empty")).
//		ElseIf(len(s) == 1).
//		Then("one element slice", nil).
//		Else(fmt.Sprintf("%d-element slice", len(s)), nil)
func If2[T1, T2 any](cond bool) Then2[T1, T2] {
	return then2[T1, T2](cond)
}

// The [Then2] branch provide methods [Then2.Then] and [Then2.ThenF]
// to pass results if condition is true and returns [Else2] branch.
type Then2[T1, T2 any] interface {
	// Takes values for true condition and returns [Else2].
	Then(T1, T2) Else2[T1, T2]
	// Takes closure for true condition and returns [Else].
	ThenF(func() (T1, T2)) Else2[T1, T2]
}

// The [Else2] provide methods:
// 	- [Else2.Else], [Else2.ElseF] to pass results if condition is false, or
//	- [Else2.ElseIf], [Else2.ElseIfF] to continue with nested [If2], or
// 	- [Else2.Panic] to stop evaluation with panic.
type Else2[T1, T2 any] interface {
	// Takes values for false condition, finish evaluation and returns results.
	Else(T1, T2) (T1, T2)

	// Takes closure for false condition, finish evaluation and returns results.
	ElseF(func() (T1, T2)) (T1, T2)

	// Takes the condition for a nested [If2], constructs it and returns its [Then2].
	ElseIf(bool) Then2[T1, T2]

	// Takes the closure what return a condition, constructs a nested [If2] and returns its [Then2].
	ElseIfF(func() bool) Then2[T1, T2]

	// If condition is true then returns values passed to [Then2], otherwise panics.
	ElsePanic(any) (T1, T2)
}

// Implements [Then2] branch
type then2[T1, T2 any] bool

// [Then2.Then]
func (t2 then2[T1, T2]) Then(v1 T1, v2 T2) Else2[T1, T2] {
	if t2 {
		return ret2[T1, T2]{v1, v2}
	}
	return else2[T1, T2]{}
}

// [Then2.ThenF]
func (t2 then2[T1, T2]) ThenF(f func() (T1, T2)) Else2[T1, T2] {
	if t2 {
		v1, v2 := f()
		return ret2[T1, T2]{v1, v2}
	}
	return else2[T1, T2]{}
}

// Implements [Else2] branch
type else2[T1, T2 any] struct{}

// [Else2.Else]
func (else2[T1, T2]) Else(v1 T1, v2 T2) (T1, T2) {
	return v1, v2
}

// [Else2.ElseF]
func (else2[T1, T2]) ElseF(f func() (T1, T2)) (T1, T2) {
	return f()
}

// [Else2.ElseIf]
func (else2[T1, T2]) ElseIf(cond bool) Then2[T1, T2] {
	return then2[T1, T2](cond)
}

// [Else2.ElseIfF]
func (else2[T1, T2]) ElseIfF(f func() bool) Then2[T1, T2] {
	return then2[T1, T2](f())
}

// [Else2.ElsePanic]
func (else2[T1, T2]) ElsePanic(v any) (T1, T2) {
	panic(v)
}

// Implements both branches ([Then2] and [Else2]) to return for succussfully completed evaluation.
type ret2[T1, T2 any] struct {
	v1 T1
	v2 T2
}

// [Else2.Else]
func (d ret2[T1, T2]) Else(T1, T2) (T1, T2) {
	return d.v1, d.v2
}

// [Else2.ElseF]
func (d ret2[T1, T2]) ElseF(func() (T1, T2)) (T1, T2) {
	return d.v1, d.v2
}

// [Else2.ElseIf]
func (d ret2[T1, T2]) ElseIf(bool) Then2[T1, T2] {
	return d
}

// [Else2.ElseIfF]
func (d ret2[T1, T2]) ElseIfF(func() bool) Then2[T1, T2] {
	return d
}

// [Else2.ElsePanic]
func (d ret2[T1, T2]) ElsePanic(any) (T1, T2) {
	return d.v1, d.v2
}

// [Then.Then]
func (d ret2[T1, T2]) Then(T1, T2) Else2[T1, T2] {
	return d
}

// [Then.ThenF]
func (d ret2[T1, T2]) ThenF(func() (T1, T2)) Else2[T1, T2] {
	return d
}
