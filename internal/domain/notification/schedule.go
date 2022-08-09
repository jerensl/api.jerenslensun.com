package notification

import (
	"time"
)

type Job struct {
	name		string
	title 		string
	messages 	string
	duration 	time.Duration
}

func (t Job) Execute(task func(title, message string))  {
	time.Sleep(t.duration)
	task(t.title, t.messages)
}

type Scheduler struct {
	Job chan Job
	Worker int
}

func NewScheduler(worker int, task func (title, message string)) *Scheduler {
	jobs := make(chan Job, worker)

	for i := 0; i < worker; i++ {
		go func() {
			for j := range jobs {
				j.Execute(task)
			}
		}()
	}

	return &Scheduler{
		Job: jobs,
		Worker: worker,
	}
}


func (s Scheduler) NewJob(name,  title,  message string, duration time.Duration) {
	job := Job{
		name: name,
		title: title,
		messages: message,
		duration: duration,
	}

	s.Job <- job
}