
# `trn` package

The package `trn` provides ternary operators for Go.

## Installation

Install with the go get command:

```shell
go get github.com/mail2nnv/ternary
```

## `If`

### Simplest `If` ternary

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
```

### Ternary `If` with nested

```go
package trn_test

import (
	"fmt"

	trn "github.com/mail2nnv/ternary"
)

func ExampleIf_elseIf() {

for i := range 4 {
	fmt.Println(
		trn.If(i == 0).Then("zero").Else(
			trn.If(i == 1).Then("one").Else(
				trn.If(i == 2).Then("two").Else("many"))))
	}
	// Output:
	// zero
	// one
	// two
	// many
}
```

### Ternary `If` with lazy evaluation

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
		trn.If(ptr != nil).
			ThenF(func() string { return fmt.Sprint(*ptr) }).
			Else("nil int"))

	s := new("string")
	fmt.Println(
		trn.If(s == nil).
			Then("nil string").
			ElseF(func() string { return *s }))

	err := fmt.Errorf("error %w", errors.ErrUnsupported)
	fmt.Println(
		trn.If(err == nil).
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

## `Switch`

### Simplest `Switch` ternary

```go
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
```

### Ternary `Switch` with lazy evaluation

```go
package trn_test

import (
	"fmt"

	trn "github.com/mail2nnv/ternary"
)

func ExampleSwitch_lazy() {
	for i := range 4 {
		fmt.Println(
			trn.Switch(i).
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
```

## Ternary and performance

Using ternary operators can improve code readability, but slightly reduces performance, especially if you have to use lazy returns from closures (`If().ThenF(…)` or `If().Then().ElseF(…)`).

See [bench results](bench/lastest.md) for particulars.
