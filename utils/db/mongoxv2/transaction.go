package mongoxv2

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TransactionFunc is a handler to manage a transaction.
type TransactionFunc func(session mongo.Session, sc mongo.SessionContext) error

type TxKey struct{}

// DoTx executes a function within a MongoDB transaction using default options.
func DoTx(ctx context.Context, f func(ctx context.Context) error) error {
	// Execute the transaction function with default options.
	return DoTxWithOpts(ctx, f, defaultOption)
}

type Option struct {
	SessionOpts *options.SessionOptions
	TxOpts      *options.TransactionOptions
}

var defaultOption = Option{}

// DoTxWithOpts executes a function within a MongoDB transaction with the specified options.
//
// If a transaction is already present in the context, it simply executes the function without starting a new transaction.
func DoTxWithOpts(ctx context.Context, f func(ctx context.Context) error, opt Option) error {
	// Check if there is already a transaction session in the context.
	if preSession := ctx.Value(TxKey{}); preSession != nil {
		// If a transaction is already present, execute the function without starting a new transaction.
		return f(ctx)
	}

	// Add a placeholder value to the context to indicate the transaction session.
	ctx = context.WithValue(ctx, TxKey{}, struct{}{})

	// Define the transaction function to be executed.
	txFunc := func(session mongo.Session, sc mongo.SessionContext) error {
		// Execute the provided function within the transaction context.
		if err := f(sc); err != nil {
			return err
		}
		// Commit the transaction upon successful execution of the function.
		return session.CommitTransaction(sc)
	}

	// Start the transaction with the given options and execute the transaction function.
	return tx(ctx, txFunc, opt)
}

// tx executes a transaction function within a MongoDB transaction with the specified options.
func tx(ctx context.Context, f TransactionFunc, opt Option) error {
	// Start a new session with the given options.
	session, err := client.StartSession(opt.SessionOpts)
	if err != nil {
		return err
	}
	// We must end the session when we're done with it.
	defer session.EndSession(ctx)

	// Start a transaction with the given options.
	if err = session.StartTransaction(opt.TxOpts); err != nil {
		return err
	}

	// Define a wrapper function around the provided transaction function.
	wrapperFn := func(sc mongo.SessionContext) error {
		return f(session, sc)
	}

	// Execute the transaction function with the session context.
	return mongo.WithSession(ctx, session, wrapperFn)
}
