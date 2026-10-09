/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn

func Switch1[T comparable](want T) SwitchBranch[T] {
	return SwitchBranch[T]{want}
}

type SwitchBranch[T comparable] struct {
	want T
}

func (s SwitchBranch[T]) Case[V any](got T, v V) Case[T, V] {
	if s.want == got {
		return Case[T, V]{resolved: true, result: v}
	}
	return Case[T, V]{want: s.want}
}

func (s SwitchBranch[T]) CaseF[V any](got func() T, v V) Case[T, V] {
	if s.want == got() {
		return Case[T, V]{resolved: true, result: v}
	}
	return Case[T, V]{want: s.want}
}

func (s SwitchBranch[T]) CaseR[R func() V, V any](got T, v R) Case[T, V] {
	if s.want == got {
		return Case[T, V]{resolved: true, result: v()}
	}
	return Case[T, V]{want: s.want}
}

func (s SwitchBranch[T]) CaseFR[R func() V, V any](got func() T, v R) Case[T, V] {
	if s.want == got() {
		return Case[T, V]{resolved: true, result: v()}
	}
	return Case[T, V]{want: s.want}
}

type Case[T comparable, V any] struct {
	resolved bool
	result   V
	want     T
}

func (c Case[T, V]) Case(got T, v V) Case[T, V] {
	if c.resolved {
		return c
	}
	if c.want == got {
		return Case[T, V]{resolved: true, result: v}
	}
	return c
}

func (c Case[T, V]) CaseF(got func() T, v V) Case[T, V] {
	if c.resolved {
		return c
	}
	if c.want == got() {
		return Case[T, V]{resolved: true, result: v}
	}
	return c
}

func (c Case[T, V]) CaseR(got T, v func() V) Case[T, V] {
	if c.resolved {
		return c
	}
	if c.want == got {
		return Case[T, V]{resolved: true, result: v()}
	}
	return c
}

func (c Case[T, V]) CaseFR(got func() T, v func() V) Case[T, V] {
	if c.resolved {
		return c
	}
	if c.want == got() {
		return Case[T, V]{resolved: true, result: v()}
	}
	return c
}

func (c Case[T, V]) Default(v V) V {
	if c.resolved {
		return c.result
	}
	return v
}

func (c Case[T, V]) DefaultR(v func() V) V {
	if c.resolved {
		return c.result
	}
	return v()
}
