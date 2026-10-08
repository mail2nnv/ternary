/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package iF

/*
package if
----------
package ìf
package íf
package îf
package ïf
package įf
----------
package iF
*/

import (
	trn "github.com/mail2nnv/ternary"
)

type C = trn.Condition

// Takes condition and returns [trn.Condition] branch.
//
// # Example:
//
// Simple return string:
//
//	s := iF.True(a == 1).
//		Then("one").
//		Else("not one")
//
// Lazy evaluate by condition:
//
//	s := iF.True(a != nil).
//		ThenF(func() string { return a.String() }).
//		Else("nil")
//
// Nested conditions:
//
//	s :=
//		iF.True(m == nil).Then("nil").Else(
//			iF.True(len(m) == 0).Then("empty").Else(
//				iF.True(len(m) == 1).Then("one item").Else(
//					fmt.Sprintf("%d items", len(m)))))
func True(v bool) C { return trn.If(v) }

// Takes condition and returns [trn.Condition] branch.
//
// # Example:
//
// Simple return string:
//
//	s := iF.Not(a == 1).
//		Then("not one").
//		Else("one")
//
// Lazy evaluate by condition:
//
//	s := iF.Not(a == nil).
//		ThenF(func() string { return a.String() }).
//		Else("nil")
//
// Nested conditions:
//
//	s := iF.Not(m == nil).Then(
//		iF.True(len(m) == 0).Then("empty").Else(
//			iF.True(len(m) == 1).Then("one item").Else(
//				fmt.Sprintf("%d items", len(m))))).
//		Else("nil")
func Not(v bool) C { return trn.If(!v) }
