package task

import (
	"enterpret/task/worker"
	"time"
)

// Pull data from the external source and ingest
func Start() {
	discource := worker.NewDiscourse()
	for {
		discource.Pull()
		time.Sleep(1 * time.Hour)
	}
}
