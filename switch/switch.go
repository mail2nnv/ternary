/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package sw

// Takes condition and returns [trn.FirstCase].
//
// # Example:
//
// Simple return string:
//
//	s := sw.Val(a).
//		Case(0).Return("zero").
//		Case(1).Return("one").
//		Case(2).Return("two").
//		Else("many")
//
// func Val[T comparable](v T) trn.FirstCase[T] {
// 	return trn.Switch(v)
// }
