/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn

// Takes condition and returns [Condition] branch.
//
// # Example:
//
// Simple return string and error:
//	s, err := If2(a == 0).
//		Then("", errors.New("zero")).
//		Else("natural", nil)
//
// Lazy evaluate by condition:
//	s, err := If2(a != nil).
//		ThenF(func() string { return a.String(), error(nil) }).
//		Else("", errors.New("nil"))
//
// Nested conditions:
//	s, l :=
// 		If(m == nil).Then("nil", 0).Else(
//			If(len(m) == 0).Then("empty", 0).Else(
//				If(len(m) == 1).Then("one key", 1).Else(
// 					fmt.Sprintf("%d keys", len(m)), len(m))))
func If2(cond bool) Condition2 {
	return Condition2(cond)
}

// The [Condition2] provide methods to construct the [Branch2].
type Condition2 bool

// Passes the values for true-[Condition2].
// Returns the [Branch2].
func (c Condition2) Then[T1, T2 any](v1 T1, v2 T2) Branch2[T1, T2] {
	return Branch2[T1, T2]{bool(c), v1, v2}
}

// Passes the closure what returns the values for true-[Condition2].
// Returns the [Branch2].
func (c Condition2) ThenF[T1, T2 any](f func() (T1, T2)) Branch2[T1, T2] {
	if c {
		v1, v2 := f()
		return Branch2[T1, T2]{c: true, v1: v1, v2: v2}
	}
	return Branch2[T1, T2]{c: false}
}

// The [Branch2] provide methods to evaluate the [Condition2].
type Branch2[T1, T2 any] struct {
	c  bool
	v1 T1
	v2 T2
}

// Pass the values for false-[Condition2], evaluates the [Condition2] and returns the results.
//
// If [Condition2] is true, then returns the values early passed to [Condition2.Then] (or to [Condition2.ThenF]),
// elsewhere return passed values.
func (t Branch2[T1, T2]) Else(v1 T1, v2 T2) (T1, T2) {
	if t.c {
		return t.v1, t.v2
	}
	return v1, v2
}

// Pass the closure for false-[Condition2], evaluates the [Condition] and returns the results.
//
// If [Condition2] is true, then returns the value2 early passed to [Condition2.Then] (or to [Condition2.ThenF]),
// elsewhere returns the passed closure results.
func (t Branch2[T1, T2]) ElseF(f func() (T1, T2)) (T1, T2) {
	if t.c {
		return t.v1, t.v2
	}
	return f()
}
