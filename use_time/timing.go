package use_time

import (
	"github.com/sirupsen/logrus"
	"time"
)

//实现定时的五分钟方式

func UseTimeAfter() {
	afterTime := time.After(time.Second * 2)
	<-afterTime
	go func() {
		logrus.Println("UseTimeAfter======")
	}()
	logrus.Println("UseTimeAfter", time.Now().Format("2006-01-02 15:04:05"))
	UseTimeAfter()
}

func UseTimeSleep() {
	time.Sleep(time.Second * 2)
	go func() {
		logrus.Println("UseTimeSleep======")
	}()
	logrus.Println("UseTimeSleep", time.Now().Format("2006-01-02 15:04:05"))
	UseTimeSleep()
}

func UseNewTimer() {
	newTime := time.NewTimer(time.Second * 2)
	<-newTime.C
	go func() {
		logrus.Println("UseNewTimer======")
	}()
	logrus.Println("UseNewTimer", time.Now().Format("2006-01-02 15:04:05"))
	UseNewTimer()
}

func UseNewTicker1() {
	newTime := time.NewTicker(time.Second * 2)
	for range newTime.C {
		go func() {
			logrus.Println("UseNewTicker1======")
		}()
		logrus.Println("UseNewTicker1", time.Now().Format("2006-01-02 15:04:05"))
	}
	UseNewTicker1()
}

func UseNewTicker2() {
	newTime := time.NewTicker(time.Second * 2)
	<-newTime.C
	go func() {
		logrus.Println("UseNewTicker2======")
	}()
	logrus.Println("UseNewTicker2", time.Now().Format("2006-01-02 15:04:05"))

	UseNewTicker2()
}

