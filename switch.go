/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn

// Takes switch value and returns [SwitchBranch].
//
// # Example:
//
// Simple return string:
//
//	var a int
//	…
//	s := Switch(a).
//		Case(1, "one").
//		Case(2, "two").
//		Default("more")
//
// Lazy evaluate value or result:
//
//	var a int
//	…
//	s := Switch(a).
//		CaseF(func() int { return 1 }, 	"one").
//		CaseR(2, func() string { return "two"} ).
//		CaseFR(func() int { return 3 }, func() string { return "three"} ).
//		DefaultR(func() string { return "more"} )
func Switch[K comparable](want K) SwitchBranch[K] {
	return SwitchBranch[K]{want}
}

// The [SwitchBranch] provide methods to construct the [Case].
type SwitchBranch[K comparable] struct {
	want K
}

// Passes the value `got` and the result `v`, which will be used
// if `got` matches the passed to [Switch] `want`.
//
// Returns the next [Case].
func (s SwitchBranch[K]) Case[V any](got K, v V) Case[K, V] {
	if s.want == got {
		return Case[K, V]{resolved: true, result: v}
	}
	return Case[K, V]{want: s.want}
}

// Passes the closure `got` and the result `v`, which will be used
// if `got` return matches the passed to [Switch] `want`.
//
// Returns the next [Case].
func (s SwitchBranch[K]) CaseK[V any](got func() K, v V) Case[K, V] {
	if s.want == got() {
		return Case[K, V]{resolved: true, result: v}
	}
	return Case[K, V]{want: s.want}
}

// Passes the value `got` and the closure `v`, which result will be used
// if `got` matches the passed to [Switch] `want`.
//
// Returns the next [Case].
func (s SwitchBranch[K]) CaseV[R func() V, V any](got K, v R) Case[K, V] {
	if s.want == got {
		return Case[K, V]{resolved: true, result: v()}
	}
	return Case[K, V]{want: s.want}
}

// Passes the closure `got` and the closure `v`, which result will be used
// if `got` return matches the passed to [Switch] `want`.
//
// Returns the next [Case].
func (s SwitchBranch[K]) CaseKV[R func() V, V any](got func() K, v R) Case[K, V] {
	if s.want == got() {
		return Case[K, V]{resolved: true, result: v()}
	}
	return Case[K, V]{want: s.want}
}

// The [Case] provide methods to construct the next [Case] and method [Default]
// to finish [SwitchBranch] evaluation.
type Case[K comparable, V any] struct {
	resolved bool
	result   V
	want     K
}

// Passes the value `got` and the result `v`, which will be used
// if `got` matches the passed to [Switch] `want`.
//
// Returns the next [Case].
func (c Case[K, V]) Case(got K, v V) Case[K, V] {
	if c.resolved {
		return c
	}
	if c.want == got {
		return Case[K, V]{resolved: true, result: v}
	}
	return c
}

// Passes the closure `got` and the result `v`, which will be used
// if `got` return matches the passed to [Switch] `want`.
//
// Returns the next [Case].
func (c Case[K, V]) CaseK(got func() K, v V) Case[K, V] {
	if c.resolved {
		return c
	}
	if c.want == got() {
		return Case[K, V]{resolved: true, result: v}
	}
	return c
}

// Passes the value `got` and the closure `v`, which result will be used
// if `got` matches the passed to [Switch] `want`.
//
// Returns the next [Case].
func (c Case[K, V]) CaseV(got K, v func() V) Case[K, V] {
	if c.resolved {
		return c
	}
	if c.want == got {
		return Case[K, V]{resolved: true, result: v()}
	}
	return c
}

// Passes the closure `got` and the closure `v`, which result will be used
// if `got` return matches the passed to [Switch] `want`.
//
// Returns the next [Case].
func (c Case[K, V]) CaseKV(got func() K, v func() V) Case[K, V] {
	if c.resolved {
		return c
	}
	if c.want == got() {
		return Case[K, V]{resolved: true, result: v()}
	}
	return c
}

// Returns the value `v` previously passed to the [Case]
// ​​whose `got` value matched the `want` value from [Switch],
// or the value `v` passed here if no such [Case] ​​exists.
func (c Case[K, V]) Default(v V) V {
	if c.resolved {
		return c.result
	}
	return v
}

// Returns the value `v` previously passed to the [Case]
// ​​whose `got` value matched the `want` value from [Switch],
// or the value returned by the closure `v` passed here
// if no such [Case] ​​exists.
func (c Case[K, V]) DefaultV(v func() V) V {
	if c.resolved {
		return c.result
	}
	return v()
}
