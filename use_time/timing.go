package use_time

import (
	"github.com/sirupsen/logrus"
	"time"
)

//实现定时的五种方式

func UseTimeAfter(d time.Duration) {
	afterTime := time.After(d)
	<-afterTime
	go func() {
		logrus.Println("UseTimeAfter======")
	}()
	logrus.Println("UseTimeAfter", time.Now().Format("2006-01-02 15:04:05"))
	UseTimeAfter(d)
}

func UseTimeSleep(d time.Duration) {
	time.Sleep(d)
	go func() {
		logrus.Println("UseTimeSleep======")
	}()
	logrus.Println("UseTimeSleep", time.Now().Format("2006-01-02 15:04:05"))
	UseTimeSleep(d)
}

func UseNewTimer(d time.Duration) {
	newTime := time.NewTimer(d)
	<-newTime.C
	go func() {
		logrus.Println("UseNewTimer======")
	}()
	logrus.Println("UseNewTimer", time.Now().Format("2006-01-02 15:04:05"))
	UseNewTimer(d)
}

func UseNewTicker1(d time.Duration) {
	newTime := time.NewTicker(d)
	for range newTime.C {
		go func() {
			logrus.Println("UseNewTicker1======")
		}()
		logrus.Println("UseNewTicker1", time.Now().Format("2006-01-02 15:04:05"))
	}
	UseNewTicker1(d)
}

func UseNewTicker2(d time.Duration) {
	newTime := time.NewTicker(d)
	<-newTime.C
	go func() {
		logrus.Println("UseNewTicker2======")
	}()
	logrus.Println("UseNewTicker2", time.Now().Format("2006-01-02 15:04:05"))

	UseNewTicker2(d)
}
