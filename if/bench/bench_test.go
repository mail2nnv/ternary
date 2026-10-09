/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package iF_test

import (
	"testing"

	iF "github.com/mail2nnv/ternary/if"
)

func Benchmark_If(b *testing.B) {

	_4 := func() int {
		// to exclude compiler magic optimization
		return len("1234")
	}

	b.Run("Native if-then-else", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			var got string

			if 2*2 == _4() {
				got = "OK"
			} else {
				got = "Fail"
			}

			if got != "OK" {
				b.Fail()
			}
		}
	})

	b.Run("Ternary if-then-else", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {

			got := iF.True(2*2 == _4()).Then("OK").Else("Fail")

			if got != "OK" {
				b.Fail()
			}
		}
	})
}
