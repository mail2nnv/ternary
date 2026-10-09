
# `sw` package

The package `sw` provides ternary switch-operator for Go.

## Installation

Install with the go get command:

```shell
go get github.com/mail2nnv/ternary/switch
```

## Examplies

### Simplest switch ternary

```go
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
```

### Ternary switch with lazy evaluation

```go
package sw_test

import (
	"errors"
	"fmt"

	sw "github.com/mail2nnv/ternary/switch"
)

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
```

## Ternary and performance

Using ternary operators can improve code readability, but slightly reduces performance, especially if you have to use lazy returns from closures (`.ThenF(…)` or `.ElseF(…)`).

See [bench results](../bench/lastest.md) for particulars.
