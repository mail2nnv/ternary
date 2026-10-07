/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn_test

import (
	"errors"
	"fmt"
	"testing"

	trn "github.com/mail2nnv/ternary"
	"github.com/stretchr/testify/require"
)

func TestIf2ThenElse(t *testing.T) {
	req := require.New(t)

	i, err := trn.If2(2 < 1).Then(0, errors.ErrUnsupported).Else(5, nil)
	req.Equal(5, i)
	req.NoError(err)

	j, err := trn.If2(2 > 1).Then(0, errors.ErrUnsupported).Else(5, nil)
	req.Equal(0, j)
	req.ErrorIs(err, errors.ErrUnsupported)
}

func TestIf2ThenFElseF(t *testing.T) {
	req := require.New(t)
	i, err := trn.If2(2 > 1).
		ThenF(func() (int, error) { return 5, nil }).
		ElseF(func() (int, error) { return 0, errors.ErrUnsupported })
	req.Equal(5, i)
	req.NoError(err)

	j, err := trn.If2(2 < 1).
		ThenF(func() (int, error) { return 5, nil }).
		ElseF(func() (int, error) { return 0, errors.ErrUnsupported })
	req.Equal(0, j)
	req.ErrorIs(err, errors.ErrUnsupported)
}

func TestIf2ThenElseNested(t *testing.T) {
	tests := []struct {
		s    []any
		want string
		err  error
	}{
		{nil, "nil", errors.ErrUnsupported},
		{[]any{}, "empty", nil},
		{[]any{1}, "one element", nil},
		{[]any{1, 2}, "2 elements", nil},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			req := require.New(t)
			got, err :=
				trn.If2(tt.s == nil).Then("nil", errors.ErrUnsupported).Else(
					trn.If2(len(tt.s) == 0).Then("empty", error(nil)).Else(
						trn.If2(len(tt.s) == 1).Then("one element", error(nil)).Else(
							fmt.Sprintf("%d elements", len(tt.s)), nil)))
			req.Equal(tt.want, got)
			if tt.err == nil {
				req.NoError(err)
			} else {
				req.ErrorIs(err, tt.err)
			}
		})
	}
}
