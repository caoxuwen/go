package bridge

import (
	b "github.com/caoxuwen/go/build"
	shared "github.com/caoxuwen/go/services/internal/bridge-compliance-shared"
	"github.com/caoxuwen/go/services/internal/bridge-compliance-shared/http/helpers"
)

// LiquidationOperationBody represents liquidation operation
type LiquidationOperationBody struct {
	Source *string
}

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op LiquidationOperationBody) ToTransactionMutator() b.TransactionMutator {
	var mutators []interface{}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.Inflation(mutators...)
}

// Validate validates if operation body is valid.
func (op LiquidationOperationBody) Validate() error {
	if op.Source != nil && !shared.IsValidAccountID(*op.Source) {
		return helpers.NewInvalidParameterError("source", "Source must be a public key (starting with `G`).")
	}

	return nil
}
