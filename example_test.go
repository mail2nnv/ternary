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

func ExampleIf() {
	i := 1
	fmt.Println(
		trn.If[string](i == 1).
			Then("one").
			Else("not one"))
	i++
	fmt.Println(
		trn.If[string](i == 1).
			Then("one").
			Else("not one"))
	// Output:
	// one
	// not one
}

func ExampleIf_thenF_elseF() {
	ptr := (*int)(nil)
	fmt.Println(
		trn.If[string](ptr != nil).
			ThenF(func() string { return fmt.Sprint(*ptr) }).
			Else("nil int"))

	s := new("string")
	fmt.Println(
		trn.If[string](s == nil).
			Then("nil string").
			ElseF(func() string { return *s }))

	err := fmt.Errorf("error %w", errors.ErrUnsupported)
	fmt.Println(
		trn.If[string](err == nil).
			Then("nil error").
			ElseIf(errors.Is(err, errors.ErrUnsupported)).
			ThenF(func() string { return err.Error() }).
			ElseF(func() string { return fmt.Sprintf("surprise: %s", err.Error()) }))

	// Output:
	// nil int
	// string
	// error unsupported operation
}

func ExampleIf_elseIf() {
	for i := range 4 {
		fmt.Println(
			trn.If[string](i == 0).
				Then("zero").
				ElseIf(i == 1).
				Then("one").
				ElseIf(i == 2).
				Then("two").
				Else("many"))
	}
	// Output:
	// zero
	// one
	// two
	// many
}

func ExampleIf_elseIfF() {
	s := ""
	for range 2 {
		fmt.Println(
			trn.If[string](s == "").
				Then("null").
				ElseIfF(func() bool { return s[0] == 'a' }).
				Then("a*").
				Else("other"))
		s += "a"
	}
	// Output:
	// null
	// a*
}

func ExampleIf_elsePanic() {
	fmt.Println(
		trn.If[string](2*2 == 4).
			Then("4").
			ElsePanic("🤪"))
	// Output:
	// 4
}
