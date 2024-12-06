package main

import (
	"log"
	"subscription/app"
	"subscription/clientx"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

var cli client.Client

func GetClient() client.Client {
	return cli
}

func main() {
	var err error
	cli, err = clientx.Init()
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer cli.Close()

	w := worker.New(cli, app.UserSubcribeSuccessTaskQueueName, worker.Options{})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(app.UserSubcribeSuccess)
	w.RegisterActivity(app.SampleActivity)
	w.RegisterActivity(app.SetScheduleRenew)

	w2 := worker.New(cli, app.RenewSubscriptionScheduleQueueName, worker.Options{})
	w2.RegisterWorkflow(app.RenewSubsriptionWorkflow)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}

	// err = w2.Run(worker.InterruptCh())
	// if err != nil {
	// 	log.Fatalln("unable to start Worker", err)
	// }
}
