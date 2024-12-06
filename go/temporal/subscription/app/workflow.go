package app

import (
	"log"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func UserSubcribeSuccess(ctx workflow.Context, data UserSubcribeSuccessData) error {
	if err := workflow.Sleep(ctx, 5 * time.Minute); err != nil {
		return err
	}
	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        500, // 0 is unlimited retries
		NonRetryableErrorTypes: []string{"InvalidAccountError", "InsufficientFundsError"},
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, SampleActivity, data).Get(ctx, &result)
	if err != nil {
		return err
	}

	log.Printf("SampleActivity result: %s\n", result)

	err = workflow.ExecuteActivity(ctx, SetScheduleRenew, data).Get(ctx, &result)
	if err != nil {
		return err
	}

	log.Printf("SetScheduleRenew result: %s\n", result)

	return nil
}

func RenewSubsriptionWorkflow(ctx workflow.Context, data RenewSubscriptionData) error {
	log.Printf("RenewSubsriptionWorkflow is running with data: %v\n", data)

	return nil
}
