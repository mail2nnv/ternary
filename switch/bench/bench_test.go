/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package sw_test

import (
	"testing"

	sw "github.com/mail2nnv/ternary/switch"
)

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
				got := sw.V(i).
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
}
