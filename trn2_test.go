/*
 * Copyright (c) 2026-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package trn_test

import (
	"errors"
	"testing"

	trn "github.com/mail2nnv/ternary"
	"github.com/stretchr/testify/require"
)

func TestIf2ThenElse(t *testing.T) {
	require := require.New(t)

	i, err := trn.If2[int, error](2 > 1).Then(5, nil).Else(0, errors.ErrUnsupported)
	require.Equal(5, i)
	require.NoError(err)

	j, err := trn.If2[int, error](2 < 1).Then(5, nil).Else(0, errors.ErrUnsupported)
	require.Equal(0, j)
	require.ErrorIs(err, errors.ErrUnsupported)
}

func TestIf2ThenRetElseRet(t *testing.T) {
	require := require.New(t)
	i, err := trn.If2[int, error](2 > 1).
		ThenF(func() (int, error) { return 5, nil }).
		ElseF(func() (int, error) { return 0, errors.ErrUnsupported })
	require.Equal(5, i)
	require.NoError(err)

	j, err := trn.If2[int, error](2 < 1).
		ThenF(func() (int, error) { return 5, nil }).
		ElseF(func() (int, error) { return 0, errors.ErrUnsupported })
	require.Equal(0, j)
	require.ErrorIs(err, errors.ErrUnsupported)
}

func Test2IfThenElseIf(t *testing.T) {
	require := require.New(t)
	s := ""
	for range 5 {
		l, err := trn.If2[int, error](s == "").
			Then(0, nil).
			ElseIf(s == "a").
			Then(1, nil).
			ElseIf(s == "aa").
			Then(2, nil).
			ElseIf(s == "aaa").
			Then(3, nil).
			ElseIf(s == "aaaa").
			ThenF(func() (int, error) { return len(s), errors.New("too long") }).
			ElsePanic("ops")

		require.Len(s, l)
		if l < 4 {
			require.NoError(err)
		} else {
			require.ErrorContains(err, "too long")
		}

		s += "a"
	}
}

func TestIf2ThenElsePanic(t *testing.T) {
	require := require.New(t)
	require.Panics(
		func() {
			_, _ = trn.If2[int, int](2*2 == 5).
				Then(0, 0).
				ElsePanic("🤪")
		})

	require.NotPanics(
		func() {
			i, j := trn.If2[int, int](2*2 == 4).
				Then(0, 0).
				ElsePanic("🤪")
			require.Zero(i)
			require.Zero(j)
		})
}
