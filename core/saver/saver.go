package saver

import (
	"log"

	"github.com/ozoncp/ocp-progress-api/core/alarmer"
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
	alarmer         alarmer.Alarmer
	progressChannel chan progress.Progress
	progressSlice   []progress.Progress
	lostNotAllData  bool
	endSignal       chan struct{}
}

func New(capacity uint, flusher flusher.Flusher, alarmer alarmer.Alarmer) Saver {
	if flusher == nil || alarmer == nil {
		return nil
	}
	return &saver{
		capacity:        capacity,
		flusher:         flusher,
		alarmer:         alarmer,
		progressChannel: make(chan progress.Progress),
		progressSlice:   make([]progress.Progress, capacity),
		lostNotAllData:  false,
		endSignal:       make(chan struct{}),
	}
}

func (s *saver) Save(progress progress.Progress) {
	s.progressChannel <- progress
}

func (s *saver) Init(lostNotAllData bool) {
	s.alarmer.Init()
	s.lostNotAllData = lostNotAllData

	go func() {
		for {
			select {
			case progressTmp := <-s.progressChannel:
				s.addProgressData(progressTmp)
			case <-s.alarmer.Alarm():
				s.flushData()
			case <-s.endSignal:
				return
			}
		}
	}()

}

func (s *saver) Close() {
	s.endSignal <- struct{}{}
	defer close(s.endSignal)
	s.flushData()
	s.alarmer.Close()
}

func (s *saver) addProgressData(progressTmp progress.Progress) {
	if len(s.progressSlice) >= int(s.capacity) {
		if s.lostNotAllData {
			s.progressSlice = s.progressSlice[1:]
		} else {
			s.progressSlice = s.progressSlice[:0] // clear slice and Keep allocated memory
		}
	}

	s.progressSlice = append(s.progressSlice, progressTmp)
}

func (s *saver) flushData() {
	flushResult := s.flusher.Flush(s.progressSlice)

	if flushResult != nil {
		log.Fatalln("Failed to flush")
		s.progressSlice = s.progressSlice[:copy(s.progressSlice, flushResult)]
	} else {
		s.progressSlice = s.progressSlice[:0]
	}
}
