
# `trn` package

The package `trn` provides ternary operators for Go.

## Installation

Install with the go get command:

```
go get github.com/mail2nnv/ternary
```

## Examplies

### Simplest ternary

```go
package trn_test

import (
	"fmt"

	trn "github.com/mail2nnv/ternary"
)

//cspell:words Println

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
```

### Ternary with elseif

```go
package trn_test

import (
	"fmt"

	trn "github.com/mail2nnv/ternary"
)

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
```

### Ternary with lazy evaluation

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
```

### Ternary with else panic

```go
package trn_test

import (
	"fmt"

	trn "github.com/mail2nnv/ternary"
)

func ExampleIf_elsePanic() {
	fmt.Println(
		trn.If[string](2*2 == 4).
			Then("4").
			ElsePanic("🤪"))
	// Output:
	// 4
}
```

## Ternary and performance

Using ternary operators can improve code readability, but reduces performance.

See [bench results](bench/bench-2026-06-19.md) for particulars.
