/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package iF

import (
	trn "github.com/mail2nnv/ternary"
)

type If = trn.IfBranch

// Takes condition and returns [If] branch.
//
// # Example:
//
// Simple return string:
//
//	s := iF.T(a == 1).
//		Then("one").
//		Else("not one")
//
// Lazy evaluate by condition:
//
//	s := iF.T(a != nil).
//		ThenF(func() string { return a.String() }).
//		Else("nil")
//
// Nested conditions:
//
//	s :=
//		iF.T(m == nil).Then("nil").Else(
//			iF.T(len(m) == 0).Then("empty").Else(
//				iF.T(len(m) == 1).Then("one item").Else(
//					fmt.Sprintf("%d items", len(m)))))
func T(v bool) If { return trn.If(v) }

// Takes condition and returns [If] branch.
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
//		iF.T(len(m) == 0).Then("empty").Else(
//			iF.T(len(m) == 1).Then("one item").Else(
//				fmt.Sprintf("%d items", len(m))))).
//		Else("nil")
func Not(v bool) If { return trn.If(!v) }
