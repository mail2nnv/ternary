/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn

// Switch(a). 							// -> Sw[int]
// 	Case(1).								// -> FirstCase[int]
// 		Return[string]("1").	// -> Return[int, string]
// 	Case(2).								// -> NextCase[int, string]
// 		Return("2").					// -> Return[int, string]
// 	Case(3).								// -> NextCase[int, string]
// 		Return("3").					// -> Return[int, string]
//	Default("more")

func Switch[T comparable](v T) Sw[T] {
	return Sw[T]{want: v}
}

type Sw[T comparable] struct {
	want T
}

func (s Sw[T]) Case(v T) FirstCase[T] {
	if s.want == v {
		return FirstCase[T]{state: resolved}
	}
	return FirstCase[T]{want: s.want}
}

type FirstCase[T comparable] struct {
	state int
	want  T
}

func (c FirstCase[T]) Return[V any](v V) Return[T, V] {
	if c.state == resolved {
		return Return[T, V]{
			state:  resolved,
			result: v,
		}
	}
	return Return[T, V]{
		want: c.want,
	}
}

type Return[T comparable, V any] struct {
	state  int
	result V
	want   T
}

func (r Return[T, V]) Case(v T) NextCase[T, V] {
	if r.state == resolved {
		return NextCase[T, V]{
			state:  resolved,
			result: r.result,
		}
	}

	if r.want == v {
		return NextCase[T, V]{
			state: temp,
		}
	}

	return NextCase[T, V]{
		want: r.want,
	}
}

func (r Return[T, V]) Default(v V) V {
	if r.state == resolved {
		return r.result
	}
	return v
}

type NextCase[T comparable, V any] struct {
	state  int
	result V
	want   T
}

func (c NextCase[T, V]) Return(v V) Return[T, V] {
	switch c.state {
	case resolved:
		return Return[T, V]{
			state:  resolved,
			result: c.result,
		}
	case temp:
		return Return[T, V]{
			state:  resolved,
			result: v,
		}
	default:
		return Return[T, V]{
			want: c.want,
		}
	}
}

const (
	unresolved int = iota
	resolved
	temp
)
