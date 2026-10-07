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

	i, err := trn.If2(2 > 1).Then(5, error(nil)).Else(0, errors.ErrUnsupported)
	require.Equal(5, i)
	require.NoError(err)

	j, err := trn.If2(2 < 1).Then(5, error(nil)).Else(0, errors.ErrUnsupported)
	require.Equal(0, j)
	require.ErrorIs(err, errors.ErrUnsupported)
}

func TestIf2ThenRetElseRet(t *testing.T) {
	require := require.New(t)
	i, err := trn.If2(2 > 1).
		ThenF(func() (int, error) { return 5, nil }).
		ElseF(func() (int, error) { return 0, errors.ErrUnsupported })
	require.Equal(5, i)
	require.NoError(err)

	j, err := trn.If2(2 < 1).
		ThenF(func() (int, error) { return 5, nil }).
		ElseF(func() (int, error) { return 0, errors.ErrUnsupported })
	require.Equal(0, j)
	require.ErrorIs(err, errors.ErrUnsupported)
}

func Test2IfThenElseIf(t *testing.T) {
	require := require.New(t)
	s := ""
	for range 5 {
		l, err := trn.If2(s == "").
			Then(0, error(nil)).
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

func TestIf2ThenElseIfThenFElseF(t *testing.T) {
	require := require.New(t)

	test := "abcd"

	s, i := trn.If2(len(test) <= 3).
		Then("short", 1).
		ElseIfF(func() bool { return test[3] == 'a' }).
		ThenF(func() (string, int) { return test[3:4], 2 }).
		ElseF(func() (string, int) { return "long", 3 })
	require.Equal("long", s)
	require.Equal(3, i)

	s1, i1 := trn.If2(len(test) <= 3).
		Then("short", 1).
		ElseIfF(func() bool { return test[3] == 'd' }).
		ThenF(func() (string, int) { return test[3:4], 2 }).
		ElseF(func() (string, int) { return "long", 3 })
	require.Equal("d", s1)
	require.Equal(2, i1)
}

func TestIf2ThenElsePanic(t *testing.T) {
	require := require.New(t)
	require.Panics(
		func() {
			_, _ = trn.If2(2*2 == 5).
				Then(0, 0).
				ElsePanic("🤪")
		})

	require.NotPanics(
		func() {
			i, j := trn.If2(2*2 == 4).
				Then(0, 0).
				ElsePanic("🤪")
			require.Zero(i)
			require.Zero(j)
		})
}
