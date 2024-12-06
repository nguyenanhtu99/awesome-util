package main

import (
	"context"
	"log"
	"subscription/app"
	"subscription/clientx"

	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object just once per process
	c, err := clientx.Init()
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	data := app.UserSubcribeSuccessData{
		UserID:  "123",
		SubsID:  1,
		Version: 1,
	}

	options := client.StartWorkflowOptions{
		ID:        "user-subscribe-success-701",
		TaskQueue: app.UserSubcribeSuccessTaskQueueName,
	}

	log.Printf("User ID: %s subscribe success subscription ID: %d, version: %d\n", data.UserID, data.SubsID, data.Version)

	we, err := c.ExecuteWorkflow(context.Background(), options, app.UserSubcribeSuccess, data)
	if err != nil {
		log.Fatalln("Unable to start the Workflow:", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	// scheduleID := fmt.Sprintf("renew_subscription_schedule_user_%s", data.UserID)
	// workflowID := "renew_subscription_workflow"
	// // Create the schedule.
	// scheduleHandle, err := clientx.GetClient().ScheduleClient().Create(context.Background(), client.ScheduleOptions{
	// 	ID: scheduleID,
	// 	Spec: client.ScheduleSpec{

	// 	},
	// 	Action: &client.ScheduleWorkflowAction{
	// 		ID:        workflowID,
	// 		Workflow:  app.RenewSubsriptionWorkflow,
	// 		TaskQueue: app.RenewSubscriptionScheduleQueueName,
	// 		Args:      []interface{}{app.RenewSubscriptionData{UserID: data.UserID}},
	// 	},
	// })
	// if err != nil {
	// 	return
	// }

	// log.Printf("Schedule ID: %s, Schedule Handle: %s\n", scheduleID, scheduleHandle)

	var result string

	err = we.Get(context.Background(), &result)

	if err != nil {
		log.Fatalln("Unable to get Workflow result:", err)
	}

	log.Println(result)
}
