package math

import (
	"math"

	"github.com/glispy/glispy/common"
	"github.com/glispy/glispy/eval"
	"github.com/glispy/glispy/types"
)

// Add will add a series of numbers
func Add(sc types.Scope, args types.List) (out types.Expression, err error) {
	var (
		n   types.Number
		num types.Number
		ok  bool
	)

	out = 0

	for _, exp := range args {
		switch val := exp.(type) {
		case types.Number:
			n += val
		case types.List:
			var exp types.Expression
			if exp, err = eval.Eval(sc, val); err != nil {
				return
			}

			if num, ok = exp.(types.Number); !ok {
				err = common.ErrExpectedNumber
				return
			}

			n += num

		default:
			err = common.ErrExpectedNumber
			return
		}
	}

	out = n
	return
}

// Multiply will multiply a series of numbers
func Multiply(sc types.Scope, args types.List) (out types.Expression, err error) {
	var (
		n   types.Number
		num types.Number
	)

	for i, exp := range args {
		if num, err = eval.GetNumber(sc, exp); err != nil {
			return
		}

		if i == 0 {
			n = num
		} else {
			n *= num
		}
	}

	out = n
	return
}

// LessThan will return if a is less than b
func LessThan(sc types.Scope, an types.Number, b types.Atom) (out types.Expression, err error) {
	var bn types.Number
	if bn, err = types.ToNumber(b); err != nil {
		return
	}

	if an < bn {
		out = "true"
	}

	return
}

// GreaterThan will return if a is greater than b
func GreaterThan(sc types.Scope, an types.Number, b types.Atom) (out types.Expression, err error) {
	var bn types.Number
	if bn, err = types.ToNumber(b); err != nil {
		return
	}

	if an > bn {
		out = "true"
	}

	return
}

// Square will return the squared value of the provided number
func Square(sc types.Scope, args types.List) (out types.Expression, err error) {
	var val types.Number
	if val, err = args.GetNumber(0); err != nil {
		return
	}

	sqrt := math.Sqrt(float64(val))
	out = types.Number(sqrt)
	return
}
