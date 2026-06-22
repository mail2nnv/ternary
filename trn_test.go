/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn_test

import (
	"testing"

	trn "github.com/mail2nnv/ternary"
	"github.com/stretchr/testify/require"
)

func TestIfThenElse(t *testing.T) {
	require := require.New(t)

	require.Equal(5, trn.If[int](2 > 1).Then(5).Else(0))
	require.Equal(0, trn.If[int](2 < 1).Then(5).Else(0))
}

func TestIfThenFElseF(t *testing.T) {
	require := require.New(t)
	require.Equal(5, trn.If[int](2 > 1).ThenF(func() int { return 5 }).ElseF(func() int { return 0 }))
	require.Equal(0, trn.If[int](2 < 1).ThenF(func() int { return 5 }).ElseF(func() int { return 0 }))
}

func TestIfThenElseIf(t *testing.T) {
	require := require.New(t)
	s := ""
	for range 5 {
		l := trn.If[int](s == "").
			Then(0).
			ElseIf(s == "a").
			Then(1).
			ElseIf(s == "aa").
			Then(2).
			ElseIf(s == "aaa").
			Then(3).
			ElseIf(s == "aaaa").
			ThenF(func() int { return len(s) }).
			ElsePanic("ops")
		require.Len(s, l)
		s += "a"
	}
}

func TestIfThenElsePanic(t *testing.T) {
	require := require.New(t)
	require.Panics(
		func() {
			_ = trn.If[int](2*2 == 5).
				Then(0).
				ElsePanic("🤪")
		})

	require.NotPanics(
		func() {
			require.Zero(
				trn.If[int](2*2 == 4).
					Then(0).
					ElsePanic("🤪"))
		})
}
