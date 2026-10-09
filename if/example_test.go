/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package iF_test

import (
	"fmt"

	iF "github.com/mail2nnv/ternary/if"
)

func ExampleT() {
	i := 1
	fmt.Println(
		iF.T(i == 1).
			Then("one").
			Else("not one"))
	i++
	fmt.Println(
		iF.T(i == 1).
			Then("one").
			Else("not one"))
	// Output:
	// one
	// not one
}

func ExampleNot() {
	i := 1
	fmt.Println(
		iF.Not(i == 1).
			Then("not one").
			Else("one"))
	i++
	fmt.Println(
		iF.Not(i == 1).
			Then("not one").
			Else("one"))
	// Output:
	// one
	// not one
}
