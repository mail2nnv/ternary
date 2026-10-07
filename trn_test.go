/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn_test

import (
	"fmt"
	"testing"

	trn "github.com/mail2nnv/ternary"
	"github.com/stretchr/testify/require"
)

func TestIfThenElse(t *testing.T) {
	require := require.New(t)

	require.Equal(5, trn.If(2 > 1).Then(5).Else(0))
	require.Equal(0, trn.If(2 < 1).Then(5).Else(0))
}

func TestIfThenFElseF(t *testing.T) {
	req := require.New(t)
	req.Equal(5, trn.If(2 > 1).ThenF(func() int { return 5 }).ElseF(func() int { return 0 }))
	req.Equal(0, trn.If(2 < 1).ThenF(func() int { return 5 }).ElseF(func() int { return 0 }))

	t.Run("Safety", func(t *testing.T) {
		tests := []struct {
			s    string
			want string
		}{
			{"", "<empty>"},
			{"abc", "a"},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("%q->%q", tt.s, tt.want), func(t *testing.T) {
				req := require.New(t)
				got := trn.If(len(tt.s) == 0).Then("<empty>").ElseF(func() string { return tt.s[0:1] })
				req.Equal(tt.want, got)

				got = trn.If(len(tt.s) > 0).ThenF(func() string { return tt.s[0:1] }).Else("<empty>")
				req.Equal(tt.want, got)
			})
		}
	})
}

func TestIfThenElseNested(t *testing.T) {
	tests := []struct {
		s    []any
		want string
	}{
		{nil, "nil"},
		{[]any{}, "empty slice"},
		{[]any{1}, "one element slice"},
		{[]any{1, 2}, "2-element slice"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			req := require.New(t)
			got :=
				trn.If(tt.s == nil).Then("nil").Else(
					trn.If(len(tt.s) == 0).Then("empty slice").Else(
						trn.If(len(tt.s) == 1).Then("one element slice").Else(
							fmt.Sprintf("%d-element slice", len(tt.s)))))
			req.Equal(tt.want, got)
		})
	}
}
