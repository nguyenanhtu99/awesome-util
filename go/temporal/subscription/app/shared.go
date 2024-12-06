package app

const (
	UserSubcribeSuccessTaskQueueName   = "USER_SUBSCRIBE_SUCCESS_TASK_QUEUE"
	RenewSubscriptionScheduleQueueName = "RENEW_SUBSCRIPTION_SCHEDULE_QUEUE"
)

type UserSubcribeSuccessData struct {
	UserID  string
	SubsID  int
	Version int
}

type RenewSubscriptionData struct {
	UserID string
}
