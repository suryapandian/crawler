package worker

type Job interface {
	Run()
	Stop()
}

func RunJobs(jobs []Job) {
	for _, job := range jobs {
		go job.Run()
	}
}

func StopJobs(jobs []Job) {
	for _, job := range jobs {
		job.Stop()
	}
}
