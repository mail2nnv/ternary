/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn

// Takes condition and returns [IfBranch] branch.
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
func If2(cond bool) IfBranch2 {
	return IfBranch2(cond)
}

// The [IfBranch2] provide methods to construct the [Then2].
type IfBranch2 bool

// Passes the values for true-[IfBranch2].
// Returns the [Then2].
func (i IfBranch2) Then[V1, V2 any](v1 V1, v2 V2) Then2[V1, V2] {
	return Then2[V1, V2]{bool(i), v1, v2}
}

// Passes the closure what returns the values for true-[IfBranch2].
// Returns the [Then2].
func (i IfBranch2) ThenF[V1, V2 any](f func() (V1, V2)) Then2[V1, V2] {
	if i {
		v1, v2 := f()
		return Then2[V1, V2]{bool: true, v1: v1, v2: v2}
	}
	return Then2[V1, V2]{bool: false}
}

// The [Then2] provide methods to evaluate the [IfBranch2].
type Then2[V1, V2 any] struct {
	bool
	v1 V1
	v2 V2
}

// Pass the values for false-[IfBranch2], evaluates the [IfBranch2] and returns the results.
//
// If [IfBranch2] is true, then returns the values early passed to [IfBranch2.Then] (or to [IfBranch2.ThenF]),
// elsewhere return passed values.
func (t Then2[V1, V2]) Else(v1 V1, v2 V2) (V1, V2) {
	if t.bool {
		return t.v1, t.v2
	}
	return v1, v2
}

// Pass the closure for false-[IfBranch2], evaluates the [IfBranch] and returns the results.
//
// If [IfBranch2] is true, then returns the value2 early passed to [IfBranch2.Then] (or to [IfBranch2.ThenF]),
// elsewhere returns the passed closure results.
func (t Then2[V1, V2]) ElseF(f func() (V1, V2)) (V1, V2) {
	if t.bool {
		return t.v1, t.v2
	}
	return f()
}
