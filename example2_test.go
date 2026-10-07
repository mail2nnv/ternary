/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn_test

import (
	"errors"
	"fmt"

	trn "github.com/mail2nnv/ternary"
)

func ExampleIf2() {
	i := 1
	fmt.Println(
		trn.If2(i == 1).
			Then("one", 1).
			Else("not one", i))
	i++
	fmt.Println(
		trn.If2(i == 1).
			Then("one", i).
			Else("not one", i))
	// Output:
	// one 1
	// not one 2
}

func ExampleIf2_thenF_elseF() {
	ptr := (*int)(nil)
	fmt.Println(
		trn.If2(ptr != nil).
			ThenF(func() (string, error) { return fmt.Sprint(*ptr), nil }).
			Else("nil int", errors.ErrUnsupported))

	s := new("string")
	fmt.Println(
		trn.If2(s == nil).
			Then("nil string", errors.ErrUnsupported).
			ElseF(func() (string, error) { return *s, nil }))

	// Output:
	// nil int unsupported operation
	// string <nil>
}
