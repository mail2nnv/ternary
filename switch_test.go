/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin


package trn
/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
*/

package trn_test

import (
	"strconv"
	"testing"

	trn "github.com/mail2nnv/ternary"
	"github.com/stretchr/testify/require"
)

func TestSwitch(t *testing.T) {
	require := require.New(t)
	require.Equal("3",
		trn.Switch(3).
			Case(1).Return("1").
			Case(2).Return("2").
			Case(3).Return("3").
			Default("more"))

	for i := range 5 {
		require.Equal(strconv.Itoa(i),
			trn.Switch(i).
				Case(0).Return("0").
				Case(1).Return("1").
				Case(2).Return("2").
				Case(3).Return("3").
				Default("4"))
	}
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
					Case(0).Return(9).
					Case(1).Return(8).
					Case(2).Return(7).
					Case(3).Return(6).
					Default(5)
				if got != 9-i {
					b.Fail()
				}
			}
		}
	})
}
