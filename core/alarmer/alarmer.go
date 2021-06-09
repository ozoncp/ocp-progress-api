package alarmer

import (
	"log"
	"time"
)

//Alarmer it is timer to creat signals
//Deprecated: It is good but useless
type Alarmer interface {
	Alarm() <-chan struct{}
	Init()
	Close()
}

type alarmer struct {
	interval time.Duration
	alarm    chan struct{}
	stop     chan struct{}
}

func New(interval time.Duration) Alarmer {
	return &alarmer{
		interval: interval,
		alarm:    make(chan struct{}),
		stop:     make(chan struct{}),
	}
}

func (a *alarmer) Init() {
	go func() {
		tick := time.NewTicker(a.interval)
		defer tick.Stop()
		defer close(a.alarm)

		for {
			select {
			case <-a.stop:
				return
			case <-tick.C:
				select {
				case a.alarm <- struct{}{}:
				default:
					log.Fatalln("Failed to send alarm")
				}
			}
		}
	}()

}

func (a *alarmer) Alarm() <-chan struct{} {
	return a.alarm
}

func (a *alarmer) Close() {
	a.stop <- struct{}{}
	close(a.stop)
}
