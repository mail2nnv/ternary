/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package sw_test

import (
	"fmt"

	sw "github.com/mail2nnv/ternary/switch"
)

func ExampleSwitch() {
	for i := range 3 {
		fmt.Println(
			sw.V(i).
				Case(1, "one").
				Case(2, "two").
				Default("zero"))
	}
	// Output:
	// zero
	// one
	// two
}

func ExampleSwitch_lazy() {
	for i := range 4 {
		fmt.Println(
			sw.V(i).
				CaseK(func() int { return 1 }, "one").
				CaseV(2, func() string { return "two" }).
				CaseKV(func() int { return 3 }, func() string { return "three" }).
				DefaultV(func() string { return "zero" }))
	}
	// Output:
	// zero
	// one
	// two
	// three
}
