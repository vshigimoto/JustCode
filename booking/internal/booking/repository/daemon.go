package repository

import (
	"booking/internal/booking/worker"
)

func (r *Repo) Daemon() {
	for {
		select {
		case <-worker.Result:
			id := <-worker.Result
			r.main.Exec("DELETE from bookrequest WHERE id=$1", id)
		}
	}
}
