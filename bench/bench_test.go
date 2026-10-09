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
				got := trn.If(cases[i].arg%2 == 0).
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
				got := trn.If(cases[i].arg%2 == 0).
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
				got := trn.If(cases[i].arg%2 == 0).
					ThenF(func() int { return cases[i].arg }).
					ElseF(func() int { return -cases[i].arg })
				if got != cases[i].want {
					b.Fail()
				}
			}
		}
	})
}

func Benchmark_Switch(b *testing.B) {

	b.Run("Native switch", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			for i := range 5 {
				var got int
				switch i {
				case 0:
					got = 9
				case 1:
					got = 8
				case 2:
					got = 7
				case 3:
					got = 6
				default:
					got = 5
				}
				if got != 9-i {
					b.Fail()
				}
			}
		}
	})

	b.Run("Ternary switch", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			for i := range 5 {
				got := trn.Switch(i).
					Case(0, 9).
					Case(1, 8).
					Case(2, 7).
					Case(3, 6).
					Default(5)
				if got != 9-i {
					b.Fail()
				}
			}
		}
	})

	b.Run("Ternary switch case closure", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			for i := range 5 {
				got := trn.Switch(i).
					CaseK(func() int { return 0 }, 9).
					CaseK(func() int { return 1 }, 8).
					CaseK(func() int { return 2 }, 7).
					CaseK(func() int { return 3 }, 6).
					Default(5)
				if got != 9-i {
					b.Fail()
				}
			}
		}
	})

	b.Run("Ternary switch return closure", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			for i := range 5 {
				got := trn.Switch(i).
					CaseV(0, func() int { return 9 }).
					CaseV(1, func() int { return 8 }).
					CaseV(2, func() int { return 7 }).
					CaseV(3, func() int { return 6 }).
					DefaultV(func() int { return 5 })
				if got != 9-i {
					b.Fail()
				}
			}
		}
	})

	b.Run("Ternary switch case and return closures", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			for i := range 5 {
				got := trn.Switch(i).
					CaseKV(func() int { return 0 }, func() int { return 9 }).
					CaseKV(func() int { return 1 }, func() int { return 8 }).
					CaseKV(func() int { return 2 }, func() int { return 7 }).
					CaseKV(func() int { return 3 }, func() int { return 6 }).
					DefaultV(func() int { return 5 })
				if got != 9-i {
					b.Fail()
				}
			}
		}
	})
}
