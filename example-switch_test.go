/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn_test

import (
	"fmt"

	trn "github.com/mail2nnv/ternary"
)

func ExampleSwitch() {
	for i := range 3 {
		fmt.Println(
			trn.Switch(i).
				Case(1, "one").
				Case(2, "two").
				Default("zero"))
	}
	// Output:
	// zero
	// one
	// two
}
