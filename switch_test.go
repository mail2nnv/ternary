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
	req := require.New(t)

	for i := range 5 {
		want := strconv.Itoa(i)
		req.Equal(want,
			trn.Switch(i).
				Case(1, "1").
				Case(2, "2").
				Case(3, "3").
				Case(4, "4").
				Default("0"))

		req.Equal(want,
			trn.Switch(i).
				CaseK(func() int { return (1) }, "1").
				CaseK(func() int { return (2) }, "2").
				CaseK(func() int { return (3) }, "3").
				CaseK(func() int { return (4) }, "4").
				Default("0"))

		req.Equal(want,
			trn.Switch(i).
				CaseV(1, func() string { return "1" }).
				CaseV(2, func() string { return "2" }).
				CaseV(3, func() string { return "3" }).
				CaseV(4, func() string { return "4" }).
				DefaultV(func() string { return "0" }))

		req.Equal(want,
			trn.Switch(i).
				CaseKV(func() int { return 1 }, func() string { return "1" }).
				CaseKV(func() int { return 2 }, func() string { return "2" }).
				CaseKV(func() int { return 3 }, func() string { return "3" }).
				CaseKV(func() int { return 4 }, func() string { return "4" }).
				DefaultV(func() string { return "0" }))
	}
}
