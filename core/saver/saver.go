package saver

import (
	"errors"
	"log"
	"time"

	"github.com/ozoncp/ocp-progress-api/core/flusher"
	"github.com/ozoncp/ocp-progress-api/core/progress"
)

type Saver interface {
	Save(progress progress.Progress)
	Init(lostNotAllData bool)
	Close()
}

type saver struct {
	capacity        uint
	flusher         flusher.Flusher
	interval        time.Duration
	progressChannel chan progress.Progress
	progressSlice   []progress.Progress
	lostNotAllData  bool
	endSignal       chan struct{}
}

func New(capacity uint, flusher flusher.Flusher, interval time.Duration) (Saver, error) {
	if flusher == nil {
		return nil, errors.New("flusher can`t be NIL")
	}
	return &saver{
		capacity:        capacity,
		flusher:         flusher,
		interval:        interval,
		progressChannel: make(chan progress.Progress),
		progressSlice:   make([]progress.Progress, capacity),
		lostNotAllData:  false,
		endSignal:       make(chan struct{}),
	}, nil
}

func (s *saver) Save(progress progress.Progress) {
	s.progressChannel <- progress
}

func (s *saver) Init(lostNotAllData bool) {
	s.lostNotAllData = lostNotAllData

	go func() {

		timer := time.NewTicker(s.interval)
		for {
			select {
			case progressTmp := <-s.progressChannel:
				s.addProgressData(progressTmp)
			case <-timer.C:
				s.flushData()
			case <-s.endSignal:
				log.Printf("SAVER: Signal for finish, finishing....  ")
				return
			}
		}
	}()

}

func (s *saver) Close() {
	defer close(s.endSignal)

	s.endSignal <- struct{}{}
	s.flushData()
}

func (s *saver) addProgressData(progressTmp progress.Progress) {
	if len(s.progressSlice) >= int(s.capacity) {
		if s.lostNotAllData {
			log.Printf("SAVER: Unable to add new data to the flush buffer, I skip some previous data ")
			s.progressSlice = s.progressSlice[1:]
		} else {
			log.Printf("SAVER: Unable to add new data to the flush buffer, I skip all previous data ")
			s.progressSlice = s.progressSlice[:0] // clear slice and Keep allocated memory
		}
	}

	s.progressSlice = append(s.progressSlice, progressTmp)
}

func (s *saver) flushData() {
	flushResult := s.flusher.Flush(s.progressSlice)

	if flushResult != nil {
		log.Print("SAVER: Failed to flush")
	}
	s.progressSlice = s.progressSlice[:copy(s.progressSlice, flushResult)]
}
