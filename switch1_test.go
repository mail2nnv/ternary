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

func TestSwitch1(t *testing.T) {
	req := require.New(t)

	for i := range 5 {
		want := strconv.Itoa(i)
		req.Equal(want,
			trn.Switch1(i).
				Case(1, "1").
				Case(2, "2").
				Case(3, "3").
				Case(4, "4").
				Default("0"))

		req.Equal(want,
			trn.Switch1(i).
				CaseF(func() int { return (1) }, "1").
				CaseF(func() int { return (2) }, "2").
				CaseF(func() int { return (3) }, "3").
				CaseF(func() int { return (4) }, "4").
				Default("0"))

		req.Equal(want,
			trn.Switch1(i).
				CaseR(1, func() string { return "1" }).
				CaseR(2, func() string { return "2" }).
				CaseR(3, func() string { return "3" }).
				CaseR(4, func() string { return "4" }).
				DefaultR(func() string { return "0" }))

		req.Equal(want,
			trn.Switch1(i).
				CaseFR(func() int { return 1 }, func() string { return "1" }).
				CaseFR(func() int { return 2 }, func() string { return "2" }).
				CaseFR(func() int { return 3 }, func() string { return "3" }).
				CaseFR(func() int { return 4 }, func() string { return "4" }).
				DefaultR(func() string { return "0" }))
	}
}

func Benchmark_Switch1(b *testing.B) {

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
				got := trn.Switch1(i).
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
				got := trn.Switch1(i).
					CaseF(func() int { return 0 }, 9).
					CaseF(func() int { return 1 }, 8).
					CaseF(func() int { return 2 }, 7).
					CaseF(func() int { return 3 }, 6).
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
				got := trn.Switch1(i).
					CaseR(0, func() int { return 9 }).
					CaseR(1, func() int { return 8 }).
					CaseR(2, func() int { return 7 }).
					CaseR(3, func() int { return 6 }).
					DefaultR(func() int { return 5 })
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
				got := trn.Switch1(i).
					CaseFR(func() int { return 0 }, func() int { return 9 }).
					CaseFR(func() int { return 1 }, func() int { return 8 }).
					CaseFR(func() int { return 2 }, func() int { return 7 }).
					CaseFR(func() int { return 3 }, func() int { return 6 }).
					DefaultR(func() int { return 5 })
				if got != 9-i {
					b.Fail()
				}
			}
		}
	})
}
