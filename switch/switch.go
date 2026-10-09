/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package sw

import (
	trn "github.com/mail2nnv/ternary"
)

type Switch[K comparable] = trn.SwitchBranch[K]

// Takes condition and returns [Switch] branch.
//
// # Example:
//
// Simple return string:
//
//	s := sw.V(a).
//		Case(1, "one").
//		Case(2, "two").
//		Default("many")
//
// Lazy evaluate value or result:
//
//	var a int
//	…
//	s := sw.V(a).
//		CaseF(func() int { return 1 }, 	"one").
//		CaseR(2, func() string { return "two"} ).
//		CaseFR(func() int { return 3 }, func() string { return "three"} ).
//		DefaultR(func() string { return "more"} )
func V[K comparable](k K) Switch[K] { return trn.Switch(k) }
