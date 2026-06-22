/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn_test

import (
	"testing"

	trn "github.com/mail2nnv/ternary"
)

func Benchmark_If(b *testing.B) {

	cases := []struct {
		arg, want int
	}{
		{0, +0},
		{1, -1},
		{2, +2},
		{3, -3},
		{4, +4},
		{5, -5},
		{6, +6},
		{7, -7},
		{8, +8},
		{9, -9},
	}

	b.Run("Native if-then-else", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			for i := range cases {
				var got int
				if cases[i].arg%2 == 0 {
					got = +cases[i].arg
				} else {
					got = -cases[i].arg
				}
				if got != cases[i].want {
					b.Fail()
				}
			}
		}
	})

	b.Run("Ternary if-then-else", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			for i := range cases {
				got := trn.If[int](cases[i].arg%2 == 0).
					Then(cases[i].arg).
					Else(-cases[i].arg)
				if got != cases[i].want {
					b.Fail()
				}
			}
		}
	})

	b.Run("Ternary if-thenF-else", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			for i := range cases {
				got := trn.If[int](cases[i].arg%2 == 0).
					ThenF(func() int { return cases[i].arg }).
					Else(-cases[i].arg)
				if got != cases[i].want {
					b.Fail()
				}
			}
		}
	})

	b.Run("Ternary if-thenF-elseF", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			for i := range cases {
				got := trn.If[int](cases[i].arg%2 == 0).
					ThenF(func() int { return cases[i].arg }).
					ElseF(func() int { return -cases[i].arg })
				if got != cases[i].want {
					b.Fail()
				}
			}
		}
	})
}
