package build

import (
	"github.com/caoxuwen/go/support/errors"
	"github.com/caoxuwen/go/xdr"
)

// Liquidation groups the creation of a new LiquidationBuilder with a call to Mutate.
func Liquidation(muts ...interface{}) (result LiquidationBuilder) {
	result.Mutate(muts...)
	return
}

// InflationBuilder represents an operation that is being built.
type LiquidationBuilder struct {
	O   xdr.Operation
	Err error
}

// Mutate applies the provided mutators to this builder's operation.
func (b *LiquidationBuilder) Mutate(muts ...interface{}) {
	for _, m := range muts {
		var err error
		switch mut := m.(type) {
		case OperationMutator:
			err = mut.MutateOperation(&b.O)
		default:
			err = errors.New("Mutator type not allowed")
		}

		if err != nil {
			b.Err = errors.Wrap(err, "LiquidationBuilder error")
			return
		}
	}
}
