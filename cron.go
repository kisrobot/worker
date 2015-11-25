package worker

import (
	"os"
	"os/exec"
)

type Cron struct {
}

func NewCronQueue() *Cron {
	return &Cron{}
}

func (Cron) Add(job QorJobInterface) error {
	binaryFile := os.Args[0]
	cmd := exec.Command(binaryFile, "--qor-job", job.GetJobID())
	if err := cmd.Start(); err == nil {
		cmd.Process.Release()
		return nil
	} else {
		return err
	}
}

func (Cron) Kill(job QorJobInterface) error {
	return nil
}

func (Cron) Remove(job QorJobInterface) error {
	return nil
}
