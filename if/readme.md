
# `iF` package

The package `iF` provides ternary if-operator for Go.

## Installation

Install with the go get command:

```shell
go get github.com/mail2nnv/ternary/if
```

## Examplies

### Simplest if ternary

```go
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
```

### Ternary if with nested

```go
package iF_test

import (
	"fmt"

	iF "github.com/mail2nnv/ternary/if"
)

func ExampleIf_elseIf() {

for i := range 4 {
	fmt.Println(
		iF.T(i == 0).Then("zero").Else(
			iF.T(i == 1).Then("one").Else(
				iF.T(i == 2).Then("two").Else("many"))))
	}
	// Output:
	// zero
	// one
	// two
	// many
}
```

### Ternary if with lazy evaluation

```go
package trn_test

import (
	"errors"
	"fmt"

	trn "github.com/mail2nnv/ternary"
)

func ExampleIf_thenF_elseF() {
	ptr := (*int)(nil)
	fmt.Println(
		iF.T(ptr != nil).
			ThenF(func() string { return fmt.Sprint(*ptr) }).
			Else("nil int"))

	s := new("string")
	fmt.Println(
		iF.T(s == nil).
			Then("nil string").
			ElseF(func() string { return *s }))

	err := fmt.Errorf("error %w", errors.ErrUnsupported)
	fmt.Println(
		iF.T(err == nil).
			Then("nil error").
			ElseIf(errors.Is(err, errors.ErrUnsupported)).
			ThenF(func() string { return err.Error() }).
			ElseF(func() string { return fmt.Sprintf("surprise: %s", err.Error()) }))

	// Output:
	// nil int
	// string
	// error unsupported operation
}
```

## Ternary and performance

Using ternary operators can improve code readability, but slightly reduces performance, especially if you have to use lazy returns from closures (`.ThenF(…)` or `.ElseF(…)`).

See [bench results](../bench/lastest.md) for particulars.
