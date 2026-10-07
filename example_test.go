/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn_test

import (
	"fmt"

	trn "github.com/mail2nnv/ternary"
)

func ExampleIf() {
	i := 1
	fmt.Println(
		trn.If(i == 1).
			Then("one").
			Else("not one"))
	i++
	fmt.Println(
		trn.If(i == 1).
			Then("one").
			Else("not one"))
	// Output:
	// one
	// not one
}

func ExampleIf_thenF_elseF() {
	ptr := (*int)(nil)
	fmt.Println(
		trn.If(ptr != nil).
			ThenF(func() string { return fmt.Sprint(*ptr) }).
			Else("nil int"))

	s := new("string")
	fmt.Println(
		trn.If(s == nil).
			Then("nil string").
			ElseF(func() string { return *s }))
	// Output:
	// nil int
	// string
}
