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
		trn.If2[string, int](i == 1).
			Then("one", 1).
			Else("not one", i))
	i++
	fmt.Println(
		trn.If2[string, int](i == 1).
			Then("one", i).
			Else("not one", i))
	// Output:
	// one 1
	// not one 2
}

func ExampleIf2_thenF_elseF() {
	ptr := (*int)(nil)
	fmt.Println(
		trn.If2[string, error](ptr != nil).
			ThenF(func() (string, error) { return fmt.Sprint(*ptr), nil }).
			Else("nil int", errors.ErrUnsupported))

	s := new("string")
	fmt.Println(
		trn.If2[string, error](s == nil).
			Then("nil string", errors.ErrUnsupported).
			ElseF(func() (string, error) { return *s, nil }))

	// Output:
	// nil int unsupported operation
	// string <nil>
}

func ExampleIf2_elseIf() {
	for i := range 4 {
		fmt.Println(
			trn.If2[string, int](i == 0).
				Then("zero", 0).
				ElseIf(i == 1).
				Then("one", 1).
				ElseIf(i == 2).
				Then("two", 2).
				Else("many", i))
	}
	// Output:
	// zero 0
	// one 1
	// two 2
	// many 3
}

func ExampleIf2_elseIfF() {
	s := ""
	for range 3 {
		fmt.Println(
			trn.If2[string, int](s == "").
				Then("null", 0).
				ElseIfF(func() bool { return s[0] == 'a' }).
				Then("a*", len(s)).
				Else("other", -1))
		s += "a"
	}
	// Output:
	// null 0
	// a* 1
	// a* 2
}

func ExampleIf2_elsePanic() {
	fmt.Println(
		trn.If2[string, error](2*2 == 4).
			Then("4", nil).
			ElsePanic("🤪"))
	// Output:
	// 4 <nil>
}
