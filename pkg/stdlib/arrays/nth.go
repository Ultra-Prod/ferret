package arrays

import (
	"context"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
)

/*
 * Returns the element of an array at a given position.
 * It is the same as anyArray[position] for positive positions, but does not support negative positions.
 */
func Nth(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 2, 2)

	if err != nil {
		return values.None, err
	}

	err = core.ValidateType(args[0], core.ArrayType)

	if err != nil {
		return values.None, err
	}

	err = core.ValidateType(args[1], core.IntType)

	if err != nil {
		return values.None, err
	}

	arr := args[0].(*values.Array)
	idx := args[1].(values.Int)

	return arr.Get(idx), nil
}
