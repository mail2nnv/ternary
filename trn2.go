/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn

// Takes condition and returns [Then2Intf] branch.
//
// # Example:
//
// Simple return string and error:
//	s, err := If2(a == 1).
//		Then("one", nil).
//		Else("not one", errors.New("too more"))
//
// Lazy evaluate by condition:
//	s, err := If2(a != nil).
//		ThenF(func() (string, error) { return a.String(), nil }).
//		Else("nil", errors.New("missed"))
//
// Nested conditions:
//	s, err := If2(s == nil).
//		Then("nil", errors.New("nil")).
//		ElseIf(len(s) == 0).
//		Then("empty slice", errors.New("empty")).
//		ElseIf(len(s) == 1).
//		Then("one element slice", nil).
//		Else(fmt.Sprintf("%d-element slice", len(s)), nil)
func If2(cond bool) Then2 {
	return Then2(cond)
}

// The [Then2Intf] branch provide methods [Then2Intf.Then] and [Then2Intf.ThenF]
// to pass two results if condition is true and returns [Else2Then2Intf] branch.
type Then2Intf[T1, T2 any] interface {
	// Takes values for true condition and returns [Else2].
	Then(T1, T2) Else2Intf[T1, T2]
	// Takes closure for true condition and returns [Else].
	ThenF(func() (T1, T2)) Else2Intf[T1, T2]
}

// The [Else2Intf] provide methods:
// 	- [Else2Intf.Else], [Else2Intf.ElseF] to pass two results if condition is false, or
//	- [Else2Intf.ElseIf], [Else2Intf.ElseIfF] to continue with nested [If2], or
// 	- [Else2Intf.Panic] to stop evaluation with panic.
type Else2Intf[T1, T2 any] interface {
	// Takes values for false condition, finish evaluation and returns results.
	Else(T1, T2) (T1, T2)

	// Takes closure for false condition, finish evaluation and returns results.
	ElseF(func() (T1, T2)) (T1, T2)

	// Takes the condition for a nested [If2], constructs it and returns its [Then2Intf].
	ElseIf(bool) Then2Intf[T1, T2]

	// Takes the closure what return a condition, constructs a nested [If2] and returns its [Then2Intf].
	ElseIfF(func() bool) Then2Intf[T1, T2]

	// If condition is true then returns values passed to [Then2], otherwise panics.
	ElsePanic(any) (T1, T2)
}

// Implements [Then2Intf] branch
type Then2 bool

// [Then2Intf.Then]
func (t2 Then2) Then[T1, T2 any](v1 T1, v2 T2) Else2Intf[T1, T2] {
	if t2 {
		return ret2[T1, T2]{v1, v2}
	}
	return else2[T1, T2]{}
}

// [Then2Intf.ThenF]
func (t2 Then2) ThenF[T1, T2 any](f func() (T1, T2)) Else2Intf[T1, T2] {
	if t2 {
		v1, v2 := f()
		return ret2[T1, T2]{v1, v2}
	}
	return else2[T1, T2]{}
}

// Implements [Else2Intf] branch
type else2[T1, T2 any] struct{}

// [Else2Intf.Else]
func (else2[T1, T2]) Else(v1 T1, v2 T2) (T1, T2) {
	return v1, v2
}

// [Else2Intf.ElseF]
func (else2[T1, T2]) ElseF(f func() (T1, T2)) (T1, T2) {
	return f()
}

// [Else2Intf.ElseIf]
func (else2[T1, T2]) ElseIf(cond bool) Then2Intf[T1, T2] {
	return else2IfThen[T1, T2](cond)
}

// [Else2Intf.ElseIfF]
func (else2[T1, T2]) ElseIfF(f func() bool) Then2Intf[T1, T2] {
	return else2IfThen[T1, T2](f())
}

// [Else2Intf.ElsePanic]
func (else2[T1, T2]) ElsePanic(v any) (T1, T2) {
	panic(v)
}

// Implements [Then2Intf] branch for [Else2Intf.ElseIf] calls
type else2IfThen[T1, T2 any] bool

// [Then2Intf.Then]
func (t else2IfThen[T1, T2]) Then(v1 T1, v2 T2) Else2Intf[T1, T2] {
	if t {
		return ret2[T1, T2]{v1, v2}
	}
	return else2[T1, T2]{}
}

// [Then2Intf.ThenF]
func (t else2IfThen[T1, T2]) ThenF(f func() (T1, T2)) Else2Intf[T1, T2] {
	if t {
		v1, v2 := f()
		return ret2[T1, T2]{v1, v2}
	}
	return else2[T1, T2]{}
}

// Implements both branches ([Then2Intf] and [Else2Intf]) to return for succussfully completed evaluation.
type ret2[T1, T2 any] struct {
	v1 T1
	v2 T2
}

// [Else2Intf.Else]
func (d ret2[T1, T2]) Else(T1, T2) (T1, T2) {
	return d.v1, d.v2
}

// [Else2Intf.ElseF]
func (d ret2[T1, T2]) ElseF(func() (T1, T2)) (T1, T2) {
	return d.v1, d.v2
}

// [Else2Intf.ElseIf]
func (d ret2[T1, T2]) ElseIf(bool) Then2Intf[T1, T2] {
	return d
}

// [Else2Intf.ElseIfF]
func (d ret2[T1, T2]) ElseIfF(func() bool) Then2Intf[T1, T2] {
	return d
}

// [Else2Intf.ElsePanic]
func (d ret2[T1, T2]) ElsePanic(any) (T1, T2) {
	return d.v1, d.v2
}

// [Then2Intf.Then]
func (d ret2[T1, T2]) Then(T1, T2) Else2Intf[T1, T2] {
	return d
}

// [Then2Intf.ThenF]
func (d ret2[T1, T2]) ThenF(func() (T1, T2)) Else2Intf[T1, T2] {
	return d
}
