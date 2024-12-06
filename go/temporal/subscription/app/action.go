package app

import (
	"context"
	"fmt"
	"log"
	"subscription/clientx"
	"time"

	"go.temporal.io/sdk/client"
)

// Define your activity function
func SampleActivity(ctx context.Context, data UserSubcribeSuccessData) (string, error) {
	log.Printf("SampleActivity is running")
	return fmt.Sprintf("Hello, %s!", data.UserID), nil
}

func SetScheduleRenew(ctx context.Context, data UserSubcribeSuccessData) (string, error) {
	scheduleID := fmt.Sprintf("renew_subscription_schedule_user_%s", data.UserID)
	workflowID := "renew_subscription_workflow"
	// Create the schedule.
	scheduleHandle, err := clientx.GetClient().ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID: scheduleID,
		Spec: client.ScheduleSpec{
			StartAt: time.Now().Add(5 * time.Minute),
		},
		Action: &client.ScheduleWorkflowAction{
			ID:        workflowID,
			Workflow:  RenewSubsriptionWorkflow,
			TaskQueue: RenewSubscriptionScheduleQueueName,
			Args:      []interface{}{RenewSubscriptionData{UserID: data.UserID}},
		},
	})
	if err != nil {
		return "", fmt.Errorf("unable to create schedule: %w", err)
	}
	log.Printf("Schedule ID: %s, Schedule Handle: %s\n", scheduleID, scheduleHandle)

	return fmt.Sprintf("Set renew subscription schedule for user: %s!", data.UserID), nil
}
