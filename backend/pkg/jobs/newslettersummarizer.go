package jobs

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/reugn/go-quartz/quartz"
)

var _ quartz.Job = &SummarizeNewslettersJob{}

type SummarizeNewslettersJob struct {
	instanceID  string
	description string
}

func NewSummarizeNewlettersJob() *SummarizeNewslettersJob {
	return &SummarizeNewslettersJob{
		instanceID:  uuid.New().String(),
		description: "Fetches all unread emails and summarizes the content and discounts using Azure OpenAI",
	}
}

func (job *SummarizeNewslettersJob) StartScheduler(ctx context.Context) error {
	sched := quartz.NewStdScheduler()
	sched.Start(ctx)
	updateIntervalExpression := "0 " + fmt.Sprint(0) + "/" + fmt.Sprint(2) + " * * * *"
	cronTrigger, _ := quartz.NewCronTrigger(updateIntervalExpression)
	err := sched.ScheduleJob(ctx, job, cronTrigger)
	return err
}

func (job *SummarizeNewslettersJob) Execute(context.Context) {
	fmt.Println("Executing SummarizeNewslettersJob")
}

func (job *SummarizeNewslettersJob) Description() string {
	return job.description
}

func (job *SummarizeNewslettersJob) Key() int {
	return quartz.HashCode(job.instanceID)
}
