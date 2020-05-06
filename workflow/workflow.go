package workflow

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type Workflow struct {
	Id         int
	D          time.Duration
	runTimer   *time.Timer
	Ctx        context.Context
	CancelFunc context.CancelFunc
}

func (req *Workflow) Start() {
	logrus.Println("start=================")
	if req.runTimer != nil && req.runTimer.Stop() {
		return
	}
	req.runTimer = time.NewTimer(req.D)
	go func() {
		select {
		case _, ok := <-req.Ctx.Done():
			logrus.Println("Workflow ctx.Done:", ok)
			return
		case <-req.runTimer.C:
			logrus.Println("Workflow Start:", req.Id)
			time.Sleep(time.Second * 1)
			return
		}
	}()
}
